package main

import (
	"fmt"

	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/realm/wdb/models"
)

type CreatureTemplate struct {
	Entry                uint32  `xorm:"'Entry'"`
	Name                 string  `xorm:"'Name'"`
	SubName              string  `xorm:"'SubName'"`
	MinLevel             uint32  `xorm:"'MinLevel'"`
	MaxLevel             uint32  `xorm:"'MaxLevel'"`
	ModelId1             uint32  `xorm:"'ModelId1'"`
	ModelId2             uint32  `xorm:"'ModelId2'"`
	ModelId3             uint32  `xorm:"'ModelId3'"`
	ModelId4             uint32  `xorm:"'ModelId4'"`
	Faction              uint32  `xorm:"'Faction'"`
	Scale                float32 `xorm:"'Scale'"`
	Family               int32   `xorm:"'Family'"`
	CreatureType         uint32  `xorm:"'CreatureType'"`
	InhabitType          uint32  `xorm:"'InhabitType'"`
	RegenerateStats      uint32  `xorm:"'RegenerateStats'"`
	RacialLeader         uint32  `xorm:"'RacialLeader'"`
	NpcFlags             uint32  `xorm:"'NpcFlags'"`
	UnitFlags            uint32  `xorm:"'UnitFlags'"`
	DynamicFlags         uint32  `xorm:"'DynamicFlags'"`
	ExtraFlags           uint32  `xorm:"'ExtraFlags'"`
	CreatureTypeFlags    uint32  `xorm:"'CreatureTypeFlags'"`
	SpeedWalk            float32 `xorm:"'SpeedWalk'"`
	SpeedRun             float32 `xorm:"'SpeedRun'"`
	UnitClass            uint32  `xorm:"'UnitClass'"`
	Rank                 uint32  `xorm:"'Rank'"`
	HealthMultiplier     float32 `xorm:"'HealthMultiplier'"`
	PowerMultiplier      float32 `xorm:"'PowerMultiplier'"`
	DamageMultiplier     float32 `xorm:"'DamageMultiplier'"`
	DamageVariance       float32 `xorm:"'DamageVariance'"`
	ArmorMultiplier      float32 `xorm:"'ArmorMultiplier'"`
	ExperienceMultiplier float32 `xorm:"'ExperienceMultiplier'"`
	MinLevelHealth       uint32  `xorm:"'MinLevelHealth'"`
	MaxLevelHealth       uint32  `xorm:"'MaxLevelHealth'"`
	MinLevelMana         uint32  `xorm:"'MinLevelMana'"`
	MaxLevelMana         uint32  `xorm:"'MaxLevelMana'"`
	MinMeleeDmg          float32 `xorm:"'MinMeleeDmg'"`
	MaxMeleeDmg          float32 `xorm:"'MaxMeleeDmg'"`
	MinRangedDmg         float32 `xorm:"'MinRangedDmg'"`
	MaxRangedDmg         float32 `xorm:"'MaxRangedDmg'"`
	Armor                uint32  `xorm:"'Armor'"`
	MeleeAttackPower     uint32  `xorm:"'MeleeAttackPower'"`
	RangedAttackPower    uint32  `xorm:"'RangedAttackPower'"`
	MeleeBaseAttackTime  uint32  `xorm:"'MeleeBaseAttackTime'"`
	RangedBaseAttackTime uint32  `xorm:"'RangedBaseAttackTime'"`
	DamageSchool         int32   `xorm:"'DamageSchool'"`
	MinLootGold          uint32  `xorm:"'MinLootGold'"`
	MaxLootGold          uint32  `xorm:"'MaxLootGold'"`
	LootId               uint32  `xorm:"'LootId'"`
	PickpocketLootId     uint32  `xorm:"'PickpocketLootId'"`
	SkinningLootId       uint32  `xorm:"'SkinningLootId'"`
	KillCredit1          uint32  `xorm:"'KillCredit1'"`
	KillCredit2          uint32  `xorm:"'KillCredit2'"`
	MechanicImmuneMask   uint32  `xorm:"'MechanicImmuneMask'"`
	SchoolImmuneMask     uint32  `xorm:"'SchoolImmuneMask'"`
	ResistanceHoly       int32   `xorm:"'ResistanceHoly'"`
	ResistanceFire       int32   `xorm:"'ResistanceFire'"`
	ResistanceNature     int32   `xorm:"'ResistanceNature'"`
	ResistanceFrost      int32   `xorm:"'ResistanceFrost'"`
	ResistanceShadow     int32   `xorm:"'ResistanceShadow'"`
	ResistanceArcane     int32   `xorm:"'ResistanceArcane'"`
	PetSpellDataId       uint32  `xorm:"'PetSpellDataId'"`
	MovementType         uint32  `xorm:"'MovementType'"`
	TrainerType          int32   `xorm:"'TrainerType'"`
	TrainerSpell         uint32  `xorm:"'TrainerSpell'"`
	TrainerClass         uint32  `xorm:"'TrainerClass'"`
	TrainerRace          uint32  `xorm:"'TrainerRace'"`
	TrainerTemplateId    uint32  `xorm:"'TrainerTemplateId'"`
	VendorTemplateId     uint32  `xorm:"'VendorTemplateId'"`
	GossipMenuId         uint32  `xorm:"'GossipMenuId'"`
	EquipmentTemplateId  uint32  `xorm:"'EquipmentTemplateId'"`
	Civilian             uint32  `xorm:"'Civilian'"`
	AIName               string  `xorm:"'AIName'"`
	ScriptName           string  `xorm:"'ScriptName'"`
}

