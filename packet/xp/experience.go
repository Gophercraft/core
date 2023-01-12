package xp

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Exploration struct {
	ZoneID     uint32
	Experience uint32
}

func (ee *Exploration) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ee.ZoneID = in.ReadUint32()
	ee.Experience = in.ReadUint32()
	return nil
}

func (ee *Exploration) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_EXPLORATION_EXPERIENCE
	out.WriteUint32(ee.ZoneID)
	out.WriteUint32(ee.Experience)
	return nil
}
