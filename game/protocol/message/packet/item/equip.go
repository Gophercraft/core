package item

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/version"
)

type AutoEquipItemRequest struct {
	SrcBag  models.ItemSlot
	SrcSlot models.ItemSlot
}

func (ae *AutoEquipItemRequest) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_AUTOEQUIP_ITEM
	out.WriteInt8(int8(ae.SrcBag))
	out.WriteInt8(int8(ae.SrcSlot))
	return nil
}

func (ae *AutoEquipItemRequest) Decode(build version.Build, in *message.Packet) error {
	ae.SrcBag = models.ItemSlot(in.ReadInt8())
	ae.SrcSlot = models.ItemSlot(in.ReadInt8())
	return nil
}
