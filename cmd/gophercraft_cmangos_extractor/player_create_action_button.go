package main

import "github.com/Gophercraft/core/realm/wdb/models"

type PlayerCreateActionButton struct {
	Race   uint8
	Class  uint8
	Button uint8
	Action uint32
	Type   uint8
}

func (PlayerCreateActionButton) TableName() string {
	return "Playercreateinfo_action"
}

func extractPlayerCreateActionButtons() {
	// Disabled due to containing incorrect results.
	var pca []PlayerCreateActionButton
	err := DB.Find(&pca)
	if err != nil {
		panic(err)
	}
	fl := openFile("DB/PlayerCreateActionButton.txt")
	printTimestamp(fl)
	wr := openTextWriter(fl)
	for _, pcab := range pca {
		if err := wr.Encode(models.PlayerCreateActionButton{
			Race:   models.Race(pcab.Race),
			Class:  models.Class(pcab.Class),
			Button: pcab.Button,
			Action: pcab.Action,
			Type:   pcab.Type,
		}); err != nil {
			panic(err)
		}
	}

	fl.Close()
}
