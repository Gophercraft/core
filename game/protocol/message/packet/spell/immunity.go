package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type CreatureImmunities struct {
	School uint32
	Value  uint32
}

func (ci *CreatureImmunities) Encode(build version.Build, out *message.Packet) error {
	out.WriteUint32(ci.School)
	out.WriteUint32(ci.Value)
	return nil
}

func (ci *CreatureImmunities) Decode(build version.Build, in *message.Packet) error {
	ci.School = in.ReadUint32()
	ci.Value = in.ReadUint32()
	return nil
}
