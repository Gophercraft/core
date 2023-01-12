package main

import (
	"github.com/Gophercraft/core/realm/wdb/models"
)

type PlayerCreateActionButton struct {
	ID     uint32 `xorm:"'id'"`
	Race   uint8
	Class  uint8
	Button uint8
	Action int32
	Type   uint8
}

func (PlayerCreateActionButton) TableName() string {
	return "Playercreateinfo_action"
}

const spellMask = 0xffffff

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
		var atype models.ActionType
		var action uint32
		actionData := pcab.Action

		if actionData < 0 {
			atype = models.ActionItem
			action = uint32(actionData * -1)
		} else {
			atype = models.ActionSpell
			action = uint32(actionData)
		}

		// actionNit := actionData & spellMask
		// actionType := (actionData >> 24) & 0xff

		// fmt.Println(actionNit, actionType)

		if err := wr.Encode(models.PlayerCreateActionButton{
			Race:   models.Race(pcab.Race),
			Class:  models.Class(pcab.Class),
			Button: pcab.Button,
			Action: action,
			Type:   atype,
		}); err != nil {
			panic(err)
		}
	}

	fl.Close()
}
