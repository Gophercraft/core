package main

import (
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/text"
)

// Base stats that apply to all members of a class.
// May be overwritten with the less general PlayerLevelStats
type PlayerClassLevelStats struct {
	Class      uint8  `xorm:"'class'"`
	Level      uint8  `xorm:"'level'"`
	BaseHealth uint32 `xorm:"'basehp'"`
	BaseMana   uint32 `xorm:"'basemana'"`
}

func (PlayerClassLevelStats) TableName() string {
	return "player_classlevelstats"
}

type PlayerLevelStats struct {
	Class     uint8  `xorm:"'class'"`
	Race      uint8  `xorm:"'race'"`
	Level     uint8  `xorm:"'level'"`
	Strength  uint32 `xorm:"'str'"`
	Agility   uint32 `xorm:"'agi'"`
	Stamina   uint32 `xorm:"'sta'"`
	Intellect uint32 `xorm:"'inte'"`
	Spirit    uint32 `xorm:"'spi'"`
}

func (PlayerLevelStats) TableName() string {
	return "player_levelstats"
}

func extractPlayerClassLevelStats() {
	fl := openFile("DB/ClassLevelStats.txt")
	printTimestamp(fl)
	wr := text.NewEncoder(fl)

	var pcls []PlayerClassLevelStats
	if err := DB.Find(&pcls); err != nil {
		panic(err)
	}

	for _, stats := range pcls {
		var pstats models.ClassLevelStats
		pstats.Class = models.Class(stats.Class)
		// pstats.Race = models.Race(stats.Race)
		pstats.Level = uint32(stats.Level)

		pstats.BaseStats = map[models.ModStat]float64{
			models.BaseHealth: float64(stats.BaseHealth),
			models.BaseMana:   float64(stats.BaseMana),
		}

		wr.Encode(&pstats)
	}

	fl.Close()
}

func extractPlayerLevelStats() {
	fl := openFile("DB/RaceClassLevelStats.txt")
	printTimestamp(fl)
	wr := text.NewEncoder(fl)

	var pcls []PlayerLevelStats
	if err := DB.Find(&pcls); err != nil {
		panic(err)
	}

	for _, stats := range pcls {
		var pstats models.RaceClassLevelStats
		pstats.Class = models.Class(stats.Class)
		// pstats.Race = models.Race(stats.Race)
		pstats.Level = uint32(stats.Level)

		pstats.BaseStats = map[models.ModStat]float64{
			models.Strength:  float64(stats.Strength),
			models.Agility:   float64(stats.Agility),
			models.Stamina:   float64(stats.Stamina),
			models.Intellect: float64(stats.Intellect),
			models.Spirit:    float64(stats.Spirit),
		}

		wr.Encode(&pstats)
	}

	fl.Close()
}
