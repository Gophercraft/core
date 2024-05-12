package area

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type ZoneUpdate struct {
	ID uint32
}

func (zu *ZoneUpdate) Decode(build version.Build, in *message.Packet) error {
	zu.ID = in.ReadUint32()
	return nil
}

func (zu *ZoneUpdate) Encode(build version.Build, out *message.Packet) error {
	out.WriteUint32(zu.ID)
	return nil
}
