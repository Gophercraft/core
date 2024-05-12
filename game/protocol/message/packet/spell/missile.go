package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type MissileTrajectoryResult struct {
	TravelTime uint32
	Pitch      float32
}

func (mtr *MissileTrajectoryResult) Encode(build version.Build, out *message.Packet) error {
	out.WriteUint32(mtr.TravelTime)
	out.WriteFloat32(mtr.Pitch)
	return nil
}

func (mt *MissileTrajectoryResult) Decode(build version.Build, in *message.Packet) error {
	mt.TravelTime = in.ReadUint32()
	mt.Pitch = in.ReadFloat32()
	return nil
}
