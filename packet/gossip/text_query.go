package gossip

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type TextQuery struct {
	Entry uint32
}

func (tq *TextQuery) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.CMSG_NPC_TEXT_QUERY
	out.WriteUint32(tq.Entry)
	return
}

func (tq *TextQuery) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	tq.Entry = in.ReadUint32()
	return
}
