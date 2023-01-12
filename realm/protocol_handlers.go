package realm

import (
	"fmt"
	"reflect"

	"github.com/Gophercraft/core/packet"
	p "github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/log"
)

type ProtocolHandlers map[p.WorldType]*ProtocolHandler

func (ws *Server) initHandlers() {
	h := ProtocolHandlers{}

	// Connection
	h.On(p.CMSG_AUTH_SESSION, Handshaking, (*Session).HandleAuthSession)
	h.On(p.CMSG_ENABLE_NAGLE, Handshaking, (*Session).HandleEnableNaglesAlgorithm)
	h.On(p.CMSG_ENTER_ENCRYPTED_MODE_ACK, Handshaking, (*Session).HandleEnterEncryptedModeAck)
	h.On(p.CMSG_PING, Handshaking, (*Session).HandlePing)
	h.On(p.CMSG_PET_NAME_QUERY, CharacterSelectMenu, (*Session).HandlePetNameQuery)
	h.On(p.CMSG_UPDATE_ACCOUNT_DATA, CharacterSelectMenu, (*Session).HandleAccountDataUpdate)
	h.On(p.CMSG_REALM_SPLIT, CharacterSelectMenu, (*Session).HandleRealmSplit)
	h.On(p.CMSG_UI_TIME_REQUEST, CharacterSelectMenu, (*Session).HandleUITimeRequest, OptionAsync)
	h.On(p.CMSG_QUERY_TIME, Handshaking, (*Session).HandleQueryTime, OptionAsync)

	// Character selection menu
	h.On(p.CMSG_CHAR_ENUM, CharacterSelectMenu, (*Session).HandleRequestCharacterList)
	h.On(p.CMSG_CHAR_DELETE, CharacterSelectMenu, (*Session).HandleDeleteCharacter)
	h.On(p.CMSG_CHAR_CREATE, CharacterSelectMenu, (*Session).HandleCreateCharacter)
	h.On(p.CMSG_PLAYER_LOGIN, CharacterSelectMenu, (*Session).HandlePlayerLogin, OptionAsync)
	h.On(p.CMSG_LOGOUT_REQUEST, InWorld, (*Session).HandleLogoutRequest)
	h.On(p.CMSG_LOGOUT_CANCEL, InWorld, (*Session).HandleLogoutCancel)

	// Social
	h.On(p.CMSG_NAME_QUERY, InWorld, (*Session).HandleNameQuery, OptionAsync)
	h.On(p.CMSG_MESSAGECHAT, InWorld, (*Session).HandleChat)
	h.On(p.CMSG_WHO, InWorld, (*Session).HandleWho)
	h.On(p.CMSG_SET_SELECTION, InWorld, (*Session).HandleTarget)
	// Vanilla: Request friend list and ignore list as individual packets.
	h.On(p.CMSG_FRIEND_LIST, InWorld, (*Session).HandleFriendListRequest)
	// 2.4.3+: All social data is merged into a single list.
	h.On(p.CMSG_CONTACT_LIST, InWorld, (*Session).HandleSocialListRequest)
	h.On(p.CMSG_ADD_FRIEND, InWorld, (*Session).HandleFriendAdd)
	h.On(p.CMSG_DEL_FRIEND, InWorld, (*Session).HandleFriendDelete)
	h.On(p.CMSG_ADD_IGNORE, InWorld, (*Session).HandleIgnoreAdd)
	h.On(p.CMSG_DEL_IGNORE, InWorld, (*Session).HandleIgnoreDelete)

	// Party/Group
	h.On(p.CMSG_GROUP_INVITE, InWorld, (*Session).HandleGroupInvite)
	h.On(p.CMSG_GROUP_ACCEPT, InWorld, (*Session).HandleGroupAccept)
	h.On(p.CMSG_GROUP_DECLINE, InWorld, (*Session).HandleGroupDecline)
	h.On(p.CMSG_GROUP_DISBAND, InWorld, (*Session).HandleGroupDisband)
	h.On(p.CMSG_REQUEST_PARTY_MEMBER_STATS, InWorld, (*Session).HandleRequestPartyMemberStats)
	h.On(p.CMSG_REQUEST_RAID_INFO, InWorld, (*Session).HandleRequestRaidInfo)

	// Movement
	h.On(p.MSG_MOVE_HEARTBEAT, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_FORWARD, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_BACKWARD, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_STRAFE_LEFT, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_STRAFE_RIGHT, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_STRAFE, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_JUMP, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_TURN_LEFT, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_TURN_RIGHT, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_TURN, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_PITCH_UP, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_PITCH_DOWN, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_SET_PITCH, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_PITCH, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_SET_RUN_MODE, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_SET_FACING, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_FALL_LAND, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_SWIM, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_SWIM, InWorld, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_WORLDPORT_ACK, InWorld, (*Session).HandleWorldportAck)
	h.On(p.CMSG_SUMMON_RESPONSE, InWorld, (*Session).HandleSummonResponse)

	h.On(p.CMSG_PLAYED_TIME, InWorld, (*Session).HandlePlayedTimeRequest)

	// Location
	h.On(p.CMSG_ZONEUPDATE, InWorld, (*Session).HandleZoneUpdate)
	h.On(p.CMSG_AREATRIGGER, InWorld, (*Session).HandleAreaTrigger)
	h.On(p.CMSG_WORLD_TELEPORT, InWorld, (*Session).HandleWorldTeleport)

	// Animation
	h.On(p.CMSG_STANDSTATECHANGE, InWorld, (*Session).HandleStandStateChange)
	h.On(p.CMSG_TEXT_EMOTE, InWorld, (*Session).HandleTextEmote)
	h.On(p.CMSG_SETSHEATHED, InWorld, (*Session).HandleSheathe)
	// Alpha quirk
	h.On(p.CMSG_SETWEAPONMODE, InWorld, (*Session).HandleSetWeaponMode)

	// Gameobjects
	h.On(p.CMSG_GAMEOBJECT_QUERY, InWorld, (*Session).HandleGameObjectQuery, OptionAsync)
	h.On(p.CMSG_GAMEOBJ_USE, InWorld, (*Session).HandleGameObjectUse)

	// Creatures
	h.On(p.CMSG_CREATURE_QUERY, InWorld, (*Session).HandleCreatureQuery, OptionAsync)
	h.On(p.CMSG_QUESTGIVER_STATUS_QUERY, InWorld, (*Session).HandleQuestgiverStatusQuery)
	h.On(p.CMSG_GOSSIP_HELLO, InWorld, (*Session).HandleGossipHello)
	h.On(p.CMSG_GOSSIP_SELECT_OPTION, InWorld, (*Session).HandleGossipSelectOption)
	h.On(p.CMSG_NPC_TEXT_QUERY, InWorld, (*Session).HandleGossipTextQuery, OptionAsync)

	// Inventory
	h.On(p.CMSG_ITEM_QUERY_SINGLE, CharacterSelectMenu, (*Session).HandleItemQuerySingle, OptionAsync)
	h.On(p.CMSG_SWAP_INV_ITEM, InWorld, (*Session).HandleSwapInventoryItem)
	h.On(p.CMSG_SWAP_ITEM, InWorld, (*Session).HandleSwapItem)
	h.On(p.CMSG_AUTOEQUIP_ITEM, InWorld, (*Session).HandleAutoEquipItem)
	h.On(p.CMSG_AUTOSTORE_BAG_ITEM, InWorld, (*Session).HandleAutoStoreBagItem)
	h.On(p.CMSG_DESTROYITEM, InWorld, (*Session).HandleDestroyItem)
	h.On(p.CMSG_SPLIT_ITEM, InWorld, (*Session).HandleSplitItem)

	// Spells/Combat
	h.On(p.CMSG_NEW_SPELL_SLOT, InWorld, (*Session).HandleNewSpellSlot)
	h.On(p.CMSG_SET_ACTION_BUTTON, InWorld, (*Session).HandleSetActionButton)
	h.On(p.CMSG_CAST_SPELL, InWorld, (*Session).HandleCast)
	h.On(p.CMSG_CANCEL_AURA, InWorld, (*Session).HandleAuraCancel)

	ws.ProtocolHandlers = h
}

