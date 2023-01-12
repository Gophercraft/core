package area

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Trigger struct {
	ID uint32
}

func (at *Trigger) Decode(build vsn.Build, in *packet.WorldPacket) error {
	at.ID = in.ReadUint32()
	return nil
}

func (at *Trigger) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_AREATRIGGER
	out.WriteUint32(at.ID)
	return nil
}
