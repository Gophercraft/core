package spell

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type AuraCancel struct {
	Spell uint32
}

func (ac *AuraCancel) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ac.Spell = in.ReadUint32()
	return nil
}
