package teleport

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/version"
)

type BindpointUpdate struct {
	Position tempest.C3Vector
	ZoneID   uint32
	MapID    uint32
}

const zoneAdded = version.V1_12_1

func (bpu *BindpointUpdate) Decode(build version.Build, in *message.Packet) (err error) {
	bpu.Position, err = tempest.DecodeC3Vector(in)
	if err != nil {
		return err
	}
	if build.AddedIn(zoneAdded) {
		bpu.ZoneID = in.ReadUint32()
	}
	bpu.MapID = in.ReadUint32()
	return nil
}

func (bpu *BindpointUpdate) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_BINDPOINTUPDATE
	if err := bpu.Position.Encode(out); err != nil {
		return err
	}
	if build.AddedIn(zoneAdded) {
		out.WriteUint32(bpu.ZoneID)
	}
	out.WriteUint32(bpu.MapID)
	return nil
}
