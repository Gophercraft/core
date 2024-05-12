package main

type PlayerCreateInfoSpell struct {
	Race  uint8
	Class uint8
	Spell uint32
	Note  string
}

func (PlayerCreateInfoSpell) TableName() string {
	return "Playercreateinfo_spell"
}
