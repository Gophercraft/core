package chat

import (
	"fmt"
	"io"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Tag uint8

const (
	TagNone Tag = iota
	TagAFK
	TagDND
	TagGM
	TagCommentator
	TagDeveloper
)

type TagDescriptor map[Tag]uint8

var TagDescriptors = map[vsn.BuildRange]TagDescriptor{
	{0, 5875}: {
		TagNone: 0,
		TagAFK:  1,
		TagDND:  2,
		TagGM:   3,
	},

	{vsn.V2_0_1, vsn.Max}: {
		TagNone:        0x00,
		TagAFK:         0x01,
		TagDND:         0x02,
		TagGM:          0x04,
		TagCommentator: 0x08, // Commentator
		TagDeveloper:   0x10,
	},
}

func (myTag *Tag) Encode(build vsn.Build, out io.Writer) error {
	var tag TagDescriptor
	if err := vsn.QueryDescriptors(build, TagDescriptors, &tag); err != nil {
		return err
	}
	code := tag[*myTag]
	if _, err := out.Write([]byte{code}); err != nil {
		return err
	}
	return nil
}

func (myTag *Tag) Decode(build vsn.Build, in io.Reader) error {
	var tag TagDescriptor
	if err := vsn.QueryDescriptors(build, TagDescriptors, &tag); err != nil {
		return err
	}
	var b [1]byte
	if _, err := in.Read(b[:]); err != nil {
		return err
	}
	for tagType, tagCode := range tag {
		if tagCode == b[0] {
			*myTag = tagType
			return nil
		}
	}
	return fmt.Errorf("packet/chat: unknown code %d", b[0])
}

type ServerMessage struct {
	Type        MsgType
	Language    Language
	ChannelName string
	PlayerRank  uint32
	Name        string
	SenderGUID  guid.GUID
	TargetGUID  guid.GUID
	TargetName  string
	Body        string
	Tag         Tag
}

func (msg *ServerMessage) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_MESSAGECHAT

	msgCode, err := ConvertMsgType(build, msg.Type)
	if err != nil {
		return err
	}

	out.WriteByte(uint8(msgCode))
	out.WriteUint32(uint32(msg.Language))

	switch {
	// Alpha
	case vsn.Range(0, 3368).Contains(build):
		msg.SenderGUID.EncodeUnpacked(build, out)
		EncodeChatString(build, out.Buffer, msg.Body)
		if err := msg.Tag.Encode(build, out); err != nil {
			return err
		}
		// Vanilla
	case vsn.Range(5875, 6005).Contains(build):
		switch msg.Type {
		case MsgCreatureWhisper, MsgRaidBossWhisper, MsgRaidBossEmote, MsgCreatureEmote:
			EncodeChatString(build, out.Buffer, msg.Name)
			msg.SenderGUID.EncodeUnpacked(build, out)
		case MsgSay, MsgParty, MsgYell:
			msg.SenderGUID.EncodeUnpacked(build, out)
			msg.SenderGUID.EncodeUnpacked(build, out)
		case MsgCreatureSay, MsgCreatureYell:
			msg.SenderGUID.EncodeUnpacked(build, out)
			EncodeChatString(build, out.Buffer, msg.Name)
			msg.TargetGUID.EncodeUnpacked(build, out)
		case MsgChannel:
			EncodeChatString(build, out.Buffer, msg.ChannelName)
			out.WriteUint32(msg.PlayerRank)
			msg.SenderGUID.EncodeUnpacked(build, out)
		default:
			msg.SenderGUID.EncodeUnpacked(build, out)
		}

		EncodeChatString(build, out.Buffer, msg.Body)
		if err := msg.Tag.Encode(build, out); err != nil {
			return err
		}
		// TBC-
	case vsn.Range(6180, 12340).Contains(build):
		msg.SenderGUID.EncodeUnpacked(build, out)
		out.WriteUint32(0)

		switch msg.Type {
		case MsgCreatureSay, MsgCreatureParty, MsgCreatureYell, MsgCreatureWhisper, MsgCreatureEmote, MsgRaidBossWhisper, MsgRaidBossEmote, MsgWhisperForeign:
			EncodeChatString(build, out.Buffer, msg.Name)
			msg.TargetGUID.EncodeUnpacked(build, out)
			if msg.Type != MsgWhisperForeign {
				if (msg.TargetGUID != guid.Nil) && !(msg.TargetGUID.HighType() == guid.Player) && !(msg.TargetGUID.HighType() == guid.Pet) {
					EncodeChatString(build, out.Buffer, msg.TargetName)
				}
			}
			EncodeChatString(build, out.Buffer, msg.Body)
			if err := msg.Tag.Encode(build, out); err != nil {
				return err
			}
		case MsgBGSystemNeutral, MsgBGSystemBlueTeam, MsgBGSystemRedTeam:
			msg.TargetGUID.EncodeUnpacked(build, out)
			if msg.TargetGUID != guid.Nil && msg.TargetGUID.HighType() != guid.Player {
				EncodeChatString(build, out.Buffer, msg.TargetName)
			}
			EncodeChatString(build, out.Buffer, msg.Body)
			if err := msg.Tag.Encode(build, out); err != nil {
				return err
			}
		default:
			if msg.Type == MsgChannel {
				out.WriteCString(msg.ChannelName)
			}

			msg.TargetGUID.EncodeUnpacked(build, out)
			EncodeChatString(build, out.Buffer, msg.Body)
			if err := msg.Tag.Encode(build, out); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("chat: unhandled build %s", build)
	}

	return nil
}

func (msg *ServerMessage) Decode(build vsn.Build, in *packet.WorldPacket) error {
	code := in.ReadByte()

	var err error
	msg.Type, err = ResolveMsgType(build, uint32(code))
	if err != nil {
		return err
	}

	msg.Language = Language(in.ReadUint32())

	switch msg.Type {
	case MsgCreatureWhisper, MsgRaidBossWhisper, MsgRaidBossEmote, MsgCreatureEmote:
		msg.Name = DecodeChatString(build, in.Buffer)
		msg.SenderGUID, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}
	case MsgSay, MsgParty, MsgYell:
		msg.SenderGUID = guid.Classic(in.ReadUint64())
		in.ReadUint64()
	case MsgCreatureSay, MsgCreatureYell:
		msg.SenderGUID = guid.Classic(in.ReadUint64())
		msg.Name = DecodeChatString(build, in.Buffer)
		msg.TargetGUID = guid.Classic(in.ReadUint64())
	case MsgChannel:
		msg.ChannelName = DecodeChatString(build, in.Buffer)
		msg.PlayerRank = in.ReadUint32()
		msg.SenderGUID = guid.Classic(in.ReadUint64())
	default:
		msg.SenderGUID = guid.Classic(in.ReadUint64())
	}

	msg.Body = DecodeChatString(build, in.Buffer)
	return nil
}