const (
	OptionFlagAsync = 1 << iota
)

type Options struct {
	OptionFlags uint32
}

type Option func(*p.WorldPacket, *Options) error

func OptionAsync(wp *p.WorldPacket, opts *Options) error {
	opts.OptionFlags |= OptionFlagAsync
	return nil
}

func (h ProtocolHandlers) On(pt p.WorldType, requiredState SessionState, function interface{}, options ...Option) {
	h[pt] = &ProtocolHandler{pt, requiredState, reflect.ValueOf(function), options}
}

type ProtocolHandler struct {
	Op            p.WorldType
	RequiredState SessionState
	Fn            reflect.Value
	Options       []Option
}

var (
	formatType  = reflect.TypeOf((*packet.Decodable)(nil)).Elem()
	sessionType = reflect.TypeOf(&Session{})
	wPacketType = reflect.TypeOf(&packet.WorldPacket{})
	wTypeType   = reflect.TypeOf(packet.WorldType(0))
)

func (session *Session) dispatchHandler(ph *ProtocolHandler, wPacket *packet.WorldPacket) {
	fnType := ph.Fn.Type()

	if fnType.NumIn() < 1 || fnType.In(0) != sessionType {
		panic(fmt.Errorf("invalid protocol handler %s, first argument must be *Session", wPacket.Type))
	}

	args := []reflect.Value{reflect.ValueOf(session)}

	if fnType.NumIn() >= 2 {
		dataType := fnType.In(1)

		for i := 1; i < fnType.NumIn(); i++ {
			// some handlers have several possible trigger opcodes, they need to know the opcode while processing
			if dataType == wTypeType {
				args = append(args, reflect.ValueOf(wPacket.Type))
			} else if dataType == wPacketType {
				args = append(args, reflect.ValueOf(wPacket))
			} else if dataType.Implements(formatType) {
				if dataType.Kind() == reflect.Ptr {
					dataType = dataType.Elem()
				}
				packetFormat := reflect.New(dataType)
				pFormat := packetFormat.Interface().(packet.Decodable)
				err := pFormat.Decode(session.Build(), wPacket)
				if err != nil {
					log.Warn(err)
					session.End()
					return
				}
				args = append(args, packetFormat)
			} else {
				panic(fmt.Errorf("realm: %s cannot use this kind of handler argument %s", wPacket.Type, dataType))
			}
		}
	}

	ph.Fn.Call(args)
}
