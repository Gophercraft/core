package teleport

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
)

type NewWorld struct {
	MapID    uint32
	Position tempest.C4Vector
}

func (nw *NewWorld) Decode(build vsn.Build, in *packet.WorldPacket) error {
	if build > 3368 {
		nw.MapID = in.ReadUint32()
	} else {
		// Alpha shenanigans
		nw.MapID = uint32(in.ReadByte())
	}

	var err error
	nw.Position, err = tempest.DecodeC4Vector(in)
	return err
}

func (nw *NewWorld) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_NEW_WORLD
	if build > 3368 {
		out.WriteUint32(nw.MapID)
	} else {
		out.WriteByte(uint8(nw.MapID))
	}
	return nw.Position.Encode(out)
}
