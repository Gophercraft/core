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

	wr := openTextFile("DB/PlayerCreateItem.txt")

	for _, pci := range pcis {
		var item models.PlayerCreateItem
		item.Equip = models.EquipInventory
		item.Race.Set(models.Race(pci.Race), true)
		item.Class.Set(models.Class(pci.Class), true)
		item.Item = fmt.Sprintf("it:%d", pci.ItemID)
		item.Amount = pci.Amount

		if err := wr.Encode(&item); err != nil {
			panic(err)
		}
	}

	wr.close()
}
