package models

import (
	"fmt"

	"github.com/Gophercraft/core/tempest"
)

// PlayerCreateInfo determines where a character is spawned at upon their first login, using their race and class.
type PlayerCreateInfo struct {
	Race     Race
	Class    Class
	Position tempest.C4Vector
	Map      uint32
	Zone     uint32
}

type EquipType uint8

const (
	// Not encoded
	EquipNone EquipType = iota
	// Equip in the next free available inventory slot
	EquipInventory
	// Item is a bag. Put this bag in the next free bag slot.
	EquipContainer
	// Item is to be equipped in the paper doll based on its inventory type
	EquipPaperDoll
)

func (et *EquipType) EncodeWord() (s string, e error) {
	switch *et {
	case EquipInventory:
		s = "Inventory"
	case EquipContainer:
		s = "Container"
	case EquipPaperDoll:
		s = "PaperDoll"
	default:
		e = fmt.Errorf("models: unknown EquipType %d", *et)
	}

	return
}

func (et *EquipType) DecodeWord(data string) (e error) {
	switch data {
	case "Inventory":
		*et = EquipInventory
	case "Container":
		*et = EquipContainer
	case "PaperDoll":
		*et = EquipPaperDoll
	default:
		e = fmt.Errorf("models: unknown EquipType %s", data)
	}

	return
}

type PlayerCreateItem struct {
	Equip EquipType
	// Slot   ItemSlot
	Race   RaceMask
	Class  ClassMask
	Item   string
	Amount uint32
}

type PlayerCreateAbility struct {
	Race   RaceMask
	Class  ClassMask
	Spell  uint32
	Note   string
	Active bool
}

type PlayerCreateActionButton struct {
	Race   Race
	Class  Class
	Button uint8
	Action uint32
	Type   ActionType
	Misc   uint8
}
