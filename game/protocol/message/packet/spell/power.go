package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type PowerType uint32

type ModifiedPower struct {
	PowerType  uint32
	PowerValue uint32
}

func (mp *ModifiedPower) Encode(build version.Build, out *message.Packet) error {
	out.WriteUint32(mp.PowerType)
	out.WriteUint32(mp.PowerValue)
	return nil
}

func (mp *ModifiedPower) Decode(build version.Build, in *message.Packet) error {
	mp.PowerType = in.ReadUint32()
	mp.PowerValue = in.ReadUint32()

	return nil
}
