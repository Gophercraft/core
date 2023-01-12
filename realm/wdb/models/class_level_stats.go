package models

type ClassLevelStats struct {
	Class Class
	Level uint32

	BaseStats map[ModStat]float64
}

type RaceClassLevelStats struct {
	Race      Race
	Class     Class
	Level     uint32
	BaseStats map[ModStat]float64
}
