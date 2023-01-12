package models

import (
	"fmt"

	"github.com/Gophercraft/core/vsn"
)

type ItemMod uint8

const (
	ItemModMana ItemMod = iota
	ItemModHealth
	ItemModAgility
	ItemModStrength
	ItemModIntellect
	ItemModSpirit
	ItemModStamina
	ItemModDefenseSkillRating
	ItemModDodgeRating
	ItemModParryRating
	ItemModBlockRating
	ItemModHitMeleeRating
	ItemModHitRangedRating
	ItemModHitSpellRating
	ItemModCritMeleeRating
	ItemModCritRangedRating
	ItemModCritSpellRating
	ItemModHitTakenMeleeRating
	ItemModHitTakenRangedRating
	ItemModHitTakenSpellRating
	ItemModCritTakenMeleeRating
	ItemModCritTakenRangedRating
	ItemModCritTakenSpellRating
	ItemModHasteMeleeRating
	ItemModHasteRangedRating
	ItemModHasteSpellRating
	ItemModHitRating
	ItemModCritRating
	ItemModHitTakenRating
	ItemModCritTakenRating
	ItemModResilienceRating
	ItemModHasteRating
	ItemModExpertiseRating
	ItemModAttackPower
	ItemModRangedAttackPower
	ItemModFeralAttackPower
	ItemModSpellHealingDone
	ItemModSpellDamageDone
	ItemModManaRegeneration
	ItemModArmorPenetrationRating
	ItemModSpellPower
	ItemModHealthRegen
	ItemModSpellPenetration
	ItemModBlockValue
)

var (
	ItemModNames = map[ItemMod]string{
		ItemModMana:                   "Mana",
		ItemModHealth:                 "Health",
		ItemModAgility:                "Agility",
		ItemModStrength:               "Strength",
		ItemModIntellect:              "Intellect",
		ItemModSpirit:                 "Spirit",
		ItemModStamina:                "Stamina",
		ItemModDefenseSkillRating:     "DefenseSkillRating",
		ItemModDodgeRating:            "DodgeRating",
		ItemModParryRating:            "ParryRating",
		ItemModBlockRating:            "BlockRating",
		ItemModHitMeleeRating:         "HitMeleeRating",
		ItemModHitRangedRating:        "HitRangedRating",
		ItemModHitSpellRating:         "HitSpellRating",
		ItemModCritMeleeRating:        "CritMeleeRating",
		ItemModCritRangedRating:       "CritRangedRating",
		ItemModCritSpellRating:        "CritSpellRating",
		ItemModHitTakenMeleeRating:    "HitTakenMeleeRating",
		ItemModHitTakenRangedRating:   "HitTakenRangedRating",
		ItemModHitTakenSpellRating:    "HitTakenSpellRating",
		ItemModCritTakenMeleeRating:   "CritTakenMeleeRating",
		ItemModCritTakenRangedRating:  "CritTakenRangedRating",
		ItemModCritTakenSpellRating:   "CritTakenSpellRating",
		ItemModHasteMeleeRating:       "HasteMeleeRating",
		ItemModHasteRangedRating:      "HasteRangedRating",
		ItemModHasteSpellRating:       "HasteSpellRating",
		ItemModHitRating:              "HitRating",
		ItemModCritRating:             "CritRating",
		ItemModHitTakenRating:         "HitTakenRating",
		ItemModCritTakenRating:        "CritTakenRating",
		ItemModResilienceRating:       "ResilienceRating",
		ItemModHasteRating:            "HasteRating",
		ItemModExpertiseRating:        "ExpertiseRating",
		ItemModAttackPower:            "AttackPower",
		ItemModRangedAttackPower:      "RangedAttackPower",
		ItemModFeralAttackPower:       "FeralAttackPower",
		ItemModSpellHealingDone:       "SpellHealingDone",
		ItemModSpellDamageDone:        "SpellDamageDone",
		ItemModManaRegeneration:       "ManaRegeneration",
		ItemModArmorPenetrationRating: "ArmorPenetrationRating",
		ItemModSpellPower:             "SpellPower",
		ItemModHealthRegen:            "HealthRegen",
		ItemModSpellPenetration:       "SpellPenetration",
		ItemModBlockValue:             "BlockValue",
	}

	ItemModLookup map[string]ItemMod
)

func init() {
	ItemModLookup = make(map[string]ItemMod, len(ItemModNames))

	for code, name := range ItemModNames {
		ItemModLookup[name] = code
	}
}

func (i *ItemMod) EncodeWord() (str string, err error) {
	var ok bool
	str, ok = ItemModNames[*i]
	if !ok {
		err = fmt.Errorf("models: ItemMod.EncodeWord, no name for %d", *i)
	}
	return
}

func (i *ItemMod) DecodeWord(data string) (err error) {
	var ok bool
	*i, ok = ItemModLookup[data]
	if !ok {
		err = fmt.Errorf("models: ItemMod.DecodeWord, no code for %s", data)
	}
	return
}

type ItemStat struct {
	Type  ItemMod
	Value int32
}

func (i *ItemMod) Resolve(build vsn.Build, u32 uint32) error {
	if build <= vsn.Alpha {
		mods := map[uint32]ItemMod{
			0: ItemModMana,
			1: ItemModHealth,
			3: ItemModAgility,
			4: ItemModStrength,
			5: ItemModIntellect,
			6: ItemModSpirit,
			7: ItemModStamina,
		}

		var ok bool
		*i, ok = mods[u32]
		if !ok {
			return fmt.Errorf("invalid code: %d", u32)
		}

		return nil
	}

	*i = ItemMod(u32)
	return nil
}

func (i *ItemMod) Uint32(build vsn.Build) (uint32, error) {
	if build <= vsn.Alpha {
		codes := map[ItemMod]uint32{
			ItemModMana:      0,
			ItemModHealth:    1,
			ItemModAgility:   3,
			ItemModStrength:  4,
			ItemModIntellect: 5,
			ItemModSpirit:    6,
			ItemModStamina:   7,
		}

		var err error

		code, ok := codes[*i]
		if !ok {
			err = fmt.Errorf("invalid item mod code for alpha")
		}

		return code, err
	}

	return uint32(*i), nil
}
