package models

import "fmt"

// ModStat enumerates a list of attributes, that can be multiplied and added to by spell effects.
// Note that these stats are for internal Gophercraft server logic. These values have no bearing on any protocol.
type ModStat uint8

const (
	WalkSpeed ModStat = iota
	RunSpeed
	SwimSpeed
	MountSpeed
	BaseHealth
	BaseMana
	Strength
	Agility
	Stamina
	Intellect
	Spirit
	CriticalStrike
	PhysicalDamageDealt // note: dealt to enemies,
	HolyDamageDealt
	FireDamageDealt
	NatureDamageDealt
	FrostDamageDealt
	ShadowDamageDealt
	ArcaneDamageDealt
	PhysicalDamageTaken // note: this is the damage taken from enemiess
	HolyDamageTaken
	FireDamageTaken
	NatureDamageTaken
	FrostDamageTaken
	ShadowDamageTaken
	ArcaneDamageTaken
	ResistPhysicalDamage
	ResistHolyDamage
	ResistFireDamage
	ResistNatureDamage
	ResistFrostDamage
	ResistShadowDamage
	ResistArcaneDamage
	Health
	DefenseRating
	DodgeRating
	ParryRating
	BlockRating
	MeleeHitRating
	RangedHitRating
	SpellHitRating
	MeleeCriticalRating
	RangedCriticalRating
	SpellCriticalRating
	MeleeHitTakenRating
	RangedHitTakenRating
	SpellHitTakenRating
	MeleeCriticalTakenRating
	RangedCriticalTakenRating
	SpellCriticalTakenRating
	MeleeHasteRating
	RangedHasteRating
	HasteSpellRating
	HitRating
	CritRating
	HitTakenRating
	CritTakenRating
	ResilienceRating
	HasteRating
	ExpertiseRating
	AttackPower
	RangedAttackPower
	FeralAttackPower
	SpellHealingDone
	SpellDamageDone
	ManaRegeneration
	ArmorPenetrationRating
	SpellPower
	HealthRegen
	SpellPenetration
	BlockValue
	NumModStats
)

var StatNames = map[ModStat]string{
	WalkSpeed:                 "WalkSpeed",
	RunSpeed:                  "RunSpeed",
	SwimSpeed:                 "SwimSpeed",
	MountSpeed:                "MountSpeed",
	BaseHealth:                "BaseHealth",
	BaseMana:                  "BaseMana",
	Stamina:                   "Stamina",
	Strength:                  "Strength",
	Agility:                   "Agility",
	Intellect:                 "Intellect",
	Spirit:                    "Spirit",
	CriticalStrike:            "CriticalStrike",
	PhysicalDamageDealt:       "PhysicalDamageDealt",
	HolyDamageDealt:           "HolyDamageDealt",
	FireDamageDealt:           "FireDamageDealt",
	NatureDamageDealt:         "NatureDamageDealt",
	FrostDamageDealt:          "FrostDamageDealt",
	ShadowDamageDealt:         "ShadowDamageDealt",
	ArcaneDamageDealt:         "ArcaneDamageDealt",
	PhysicalDamageTaken:       "PhysicalDamageTaken",
	HolyDamageTaken:           "HolyDamageTaken",
	FireDamageTaken:           "FireDamageTaken",
	NatureDamageTaken:         "NatureDamageTaken",
	FrostDamageTaken:          "FrostDamageTaken",
	ShadowDamageTaken:         "ShadowDamageTaken",
	ArcaneDamageTaken:         "ArcaneDamageTaken",
	ResistPhysicalDamage:      "ResistPhysicalDamage",
	ResistHolyDamage:          "ResistHolyDamage",
	ResistFireDamage:          "ResistFireDamage",
	ResistNatureDamage:        "ResistNatureDamage",
	ResistFrostDamage:         "ResistFrostDamage",
	ResistShadowDamage:        "ResistShadowDamage",
	ResistArcaneDamage:        "ResistArcaneDamage",
	Health:                    "Health",
	DefenseRating:             "DefenseRating",
	DodgeRating:               "DodgeRating",
	ParryRating:               "ParryRating",
	BlockRating:               "BlockRating",
	MeleeHitRating:            "MeleeHitRating",
	RangedHitRating:           "RangedHitRating",
	SpellHitRating:            "SpellHitRating",
	MeleeCriticalRating:       "MeleeCriticalRating",
	RangedCriticalRating:      "RangedCriticalRating",
	SpellCriticalRating:       "SpellCriticalRating",
	MeleeHitTakenRating:       "MeleeHitTakenRating",
	RangedHitTakenRating:      "RangedHitTakenRating",
	SpellHitTakenRating:       "SpellHitTakenRating",
	MeleeCriticalTakenRating:  "MeleeCriticalTakenRating",
	RangedCriticalTakenRating: "RangedCriticalTakenRating",
	SpellCriticalTakenRating:  "SpellCriticalTakenRating",
	MeleeHasteRating:          "MeleeHasteRating",
	RangedHasteRating:         "RangedHasteRating",
	HasteSpellRating:          "HasteSpellRating",
	HitRating:                 "HitRating",
	CritRating:                "CritRating",
	HitTakenRating:            "HitTakenRating",
	CritTakenRating:           "CritTakenRating",
	ResilienceRating:          "ResilienceRating",
	HasteRating:               "HasteRating",
	ExpertiseRating:           "ExpertiseRating",
	AttackPower:               "AttackPower",
	RangedAttackPower:         "RangedAttackPower",
	FeralAttackPower:          "FeralAttackPower",
	SpellHealingDone:          "SpellHealingDone",
	SpellDamageDone:           "SpellDamageDone",
	ManaRegeneration:          "ManaRegeneration",
	ArmorPenetrationRating:    "ArmorPenetrationRating",
	SpellPower:                "SpellPower",
	HealthRegen:               "HealthRegen",
	SpellPenetration:          "SpellPenetration",
	BlockValue:                "BlockValue",
}

var StatLookup = map[string]ModStat{}

func init() {
	for k, v := range StatNames {
		StatLookup[v] = k
	}
}

func (ms *ModStat) DecodeWord(in string) error {
	var ok bool
	*ms, ok = StatLookup[in]
	if !ok {
		return fmt.Errorf("models: no ModStat for %s", in)
	}
	return nil
}

func (ms *ModStat) EncodeWord() (string, error) {
	str, ok := StatNames[*ms]
	if !ok {
		return "", fmt.Errorf("models: no name for ModStat %d", *ms)
	}

	return str, nil
}

type ModStatMask uint64

func (mm *ModStatMask) Set(ms ModStat, v bool) {
	var flag ModStatMask
	flag = 1 << ModStatMask(ms)

	if v {
		*mm |= flag
	} else {
		*mm &= ^flag
	}
}

func (mm *ModStatMask) Add(ms ModStat) {
	mm.Set(ms, true)
}

func (mm *ModStatMask) Has(ms ModStat) bool {
	var flag ModStatMask
	flag = 1 << ModStatMask(ms)
	return *mm&flag != 0
}

type ModOp uint8

const (
	// Mod by adding a value
	ModMult ModOp = 1 << iota
	// Mod with multiplication
	ModPlus
)

type ClientStat uint8

const (
	ClientStatStrength ClientStat = iota
	ClientStatAgility
	ClientStatStamina
	ClientStatIntellect
	ClientStatSpirit
)

func (ms ModStat) IsClientStat() bool {
	return ms >= Strength && ms <= Spirit
}

func (ms ModStat) ClientStat() ClientStat {
	if ms.IsClientStat() == false {
		panic(ms)
	}

	return ClientStat(ms - Strength)
}
