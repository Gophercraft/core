package xp

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type Exploration struct {
	ZoneID     uint32
	Experience uint32
}

func (ee *Exploration) Decode(build version.Build, in *message.Packet) error {
	ee.ZoneID = in.ReadUint32()
	ee.Experience = in.ReadUint32()
	return nil
}

func (ee *Exploration) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_EXPLORATION_EXPERIENCE
	out.WriteUint32(ee.ZoneID)
	out.WriteUint32(ee.Experience)
	return nil
}
