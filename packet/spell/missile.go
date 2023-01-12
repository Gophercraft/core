package spell

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type MissileTrajectoryResult struct {
	TravelTime uint32
	Pitch      float32
}

func (mtr *MissileTrajectoryResult) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteUint32(mtr.TravelTime)
	out.WriteFloat32(mtr.Pitch)
	return nil
}

func (mt *MissileTrajectoryResult) Decode(build vsn.Build, in *packet.WorldPacket) error {
	mt.TravelTime = in.ReadUint32()
	mt.Pitch = in.ReadFloat32()
	return nil
}
