package models

import "time"

// Character describes a Player/Session's character attributes.
type Character struct {
	ID              uint64 `json:"id" xorm:"'id' pk autoincr"`
	GameAccount     uint64 `json:"gameAccount"`
	Name            string `json:"name"`
	Faction         uint32 `json:"faction"`
	FirstLogin      bool
	Level           uint32 `json:"level"`
	XP              uint32
	RealmID         uint64  `json:"realmID" xorm:"'realm_id'"`
	Race            Race    `json:"race"`
	Class           Class   `json:"class"`
	BodyType        uint8   `json:"gender"`
	Skin            uint8   `json:"skin"`
	Face            uint8   `json:"face"`
	HairStyle       uint8   `json:"hairStyle"`
	HairColor       uint8   `json:"hairColor"`
	FacialHair      uint8   `json:"facialHair"`
	Coinage         Money   `json:"coinage"`
	Zone            uint32  `json:"zone"`
	Map             uint32  `json:"map"`
	X               float32 `json:"x"`
	Y               float32 `json:"y"`
	Z               float32 `json:"z"`
	O               float32 `json:"o"`
	Leader          uint64
	Guild           uint64
	HideHelm        bool
	HideCloak       bool
	Health          uint32
	Mana            uint32
	Ghost           bool
	TotalPlayedTime time.Duration `xorm:"'playtime_total'"`
	LevelPlayedTime time.Duration `xorm:"'playtime_level'"`
}

type PropID string

type CharacterProp struct {
	ID     uint64
	PropID PropID
	Value  string
}

type CharacterAchievement struct {
	ID            uint64
	AchievementID uint32
}

// Item describes a *spawned* item. For the item's constant attributes, refer to ItemTemplate.
type Item struct {
	ID           uint64        `xorm:"'id' pk autoincr"`
	Creator      uint64        `xorm:"'creator'"` // player UID
	ItemType     InventoryType `xorm:"'item_type'"`
	ItemID       string        `xorm:"'item_id'"`
	DisplayID    uint32        `xorm:"'display_id'"`
	StackCount   uint32        `xorm:"'stack_count'"`
	Enchantments []uint32
	Charges      []int32 `xorm:"'charges'"`
}

// Inventory describes the positions of items/item stacks in a player's inventory.
type Inventory struct {
	ItemID uint64   `xorm:"'item_id' pk"`
	Player uint64   `xorm:"'player' index"`
	Bag    ItemSlot `xorm:"'bag'"`
	Slot   ItemSlot `xorm:"'slot'"`
}

// Contact describes the friend and ignore statuses of a Player in relation to another Player.
type Contact struct {
	Player   uint64 `xorm:"'player' index"`
	Friend   uint64 `xorm:"'friend'"`
	Friended bool   `xorm:"'friended'"`
	Ignored  bool   `xorm:"'ignored'"`
	Muted    bool   `xorm:"'muted'"`
	Note     string `xorm:"'note'"`
}

// LearnedAbility lists all the abilities/spells a player has learned.
type LearnedAbility struct {
	Player uint64 `xorm:"'player' index"`
	Spell  uint32 `xorm:"'spell'"`
	Active bool   `xorm:"'active'"`
	Slot   int    `xorm:"'slot'"`
}

// ExploredZone lists a player, and the zones which that player has explored in their map.
type ExploredZone struct {
	Player uint64 `xorm:"'player' index"`
	ZoneID uint32 `xorm:"'zone_id'"` // The actual zone ID, not the flag.
}
