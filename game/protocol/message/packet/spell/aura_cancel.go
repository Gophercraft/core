package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type AuraCancel struct {
	Spell uint32
}

func (ac *AuraCancel) Decode(build version.Build, in *message.Packet) error {
	ac.Spell = in.ReadUint32()
	return nil
}
