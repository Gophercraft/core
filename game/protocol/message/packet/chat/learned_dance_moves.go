package chat

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type LearnedDanceMoves struct {
	Mask [2]uint32
}

func (ldm *LearnedDanceMoves) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_LEARNED_DANCE_MOVES
	out.WriteUint32(ldm.Mask[0])
	out.WriteUint32(ldm.Mask[1])
	return nil
}

func (ldm *LearnedDanceMoves) Decode(build version.Build, in *message.Packet) error {
	ldm.Mask[0] = in.ReadUint32()
	ldm.Mask[1] = in.ReadUint32()
	return nil
}
