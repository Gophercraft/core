package cinematic

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type TriggerSequence struct {
	ID int32
}

func (ts *TriggerSequence) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_TRIGGER_CINEMATIC
	out.WriteInt32(ts.ID)
	return nil
}
