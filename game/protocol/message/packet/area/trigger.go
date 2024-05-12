package area

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type Trigger struct {
	ID uint32
}

func (at *Trigger) Decode(build version.Build, in *message.Packet) error {
	at.ID = in.ReadUint32()
	return nil
}

func (at *Trigger) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_AREATRIGGER
	out.WriteUint32(at.ID)
	return nil
}
