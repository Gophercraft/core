package realm

func (s *Session) Health() uint32 {
	return s.Get("Health").Uint32()
}

func (s *Session) MaxHealth() uint32 {
	return s.Get("MaxHealth").Uint32()
}

func (s *Session) GetPowerType() uint8 {
	return s.Get("Power").Byte()
}

const (
	Mana = iota
	Rage
	Focus
	Energy
	Happiness
)

func (s *Session) Power() uint32 {
	switch s.GetPowerType() {
	case Mana:
		return s.Get("Mana").Uint32()
	case Rage:
		return s.Get("Rage").Uint32()
	case Focus:
		return s.Get("Focus").Uint32()
	case Energy:
		return s.Get("Energy").Uint32()
	}

	panic(s.GetPowerType())
}

func (s *Session) MaxPower() uint32 {
	switch s.GetPowerType() {
	case Mana:
		return s.Get("MaxMana").Uint32()
	case Rage:
		return s.Get("MaxRage").Uint32()
	case Focus:
		return s.Get("MaxFocus").Uint32()
	case Energy:
		return s.Get("MaxEnergy").Uint32()
	}

	panic(s.GetPowerType())
}

func (s *Session) VehicleSeatID() uint32 {
	return 0
}

func (s *Session) Pet() *Creature {
	return nil
}
