package realm

import (
	"fmt"
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/account"
	"github.com/Gophercraft/core/packet/chat"
	"github.com/Gophercraft/core/packet/login"
	"github.com/Gophercraft/core/packet/social"
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/core/packet/synctime"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/packet/voip"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

var (
	DefaultSpeeds = update.Speeds{
		update.Walk:           2.5,
		update.Run:            7,
		update.RunBackward:    4.5,
		update.Swim:           4.722222,
		update.SwimBackward:   2.5,
		update.Turn:           3.141594,
		update.Flight:         7.0,
		update.FlightBackward: 4.7222,
		update.Pitch:          3.14,
	}
)

func (s *Session) SendSystemFeatures() {
	if s.Build().AddedIn(vsn.V2_4_3) {
		s.Send(&voip.FeatureSystemStatus{
			Status:       voip.AutoIgnore,
			VoiceEnabled: true,
		})
	}
}

func (s *Session) SetupOnLogin() {
	// s.SendNameQueryResponseFor(s.Char)
	if s.Build().AddedIn(3694) {
		s.SendVerifyLoginPacket()
		s.SendAccountDataTimes()
	}

	if s.Build().AddedIn(3592) {
		s.SendTutorialFlags()
	}

	s.SetTimeSpeed()

	s.SendSystemFeatures()

	s.BindpointUpdate()

	if s.Build().AddedIn(vsn.V2_0_1) {
		s.Send(&chat.MOTD{
			Lines: []string{
				"G O P H E R C R A F T",
				fmt.Sprintf("Version %s", vsn.GophercraftVersion),
			},
		})
	}

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SendRestStart()
	}

	s.SendSpellList()
	s.SendActionButtons()

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SendInitFactions()
		s.SendInitWorldStates()
	}

	s.SpawnPlayer()

	s.SyncTime()

	s.BroadcastStatus(social.FriendOnline)
	s.InitGroup()
	s.SendSocialList()

	// Show cinematic sequence on first login
	if s.Char.FirstLogin && s.Server.BoolVar("Char.StartingCinematic") {
		var race *dbdefs.Ent_ChrRaces
		s.DB().Lookup(wdb.BucketKeyUint32ID, uint32(s.Char.Race), &race)

		if race != nil && race.CinematicSequenceID != 0 {
			p := packet.NewWorldPacket(packet.SMSG_TRIGGER_CINEMATIC)
			p.WriteInt32(race.CinematicSequenceID)
			s.SendPacket(p)
		}

		// Don't show same cinematic twice.
	}

	s.Char.FirstLogin = false
	if _, err := s.DB().Cols("first_login").Update(s.Char); err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(500 * time.Millisecond)

		s.SystemChat("|TInterface\\OptionsFrame\\NvidiaLogo:128:300:0:0:128:64:0:0:0:0|t")
		s.SystemChat("|TInterface\\Icons\\rats:256:512:0:0:128:64:0:0:0:0|t")
	}()

	s.JoinedWorldAt = time.Now()
	s.LevelJoinTime = time.Now()

	// s.SendLoginSpell()
}

func (s *Session) SendInitWorldStates() {
	p := packet.NewWorldPacket(packet.SMSG_INIT_WORLD_STATES)
	p.WriteUint32(0)
	s.SendPacket(p)
}

func (s *Session) SetTimeSpeed() {
	var sts login.SetTimeSpeed
	sts.Time = time.Now()
	sts.Speed = login.DefaultTimeSpeed

	s.Send(&sts)

	log.Println("Send gamespeed")
}

func (s *Session) SendRestStart() {
	if s.Build().RemovedIn(vsn.V3_3_5a) {
		v := packet.NewWorldPacket(packet.SMSG_SET_REST_START)
		v.WriteUint32(0)
		s.SendPacket(v)
	}
}

func (s *Session) SendAccountDataTimes() {
	var dt account.DataTimes
	dt.Time = time.Now()
	dt.Mask = account.All
	dt.Activated = true
	s.Send(&dt)
}

func (s *Session) SendTutorialFlags() {
	var flags account.TutorialFlags
	flags.SetAll()
	s.Send(&flags)
	log.Println("Tutorial flags sent.")
}

func (s *Session) HandleAccountDataUpdate(upd8 *account.ClientUpdateData) {
	// log.Dump("data", data)
}

func (s *Session) GetPlayerRace() models.Race {
	return models.Race(s.Get("Race").Byte())
}

func (s *Session) GetLevel() int {
	return int(s.Get("Level").Uint32())
}

func (s *Session) IsTrackedGUID(g guid.GUID) bool {
	for _, v := range s.TrackedGUIDs {
		if v == g {
			return true
		}
	}
	return false
}

func (s *Session) ChangeDefaultSpeeds(modifier float32) {
	s.MoveSpeeds = make(update.Speeds)
	for speedType, speed := range DefaultSpeeds {
		s.MoveSpeeds[speedType] = speed * modifier
	}
}

func (s *Session) SyncTime() {
	if s.Build().AddedIn(vsn.V2_0_1) {
		log.Println("Synced time with client")
		s.Send(&synctime.Request{
			ServerTimeMs: s.Server.UptimeMS(),
		})
	}
}

