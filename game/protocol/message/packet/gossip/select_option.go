package gossip

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type SelectOption struct {
	ID     guid.GUID
	Option uint32
}

func (so *SelectOption) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.CMSG_GOSSIP_SELECT_OPTION
	err = so.ID.EncodeUnpacked(build, out)
	out.WriteUint32(so.Option)
	return
}

func (so *SelectOption) Decode(build version.Build, in *message.Packet) (err error) {
	so.ID, err = guid.DecodeUnpacked(build, in)
	so.Option = in.ReadUint32()
	return
}
