package realm

import (
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/packet/spell"
)

type SpellEffectData struct {
	Map         *Map
	Spell       *dbdefs.Ent_Spell
	EffectIndex int
	Caster      Unit
	Target      Unit
	Aura        *Aura
}

type SpellEffect func(*SpellEffectData)

func (s *Server) initSpellEffects() {
	s.SpellEffects = make([]SpellEffect, spell.NumEffects)
	s.SpellEffects[spell.EffectApplyAura] = EffectApplyAura
}
