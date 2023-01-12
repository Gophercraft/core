package packet

import (
	"github.com/superp00t/etc"
)

type WorldPacket struct {
	Type WorldType
	*etc.Buffer
}

func NewWorldPacket(t WorldType) *WorldPacket {
	return &WorldPacket{t, etc.NewBuffer()}
}

func (wp *WorldPacket) ReadBool32() bool {
	b32 := wp.ReadUint32()
	return b32 == 1
}

func (wp *WorldPacket) WriteBool32(b bool) {
	if b {
		wp.WriteUint32(1)
	} else {
		wp.WriteUint32(0)
	}
}
