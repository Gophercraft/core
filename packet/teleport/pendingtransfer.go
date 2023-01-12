package teleport

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type TransferPending struct {
	MapID uint32
}

func (tp *TransferPending) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_TRANSFER_PENDING
	out.WriteUint32(tp.MapID)
	return nil
}

func (tp *TransferPending) Decode(build vsn.Build, in *packet.WorldPacket) error {
	tp.MapID = in.ReadUint32()
	return nil
}
