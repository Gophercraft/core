package models

type FactionStanding struct {
	Player  uint64
	Faction uint32
	Value   float32
	AtWar   bool
}
