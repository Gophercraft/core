package raid

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type InstanceInfo struct {
	Instances []Instance
}

type Instance struct {
	MapID     uint32
	ResetTime uint32
	ID        uint32
}

func (inf *InstanceInfo) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_RAID_INSTANCE_INFO

	out.WriteUint32(uint32(len(inf.Instances)))

	for _, in := range inf.Instances {
		out.WriteUint32(in.MapID)
		out.WriteUint32(in.ResetTime)
		out.WriteUint32(in.ID)
	}

	return nil
}
