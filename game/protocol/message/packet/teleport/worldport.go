package teleport

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/version"
)

type Worldport struct {
	Time  uint32
	MapID uint32
	Pos   tempest.C4Vector
}

func (wt *Worldport) Decode(build version.Build, in *message.Packet) error {
	wt.Time = in.ReadUint32()

	if build > 3368 {
		wt.MapID = in.ReadUint32()
	} else {
		// Alpha shenanigans
		wt.MapID = uint32(in.ReadUint8())
	}

	var err error
	wt.Pos, err = tempest.DecodeC4Vector(in)
	return err
}

func (wt *Worldport) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_WORLD_TELEPORT
	out.WriteUint32(wt.Time)
	if build > 3368 {
		out.WriteUint32(wt.MapID)
	} else {
		out.WriteUint8(uint8(wt.MapID))
	}
	return wt.Pos.Encode(out)
}
