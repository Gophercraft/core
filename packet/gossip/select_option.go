package gossip

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type SelectOption struct {
	ID     guid.GUID
	Option uint32
}

func (so *SelectOption) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.CMSG_GOSSIP_SELECT_OPTION
	err = so.ID.EncodeUnpacked(build, out)
	out.WriteUint32(so.Option)
	return
}

func (so *SelectOption) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	so.ID, err = guid.DecodeUnpacked(build, in)
	so.Option = in.ReadUint32()
	return
}
