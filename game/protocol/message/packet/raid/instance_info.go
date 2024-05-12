package raid

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type InstanceInfo struct {
	Instances []Instance
}

type Instance struct {
	MapID     uint32
	ResetTime uint32
	ID        uint32
}

func (inf *InstanceInfo) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_RAID_INSTANCE_INFO

	out.WriteUint32(uint32(len(inf.Instances)))

	for _, in := range inf.Instances {
		out.WriteUint32(in.MapID)
		out.WriteUint32(in.ResetTime)
		out.WriteUint32(in.ID)
	}

	return nil
}
