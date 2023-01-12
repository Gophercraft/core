package realm

import (
	"errors"

	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb/models"
)

func (m *Map) UnitIsAlive(target Unit) bool {
	if target == nil {
		return false
	}

	values := target.Values()
	health := values.Get("Health")
	return health.Uint32() != 0
}

func (m *Map) UnitIsDead(target Unit) bool {
	return !m.UnitIsAlive(target)
}

// Resurrect only
func (m *Map) Resurrect(theDead Unit, health uint32) error {
	theDead.Values().Get("Health").SetUint32(health)
	m.PropagateObjectChanges(theDead)
	return nil
}

// Resurrect + heal
func (m *Map) Revive(target Unit) error {
	if !m.UnitIsAlive(target) {
		if err := m.Resurrect(target, 1); err != nil {
			return err
		}
	}

	return m.HealAll(target)
}

// Heals all wounds
func (m *Map) HealAll(target Unit) error {
	if m.UnitIsDead(target) {
		return errors.New("realm: cannot heal a dead target!")
	}

	values := target.Values()

	health := values.Get("Health")
	maxHealth := values.Get("MaxHealth")

	health.SetUint32(maxHealth.Uint32())

	powerType := models.Power(values.Get("Power").Byte())

	var (
		power    *update.Value
		maxPower *update.Value
	)

	switch powerType {
	case models.PowerMana:
		power = values.Get("Mana")
		maxPower = values.Get("MaxMana")
	case models.PowerRage:
		power = values.Get("Rage")
		maxPower = values.Get("MaxRage")
	case models.PowerFocus:
		power = values.Get("Energy")
		maxPower = values.Get("MaxEnergy")
	case models.PowerEnergy:
		power = values.Get("Energy")
		maxPower = values.Get("MaxEnergy")
		// case models.PowerHappiness:
	}

	if power != nil {
		power.SetUint32(maxPower.Uint32())
	}

	m.PropagateObjectChanges(target)
	return nil
}

// returns the amount of health recharged in a single heartbeat.
func (s *Session) rechargeHealth() uint32 {
	return 10
}

func (s *Session) rechargeHeartbeat() {
	isRecharging := !s.InCombat()

	if isRecharging {
		health := s.Get("Health")
		prevHealth := health.Uint32()
		maxHealth := s.Get("MaxHealth").Uint32()
		if prevHealth == maxHealth {
			return
		}

		nextHealth := prevHealth + s.rechargeHealth()
		if nextHealth > maxHealth {
			nextHealth = maxHealth
		}

		health.SetUint32(nextHealth)
		s.UpdatePlayer()
	}
}
