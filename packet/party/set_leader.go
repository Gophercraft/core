package party

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type SetLeader struct {
	Name string
}

func (set *SetLeader) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_GROUP_SET_LEADER
	out.WriteCString(set.Name)
	return nil
}

func (set *SetLeader) Decode(build vsn.Build, in *packet.WorldPacket) error {
	set.Name = in.ReadCString()
	return nil
}
