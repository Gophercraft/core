package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/version"
)

type UseItem struct {
	Bag        models.ItemSlot
	Slot       models.ItemSlot
	SpellCount uint8
	CastCount  uint8
	Targets    TargetData
}

func (use *UseItem) Decode(build version.Build, in *message.Packet) error {
	use.Bag = models.ItemSlot(in.ReadInt8())
	use.Slot = models.ItemSlot(in.ReadUint8())

	use.SpellCount = in.ReadUint8()
	use.CastCount = in.ReadUint8()

	return use.Targets.Decode(build, in)
}
