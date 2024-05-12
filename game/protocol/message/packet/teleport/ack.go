package teleport

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/version"
)

type Ack struct {
	GUID guid.GUID

	Info *update.MovementInfo
}

func (ta *Ack) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.MSG_MOVE_TELEPORT_ACK
	ta.GUID.EncodePacked(build, out)
	out.WriteUint32(0)
	if err := update.EncodeMovementInfo(build, out.Buffer, ta.Info); err != nil {
		return err
	}
	return nil
}

func (ta *Ack) Decode(build version.Build, in *message.Packet) (err error) {
	ta.GUID, err = guid.DecodePacked(build, in)
	if err != nil {
		return
	}
	in.ReadUint32()
	if ta.Info, err = update.DecodeMovementInfo(build, in); err != nil {
		return err
	}
	return nil
}
