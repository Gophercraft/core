package spell

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type NewSlot struct {
	Spell uint32
	Index int32
}

func (ns *NewSlot) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ns.Spell = in.ReadUint32()
	ns.Index = in.ReadInt32()
	return nil
}
