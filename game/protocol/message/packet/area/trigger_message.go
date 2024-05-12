package area

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
	"github.com/superp00t/etc"
)

type TriggerMessage struct {
	Message string
}

func (tm *TriggerMessage) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_AREA_TRIGGER_MESSAGE
	out.WriteUint32(uint32(len(tm.Message) + 1))
	out.WriteCString(tm.Message)
	return nil
}

func (tm *TriggerMessage) Decode(build version.Build, in *message.Packet) error {
	prefix := int(in.ReadUint32())
	data := etc.OfBytes(in.ReadBytes(prefix))
	tm.Message = data.ReadCString()
	return nil
}
