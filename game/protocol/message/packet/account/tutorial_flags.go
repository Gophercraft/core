package account

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type TutorialFlags struct {
	Mask [8]uint32
}

func (tf *TutorialFlags) SetBit(index int, val bool) {
	max := 32 * len(tf.Mask)
	if index > max {
		return
	}

	blockOffset := int(index / 32)
	bitOffset := int(index % 32)

	if val {
		tf.Mask[blockOffset] |= (1 << bitOffset)
	} else {
		tf.Mask[blockOffset] &= ^(1 << bitOffset)
	}
}

func (tf *TutorialFlags) SetAll() {
	for i := 0; i < len(tf.Mask); i++ {
		tf.Mask[i] = 0xFFFFFFFF
	}
}

func (tf *TutorialFlags) maskSize(build version.Build) int {
	return 8
}

func (tf *TutorialFlags) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_TUTORIAL_FLAGS

	maskLen := tf.maskSize(build)

	for i := 0; i < maskLen; i++ {
		out.WriteUint32(tf.Mask[i])
	}

	return nil
}
