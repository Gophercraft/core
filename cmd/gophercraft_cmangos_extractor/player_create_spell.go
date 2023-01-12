package main

import "github.com/Gophercraft/core/realm/wdb/models"

type PlayerCreateInfoSpell struct {
	Race  uint8
	Class uint8
	Spell uint32
	Note  string
}

func (PlayerCreateInfoSpell) TableName() string {
	return "Playercreateinfo_spell"
}

func extractPlayerCreateSpells() {
	var pcs []PlayerCreateInfoSpell
	err := DB.Find(&pcs)
	if err != nil {
		panic(err)
	}
	fl := openFile("DB/PlayerCreateAbility.txt")
	printTimestamp(fl)
	wr := openTextWriter(fl)
	for _, pcab := range pcs {
		ab := models.PlayerCreateAbility{
			Spell: pcab.Spell,
			Note:  pcab.Note,
		}

		ab.Race.Set(models.Race(pcab.Race), true)
		ab.Class.Set(models.Class(pcab.Class), true)

		if err := wr.Encode(&ab); err != nil {
			panic(err)
		}
	}

	fl.Close()
}
