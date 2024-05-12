package main

import "github.com/Gophercraft/core/realm/wdb/models"

var races = []models.Race{
	1, 2, 3, 4,
	5, 6, 7, 8,
}

var classes = []models.Class{
	1,
	2,
	3,
	4,
	5,
	7,
	8,
	9,
	11,
}

var raceName = map[models.Race]string{
	1: "Human",
	2: "Orc",
	3: "Dwarf",
	4: "NightElf",
	5: "Undead",
	6: "Tauren",
	7: "Gnome",
	8: "Troll",
}

var className = map[models.Class]string{
	1:  "Warrior",
	2:  "Paladin",
	3:  "Hunter",
	4:  "Rogue",
	5:  "Priest",
	7:  "Shaman",
	8:  "Mage",
	9:  "Warlock",
	11: "Druid",
}
