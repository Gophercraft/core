package teleport

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/version"
)

type NewWorld struct {
	MapID    uint32
	Position tempest.C4Vector
}

func (nw *NewWorld) Decode(build version.Build, in *message.Packet) error {
	if build > 3368 {
		nw.MapID = in.ReadUint32()
	} else {
		// Alpha shenanigans
		nw.MapID = uint32(in.ReadUint8())
	}

	var err error
	nw.Position, err = tempest.DecodeC4Vector(in)
	return err
}

func (nw *NewWorld) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_NEW_WORLD
	if build > 3368 {
		out.WriteUint32(nw.MapID)
	} else {
		out.WriteUint8(uint8(nw.MapID))
	}
	return nw.Position.Encode(out)
}
