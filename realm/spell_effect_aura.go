package realm

import (
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

// ModEffect. Stores mods for all the stats in StatMask
type ModEffect struct {
	Op       models.ModOp
	StatMask models.ModStatMask
	Value    float64
}

// Applies the effect of an aura.
func EffectApplyAura(dat *SpellEffectData) {
	target := dat.Target
	if target == nil {
		return
	}

	auraEffectKind := spell.AuraEffect(dat.Spell.EffectAura[dat.EffectIndex])

	if auraEffectKind == spell.AuraEffectNone {
		return
	}

	if len(dat.Map.Phase.Server.AuraEffects) <= int(auraEffectKind) {
		log.Warn("Could not get aura effect (out of bounds)", auraEffectKind)
		return
	}

	auraEffect := dat.Map.Phase.Server.AuraEffects[auraEffectKind]
	if auraEffect == nil {
		log.Warn("Could not get aura effect", auraEffectKind)
	} else {
		auraEffect(dat, dat.Target, dat.Aura)
	}

	// var au Aura
	// // au.ID = uint32(dat.Spell.EffectAura[dat.EffectIndex])
	// au.ID = uint32(dat.Spell.ID)
	// au.Applications = 1
	// // au.Level =

	// if err := dat.Map.ApplyAura(&au, target); err != nil {
	// 	panic(err)
	// }
}

type AuraEffectData struct {
}

type AuraEffect func(sed *SpellEffectData, target Unit, au *Aura)

func (s *Server) initAuraEffects() {
	s.AuraEffects = make([]AuraEffect, spell.NumAuraEffects)
	s.AuraEffects[spell.AuraEffectModResistance] = AuraEffectModResistance
}

func resistStat(resistance models.Resistance) models.ModStat {
	switch resistance {
	case models.ResistPhysical:
		return models.ResistPhysicalDamage
	case models.ResistHoly:
		return models.ResistHolyDamage
	case models.ResistFire:
		return models.ResistFireDamage
	case models.ResistNature:
		return models.ResistNatureDamage
	case models.ResistFrost:
		return models.ResistFrostDamage
	case models.ResistShadow:
		return models.ResistShadowDamage
	case models.ResistArcane:
		return models.ResistArcaneDamage
	default:
		panic(resistance)
	}
}

func (s *Server) buildModResistStatMask(sed *SpellEffectData) (mask models.ModStatMask) {
	build := s.Build()

	if build <= vsn.V3_0_2 {
		misc := sed.Spell.EffectMiscValue[sed.EffectIndex]
		for i := models.ResistPhysical; i < models.NumResists; i++ {
			flag := int32(1) << int32(i)

			if misc&flag != 0 {
				mask.Add(resistStat(i))
			}
		}
	}

	return
}

func AuraEffectModResistance(sed *SpellEffectData, target Unit, au *Aura) {
	// state := target.GetAuras()

	modEffect := &ModEffect{
		Op:       models.ModPlus,
		StatMask: sed.Map.Phase.Server.buildModResistStatMask(sed),
		Value:    1 + float64(sed.Spell.EffectBasePoints[sed.EffectIndex]),
	}

	au.ModEffects = append(au.ModEffects, modEffect)
}
