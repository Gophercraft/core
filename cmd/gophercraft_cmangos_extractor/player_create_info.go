package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/log"
)

type PlayerCreateInfo struct {
	Race  uint8   `xorm:"'race'"`
	Class uint8   `xorm:"'class'"`
	Map   uint32  `xorm:"'map'"`
	Zone  uint32  `xorm:"'zone'"`
	X     float32 `xorm:"'position_x'"`
	Y     float32 `xorm:"'position_y'"`
	Z     float32 `xorm:"'position_z'"`
	O     float32 `xorm:"'orientation'"`
}

func (pci PlayerCreateInfo) TableName() string {
	return "Playercreateinfo"
}

func extractPlayerCreateInfoRaceClass(race models.Race, class models.Class) {
	var pCreateInfo PlayerCreateInfo
	found, err := DB.Where("race = ?", race).Where("class = ?", class).Get(&pCreateInfo)
	if err != nil {
		log.Warn(err)
		return
	}

	if !found {
		log.Warn("Could not find race class combo for", raceName[race], className[class])
		return
	}

	createInfo := new(models.PlayerCreateInfo)
	createInfo.Race = race
	createInfo.Class = class
	createInfo.Placement.Map = pCreateInfo.Map
	createInfo.Placement.Zone = pCreateInfo.Zone
	createInfo.Placement.Position = tempest.C4Vector{
		pCreateInfo.X,
		pCreateInfo.Y,
		pCreateInfo.Z,
		pCreateInfo.O,
	}

	// Get items
	// var pCreateItems []PlayerCreateItem
	// err = DB.Where("race = ?", race).Where("class = ?", class).Find(&pCreateItems)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, pCreateItem := range pCreateItems {
	// 	itemTemplate, ok := allItems[fmt.Sprintf("it:%d", pCreateItem.ItemID)]
	// 	if !ok {
	// 		panic(pCreateItem.ItemID)
	// 	}

	// 	var item models.PlayerCreateItem
	// 	switch itemTemplate.InventoryType {
	// 	case models.IT_Unequippable:
	// 		item.Equip = models.EquipInventory
	// 	case models.IT_Bag:
	// 		item.Equip = models.EquipContainer
	// 	default:
	// 		item.Equip = models.EquipPaperDoll
	// 	}

	// 	item.Item = itemTemplate.ID
	// 	item.Amount = pCreateItem.Amount

	// 	createInfo.Items = append(createInfo.Items, item)
	// }

	var pCreateSpells []PlayerCreateInfoSpell
	err = DB.Where("race = ?", race).Where("class = ?", class).Find(&pCreateSpells)
	if err != nil {
		panic(err)
	}

	for _, pCreateSpell := range pCreateSpells {
		createInfo.Abilities = append(createInfo.Abilities, models.PlayerCreateAbility{
			Spell: pCreateSpell.Spell,
			Note:  pCreateSpell.Note,
		})
	}

	var pCreateActionButtons []PlayerCreateActionButton
	err = DB.Where("race = ?", race).Where("class = ?", class).Find(&pCreateActionButtons)
	if err != nil {
		panic(err)
	}

	for _, pCreateActionButton := range pCreateActionButtons {
		actionbutton := models.PlayerCreateActionButton{
			Button: pCreateActionButton.Button,
			Action: pCreateActionButton.Action,
		}

		switch pCreateActionButton.Type {
		case 0:
			actionbutton.Type = models.ActionSpell
		case 64:
			actionbutton.Type = models.ActionMacro
		case 128:
			actionbutton.Type = models.ActionItem
		default:
			panic(pCreateActionButton.Type)
		}

		createInfo.ActionButtons = append(createInfo.ActionButtons, actionbutton)
	}

	file := openTextFile(fmt.Sprintf("DB/PlayerCreateInfo/%s_%s.txt", raceName[race], className[class]))
	if err := file.Encode(createInfo); err != nil {
		panic(err)
	}
	file.close()
}

func extractPlayerCreateInfo() {
	os.MkdirAll(filepath.Join("DB", "PlayerCreateInfo"), 0700)

	for _, race := range races {
		for _, class := range classes {
			extractPlayerCreateInfoRaceClass(race, class)
		}
	}

	// fl := openFile("DB/PlayerCreateInfo.txt")
	// wr := text.NewEncoder(fl)

	// var pCreateInfo []PlayerCreateInfo

	// err := DB.Find(&pCreateInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, pci := range pCreateInfo {
	// 	mod := &models.PlayerCreateInfo{
	// 		Race:  models.Race(pci.Race),
	// 		Class: models.Class(pci.Class),
	// 		Position: tempest.C4Vector{
	// 			pci.X,
	// 			pci.Y,
	// 			pci.Z,
	// 			pci.O,
	// 		},
	// 		Map:  pci.Map,
	// 		Zone: pci.Zone,
	// 	}

	// 	if err := wr.Encode(mod); err != nil {
	// 		panic(err)
	// 	}
	// }

	// fl.Close()
}
