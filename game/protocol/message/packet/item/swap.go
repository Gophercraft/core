package item

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/version"
)

type SwapBackpackRequest struct {
	SrcSlot models.ItemSlot
	DstSlot models.ItemSlot
}

func (swp *SwapBackpackRequest) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_SWAP_INV_ITEM
	out.WriteInt8(int8(swp.SrcSlot))
	out.WriteInt8(int8(swp.DstSlot))
	return nil
}

func (swp *SwapBackpackRequest) Decode(build version.Build, in *message.Packet) error {
	swp.SrcSlot = models.ItemSlot(in.ReadInt8())
	swp.DstSlot = models.ItemSlot(in.ReadInt8())
	return nil
}

type SwapRequest struct {
	DstBag, DstSlot models.ItemSlot
	SrcBag, SrcSlot models.ItemSlot
}

func (swp *SwapRequest) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_SWAP_ITEM
	out.WriteInt8(int8(swp.DstBag))
	out.WriteInt8(int8(swp.DstSlot))
	out.WriteInt8(int8(swp.SrcBag))
	out.WriteInt8(int8(swp.SrcSlot))
	return nil
}

func (swp *SwapRequest) Decode(build version.Build, in *message.Packet) error {
	swp.DstBag = models.ItemSlot(in.ReadInt8())
	swp.DstSlot = models.ItemSlot(in.ReadInt8())
	swp.SrcBag = models.ItemSlot(in.ReadInt8())
	swp.SrcSlot = models.ItemSlot(in.ReadInt8())
	return nil
}

type AutoStoreBag struct {
	SrcBag, SrcSlot models.ItemSlot
	DstBag          models.ItemSlot
}

func (au *AutoStoreBag) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_AUTOSTORE_BAG_ITEM
	out.WriteInt8(int8(au.SrcBag))
	out.WriteInt8(int8(au.SrcSlot))
	out.WriteInt8(int8(au.DstBag))
	return nil
}

func (au *AutoStoreBag) Decode(build version.Build, in *message.Packet) error {
	au.SrcBag = models.ItemSlot(in.ReadInt8())
	au.SrcSlot = models.ItemSlot(in.ReadInt8())
	au.DstBag = models.ItemSlot(in.ReadInt8())
	return nil
}
