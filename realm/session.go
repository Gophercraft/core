package realm

import (
	"fmt"
	"sync"
	"time"

	"github.com/Gophercraft/core/bnet"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet/area"
	"github.com/Gophercraft/core/packet/auth"
	"github.com/Gophercraft/core/packet/query"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"

	"github.com/Gophercraft/core/guid"
	"github.com/superp00t/etc"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/packet"
)

type SessionState int8

const Ended SessionState = -1

const (
	Handshaking SessionState = iota
	Authed
	Waiting
	CharacterSelectMenu
	InWorld
)

type PlayerSession struct {
	*update.ValuesBlock
	MovementManager

	Char           *models.Character
	CharacterProps []models.CharacterProp

	GuardInventory sync.Mutex
	Inventory      []models.Inventory
	Items          map[guid.GUID]*Item
	// In-world data
	JoinedWorldAt time.Time
	LevelJoinTime time.Time
	CurrentPhase  string
	CurrentMap    *Map
	CurrentArea   uint32
	// CurrentChunkIndex *terrain.TileChunkLookupIndex
	ZoneID       uint32
	AuraState    AuraState
	SpellManager *SpellManager

	// currently tracked objects
	Camera            Sight
	GuardTrackedGUIDs sync.Mutex
	TrackedGUIDs      []guid.GUID
	Flying            bool

	// Social
	Group           *Group
	GroupInvite     guid.GUID
	TeleportSummons *Summons

	gameMode models.GameMode
}

type Session struct {
	// Account data
	Server        *Server
	Connection    *packet.Connection
	AuthChallenge *auth.Challenge
	GuardSession  sync.Mutex
	ending        bool
	state         SessionState
	Tier          rpcnet.Tier
	Locale        i18n.Locale
	Account       uint64
	GameAccount   uint64
	SessionKey    []byte
	Processes     []*Process
	PacketPipe    chan *packet.WorldPacket
	// In-world Data: better kept as pointer
	// Otherwise we're allocating EVERYTHING when someone opens a socket
	*PlayerSession
}

func (s *Session) Log(msg ...any) {

	text := fmt.Sprintln(
		append([]any{s.String()}, msg...))
	log.DefaultLogger.LogLine(&log.Line{
		Time:     time.Now(),
		Text:     text,
		Category: "session",
	})
}

func (s *Session) ClientOrigin() string {
	return s.Connection.Conn.RemoteAddr().String()
}

// error!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// error!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// error!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// error!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// error!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (s *Session) State() SessionState {
	// s.GuardSession.Lock()
	state := s.state
	// s.GuardSession.Unlock()
	return state
}

func (s *Session) HasState(ss SessionState) bool {
	if s.State() == Ended {
		if ss == Ended {
			return true
		}
		return false
	}
	return s.State() >= ss
}

func (s *Session) SetState(ss SessionState) {
	s.GuardSession.Lock()
	s.state = ss

	for i := len(s.Processes) - 1; i >= 0; i-- {
		proc := s.Processes[i]
		if proc.CreationState > s.state {
			s.killProcess(i)
		}
	}

	s.GuardSession.Unlock()
}

func (session *Session) SendAuthSuccess() {
	exp := uint8(session.Build().Exp())
	var as auth.Response
	as.Result = bnet.ERROR_OK
	as.SuccessInfo = &auth.SuccessInfo{
		ActiveExpansionLevel:  exp,
		AccountExpansionLevel: exp,
		// AvailableClasses:      session.Server.AvailableClasses(),
	}
	session.Send(&as)
}

func (s *Session) TypeID() guid.TypeID {
	// activeplayer
	return guid.TypePlayer
}

func (s *Session) GUID() guid.GUID {
	if s.Char == nil {
		return guid.Nil
	}
	return guid.RealmSpecific(guid.Player, s.Server.RealmID(), s.Char.ID)
}

func (s *PlayerSession) Values() *update.ValuesBlock {
	return s.ValuesBlock
}

func (s *Session) HandlePing(pi *packet.Ping) {
	log.Println("Ping: ", pi.Ping, "Latency", pi.Latency)
	s.Send(&packet.Pong{
		Ping: pi.Ping,
	})
}

func (s *Session) DB() *wdb.Core {
	return s.Server.DB
}

func (s *Session) InitPlayerSession(char *models.Character) {
	if s.PlayerSession == nil {
		s.PlayerSession = &PlayerSession{
			Char: char,
		}
	} else {
		s.PlayerSession.Char = char
	}
}

func (s *Session) HandlePetNameQuery(p *query.PetName) {
	var resp query.PetNameResponse

	s.Send(&resp)
	log.Println("Sent pet response")
}

func (s *Session) Alertf(format string, args ...interface{}) {
	s.SendAlertText(fmt.Sprintf(format, args...))
}

func (s *Session) SendAlertText(data string) {
	s.Send(&area.TriggerMessage{
		Message: data,
	})
}

func (s *Session) Map() *Map {
	return s.CurrentMap
}

func (s *Session) MapID() uint32 {
	return s.CurrentMap.ID
}

func (s *Session) Phase() *Phase {
	return s.Server.Phase(s.CurrentPhase)
}

func (s *Session) GetPlayerClass() models.Class {
	return models.Class(s.ValuesBlock.Get("Class").Byte())
}

func (s *Session) Config() *config.World {
	return s.Server.Config
}

func (s *Session) HandleRealmSplit(split *etc.Buffer) {
	splitReq := split.ReadInt32() // realm ID perhaps?
	log.Println("User requested realm split", splitReq)

	response := packet.NewWorldPacket(packet.SMSG_REALM_SPLIT)
	response.WriteInt32(splitReq)
	response.WriteInt32(0) // split state
	response.WriteCString("01/01/01")

	s.SendPacket(response)
}

func (s *Session) HandleUITimeRequest() {
	resp := packet.NewWorldPacket(packet.SMSG_UI_TIME)
	resp.WriteUint32(uint32(time.Now().Unix()))
	s.SendPacket(resp)
}

func (s *Session) DebugGUID(dbg guid.GUID) string {
	switch dbg.HighType() {
	case guid.Player:
		plyr, err := s.Server.GetSessionByGUID(dbg)
		if err != nil {
			return dbg.String()
		}
		return dbg.String() + " (" + plyr.PlayerName() + ")"
	case guid.Item:
		it, ok := s.Items[dbg]
		if !ok {
			return dbg.String()
		}

		var tpl *models.ItemTemplate
		s.DB().Lookup(wdb.BucketKeyStringID, it.ItemID, &tpl)
		if tpl == nil {
			return dbg.String()
		}

		return fmt.Sprintf("%s (%s)", dbg, tpl.Name.GetLocalized(s.Locale))
	default:
		return dbg.String()
	}
}

func (s *Session) TimeInWorld() (time.Duration, error) {
	if !s.HasState(InWorld) {
		return 0, fmt.Errorf("realm: TimeInWorld(): not currently in world")
	}

	return time.Since(s.JoinedWorldAt), nil
}
