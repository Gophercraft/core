package chat

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type MsgType uint8

const (
	MsgAddon = iota
	MsgSay
	MsgParty
	MsgRaid
	MsgGuild
	MsgOfficer
	MsgYell
	MsgWhisper
	MsgWhisperForeign
	MsgWhisperInform
	MsgEmote
	MsgTextEmote
	MsgSystem
	MsgCreatureSay
	MsgCreatureYell
	MsgCreatureWhisper
	MsgCreatureEmote
	MsgChannel
	MsgChannelJoin
	MsgChannelLeave
	MsgChannelList
	MsgChannelNotice
	MsgChannelNoticeUser
	MsgAFK
	MsgDND
	MsgIgnored
	MsgSkill
	MsgLoot
	MsgMoney
	MsgOpening
	MsgBGSystemNeutral
	MsgBGSystemBlueTeam
	MsgBGSystemRedTeam
	MsgRaidLeader
	MsgRaidWarning
	MsgRaidBossWhisper
	MsgRaidBossEmote
	MsgBattleground
	MsgBattlegroundLeader
	MsgCreatureParty
	MsgTradeSkills
	MsgPetInfo
	MsgCombatMiscInfo
	MsgCombatXPGain
	MsgCombatHonorGain
	MsgCombatFactionGain
	MsgFiltered
	MsgRestricted
	MsgBNet
	MsgAchievement
	MsgGuildAchievement
	MsgArenaPoints
	MsgPartyLeader
	MsgBNetWhisper
	MsgBNetWhisperInform
	MsgBNetConversation
	MsgBNetConversationNotice
	MsgBNetConversationList
	MsgBNetInlineToastAlert
	MsgBNetInlineToastBroadcast
	MsgBNetInlineToastBroadcastInform
	MsgBNetInlineToastConversation
	MsgBNetWhisperPlayerOffline
	MsgCombatGuildXPGain
	MsgCurrency
)

type MsgTypeDescriptor map[MsgType]uint32

