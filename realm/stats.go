package realm

import (
	"fmt"

	"github.com/Gophercraft/core/realm/wdb/models"
)

// func (player *PlayerSession) CalcBaseHealth() uint32 {
// 	playerLevel := player.Char.Level

// 	var levelFactor float32 = 0

// 	if playerLevel == 0 {
// 		levelFactor = 1
// 	} else {
// 		levelFactor = float32(playerLevel)
// 	}

// }

// func (player *PlayerSession) CalcMaxHealth() {

// }

func (player *Session) GetAuras() *AuraState {
	return &player.AuraState
}

func (player *Session) InitBaseStats() error {
	state := player.GetAuras()
	state.BaseStats = make(map[models.ModStat]float64)

	race := player.GetCharacterRace()
	class := player.GetCharacterClass()
	level := player.GetCharacterLevel()

	var foundClassLevel, foundRaceClass bool
	var maxLevel uint32

	// First, query to get our class base stats.
	player.DB().Range(func(cbs *models.ClassLevelStats) {
		if cbs.Class == class && cbs.Level == level {
			for k, v := range cbs.BaseStats {
				state.BaseStats[k] = v
			}
			foundClassLevel = true
		}

		if cbs.Level > maxLevel {
			maxLevel = cbs.Level
		}
	})

	// We did not find an exact match for our level-class combo.
	// We can still infer information based on what data we have
	if !foundClassLevel {
		player.DB().Range(func(cbs *models.ClassLevelStats) {
			if cbs.Class == class && cbs.Level == maxLevel {
				// If your level is over the max limit, we can still modify stats with non-confirmed values
				if level > maxLevel {
					// levelsAfter := level - maxLevel

					for k, v := range cbs.BaseStats {
						if k <= models.WalkSpeed || k >= models.MountSpeed {
							stat := v
							avgStatIncrement := stat / float64(cbs.Level)
							prev := state.BaseStats[k]
							if v > prev {
								state.BaseStats[k] = avgStatIncrement * float64(level)
							}
						}
					}
					foundClassLevel = true
				}
			}
		})
	}

	if !foundClassLevel {
		return fmt.Errorf("realm: fatal error InitBaseStats: can't find data for class %d", class)
	}

	// Search for exact stat info
	player.DB().Range(func(rcls *models.RaceClassLevelStats) {
		if rcls.Race == race && rcls.Class == class && rcls.Level == level {
			for k, v := range rcls.BaseStats {
				state.BaseStats[k] = v
			}
			foundRaceClass = true
		}
	})

	// Failing that, just look for appropriate stat info for any member of the class
	if !foundRaceClass {
		player.DB().Range(func(rcls *models.RaceClassLevelStats) {
			if rcls.Class == class && rcls.Level == level {
				for k, v := range rcls.BaseStats {
					prev := state.BaseStats[k]
					if v > prev {
						state.BaseStats[k] = v
					}
				}
				foundRaceClass = true
			}
		})
	}

	// Still nothing for this level. Okay, look for max level and extrapolate what should happen.
	if !foundRaceClass {
		player.DB().Range(func(rcls *models.RaceClassLevelStats) {
			if rcls.Class == class && rcls.Level == maxLevel {
				// If your level is over the max limit, we can still modify stats with non-confirmed values
				if level > maxLevel {
					// levelsAfter := level - maxLevel

					for k, v := range rcls.BaseStats {
						if k <= models.WalkSpeed || k >= models.MountSpeed {
							stat := v
							avgStatIncrement := stat / float64(rcls.Level)
							state.BaseStats[k] = avgStatIncrement * float64(level)
						}
					}

					foundRaceClass = true
				}
			}
		})
	}

	if !foundRaceClass {
		return fmt.Errorf("realm: fatal error InitBaseStats: can't find raceclasslevel data for class %d", class)
	}

	return nil
}

// func (m *Map) applyStatValue(unit Unit, ms models.ModStat, value float64) {
// 	values := unit.Values()

// 	switch ms {
// 	case models.BaseHealth:

// 	case models.BaseMana:
// 	case models.Stamina:
// 	case models.Strength:
// 	case models.Agility:
// 	case models.Intellect:
// 	case models.CriticalStrike:
// 	case models.PhysicalDamage:
// 	case models.HolyDamage:
// 	case models.FireDamage:
// 	case models.NatureDamage:
// 	case models.FrostDamage:
// 	case models.ShadowDamage:
// 	case models.ArcaneDamage:
// 	case models.ResistPhysicalDamage:
// 	case models.ResistHolyDamage:
// 	case models.ResistFireDamage:
// 	case models.ResistNatureDamage:
// 	case models.ResistFrostDamage:
// 	case models.ResistShadowDamage:
// 	case models.ResistArcaneDamage:
// 	// case models.NumResistsDamage:
// 	default:
// 		panic(ms)
// 	}
// }

