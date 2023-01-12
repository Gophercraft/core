package chat

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type LearnedDanceMoves struct {
	Mask [2]uint32
}

func (ldm *LearnedDanceMoves) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_LEARNED_DANCE_MOVES
	out.WriteUint32(ldm.Mask[0])
	out.WriteUint32(ldm.Mask[1])
	return nil
}

func (ldm *LearnedDanceMoves) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ldm.Mask[0] = in.ReadUint32()
	ldm.Mask[1] = in.ReadUint32()
	return nil
}
