package realm

import (
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/core/realm/wdb/models"
)

func (m *Map) PlaySpellVisualKind(unit Unit, spellID uint32, class models.VisualClass) {
	var spelldata *dbdefs.Ent_Spell
	var visualdata *dbdefs.Ent_SpellVisual
	var v int32

	m.LookupSpellInfo(spellID, &spelldata, &visualdata, nil)

	if visualdata != nil {
		switch class {
		case models.VisualPrecast:
			v = visualdata.PrecastKit
		case models.VisualCast:
			v = visualdata.CastKit
		case models.VisualImpact:
			v = visualdata.ImpactKit
		case models.VisualState:
			v = visualdata.StateKit
		case models.VisualStateDone:
			v = visualdata.StateDoneKit
		case models.VisualChannel:
			v = visualdata.ChannelKit
		default:
			panic(class)
		}
	}

	if v != 0 {
		m.PlaySpellVisual(unit, uint32(v))
	}
}

func (m *Map) PlaySpellVisual(unit Unit, visualKit uint32) {
	nearSessions := m.GetObjectsNearPosition(unit.Movement().Position.C3()).Sessions()
	nearSessions.Iter(func(s *Session) {
		s.Send(&spell.PlayVisual{
			ID:             unit.GUID(),
			SpellVisualKit: visualKit,
		})
	})
}
