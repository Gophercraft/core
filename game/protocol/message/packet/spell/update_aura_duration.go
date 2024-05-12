package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type AuraDurationUpdate struct {
	Slot     uint8
	Duration uint32
}

func (adu *AuraDurationUpdate) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_UPDATE_AURA_DURATION
	out.WriteUint8(adu.Slot)
	out.WriteUint32(adu.Duration)
	return nil
}
