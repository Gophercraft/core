package gossip

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Hello struct {
	ID guid.GUID
}

func (h *Hello) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.CMSG_GOSSIP_HELLO
	err = h.ID.EncodeUnpacked(build, out)
	return
}

func (h *Hello) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	h.ID, err = guid.DecodeUnpacked(build, in)
	return
}
