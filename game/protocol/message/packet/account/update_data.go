package account

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type UpdateData struct {
	AccountDataType int32
	Time            uint32
}

func (upd8 *UpdateData) Decode(build version.Build, in *message.Packet) error {
	upd8.AccountDataType = in.ReadInt32()

	return nil
}

type ClientUpdateData struct {
	UpdateData
}

func (clupd8 *ClientUpdateData) Decode(build version.Build, in *message.Packet) error {
	if err := clupd8.UpdateData.Decode(build, in); err != nil {
		return err
	}
	return nil
}
