package item

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type SplitRequest struct {
	SrcBag, SrcSlot models.ItemSlot
	DstBag, DstSlot models.ItemSlot
	Count           uint32
}

func (isr *SplitRequest) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_SPLIT_ITEM
	out.WriteInt8(int8(isr.SrcBag))
	out.WriteInt8(int8(isr.SrcSlot))
	out.WriteInt8(int8(isr.DstBag))
	out.WriteInt8(int8(isr.DstSlot))
	out.WriteUint32(isr.Count)
	return nil
}

func (isr *SplitRequest) Decode(build vsn.Build, in *packet.WorldPacket) error {
	isr.SrcBag = models.ItemSlot(in.ReadInt8())
	isr.SrcSlot = models.ItemSlot(in.ReadInt8())
	isr.DstBag = models.ItemSlot(in.ReadInt8())
	isr.DstSlot = models.ItemSlot(in.ReadInt8())
	isr.Count = in.ReadUint32()
	return nil
}