var MsgTypeDescriptors = map[vsn.BuildRange]MsgTypeDescriptor{
	vsn.Range(0, 3368): {
		MsgSay:               0x00,
		MsgParty:             0x01,
		MsgGuild:             0x02,
		MsgOfficer:           0x03,
		MsgYell:              0x04,
		MsgWhisper:           0x05,
		MsgWhisperInform:     0x06,
		MsgEmote:             0x07,
		MsgTextEmote:         0x08,
		MsgSystem:            0x09,
		MsgCreatureSay:       0x0A,
		MsgCreatureYell:      0x0B,
		MsgCreatureEmote:     0x0C,
		MsgChannel:           0x0D,
		MsgChannelJoin:       0x0E,
		MsgChannelLeave:      0xF,
		MsgChannelList:       0x10,
		MsgChannelNotice:     0x11,
		MsgChannelNoticeUser: 0x12,
		MsgAFK:               0x13,
		MsgDND:               0x14,
		MsgIgnored:           0x16,
		MsgSkill:             0x17,
		MsgLoot:              0x18,
	},

	vsn.Range(5875, 6005): {
		MsgAddon:              0xFFFFFFFF,
		MsgSay:                0x00,
		MsgParty:              0x01,
		MsgRaid:               0x02,
		MsgGuild:              0x03,
		MsgOfficer:            0x04,
		MsgYell:               0x05,
		MsgWhisper:            0x06,
		MsgWhisperInform:      0x07,
		MsgEmote:              0x08,
		MsgTextEmote:          0x09,
		MsgSystem:             0x0A,
		MsgCreatureSay:        0x0B,
		MsgCreatureYell:       0x0C,
		MsgCreatureEmote:      0x0D,
		MsgChannel:            0x0E,
		MsgChannelJoin:        0x0F,
		MsgChannelLeave:       0x10,
		MsgChannelList:        0x11,
		MsgChannelNotice:      0x12,
		MsgChannelNoticeUser:  0x13,
		MsgAFK:                0x14,
		MsgDND:                0x15,
		MsgIgnored:            0x16,
		MsgSkill:              0x17,
		MsgLoot:               0x18,
		MsgCreatureWhisper:    0x1A,
		MsgBGSystemNeutral:    0x52,
		MsgBGSystemBlueTeam:   0x53,
		MsgBGSystemRedTeam:    0x54,
		MsgRaidLeader:         0x57,
		MsgRaidWarning:        0x58,
		MsgRaidBossWhisper:    0x59,
		MsgRaidBossEmote:      0x5A,
		MsgBattleground:       0x5C,
		MsgBattlegroundLeader: 0x5D,
	},

	vsn.Range(6180, 8606): {
		MsgAddon:              0xFFFFFFFF,
		MsgSystem:             0x00,
		MsgSay:                0x01,
		MsgParty:              0x02,
		MsgRaid:               0x03,
		MsgGuild:              0x04,
		MsgOfficer:            0x05,
		MsgYell:               0x06,
		MsgWhisper:            0x07,
		MsgWhisperForeign:     0x08,
		MsgWhisperInform:      0x09,
		MsgEmote:              0x0A,
		MsgTextEmote:          0x0B,
		MsgCreatureSay:        0x0C,
		MsgCreatureParty:      0x0D,
		MsgCreatureYell:       0x0E,
		MsgCreatureWhisper:    0x0F,
		MsgCreatureEmote:      0x10,
		MsgChannel:            0x11,
		MsgChannelJoin:        0x12,
		MsgChannelLeave:       0x13,
		MsgChannelList:        0x14,
		MsgChannelNotice:      0x15,
		MsgChannelNoticeUser:  0x16,
		MsgAFK:                0x17,
		MsgDND:                0x18,
		MsgIgnored:            0x19,
		MsgSkill:              0x1A,
		MsgLoot:               0x1B,
		MsgMoney:              0x1C,
		MsgOpening:            0x1D,
		MsgTradeSkills:        0x1E,
		MsgPetInfo:            0x1F,
		MsgCombatMiscInfo:     0x20,
		MsgCombatXPGain:       0x21,
		MsgCombatHonorGain:    0x22,
		MsgCombatFactionGain:  0x23,
		MsgBGSystemNeutral:    0x24,
		MsgBGSystemBlueTeam:   0x25,
		MsgBGSystemRedTeam:    0x26,
		MsgRaidLeader:         0x27,
		MsgRaidWarning:        0x28,
		MsgRaidBossEmote:      0x29,
		MsgRaidBossWhisper:    0x2A,
		MsgFiltered:           0x2B,
		MsgBattleground:       0x2C,
		MsgBattlegroundLeader: 0x2D,
		MsgRestricted:         0x2E,
	},

	vsn.Range(9056, 12340): {
		MsgAddon:                          0xFFFFFFFF,
		MsgSystem:                         0x00,
		MsgSay:                            0x01,
		MsgParty:                          0x02,
		MsgRaid:                           0x03,
		MsgGuild:                          0x04,
		MsgOfficer:                        0x05,
		MsgYell:                           0x06,
		MsgWhisper:                        0x07,
		MsgWhisperForeign:                 0x08,
		MsgWhisperInform:                  0x09,
		MsgEmote:                          0x0A,
		MsgTextEmote:                      0x0B,
		MsgCreatureSay:                    0x0C,
		MsgCreatureParty:                  0x0D,
		MsgCreatureYell:                   0x0E,
		MsgCreatureWhisper:                0x0F,
		MsgCreatureEmote:                  0x10,
		MsgChannel:                        0x11,
		MsgChannelJoin:                    0x12,
		MsgChannelLeave:                   0x13,
		MsgChannelList:                    0x14,
		MsgChannelNotice:                  0x15,
		MsgChannelNoticeUser:              0x16,
		MsgAFK:                            0x17,
		MsgDND:                            0x18,
		MsgIgnored:                        0x19,
		MsgSkill:                          0x1A,
		MsgLoot:                           0x1B,
		MsgMoney:                          0x1C,
		MsgOpening:                        0x1D,
		MsgTradeSkills:                    0x1E,
		MsgPetInfo:                        0x1F,
		MsgCombatMiscInfo:                 0x20,
		MsgCombatXPGain:                   0x21,
		MsgCombatHonorGain:                0x22,
		MsgCombatFactionGain:              0x23,
		MsgBGSystemNeutral:                0x24,
		MsgBGSystemBlueTeam:               0x25,
		MsgBGSystemRedTeam:                0x26,
		MsgRaidLeader:                     0x27,
		MsgRaidWarning:                    0x28,
		MsgRaidBossEmote:                  0x29,
		MsgRaidBossWhisper:                0x2A,
		MsgFiltered:                       0x2B,
		MsgBattleground:                   0x2C,
		MsgBattlegroundLeader:             0x2D,
		MsgRestricted:                     0x2E,
		MsgBNet:                           0x2F,
		MsgAchievement:                    0x30,
		MsgGuildAchievement:               0x31,
		MsgArenaPoints:                    0x32,
		MsgPartyLeader:                    0x33,
		MsgBNetWhisper:                    0x35,
		MsgBNetWhisperInform:              0x36,
		MsgBNetConversation:               0x37,
		MsgBNetConversationNotice:         0x38,
		MsgBNetConversationList:           0x39,
		MsgBNetInlineToastAlert:           0x3A,
		MsgBNetInlineToastBroadcast:       0x3B,
		MsgBNetInlineToastBroadcastInform: 0x3C,
		MsgBNetInlineToastConversation:    0x3D,
		MsgBNetWhisperPlayerOffline:       0x3E,
		MsgCombatGuildXPGain:              0x3F,
		MsgCurrency:                       0x40,
	},
}

func ConvertMsgType(build vsn.Build, value MsgType) (uint32, error) {
	var desc MsgTypeDescriptor
	if err := vsn.QueryDescriptors(build, MsgTypeDescriptors, &desc); err != nil {
		return 0, err
	}

	code, ok := desc[value]
	if !ok {
		return 0, fmt.Errorf("chat: no MsgType value found for %v", value)
	}

	return code, nil
}

func ResolveMsgType(build vsn.Build, code uint32) (MsgType, error) {
	var desc MsgTypeDescriptor
	if err := vsn.QueryDescriptors(build, MsgTypeDescriptors, &desc); err != nil {
		return 0, err
	}

	for k, v := range desc {
		if v == code {
			return k, nil
		}
	}

	return 0, fmt.Errorf("chat: no MsgType found for 0x%04X in descriptor %s", code, build)
}

type Error struct {
	Type packet.WorldType
	Name string
}

func (e *Error) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = e.Type
	out.WriteCString(e.Name)
	return nil
}

func (e *Error) Decode(build vsn.Build, in *packet.WorldPacket) error {
	e.Type = in.Type
	e.Name = in.ReadCString()
	return nil
}
