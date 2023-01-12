package models

import (
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet/update"
)

type ItemBind uint8

const (
	ItemUnbound ItemBind = iota
	ItemBindOnPickup
	ItemBindOnEquip
	ItemBindOnUse
	ItemQuestItem
	ItemQuestItem2
)

type ItemDamage struct {
	Type uint8
	Min  float32
	Max  float32
}

type ItemSpell struct {
	ID               uint32
	Trigger          uint32
	Charges          int32
	PPMRate          float32
	Cooldown         int64
	Category         uint32
	CategoryCooldown int64
}

type ItemTemplate struct {
	Entry                     uint32    `json:",omitempty" xorm:"'entry' bigint pk" csv:"-"`
	ID                        string    `json:",omitempty" xorm:"'id' index"`
	Name                      i18n.Text `json:",omitempty" xorm:"'name' index"`
	Class                     uint32    `json:",omitempty" xorm:"'class'"`
	Subclass                  uint32    `json:",omitempty" xorm:"'subclass'"`
	SoundOverrideSubclass     uint32
	DisplayID                 uint32           `json:",omitempty" xorm:"'PaperDoll_id'"`
	Quality                   ItemQuality      `json:",omitempty" xorm:"'quality'"`
	Flags                     update.ItemFlags `json:",omitempty" xorm:"'flags'"`
	BuyCount                  uint8            `json:",omitempty" xorm:"'buy_count'"`
	BuyPrice                  Money            `json:",omitempty" xorm:"'buy_price'"`
	SellPrice                 Money            `json:",omitempty" xorm:"'sell_price'"`
	InventoryType             InventoryType    `json:",omitempty" xorm:"'inv_type'"`
	AllowableClass            ClassMask        `json:",omitempty" xorm:"'allowable_class'"`
	AllowableRace             RaceMask         `json:",omitempty" xorm:"'allowable_race'"`
	ItemLevel                 uint32           `json:",omitempty" xorm:"'item_level'"`
	RequiredLevel             uint8            `json:",omitempty" xorm:"'required_level'"`
	RequiredSkill             uint32           `json:",omitempty" xorm:"'required_skill'"`
	RequiredSkillRank         uint32           `json:",omitempty" xorm:"'required_skill_rank'"`
	RequiredSpell             uint32           `json:",omitempty" xorm:"'required_spell'"`
	RequiredHonorRank         uint32           `json:",omitempty" xorm:"'required_honor_rank'"`
	RequiredCityRank          uint32           `json:",omitempty" xorm:"'required_city_rank'"`
	RequiredReputationFaction uint32           `json:",omitempty" xorm:"'required_reputation_faction'"`
	RequiredReputationRank    uint32           `json:",omitempty" xorm:"'required_reputation_rank'"`
	MaxCount                  uint32           `json:",omitempty" xorm:"'max_count'"`
	Stackable                 uint32           `json:",omitempty" xorm:"'stackable'"`
	ContainerSlots            uint8            `json:",omitempty" xorm:"'container_slots'"`
	Stats                     []ItemStat       `json:",omitempty" xorm:"'stats'"`
	Damage                    []ItemDamage     `json:",omitempty" xorm:"'dmg'"`
	Armor                     uint32           `json:",omitempty" xorm:"'armor'"`
	HolyRes                   uint32           `json:",omitempty" xorm:"'holy_res'"`
	FireRes                   uint32           `json:",omitempty" xorm:"'fire_res'"`
	NatureRes                 uint32           `json:",omitempty" xorm:"'nature_res'"`
	FrostRes                  uint32           `json:",omitempty" xorm:"'frost_res'"`
	ShadowRes                 uint32           `json:",omitempty" xorm:"'shadow_res'"`
	ArcaneRes                 uint32           `json:",omitempty" xorm:"'arcane_res'"`
	Delay                     uint32           `json:",omitempty" xorm:"'delay'"`
	AmmoType                  uint32           `json:",omitempty" xorm:"'ammo_type'"`
	RangedModRange            float32          `json:",omitempty" xorm:"'ranged_mod_range'"`
	Spells                    []ItemSpell      `json:",omitempty" xorm:"'spells'"`
	Bonding                   ItemBind         `json:",omitempty" xorm:"'bonding'"`
	Description               i18n.Text        `json:",omitempty" xorm:"'description' longtext"`
	PageText                  uint32           `json:",omitempty" xorm:"'page_text"`
	LanguageID                uint32           `json:",omitempty" xorm:"'language_id'"`
	PageMaterial              uint32           `json:",omitempty" xorm:"'page_material'"`
	StartQuest                uint32           `json:",omitempty" xorm:"'start_quest'"`
	LockID                    uint32           `json:",omitempty" xorm:"'lock_id'"`
	Material                  int32            `json:",omitempty" xorm:"'material'"`
	Sheath                    uint32           `json:",omitempty" xorm:"'sheath'"`
	RandomProperty            uint32           `json:",omitempty" xorm:"'random_property'"`
	RandomSuffix              uint32           `json:",omitempty" xorm:"'random_suffix'"`
	Block                     uint32           `json:",omitempty" xorm:"'block'"`
	Itemset                   uint32           `json:",omitempty" xorm:"'itemset'"`
	MaxDurability             uint32           `json:",omitempty" xorm:"'max_durability'"`
	Area                      uint32           `json:",omitempty" xorm:"'area'"`
	Map                       int32            `json:",omitempty" xorm:"'map'"`
	BagFamily                 int32            `json:",omitempty" xorm:"'bag_family'"`
	TotemCategory             int32            `json:",omitempty" xorm:"'totem_category'"`
	Socket                    []ItemSocket     `json:",omitempty" xorm:"'sockets'"`
	SocketBonus               uint32
	GemProperties             int32   `json:",omitempty" xorm:"'gem_properties'"`
	RequiredDisenchantSkill   int32   `json:",omitempty" xorm:"'required_disenchant_skill'"`
	ArmorDamageModifier       float32 `json:",omitempty" xorm:"'armor_damage_modifier'"`
	ItemLimitCategory         uint32  `json:",omitempty" xorm:"'item_limit_category'"`
	HolidayID                 uint32
	ScriptName                string `json:",omitempty" xorm:"'script_name'"`
	DisenchantID              uint32 `json:",omitempty" xorm:"'disenchant_id'"`
	FoodType                  uint8  `json:",omitempty" xorm:"'food_type'"`
	MinMoneyLoot              Money  `json:",omitempty" xorm:"'min_money_loot'"`
	MaxMoneyLoot              Money  `json:",omitempty" xorm:"'max_money_loot'"`
	Duration                  int32  `json:",omitempty" xorm:"'duration'"`
	ExtraFlags                uint8  `json:",omitempty" xorm:"'extra_flags'"`
}
