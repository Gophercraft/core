package spell

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type CreatureImmunities struct {
	School uint32
	Value  uint32
}

func (ci *CreatureImmunities) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteUint32(ci.School)
	out.WriteUint32(ci.Value)
	return nil
}

func (ci *CreatureImmunities) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ci.School = in.ReadUint32()
	ci.Value = in.ReadUint32()
	return nil
}
