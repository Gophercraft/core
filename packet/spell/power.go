package spell

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type PowerType uint32

type ModifiedPower struct {
	PowerType  uint32
	PowerValue uint32
}

func (mp *ModifiedPower) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteUint32(mp.PowerType)
	out.WriteUint32(mp.PowerValue)
	return nil
}

func (mp *ModifiedPower) Decode(build vsn.Build, in *packet.WorldPacket) error {
	mp.PowerType = in.ReadUint32()
	mp.PowerValue = in.ReadUint32()

	return nil
}
