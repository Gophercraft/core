package gossip

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type Hello struct {
	ID guid.GUID
}

func (h *Hello) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.CMSG_GOSSIP_HELLO
	err = h.ID.EncodeUnpacked(build, out)
	return
}

func (h *Hello) Decode(build version.Build, in *message.Packet) (err error) {
	h.ID, err = guid.DecodeUnpacked(build, in)
	return
}
