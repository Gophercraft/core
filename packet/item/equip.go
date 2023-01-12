package item

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type AutoEquipItemRequest struct {
	SrcBag  models.ItemSlot
	SrcSlot models.ItemSlot
}

func (ae *AutoEquipItemRequest) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_AUTOEQUIP_ITEM
	out.WriteInt8(int8(ae.SrcBag))
	out.WriteInt8(int8(ae.SrcSlot))
	return nil
}

func (ae *AutoEquipItemRequest) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ae.SrcBag = models.ItemSlot(in.ReadInt8())
	ae.SrcSlot = models.ItemSlot(in.ReadInt8())
	return nil
}
