package models

// Determines what the base stats of a race are.
type RaceStats struct {
	Race Race
	// Level uint32
	BaseStats map[ModStat]float64
}

type ClassStats struct {
	Class Class
	// BaseStats        map[ModStat]float64
	//
	LevelBonus map[ModStat]float64
}

// Determines fixed values for a class member at a certain level
type BaseClassLevelStats struct {
	Class     Class
	Level     uint32
	BaseStats map[ModStat]float64
}
