package area

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type ZoneUpdate struct {
	ID uint32
}

func (zu *ZoneUpdate) Decode(build vsn.Build, in *packet.WorldPacket) error {
	zu.ID = in.ReadUint32()
	return nil
}

func (zu *ZoneUpdate) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteUint32(zu.ID)
	return nil
}
