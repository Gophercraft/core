package account

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type UpdateData struct {
	AccountDataType int32
	Time            uint32
}

func (upd8 *UpdateData) Decode(build vsn.Build, in *packet.WorldPacket) error {
	upd8.AccountDataType = in.ReadInt32()

	return nil
}

type ClientUpdateData struct {
	UpdateData
}

func (clupd8 *ClientUpdateData) Decode(build vsn.Build, in *packet.WorldPacket) error {
	if err := clupd8.UpdateData.Decode(build, in); err != nil {
		return err
	}
	return nil
}
