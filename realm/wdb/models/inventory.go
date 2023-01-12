package models

import "fmt"

type InventoryType int8

const (
	// InventoryTypes
	IT_Empty        = -1
	IT_Unequippable = 0
	IT_Head         = 1
	IT_Neck         = 2
	IT_Shoulder     = 3
	IT_Shirt        = 4
	IT_Chest        = 5
	IT_Waist        = 6
	IT_Legs         = 7
	IT_Feet         = 8
	IT_Wrists       = 9
	IT_Hands        = 10
	IT_Finger       = 11
	IT_Trinket      = 12
	IT_Weapon       = 13
	IT_Shield       = 14
	IT_Ranged       = 15
	IT_Back         = 16
	IT_TwoHand      = 17
	IT_Bag          = 18
	IT_Tabard       = 19
	IT_Robe         = 20
	IT_MainHand     = 21
	IT_OffHand      = 22
	IT_Holdable     = 23
	IT_Ammo         = 24
	IT_Thrown       = 25
	IT_Gun          = 26
	IT_Quiver       = 27
	IT_Relic        = 28
)

var (
	InventoryType_names = map[InventoryType]string{
		// InventoryTypes
		IT_Empty:        "Empty",
		IT_Unequippable: "Unequippable",
		IT_Head:         "Head",
		IT_Neck:         "Neck",
		IT_Shoulder:     "Shoulder",
		IT_Shirt:        "Shirt",
		IT_Chest:        "Chest",
		IT_Waist:        "Waist",
		IT_Legs:         "Legs",
		IT_Feet:         "Feet",
		IT_Wrists:       "Wrists",
		IT_Hands:        "Hands",
		IT_Finger:       "Finger",
		IT_Trinket:      "Trinket",
		IT_Weapon:       "Weapon",
		IT_Shield:       "Shield",
		IT_Ranged:       "Ranged",
		IT_Back:         "Back",
		IT_TwoHand:      "TwoHand",
		IT_Bag:          "Bag",
		IT_Tabard:       "Tabard",
		IT_Robe:         "Robe",
		IT_MainHand:     "MainHand",
		IT_OffHand:      "OffHand",
		IT_Holdable:     "Holdable",
		IT_Ammo:         "Ammo",
		IT_Thrown:       "Thrown",
		IT_Gun:          "Gun",
		IT_Quiver:       "Quiver",
		IT_Relic:        "Relic",
	}

	InventoryType_lookup map[string]InventoryType

	ItemDisplaySlots = map[InventoryType]ItemSlot{
		IT_Shield:   PaperDoll_OffHand,
		IT_Robe:     PaperDoll_Chest,
		IT_Head:     PaperDoll_Head,
		IT_Neck:     PaperDoll_Neck,
		IT_Shoulder: PaperDoll_Shoulder,
		IT_Shirt:    PaperDoll_Shirt,
		IT_Chest:    PaperDoll_Chest,
		IT_Waist:    PaperDoll_Waist,
		IT_Legs:     PaperDoll_Legs,
		IT_Feet:     PaperDoll_Feet,
		IT_Wrists:   PaperDoll_Wrist,
		IT_Hands:    PaperDoll_Hands,
		IT_Finger:   PaperDoll_Finger1,
		IT_Trinket:  PaperDoll_Trinket1,
		IT_Back:     PaperDoll_Back,
		IT_TwoHand:  PaperDoll_MainHand,
		IT_MainHand: PaperDoll_MainHand,
		IT_Holdable: PaperDoll_OffHand,
		IT_OffHand:  PaperDoll_OffHand,
		IT_Ranged:   PaperDoll_Ranged,
		IT_Gun:      PaperDoll_Ranged,
		IT_Tabard:   PaperDoll_Tabard,
		IT_Quiver:   PaperDoll_Bag1,
		IT_Bag:      PaperDoll_Bag1,
		IT_Thrown:   PaperDoll_Ranged,
		IT_Ammo:     PaperDoll_Ammo,
	}
)

func init() {
	InventoryType_lookup = make(map[string]InventoryType, len(InventoryType_names))

	for k, v := range InventoryType_names {
		InventoryType_lookup[v] = k
	}
}

func (it InventoryType) String() string {
	str, err := it.EncodeWord()
	if err != nil {
		return fmt.Sprintf("InventoryType(%d)", it)
	}
	return str
}

func (it *InventoryType) EncodeWord() (string, error) {
	dat, ok := InventoryType_names[*it]
	if !ok {
		return "", fmt.Errorf("%d", *it)
	}
	return dat, nil
}

func (it *InventoryType) DecodeWord(dat string) (err error) {
	var ok bool
	*it, ok = InventoryType_lookup[dat]
	if !ok {
		err = fmt.Errorf("%s", dat)
	}
	return
}

func (it *InventoryType) PaperDollSlot() (ItemSlot, error) {
	slot, ok := ItemDisplaySlots[*it]
	if !ok {
		return 0, fmt.Errorf("models: cannot find paper doll slot for %s", it)
	}

	return slot, nil
}
