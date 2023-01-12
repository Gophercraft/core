package teleport

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
)

type Worldport struct {
	Time  uint32
	MapID uint32
	Pos   tempest.C4Vector
}

func (wt *Worldport) Decode(build vsn.Build, in *packet.WorldPacket) error {
	wt.Time = in.ReadUint32()

	if build > 3368 {
		wt.MapID = in.ReadUint32()
	} else {
		// Alpha shenanigans
		wt.MapID = uint32(in.ReadByte())
	}

	var err error
	wt.Pos, err = tempest.DecodeC4Vector(in)
	return err
}

func (wt *Worldport) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_WORLD_TELEPORT
	out.WriteUint32(wt.Time)
	if build > 3368 {
		out.WriteUint32(wt.MapID)
	} else {
		out.WriteByte(uint8(wt.MapID))
	}
	return wt.Pos.Encode(out)
}
