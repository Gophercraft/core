package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type NewSlot struct {
	Spell uint32
	Index int32
}

func (ns *NewSlot) Decode(build version.Build, in *message.Packet) error {
	ns.Spell = in.ReadUint32()
	ns.Index = in.ReadInt32()
	return nil
}
