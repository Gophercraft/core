package realm

import (
	"fmt"
	"sync"
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

type Aura struct {
	ID           uint32
	Flag         update.AuraFlags
	Level        uint32
	Applications uint32
	Caster       guid.GUID
	ModEffects   []*ModEffect
	AppliedTime  time.Time
	ExpiryTime   time.Time
	ExpiryTimer  *time.Timer
}

func (au *Aura) Duration() time.Duration {
	return au.ExpiryTime.Sub(au.AppliedTime)
}

// Aura state holds both base and real stats.
// From this data, the update fields are set.
// Update fields are not used as a direct storage for these stats.
type AuraState struct {
	Protect      sync.Mutex
	BaseStats    map[models.ModStat]float64
	AppliedStats map[models.ModStat]float64
	Auras        []*Aura
	AuraExpire   chan *Aura
}

func (m *Map) SyncAuraState(unit Unit) error {
	state := unit.GetAuras()

	build := m.Phase.Server.Build()

	if build.RemovedIn(vsn.V3_0_2) {
		values := unit.Values()

		var (
			auras            *update.Value
			auraSize         int
			auraLevels       *update.Value
			hasLevels        bool
			auraFlags        *update.Value
			auraApplications *update.Value
			hasApplications  bool
		)

		hasLevels = build >= vsn.V1_12_1
		hasApplications = build >= vsn.V1_12_1

		auras = values.Get("Auras")
		auraSize = auras.Len()
		auraFlags = values.Get("AuraFlags")

		if hasLevels {
			auraLevels = values.Get("AuraLevels")
		}

		if hasApplications {
			auraApplications = values.Get("AuraApplications")
		}

		for x := 0; x < auraSize; x++ {
			stateHasIndex := x < len(state.Auras)
			aura := auras.Index(x)

			if stateHasIndex {
				au := state.Auras[x]
				if aura.Uint32() != au.ID {
					auras.Index(x).SetUint32(au.ID)
				}
				if hasLevels {
					auraLevel := auraLevels.Index(x)
					if auraLevel.Byte() != byte(au.Level) {
						auraLevel.SetByte(byte(au.Level))
					}
				}
				if hasApplications {
					auraApplications.Index(x).SetByte(byte(au.Applications) - 1)
				}
				update.SetAuraFlag(auraFlags, x, au.Flag)
			} else {
				if aura.Uint32() != 0 {
					aura.SetUint32(0)
				}
				if hasLevels {
					auraLevel := auraLevels.Index(x)
					if auraLevel.Byte() != 0 {
						auraLevel.SetByte(0)
					}
				}
				if hasApplications {
					auraApplications.Index(x).SetByte(0)
				}
				if err := update.SetAuraFlag(auraFlags, x, 0); err != nil {
					return err
				}
			}
		}

		m.PropagateObjectChanges(unit)

		// m.VisibleObjects(unit).Sessions().Iter(func(s *Session) {
		// 	for x := 0; x < auraSize; x++ {

		// })
	} else {
		return fmt.Errorf("realm: cannot sync aura state in this version yet :(")
	}

	return nil
}

// func (m *Map) ApplyAura(au *Aura, target WorldObject) error {

// 	for a := 0; a < auraSize; a++ {
// 		aura := auras.Index(a)
// 		application := auraApplications.Index(a)
// 		level := auraLevels.Index(a)
// 		if aura.Uint32() == au.ID {
// 			room := 255 - application.Byte()
// 			if room <= au.Applications { // a snug fit
// 				application.SetByte(application.Byte() + au.Applications)
// 			} else { // We have reached the max number of applications
// 				application.SetByte(255)
// 			}
// 			level.SetByte(au.Level)
// 			break
// 		}

// 		if aura.Uint32() == 0 {
// 			aura.SetUint32(au.ID)
// 			level.SetByte(au.Level)
// 			application.SetByte(au.Applications)
// 			break
// 		}
// 	}

// 	m.PropagateChanges(target.GUID())

// 	_ = auraSize
// 	_ = auraFlags
// 	_ = auraLevels

// 	return nil
// }

// removes a finite or infinite amount of aura applications.
// if you are interested, show visuals for the
func (m *Map) Unaura(target Unit, id uint32, count int32, visuals bool) error {
	state := target.GetAuras()
	state.Protect.Lock()

	defer func() {
		if visuals {
			m.PlaySpellVisualKind(target, id, models.VisualStateDone)
		}
	}()

	var newAuras []*Aura

	if count < 0 {
		for _, au := range state.Auras {
			if au.ID != id {
				newAuras = append(newAuras, au)
			} else {
				if au.ExpiryTimer != nil {
					au.ExpiryTimer.Stop()
				}
			}
		}

		state.Auras = newAuras
		state.Protect.Unlock()

		return m.SyncAuraState(target)
	}

	remaining := uint32(count)

	for i := 0; i < len(state.Auras); i++ {
		if remaining == 0 {
			break
		}

		au := state.Auras[i]

		pAppsRemoved := au.Applications
		if pAppsRemoved > remaining {
			pAppsRemoved = remaining
		}

		if au.Applications > 0 {
			au.Applications -= pAppsRemoved
			remaining -= pAppsRemoved
		}

		if au.Applications != 0 {
			newAuras = append(newAuras, au)
		} else {
			if au.ExpiryTimer != nil {
				au.ExpiryTimer.Stop()
				au.ExpiryTimer = nil
			}
		}
	}

	state.Auras = newAuras
	state.Protect.Unlock()

	return m.SyncAuraState(target)
}

type AuraApplication struct {
	CastTime     time.Time
	ID           uint32
	Level        int32
	Caster       Unit
	Applications uint32
}

func (m *Map) ApplyAura(target Unit, aa *AuraApplication) error {
	var (
		spellID      uint32 = aa.ID
		spellLevel   uint32 = uint32(aa.Level)
		applications        = aa.Applications
		castTime            = aa.CastTime
		durationMs   int32  = 0
	)

	if applications == 0 {
		applications = 1
	}

	if castTime.IsZero() {
		castTime = time.Now()
	}

	s := m.Phase.Server

	var spellDef *dbdefs.Ent_Spell

	s.DB.Lookup(wdb.BucketKeyUint32ID, spellID, &spellDef)
	if spellDef == nil {
		return fmt.Errorf("realm: cannot apply aura that does not exist: %d", spellID)
	}

	var spellDuration *dbdefs.Ent_SpellDuration
	s.DB.Lookup(wdb.BucketKeyUint32ID, uint32(spellDef.DurationIndex), &spellDuration)
	if spellDuration != nil {
		durationMs = spellDuration.Duration
		leveledDuration := int32(spellLevel) * spellDuration.DurationPerLevel
		durationMs += leveledDuration
		if durationMs > spellDuration.MaxDuration {
			durationMs = spellDuration.MaxDuration
		}
	}

	if durationMs == 0 {
		log.Warn("Spell cast:", aa.ID, "duration 0")
		return nil
	}

	state := target.GetAuras()

	duration := time.Duration(durationMs) * time.Millisecond
	expirationDate := time.Now().Add(duration)

	var (
		index  int
		insert bool = true
		au     *Aura
	)

	for i := range state.Auras {
		au = state.Auras[i]
		if au.ID == spellID {
			au.Applications += applications
			au.Level = uint32(spellLevel)

			if au.ExpiryTimer != nil {
				au.ExpiryTimer.Reset(duration)
			}

			au.AppliedTime = castTime
			au.ExpiryTime = expirationDate
			insert = false
			index = i
			break
		}
	}

	if insert {
		au = new(Aura)
		au.ID = spellID
		au.Flag = update.AuraMaskAll
		au.Level = spellLevel
		au.Applications = applications
		au.AppliedTime = castTime
		au.ExpiryTime = expirationDate

		if state.AuraExpire != nil {
			au.ExpiryTimer = time.AfterFunc(duration, func() {
				if state.AuraExpire != nil {
					state.AuraExpire <- au
				}
			})
		}
	}

	var sed SpellEffectData
	sed.Caster = aa.Caster
	sed.Target = target
	sed.Map = m
	sed.Spell = spellDef
	sed.Aura = au

	for i, effect := range sed.Spell.Effect {
		if effect == 0 {
			continue
		}

		handler := m.Phase.Server.SpellEffects[effect]
		if handler == nil {
			continue
		}

		sed.EffectIndex = i
		handler(&sed)
	}

	if insert {
		index = len(state.Auras)
		state.Auras = append(state.Auras, au)
	}

	if err := m.SyncAuraState(target); err != nil {
		return err
	}

	if err := m.RefreshStats(target); err != nil {
		return err
	}

	if sesh, ok := target.(*Session); ok {
		sesh.UpdateAuraDuration(index)
	}

	return nil
}
