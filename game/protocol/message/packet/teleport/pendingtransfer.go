package teleport

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type TransferPending struct {
	MapID uint32
}

func (tp *TransferPending) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_TRANSFER_PENDING
	out.WriteUint32(tp.MapID)
	return nil
}

func (tp *TransferPending) Decode(build version.Build, in *message.Packet) error {
	tp.MapID = in.ReadUint32()
	return nil
}
