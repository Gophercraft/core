package area

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc"
)

type TriggerMessage struct {
	Message string
}

func (tm *TriggerMessage) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_AREA_TRIGGER_MESSAGE
	out.WriteUint32(uint32(len(tm.Message) + 1))
	out.WriteCString(tm.Message)
	return nil
}

func (tm *TriggerMessage) Decode(build vsn.Build, in *packet.WorldPacket) error {
	prefix := int(in.ReadUint32())
	data := etc.OfBytes(in.ReadBytes(prefix))
	tm.Message = data.ReadCString()
	return nil
}
