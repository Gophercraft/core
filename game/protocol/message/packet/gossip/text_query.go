package gossip

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type TextQuery struct {
	Entry uint32
}

func (tq *TextQuery) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.CMSG_NPC_TEXT_QUERY
	out.WriteUint32(tq.Entry)
	return
}

func (tq *TextQuery) Decode(build version.Build, in *message.Packet) (err error) {
	tq.Entry = in.ReadUint32()
	return
}
