package realm

import (
	"fmt"
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

func (s *Session) CanCast(cc *spell.Cast) (err packet.Encodable) {
	if s.gameMode == models.GameMode_God {
		// If you're God, bypass all checks.
		return nil
	}

	var spelldata *dbdefs.Ent_Spell
	s.DB().Lookup(wdb.BucketKeyUint32ID, cc.Spell, &spelldata)

	if spelldata == nil {
		err = &spell.CastResult{
			Status:  spell.SpellUnavailable,
			SpellID: cc.Spell,
		}
		return
	}

	if !s.KnowsAbility(cc.Spell) {
		err = &spell.CastResult{
			Status:  spell.SpellUnavailable,
			SpellID: cc.Spell,
		}
		return
	}

	attr, erro := spell.LoadAttributes(s.Build(), spelldata)
	if erro != nil {
		panic(erro)
	}

	// if attr.Enabled(uint32(spell.Attr_))

	// cannot cast passive
	if attr.Enabled(uint32(spell.Attr_Passive)) {
		err = &spell.CastResult{
			Status:  spell.SpellUnavailable,
			SpellID: cc.Spell,
		}
		return
	}

	return nil
}

func (s *Session) HandleCast(cc *spell.Cast) {
	log.Dump("Cast", cc)

	if s.SpellManager != nil {
		s.SpellManager.StartCast <- cc
	}
}

func (m *Map) StartCast(cc *spell.Cast, caster Unit) (castTime time.Duration, err error) {
	var spelldata *dbdefs.Ent_Spell
	var casttime *dbdefs.Ent_SpellCastTimes
	m.LookupSpellInfo(cc.Spell, &spelldata, nil, &casttime)

	if spelldata == nil {
		err = fmt.Errorf("realm: you should verify spellID is correct before calling (*Map).StartCast: spell %d not found", cc.Spell)
		return
	}

	pos := caster.Movement().Position

	target := cc.Target

	if target == guid.Nil {
		if len(spelldata.ImplicitTargetA) == 0 {
			target = caster.GUID()
		} else {
			implicitTarget := models.ImplictTarget(spelldata.ImplicitTargetA[0])
			switch implicitTarget {
			case models.TargetUnitCaster:
				target = caster.GUID()
			// case models.TargetUnitNearbyEnemy:
			// 	target =
			// 	models.TargetUnitNearbyAlly
			// 	models.TargetUnitNearbyParty
			// 	models.TargetUnitPet
			// 	models.TargetUnitTargetEnemy
			// 	models.TargetUnitSrcAreaEntry
			// 	models.TargetUnitDestAreaEntry
			default:
				log.Warn("Unknown implicit target", target)
			}
		}
	}

	if target == guid.Nil {
		target = caster.GUID()
		if caster.GUID() == guid.Nil {
			log.Warn("Nil caster?")
		}
	}

	targetObject := m.GetObject(target)

	if targetObject == nil {
		err = fmt.Errorf("realm: StartCast: target lacks an object, %s", target)
		return
	}

	// var targetUnit Unit

	// if targetObject != nil {
	// 	pos = targetObject.Movement().Position

	// 	var ok bool
	// 	// TODO: Figure out how we can cast on GameObjects
	// 	targetUnit, ok = targetObject.(Unit)
	// 	if !ok {
	// 		return fmt.Errorf("realm: Cast(), %s is not a Unit", target)
	// 	}
	// }

	m.PlaySpellVisualKind(caster, cc.Spell, models.VisualPrecast)

	nearRange := tempest.CAaSphere{
		Position: pos.C3(),
		Radius:   m.VisibilityDistance(),
	}

	nearSessions := m.GetObjectsInRange(nearRange).Sessions()

	start := &spell.Start{}
	start.Data = cc.Data

	start.Source = caster.GUID()
	start.Caster = caster.GUID()

	nearSessions.Iter(func(s *Session) {
		s.Send(start)
	})

	if casttime != nil {
		castTime = time.Duration(casttime.Base) * time.Millisecond

		eachLevel := time.Duration(caster.Values().Get("Level").Uint32() - 1)

		castTime += (time.Duration(casttime.PerLevel) * eachLevel)

		minimumCastTime := time.Duration(casttime.Minimum) * time.Millisecond

		if castTime < minimumCastTime {
			castTime = minimumCastTime
		}
	}

	return
}

func (m *Map) GoCast(cc *spell.Cast, caster Unit) error {

	var spelldata *dbdefs.Ent_Spell
	m.Phase.Server.DB.Lookup(wdb.BucketKeyUint32ID, cc.Spell, &spelldata)
	if spelldata == nil {
		return fmt.Errorf("realm: you should verify spellID is correct before calling (*Map).Cast: spell %d not found", cc.Spell)
	}

	pos := caster.Movement().Position

	target := cc.Target

	if target == guid.Nil {
		if len(spelldata.ImplicitTargetA) == 0 {
			target = caster.GUID()
		} else {
			implicitTarget := models.ImplictTarget(spelldata.ImplicitTargetA[0])
			switch implicitTarget {
			case models.TargetUnitCaster:
				target = caster.GUID()
			case models.TargetUnitTargetAlly:
				target = caster.GUID()
			// case models.TargetUnitNearbyEnemy:
			// 	target =
			// 	models.TargetUnitNearbyAlly
			// 	models.TargetUnitNearbyParty
			// 	models.TargetUnitPet
			// 	models.TargetUnitTargetEnemy
			// 	models.TargetUnitSrcAreaEntry
			// 	models.TargetUnitDestAreaEntry
			default:
				log.Warn("Unhandled implicit target in spell", spelldata.ID)
			}
		}
	}

	targetObject := m.GetObject(target)

	if targetObject == nil {
		return fmt.Errorf("realm: spell manager: target lacks an object, %s", target)
	}

	var targetUnit Unit

	if targetObject != nil {
		pos = targetObject.Movement().Position

		var ok bool
		// TODO: Figure out how we can cast on GameObjects
		targetUnit, ok = targetObject.(Unit)
		if !ok {
			return fmt.Errorf("realm: Cast(), %s is not a Unit", target)
		}
	}

	m.PlaySpellVisualKind(caster, cc.Spell, models.VisualCast)

	nearRange := tempest.CAaSphere{
		Position: pos.C3(),
		Radius:   m.VisibilityDistance(),
	}

	nearSessions := m.GetObjectsInRange(nearRange).Sessions()

	spellGo := &spell.Go{}
	spellGo.Spell = cc.Spell

	spellGo.Source = caster.GUID()
	spellGo.Caster = caster.GUID()

	spellGo.TargetFlags |= spell.HasTargetMask
	spellGo.Target = targetUnit.GUID()

	if spellGo.Target == guid.Nil {
		spellGo.Target = caster.GUID()
	}

	// castLog := &spell.AuraCastLog{
	// 	Caster: caster.GUID(),
	// 	Target: target,
	// 	Spell:  cc.Spell,
	// }

	var applyAura bool

	// Check if this spell is a type of spell that adds itself as an aura
	for _, eff := range spelldata.Effect {
		switch spell.Effect(eff) {
		case spell.EffectApplyAura:
			applyAura = true
		}
	}

	if applyAura {
		spellGo.HitTargets = []guid.GUID{
			spellGo.Target,
		}
		// spellGo.Target =
	}

	nearSessions.Iter(func(s *Session) {
		s.Send(spellGo)
	})

	// If so, load all its effects into its own aura
	if applyAura {
		aa := &AuraApplication{
			ID:           cc.Spell,
			Applications: 1,
			Level:        int32(caster.Values().Get("Level").Uint32()),
		}

		m.PlaySpellVisualKind(targetUnit, cc.Spell, models.VisualImpact)
		m.PlaySpellVisualKind(targetUnit, cc.Spell, models.VisualState)

		if err := m.ApplyAura(targetUnit, aa); err != nil {
			return err
		}
	}

	if err := m.RefreshStats(targetUnit); err != nil {
		return err
	}

	// attr, erro := spell.LoadAttributes(m.Config().Version, spelldata)
	// if erro != nil {
	// 	panic(erro)
	// }

	// if attr.Enabled(spell.AttrSelf) {

	// }

	if m.Phase.Server.Build().AddedIn(vsn.V1_12_1) {
		execute := &spell.LogExecute{
			Caster: caster.GUID(),
			Spell:  cc.Spell,
		}

		// execute.Effects = append(execute.Effects, spell.LoggedEffect{
		// 	Kind: spell.EffectApplyAura,
		// 	Applications: []spell.LogEffectApplication{
		// 		{
		// 			Target: target,
		// 		},
		// 	},
		// })

		nearSessions.Iter(func(s *Session) {
			s.Send(execute)
		})
	}

	return nil
}
