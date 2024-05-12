package party

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type SetLeader struct {
	Name string
}

func (set *SetLeader) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_GROUP_SET_LEADER
	out.WriteCString(set.Name)
	return nil
}

func (set *SetLeader) Decode(build version.Build, in *message.Packet) error {
	set.Name = in.ReadCString()
	return nil
}