func (s *Server) ClassPowerType(class models.Class) spell.PowerType {
	var cls *dbdefs.Ent_ChrClasses
	s.DB.Lookup(wdb.BucketKeyUint32ID, uint32(class), &cls)
	if cls == nil {
		panic(fmt.Errorf("class not found %d", class))
	}
	return spell.PowerType(cls.DisplayPower)
}

// SpawnPlayer initializes the player into the object manager and sends the packets needed to log the client into the world.
func (s *Session) SpawnPlayer() {
	log.Println("Locking guard player list")
	s.Server.GuardPlayerList.Lock()
	s.Server.PlayerList[s.PlayerName()] = s
	log.Println("releasing guard player list")
	s.Server.GuardPlayerList.Unlock()

	var exploredZones []models.ExploredZone
	s.DB().Where("player = ?", s.PlayerID()).Find(&exploredZones)

	// fill out attribute fields
	s.MovementInfo = &update.MovementInfo{
		Flags: 0,
		Time:  s.Server.UptimeMS(),
		Position: tempest.C4Vector{
			X: s.Char.X,
			Y: s.Char.Y,
			Z: s.Char.Z,
			W: s.Char.O,
		},
	}

	s.ChangeDefaultSpeeds(1.0)

	var err error
	s.ValuesBlock, err = update.NewValuesBlock(
		s.Build(),
		guid.TypeMaskObject|guid.TypeMaskUnit|guid.TypeMaskPlayer,
	)

	if err != nil {
		panic(err)
	}

	log.Println("Initializing inventory manager")
	s.InitInventoryManager()
	log.Println("setting values")

	s.SetGUID("GUID", s.GUID())

	s.SetByte("Power", uint8(s.Server.ClassPowerType(s.Char.Class)))

	s.SetFloat32("ScaleX", 1.0)

	s.SetUint32("Health", s.Char.Health)
	s.SetUint32("MaxHealth", 80)
	s.SetUint32("Mana", 4143)
	s.SetUint32("MaxMana", 4143)
	s.SetUint32("Energy", 100)
	s.SetUint32("MaxRage", 1000)
	s.SetUint32("MaxEnergy", 100)
	s.SetUint32("Level", uint32(s.Char.Level))
	s.SetUint32("FactionTemplate", 1)

	if s.Build() == vsn.Alpha {
		s.SetUint32("NumInvSlots", 0x89)
	}

	if s.Build().AddedIn(vsn.V2_4_3) {
		s.SetUint32("MaxLevel", uint32(s.GetMaxLevel()))
	}

	s.SetByte("Race", uint8(s.Char.Race))
	s.SetByte("Class", uint8(s.Char.Class))
	s.SetByte("Gender", uint8(s.Char.BodyType))

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SetByte("PlayerGender", uint8(s.Char.BodyType))
	}

	// Player flags
	s.SetBit("PlayerControlled", true)
	// s.SetBit("Resting", true)
	s.SetBit("AurasVisible", true)

	s.SetUint32("BaseAttackTime", 2900)
	s.SetUint32("OffhandAttackTime", 2000)

	s.SetFloat32("BoundingRadius", 1.0)
	s.SetFloat32("CombatReach", 1.0)

	s.SetUint32("DisplayID", s.Server.GetNative(s.Char.Race, s.Char.BodyType))
	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SetUint32("NativeDisplayID", s.Server.GetNative(models.Race(s.Char.Race), s.Char.BodyType))

		for _, ez := range exploredZones {
			var area *dbdefs.Ent_AreaTable
			s.DB().Lookup(wdb.BucketKeyUint32ID, ez.ZoneID, &area)
			if area != nil {
				s.SetExplorationFlag(uint32(area.AreaBit))
			}
		}
	}

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SetFloat32("MinDamage", 50)
		s.SetFloat32("MaxDamage", 50)

		s.SetUint32("MinOffhandDamage", 50)
		s.SetUint32("MaxOffhandDamage", 50)
	} else {
		s.SetFloat32("Damage", 50)
	}

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SetByte("LoyaltyLevel", 0xEE)
	}

	s.SetFloat32("ModCastSpeed", 30)

	s.SetUint32("BaseMana", 60)
	// todo: replace with bit fields
	// s.SetBit("AuraByteFlagSupportable, true)
	// s.SetBit("AuraByteFlagNoDispel, true)

	auraByteFlags := s.Get("AuraByteFlags")

	if auraByteFlags != nil {
		auraByteFlags.SetByte(0x08 | 0x20 | 0x10)
	}

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SetInt32("AttackPower", 20)
		s.SetInt32("AttackPowerMods", 0)

		s.SetInt32("RangedAttackPower", 1)
		s.SetInt32("RangedAttackPowerMods", 0)

		s.SetFloat32("MinRangedDamage", 0)
		s.SetFloat32("MaxRangedDamage", 0)
	}

	s.SetByte("Skin", s.Char.Skin)
	s.SetByte("Face", s.Char.Face)
	s.SetByte("HairStyle", s.Char.HairStyle)
	s.SetByte("HairColor", s.Char.HairColor)

	s.SetByte("FacialHair", s.Char.FacialHair)
	s.SetByte("BankBagSlotCount", 8)

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SetByte("RestState", 0x01)
	}

	stats := s.Get("Stats")

	for i := 0; i < 5; i++ {
		stats.Index(i).SetUint32(10)
	}

	s.SetByte("Gender", s.Char.BodyType)

	s.SetUint32("XP", s.Char.XP)
	s.SetUint32("NextLevelXP", s.GetNextLevelXP())

	s.Get("CharacterPoints").Index(0).SetUint32(51)
	s.Get("CharacterPoints").Index(1).SetUint32(2)

	s.SetFloat32("BlockPercentage", 4.0)
	s.SetFloat32("DodgePercentage", 4.0)
	s.SetFloat32("ParryPercentage", 4.0)

	if s.Build().AddedIn(vsn.V1_12_1) {
		s.SetFloat32("CritPercentage", 4.0)
		s.SetUint32("RestStateExperience", 200)
	}

	if s.Build().AddedIn(vsn.V2_4_3) {
		s.SetFloat32("RangedCritPercentage", 4.0)
	}

	s.SetInt32("Coinage", int32(s.Char.Coinage))

	if s.Build().AddedIn(vsn.V2_4_3) {
		info0 := s.Get("SkillInfos").Index(0)
		info1 := s.Get("SkillInfos").Index(1)

		info0.Field("ID").SetUint16(98)
		info0.Field("SkillLevel").SetUint16(300)
		info0.Field("SkillCap").SetUint16(300)
		info0.Field("Bonus").SetUint16(0)

		info1.Field("ID").SetUint16(109)
		info1.Field("SkillLevel").SetUint16(300)
		info1.Field("SkillCap").SetUint16(300)
		info1.Field("Bonus").SetUint16(0)
	}

	if s.Build().AddedIn(vsn.V3_3_5a) {
		s.SetBit("RegeneratePower", true)
	}

	watchedFactionIndex := s.ValuesBlock.Get("WatchedFactionIndex")

	if watchedFactionIndex != nil {
		watchedFactionIndex.SetInt32(-1)
	}

	// s.SetInt32("WatchedFactionIndex", -1)

	// var stupidity bool
	// stupidity = true // comment me out

	// if stupidity {
	// 	s.Get("Auras").Index(0).SetUint32(168)
	// 	s.Get("AuraLevels").Index(0).SetByte(1)
	// 	s.Get("AuraApplications").Index(0).SetByte(1)

	// 	update.SetAuraFlag(s.Get("AuraFlags"), 0, update.AuraMaskAll))

	// 	if err := s.Server.SetAuraFlag(s, 0, update.AuraMaskAll); err != nil {
	// 		panic(err)
	// 	}
	// }

	if err := s.InitBaseStats(); err != nil {
		log.Fatal(err)
	}

	if err := s.Server.ApplyStats(s); err != nil {
		log.Fatal(err)
	}

	if s.Char.FirstLogin {
		maxHP := s.Get("MaxHealth").Uint32()
		s.SetHealth(maxHP)
		s.SavePowers()
		// s.DB().Cols("health").Update(s.Char)
	}

	log.Println("values set")

	// s.ClearChanges()

	s.CurrentPhase = "main"
	s.CurrentMap = s.Server.Phase(s.CurrentPhase).Map(s.Char.Map)
	s.ZoneID = s.Char.Zone

	// send player create packet of themself
	s.SetState(InWorld)
	// log.Println("sending spawn packet")
	// s.SendObjectCreate(s)
	// log.Println("Sent spawn packet.")

	s.initMovementManager()

	cMap := s.Map()

	// add our player to map, and notify nearby players of their presence
	cMap.AddObject(s)

	// // notify our player of nearby objects.
	// nearbyObjects := cMap.VisibleObjects(s)
	// createObjects := make([]Object, len(nearbyObjects))

	// for i := range nearbyObjects {
	// 	createObjects[i] = nearbyObjects[i]
	// }

	// for _, wo := range nearbyObjects {
	// 	s.TrackedGUIDs = append(s.TrackedGUIDs, wo.GUID())
	// }
	// s.SendObjectCreate(createObjects...)
	s.initSpellManager()
}

func (ws *Server) RemovePlayerFromList(name string) {
	ws.GuardPlayerList.Lock()
	delete(ws.PlayerList, name)
	ws.GuardPlayerList.Unlock()
}

func (s *Session) CleanupPlayer() {
	s.Server.RemovePlayerFromList(s.PlayerName())
	s.BroadcastStatus(social.FriendOffline)

	if s.HasState(InWorld) {
		s.Map().RemoveObject(s.GUID())
	}

	if s.Group != nil {
		s.Group.UpdateList()
		s.Group = nil
	}

	s.savePlayedTime()

	// s.MovementInfo = nil
	// s.CurrentMap = 0
	// s.CurrentPhase = ""
	// s.Char = nil
	// s.TrackedGUIDs = nil
}