func extractCreatures() {
	var ctt []CreatureTemplate
	err := DB.Find(&ctt)
	if err != nil {
		panic(err)
	}

	cfl := openFile("DB/CreatureTemplate.txt")
	wr := openTextWriter(cfl)
	for _, cr := range ctt {
		ct := models.CreatureTemplate{
			ID:                  fmt.Sprintf("cr:%d", cr.Entry),
			Entry:               cr.Entry,
			Name:                i18n.GetEnglish(cr.Name),
			SubName:             i18n.GetEnglish(cr.SubName),
			MinLevel:            cr.MinLevel,
			MaxLevel:            cr.MaxLevel,
			Faction:             cr.Faction,
			Scale:               cr.Scale,
			CreatureType:        cr.CreatureType,
			InhabitType:         cr.InhabitType,
			RegenerateStats:     cr.RegenerateStats,
			RacialLeader:        cr.RacialLeader == 1,
			Gossip:              cr.NpcFlags&1 != 0,
			QuestGiver:          cr.NpcFlags&2 != 0,
			Vendor:              cr.NpcFlags&4 != 0,
			FlightMaster:        cr.NpcFlags&8 != 0,
			Trainer:             cr.NpcFlags&16 != 0,
			SpiritHealer:        cr.NpcFlags&32 != 0,
			SpiritGuide:         cr.NpcFlags&64 != 0,
			Innkeeper:           cr.NpcFlags&128 != 0,
			Banker:              cr.NpcFlags&256 != 0,
			Petitioner:          cr.NpcFlags&512 != 0,
			TabardDesigner:      cr.NpcFlags&1024 != 0,
			BattleMaster:        cr.NpcFlags&2048 != 0,
			Auctioneer:          cr.NpcFlags&4096 != 0,
			StableMaster:        cr.NpcFlags&8192 != 0,
			Repairer:            cr.NpcFlags&16384 != 0,
			OutdoorPVP:          cr.NpcFlags&536870912 != 0,
			ServerControlled:    cr.UnitFlags&0x1 != 0,
			NonAttackable:       cr.UnitFlags&0x2 != 0,
			RemoveClientControl: cr.UnitFlags&0x4 != 0,
			PlayerControlled:    cr.UnitFlags&0x8 != 0,
			Rename:              cr.UnitFlags&0x10 != 0,
			PetAbandon:          cr.UnitFlags&0x20 != 0,
			OOCNotAttackable:    cr.UnitFlags&0x100 != 0,
			Passive:             cr.UnitFlags&0x200 != 0,
			PVP:                 cr.UnitFlags&0x1000 != 0,
			IsSilenced:          cr.UnitFlags&0x2000 != 0,
			IsPersuaded:         cr.UnitFlags&0x4000 != 0,
			Swimming:            cr.UnitFlags&0x8000 != 0,
			RemoveAttackIcon:    cr.UnitFlags&0x10000 != 0,
			IsPacified:          cr.UnitFlags&0x20000 != 0,
			IsStunned:           cr.UnitFlags&0x40000 != 0,
			InCombat:            cr.UnitFlags&0x80000 != 0,
			InTaxiFlight:        cr.UnitFlags&0x100000 != 0,
			Disarmed:            cr.UnitFlags&0x200000 != 0,
			Confused:            cr.UnitFlags&0x400000 != 0,
			Fleeing:             cr.UnitFlags&0x800000 != 0,
			Possessed:           cr.UnitFlags&0x1000000 != 0,
			NotSelectable:       cr.UnitFlags&0x2000000 != 0,
			Skinnable:           cr.UnitFlags&0x4000000 != 0,
			AurasVisible:        cr.UnitFlags&0x8000000 != 0,
			Sheathe:             cr.UnitFlags&0x40000000 != 0,
			NoKillReward:        cr.UnitFlags&0x80000000 != 0,

			Lootable:              cr.DynamicFlags&1 != 0,
			TrackUnit:             cr.DynamicFlags&2 != 0,
			Tapped:                cr.DynamicFlags&4 != 0,
			TappedByPlayer:        cr.DynamicFlags&8 != 0,
			SpecialInfo:           cr.DynamicFlags&16 != 0,
			VisuallyDead:          cr.DynamicFlags&32 != 0,
			TappedByAllThreatList: cr.DynamicFlags&128 != 0,

			InstanceBind:             cr.ExtraFlags&0x1 != 0,      // creature kill bind instance with killer and killer’s group
			NoAggroOnSight:           cr.ExtraFlags&0x2 != 0,      // no aggro (ignore faction/reputation hostility)
			NoParry:                  cr.ExtraFlags&0x4 != 0,      // creature can’t parry
			NoParryHasten:            cr.ExtraFlags&0x8 != 0,      // creature can’t counter-attack at parry
			NoBlock:                  cr.ExtraFlags&0x10 != 0,     //	creature can’t block
			NoCrush:                  cr.ExtraFlags&0x20 != 0,     // creature can’t do crush attacks
			NoXPAtKill:               cr.ExtraFlags&0x40 != 0,     // creature kill not provide XP
			Invisible:                cr.ExtraFlags&0x80 != 0,     // creature is always invisible for player (mostly trigger creatures)
			NotTauntable:             cr.ExtraFlags&0x100 != 0,    // creature is immune to taunt auras and effect attack me
			AggroZone:                cr.ExtraFlags&0x200 != 0,    // creature sets itself in combat with zone on aggro
			Guard:                    cr.ExtraFlags&0x400 != 0,    // creature is a guard
			NoCallAssist:             cr.ExtraFlags&0x800 != 0,    // creature shouldn’t call for assistance on aggro
			Active:                   cr.ExtraFlags&0x1000 != 0,   //creature is active object. Grid of this creature will be loaded and creature set as active
			ForceEnableMMap:          cr.ExtraFlags&0x2000 != 0,   // creature is forced to use MMaps
			ForceDisableMMap:         cr.ExtraFlags&0x4000 != 0,   // creature is forced to NOT use MMaps
			WalkInWater:              cr.ExtraFlags&0x8000 != 0,   // creature is forced to walk in water even it can swim
			Civilian:                 cr.ExtraFlags&0x10000 != 0,  // CreatureInfo→civilian substitute (for expansions as Civilian Colum was removed)
			NoMelee:                  cr.ExtraFlags&0x20000 != 0,  // creature can’t melee
			FarView:                  cr.ExtraFlags&0x40000 != 0,  // creature with far view
			ForceAttackingCapability: cr.ExtraFlags&0x80000 != 0,  // SetForceAttackingCapability(true); for nonattackable, nontargetable creatures that should be able to attack nontheless
			IgnoreUsedPosition:       cr.ExtraFlags&0x100000 != 0, // ignore creature when checking used positions around target
			CountSpawns:              cr.ExtraFlags&0x200000 != 0, // count creature spawns in Map*
			HasteSpellImmunity:       cr.ExtraFlags&0x400000 != 0, // immunity to COT or Mind Numbing Poison – very common in instances

			Tameable:                cr.CreatureTypeFlags&1 != 0, // Makes the mob tameable (must also be a beast and have family set)
			VisibleToGhosts:         cr.CreatureTypeFlags&2 != 0, // Sets Creatures that can ALSO be seen when player is a ghost. Used in CanInteract function by client, can’t be attacked
			BossLevel:               cr.CreatureTypeFlags&4 != 0,
			DontPlayWoundParryAnim:  cr.CreatureTypeFlags&8 != 0,
			HideFactionTooltip:      cr.CreatureTypeFlags&16 != 0, // Controls something in client tooltip related to creature faction
			SpellAttackable:         cr.CreatureTypeFlags&64 != 0,
			DeadInteract:            cr.CreatureTypeFlags&128 != 0,
			HerbLoot:                cr.CreatureTypeFlags&256 != 0, // Uses Skinning Loot Field
			MiningLoot:              cr.CreatureTypeFlags&512 != 0, // Makes Mob Corpse Mineable – Uses Skinning Loot Field
			DontLogDeath:            cr.CreatureTypeFlags&1024 != 0,
			MountedCombat:           cr.CreatureTypeFlags&2048 != 0,
			CanAssist:               cr.CreatureTypeFlags&4096 != 0, //	Can aid any player or group in combat. Typically seen for escorting NPC’s
			PetHasActionBar:         cr.CreatureTypeFlags&8192 != 0, // 	checked from calls in Lua_PetHasActionBar
			MaskUID:                 cr.CreatureTypeFlags&16384 != 0,
			EngineerLoot:            cr.CreatureTypeFlags&32768 != 0, //	Makes Mob Corpse Engineer Lootable – Uses Skinning Loot Field
			ExoticPet:               cr.CreatureTypeFlags&65536 != 0, // Tamable as an exotic pet. Normal tamable flag must also be set.
			UseDefaultCollisionBox:  cr.CreatureTypeFlags&131072 != 0,
			IsSiegeWeapon:           cr.CreatureTypeFlags&262144 != 0,
			ProjectileCollision:     cr.CreatureTypeFlags&524288 != 0,
			HideNameplate:           cr.CreatureTypeFlags&1048576 != 0,
			DontPlayMountedAnim:     cr.CreatureTypeFlags&2097152 != 0,
			IsLinkAll:               cr.CreatureTypeFlags&4194304 != 0,
			InteractOnlyWithCreator: cr.CreatureTypeFlags&8388608 != 0,
			ForceGossip:             cr.CreatureTypeFlags&134217728 != 0,

			SpeedWalk:            cr.SpeedWalk,
			SpeedRun:             cr.SpeedRun,
			UnitClass:            cr.UnitClass,
			Rank:                 cr.Rank,
			HealthMultiplier:     cr.HealthMultiplier,
			PowerMultiplier:      cr.PowerMultiplier,
			DamageMultiplier:     cr.DamageMultiplier,
			DamageVariance:       cr.DamageVariance,
			ArmorMultiplier:      cr.ArmorMultiplier,
			ExperienceMultiplier: cr.ExperienceMultiplier,
			MinLevelHealth:       cr.MinLevelHealth,
			MaxLevelHealth:       cr.MaxLevelHealth,
			MinLevelMana:         cr.MinLevelMana,
			MaxLevelMana:         cr.MaxLevelMana,
			MinMeleeDmg:          cr.MinMeleeDmg,
			MaxMeleeDmg:          cr.MaxMeleeDmg,
			MinRangedDmg:         cr.MinRangedDmg,
			MaxRangedDmg:         cr.MaxRangedDmg,
			Armor:                cr.Armor,
			MeleeAttackPower:     cr.MeleeAttackPower,
			RangedAttackPower:    cr.RangedAttackPower,
			MeleeBaseAttackTime:  cr.MeleeBaseAttackTime,
			RangedBaseAttackTime: cr.RangedBaseAttackTime,
			DamageSchool:         cr.DamageSchool,
			MinLootGold:          models.Money(cr.MinLootGold),
			MaxLootGold:          models.Money(cr.MaxLootGold),
			LootId:               cr.LootId,
			PickpocketLootId:     cr.PickpocketLootId,
			SkinningLootId:       cr.SkinningLootId,
			KillCredit1:          cr.KillCredit1,
			KillCredit2:          cr.KillCredit2,
			MechanicImmuneMask:   cr.MechanicImmuneMask,
			SchoolImmuneMask:     cr.SchoolImmuneMask,
			ResistanceHoly:       cr.ResistanceHoly,
			ResistanceFire:       cr.ResistanceFire,
			ResistanceNature:     cr.ResistanceNature,
			ResistanceFrost:      cr.ResistanceFrost,
			ResistanceShadow:     cr.ResistanceShadow,
			ResistanceArcane:     cr.ResistanceArcane,
			PetSpellDataId:       cr.PetSpellDataId,
			MovementType:         cr.MovementType,
			TrainerType:          cr.TrainerType,
			TrainerSpell:         cr.TrainerSpell,
			TrainerClass:         cr.TrainerClass,
			TrainerRace:          cr.TrainerRace,
			TrainerTemplateId:    cr.TrainerTemplateId,
			VendorTemplateId:     cr.VendorTemplateId,
			GossipMenuId:         fmt.Sprintf("%d", cr.GossipMenuId),
			EquipmentTemplateId:  cr.EquipmentTemplateId,
			DishonourableKill:    cr.Civilian == 1,
			AIName:               cr.AIName,
			ScriptName:           cr.ScriptName,
		}

		if cr.ModelId1 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId1)
		}

		if cr.ModelId2 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId2)
		}

		if cr.ModelId3 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId3)
		}

		if cr.ModelId4 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId4)
		}

		if cr.Family != 0 {
			switch cr.Family {
			case 1:
				ct.Family = "Wolf"
			case 2:
				ct.Family = "Cat"
			case 3:
				ct.Family = "Spider"
			case 4:
				ct.Family = "Bear"
			case 5:
				ct.Family = "Boar"
			case 6:
				ct.Family = "Crocolisk"
			case 7:
				ct.Family = "Carrion Bird"
			case 8:
				ct.Family = "Crab"
			case 9:
				ct.Family = "Gorilla"
			// case 10:
			// 	ct.Family = "Horse (Custom)"
			case 11:
				ct.Family = "Raptor"
			case 12:
				ct.Family = "Tallstrider"
			case 15:
				ct.Family = "Felhunter"
			case 16:
				ct.Family = "Voidwalker"
			case 17:
				ct.Family = "Succubus"
			case 19:
				ct.Family = "Doomguard"
			case 20:
				ct.Family = "Scorpid"
			case 21:
				ct.Family = "Turtle"
			case 23:
				ct.Family = "Imp"
			case 24:
				ct.Family = "Bat"
			case 25:
				ct.Family = "Hyena"
			case 26:
				ct.Family = "Owl"
			case 27:
				ct.Family = "Wind Serpent"
			case 28:
				ct.Family = "Remote Control"
			default:
				panic(fmt.Errorf("unknown id: %d", cr.Family))
			}
		}

		if err := wr.Encode(ct); err != nil {
			panic(err)
		}
	}

	cfl.Close()
}
