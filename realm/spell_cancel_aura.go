package realm

import "github.com/Gophercraft/core/packet/spell"

func (s *Session) CanUnaura(spellID uint32) bool {
	if s.God() {
		return true
	}

	return true
}

func (s *Session) HandleAuraCancel(ac *spell.AuraCancel) {
	if s.CanUnaura(ac.Spell) {
		s.Map().Unaura(s, ac.Spell, -1, true)
	}
}
