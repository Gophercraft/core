package models

import (
	"fmt"
	"strconv"

	"github.com/Gophercraft/core/tempest"
)

// PlayerCreateInfo determines where a character is spawned at upon their first login, using their race and class.
type PlayerCreateInfo struct {
	Race      RaceMask
	Class     ClassMask
	Placement PlayerCreatePlacement
}

type PlayerCreatePlacement struct {
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
	// Item will be created only if Race AND Class match
	Race         RaceMask
	Class        ClassMask
	BodySpecific bool         // If true, item will only be added if BodyType has character's body type
	BodyType     BodyTypeMask // Note: Ignored if Equip type is not EquipPaperDoll
	Item         string
	Amount       uint32
}

type PlayerCreateAbilities struct {
	Race   RaceMask
	Class  ClassMask
	Spells []uint32
}

type PlayerCreateActionButtons struct {
	Race          RaceMask
	Class         ClassMask
	ActionButtons []PlayerCreateActionButton
}

type ActionPlacement int

const (
	// Place action in first free slot (must be used after cardinal placements)
	ActionPlaceLeft ActionPlacement = -1
	// Place action in first free slot counting backward from 12
	ActionPlaceRight ActionPlacement = -2
)

func (ap *ActionPlacement) EncodeWord() (string, error) {
	switch *ap {
	case ActionPlaceLeft:
		return "Left", nil
	case ActionPlaceRight:
		return "Right", nil
	default:
		return strconv.Itoa(int(*ap)), nil
	}
}

func (ap *ActionPlacement) DecodeWord(str string) (err error) {
	switch str {
	case "Left":
		*ap = ActionPlaceLeft
		return
	case "Right":
		*ap = ActionPlaceRight
		return
	default:
		var i int
		i, err = strconv.Atoi(str)
		if err != nil {
			return
		}
		*ap = ActionPlacement(i)
		return
	}
}

type PlayerCreateActionButton struct {
	Type   ActionType
	Place  ActionPlacement
	Action uint32
}
