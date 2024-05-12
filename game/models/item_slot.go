package models

import (
	"fmt"

	"github.com/Gophercraft/core/version"
)

type ItemSlot int

const (
	Backpack ItemSlot = -1
	Bag1     ItemSlot = 0
	Bag2     ItemSlot = 1
	Bag3     ItemSlot = 2
	Bag4     ItemSlot = 3
)

const (
	PaperDoll_Ammo     ItemSlot = -1
	PaperDoll_Head     ItemSlot = 0
	PaperDoll_Neck     ItemSlot = 1
	PaperDoll_Shoulder ItemSlot = 2
	PaperDoll_Shirt    ItemSlot = 3
	PaperDoll_Chest    ItemSlot = 4
	PaperDoll_Waist    ItemSlot = 5
	PaperDoll_Legs     ItemSlot = 6
	PaperDoll_Feet     ItemSlot = 7
	PaperDoll_Wrist    ItemSlot = 8
	PaperDoll_Hands    ItemSlot = 9
	PaperDoll_Finger1  ItemSlot = 10
	PaperDoll_Finger2  ItemSlot = 11
	PaperDoll_Trinket1 ItemSlot = 12
	PaperDoll_Trinket2 ItemSlot = 13
	PaperDoll_Back     ItemSlot = 14
	PaperDoll_MainHand ItemSlot = 15
	PaperDoll_OffHand  ItemSlot = 16
	PaperDoll_Ranged   ItemSlot = 17
	PaperDoll_Tabard   ItemSlot = 18
	PaperDoll_Bag1     ItemSlot = 19
	PaperDoll_Bag2     ItemSlot = 20
	PaperDoll_Bag3     ItemSlot = 21
	PaperDoll_Bag4     ItemSlot = 22
)

func (slot ItemSlot) String() string {
	if slot == Backpack {
		return "Backpack"
	}
	return fmt.Sprintf("%d", slot)
}

func GetStartItemSlot(build version.Build) ItemSlot {
	return 23
}
