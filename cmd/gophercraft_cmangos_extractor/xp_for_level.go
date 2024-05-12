package main

import "github.com/Gophercraft/core/realm/wdb/models"

type ExperienceForLevel struct {
	Lvl uint32 `xorm:"'lvl'"`
	XP  uint32 `xorm:"'xp_for_next_level'"`
}

func (exp ExperienceForLevel) TableName() string {
	return "player_xp_for_level"
}

func extractExpForLevel() {
	var levelXP []ExperienceForLevel

	err := DB.Find(&levelXP)
	if err != nil {
		panic(err)
	}

	wr := openTextFile("DB/LevelExperience.txt")
	exp := models.LevelExperience{}

	for _, lexp := range levelXP {
		exp[lexp.Lvl] = lexp.XP
	}

	if err := wr.Encode(exp); err != nil {
		panic(err)
	}

	wr.close()
}
