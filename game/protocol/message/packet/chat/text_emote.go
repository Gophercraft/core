package chat

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
	"github.com/superp00t/etc"
)

type TextEmoteRequest struct {
	Text  uint32
	Emote uint32
	GUID  guid.GUID
}

func (em *TextEmoteRequest) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_TEXT_EMOTE
	out.WriteUint32(em.Text)
	out.WriteUint32(em.Emote)
	em.GUID.EncodeUnpacked(build, out)
	return nil
}

func (em *TextEmoteRequest) Decode(build version.Build, in *message.Packet) (err error) {
	em.Text = in.ReadUint32()
	em.Emote = in.ReadUint32()
	em.GUID, err = guid.DecodeUnpacked(build, in)
	return err
}

type TextEmote struct {
	GUID  guid.GUID
	Text  uint32
	Emote uint32
	Name  string
}

func (em *TextEmote) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_TEXT_EMOTE
	em.GUID.EncodeUnpacked(build, out)
	out.WriteUint32(em.Text)
	out.WriteUint32(em.Emote)
	nameBytes := append(
		[]byte(em.Name),
		0,
	)
	out.WriteUint32(uint32(len(nameBytes)))
	out.Write(nameBytes)
	return nil
}

func (em *TextEmote) Decode(build version.Build, in *message.Packet) (err error) {
	em.GUID, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return
	}

	em.Text = in.ReadUint32()
	em.Emote = in.ReadUint32()
	nameLength := in.ReadUint()
	em.Name = etc.OfBytes(in.ReadBytes(int(nameLength))).ReadCString()

	return
}

type Emote struct {
	Emote uint32
	GUID  guid.GUID
}

func (em *Emote) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_EMOTE
	out.WriteUint32(em.Emote)
	em.GUID.EncodeUnpacked(build, out)
	return nil
}

func (em *Emote) Decode(build version.Build, in *message.Packet) (err error) {
	em.Emote = in.ReadUint32()
	em.GUID, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return
	}
	return
}
