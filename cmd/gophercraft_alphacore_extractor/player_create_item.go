package main

import (
	"fmt"

	"github.com/Gophercraft/core/realm/wdb/models"
)

type PlayerCreateItem struct {
	Race   uint8  `xorm:"'race'"`
	Class  uint8  `xorm:"'class'"`
	ItemID uint32 `xorm:"'itemid'"`
	Amount uint32 `xorm:"'amount'"`
}

func (PlayerCreateItem) TableName() string {
	return "Playercreateinfo_item"
}

func extractPlayerCreateItems() {
	var pcis []PlayerCreateItem
	err := DB.Find(&pcis)
	if err != nil {
		panic(err)
	}

	fl := openFile("DB/PlayerCreateItem.txt")
	printTimestamp(fl)

	wr := openTextWriter(fl)
	for _, pci := range pcis {
		var item models.PlayerCreateItem
		item.Item = fmt.Sprintf("it:%d", pci.ItemID)

		itt, ok := allItems[item.Item]
		if !ok {
			panic(item.Item)
		}

		wr.WriteComment(itt.Name.String())

		switch itt.InventoryType {
		case models.IT_Unequippable:
			item.Equip = models.EquipInventory
		case models.IT_Bag:
			item.Equip = models.EquipContainer
		default:
			item.Equip = models.EquipPaperDoll
		}

		item.Race.Set(models.Race(pci.Race), true)
		item.Class.Set(models.Class(pci.Class), true)
		item.Amount = pci.Amount
		if item.Amount == 0 {
			item.Amount = 1
		}

		if err := wr.Encode(&item); err != nil {
			panic(err)
		}
	}

	fl.Close()
}