func (s *Server) setResistanceField(unit Unit, resistance models.Resistance, value uint32) {
	index := models.ResistIndex(s.Build(), resistance)

	if index == -1 {
		return
	}

	values := unit.Values()

	resistances := values.Get("Resistances")
	resistances.Index(index).SetUint32(value)
}

func (s *Server) GetUnitPowerType(unit Unit) models.Power {
	return models.Power(unit.Values().Get("Power").Byte())
}

// Take a Unit's base stats and apply mods to them, value and movement blocks accordingly
func (s *Server) ApplyStats(unit Unit) error {
	state := unit.GetAuras()

	appliedStats := make(map[models.ModStat]float64)

	for st, v := range state.BaseStats {
		appliedStats[st] = v
	}

	// speeds := unit.Movement().Speeds

	for _, au := range state.Auras {
		for _, mod := range au.ModEffects {
			for ms := models.ModStat(0); ms < models.NumModStats; ms++ {
				if mod.StatMask.Has(ms) {

					switch mod.Op {
					case models.ModPlus:
						appliedStats[ms] = appliedStats[ms] + mod.Value
					case models.ModMult:
						appliedStats[ms] = appliedStats[ms] * mod.Value
					}
				}
			}
		}
	}

	// Start application of value updates
	values := unit.Values()

	baseHealth := values.Get("BaseHealth")
	if baseHealth != nil {
		baseHealth.SetUint32(uint32(appliedStats[models.BaseHealth]))
	}

	switch s.GetUnitPowerType(unit) {
	case models.PowerMana:
		baseMana := values.Get("BaseMana")
		if baseMana != nil {
			baseMana.SetUint32(uint32(appliedStats[models.BaseMana]))
		}

		// Compute mana based on Intellect
		maxMana := uint32(appliedStats[models.BaseMana])
		manaBoost := uint32(appliedStats[models.Intellect]) * 15
		maxMana += manaBoost
	case models.PowerRage:
		maxRage := values.Get("MaxRage")
		maxRage.SetUint32(100)
	case models.PowerFocus:
		maxFocus := values.Get("MaxFocus")
		maxFocus.SetUint32(100)
	case models.PowerEnergy:
		maxEnergy := values.Get("MaxEnergy")
		maxEnergy.SetUint32(100)
	}

	// TODO: configure a better formula

	// Compute health based on Stamina
	maxHealth := uint32(appliedStats[models.BaseHealth])
	staminaBoost := uint32(appliedStats[models.Stamina] * 10)
	maxHealth += staminaBoost

	health := values.Get("Health")

	mxHealth := values.Get("MaxHealth")
	if mxHealth != nil {
		mxHealth.SetUint32(maxHealth)
	}

	// Set resistances
	for i := models.Resistance(0); i < models.NumResists; i++ {
		var st float64
		switch i {
		case models.ResistPhysical:
			st = appliedStats[models.ResistPhysicalDamage]
		case models.ResistHoly:
			st = appliedStats[models.ResistHolyDamage]
		case models.ResistFire:
			st = appliedStats[models.ResistFireDamage]
		case models.ResistNature:
			st = appliedStats[models.ResistNatureDamage]
		case models.ResistFrost:
			st = appliedStats[models.ResistFrostDamage]
		case models.ResistShadow:
			st = appliedStats[models.ResistShadowDamage]
		case models.ResistArcane:
			st = appliedStats[models.ResistArcaneDamage]
		}
		s.setResistanceField(unit, i, uint32(st))
	}

	// Set modifying stats
	clientstats := values.Get("Stats")
	if clientstats != nil {
		for k, v := range appliedStats {
			if k.IsClientStat() {
				clientstat := clientstats.Index(int(k.ClientStat()))
				clientstat.SetUint32(uint32(v))
			}
		}
	}

	if health != nil {
		if health.Uint32() > maxHealth {
			health.SetUint32(maxHealth)
		}
	}

	state.AppliedStats = appliedStats

	return nil
}

func (m *Map) RefreshStats(unit Unit) error {
	if err := m.Phase.Server.ApplyStats(unit); err != nil {
		return err
	}

	m.PropagateObjectChanges(unit)

	return nil
}

func (s *Session) SavePowers() {
	s.DB().Cols("health").Update(s.Char)
}

func (s *Session) SetHealth(u uint32) {
	s.Get("Health").SetUint32(u)
	s.Char.Health = u
}
