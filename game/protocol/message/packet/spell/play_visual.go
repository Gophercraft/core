package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type PlayVisual struct {
	ID             guid.GUID
	SpellVisualKit uint32
}

func (p *PlayVisual) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_PLAY_SPELL_VISUAL
	if err := p.ID.EncodeUnpacked(build, out); err != nil {
		return err
	}
	out.WriteUint32(p.SpellVisualKit)
	return nil
}
