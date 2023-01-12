package models

import (
	"github.com/Gophercraft/core/i18n"
)

type CreatureTemplate struct {
	ID              string
	Entry           uint32    `xorm:"'entry'"`
	Name            i18n.Text `xorm:"'name'"`
	SubName         i18n.Text `xorm:"'sub_name'"`
	MinLevel        uint32    `xorm:"'min_level'"`
	MaxLevel        uint32    `xorm:"'max_level'"`
	DisplayIDs      []uint32  `xorm:"'display_ids'"`
	Faction         uint32    `xorm:"'faction'"`
	Scale           float32   `xorm:"'scale'"`
	Family          string    `xorm:"'family'"`
	CreatureType    uint32    `xorm:"'creature_type'"`
	InhabitType     uint32    `xorm:"'inhabit_type'"`
	RegenerateStats uint32    `xorm:"'regenerateStats'"`
	RacialLeader    bool      `xorm:"'racialLeader'"`
	// NpcFlags: should not tied to any particular version.
	Gossip         bool
	QuestGiver     bool
	Vendor         bool
	FlightMaster   bool
	Trainer        bool
	SpiritHealer   bool
	SpiritGuide    bool
	Innkeeper      bool
	Banker         bool
	Petitioner     bool
	TabardDesigner bool
	BattleMaster   bool
	Auctioneer     bool
	StableMaster   bool
	Repairer       bool
	OutdoorPVP     bool

	// UnitFlags: should not tied to any particular version.
	ServerControlled    bool // 0x1
	NonAttackable       bool // 0x2
	RemoveClientControl bool // 0x4
	PlayerControlled    bool // 0x8
	Rename              bool // 0x10
	PetAbandon          bool // 0x20
	OOCNotAttackable    bool // 0x100
	Passive             bool // 0x200
	PVP                 bool // 0x1000
	IsSilenced          bool // 0x2000
	IsPersuaded         bool // 0x4000
	Swimming            bool // 0x8000
	RemoveAttackIcon    bool // 0x10000
	IsPacified          bool // 0x20000
	IsStunned           bool // 0x40000
	InCombat            bool // 0x80000
	InTaxiFlight        bool // 0x100000
	Disarmed            bool // 0x200000
	Confused            bool // 0x400000
	Fleeing             bool // 0x800000
	Possessed           bool // 0x1000000
	NotSelectable       bool // 0x2000000
	Skinnable           bool // 0x4000000
	AurasVisible        bool // 0x8000000
	Sheathe             bool // 0x40000000
	NoKillReward        bool // 0x80000000

	// DynamicFlags: should not tied to any particular version.
	Lootable              bool
	TrackUnit             bool
	Tapped                bool
	TappedByPlayer        bool
	SpecialInfo           bool
	VisuallyDead          bool
	TappedByAllThreatList bool

	// Extra flags
	InstanceBind             bool // creature kill bind instance with killer and killer’s group
	NoAggroOnSight           bool // no aggro (ignore faction/reputation hostility)
	NoParry                  bool // creature can’t parry
	NoParryHasten            bool // creature can’t counter-attack at parry
	NoBlock                  bool //	creature can’t block
	NoCrush                  bool // creature can’t do crush attacks
	NoXPAtKill               bool // creature kill not provide XP
	Invisible                bool // creature is always invisible for player (mostly trigger creatures)
	NotTauntable             bool // creature is immune to taunt auras and effect attack me
	AggroZone                bool // creature sets itself in combat with zone on aggro
	Guard                    bool // creature is a guard
	NoCallAssist             bool // creature shouldn’t call for assistance on aggro
	Active                   bool //creature is active object. Grid of this creature will be loaded and creature set as active
	ForceEnableMMap          bool // creature is forced to use MMaps
	ForceDisableMMap         bool // creature is forced to NOT use MMaps
	WalkInWater              bool // creature is forced to walk in water even it can swim
	Civilian                 bool // CreatureInfo→civilian substitute (for expansions as Civilian Colum was removed)
	NoMelee                  bool // creature can’t melee
	FarView                  bool // creature with far view
	ForceAttackingCapability bool // SetForceAttackingCapability(true); for nonattackable, nontargetable creatures that should be able to attack nontheless
	IgnoreUsedPosition       bool // ignore creature when checking used positions around target
	CountSpawns              bool // count creature spawns in Map*
	HasteSpellImmunity       bool // immunity to COT or Mind Numbing Poison – very common in instances

	// CreatureTypeFlags    uint32     `xorm:"'creatureTypeFlags'"`
	Tameable                bool // Makes the mob tameable (must also be a beast and have family set)
	VisibleToGhosts         bool // Sets Creatures that can ALSO be seen when player is a ghost. Used in CanInteract function by client, can’t be attacked
	BossLevel               bool
	DontPlayWoundParryAnim  bool
	HideFactionTooltip      bool // Controls something in client tooltip related to creature faction
	SpellAttackable         bool
	DeadInteract            bool // Player can interact with the creature if its dead (not player dead)
	HerbLoot                bool // Uses Skinning Loot Field
	MiningLoot              bool // Makes Mob Corpse Mineable – Uses Skinning Loot Field
	DontLogDeath            bool // Does not combatlog death.
	MountedCombat           bool
	CanAssist               bool // Can aid any player or group in combat. Typically seen for escorting NPC’s
	PetHasActionBar         bool // checked from calls in Lua_PetHasActionBar
	MaskUID                 bool
	EngineerLoot            bool // Makes Mob Corpse Engineer Lootable – Uses Skinning Loot Field
	ExoticPet               bool // Tamable as an exotic pet. Normal tamable flag must also be set.
	UseDefaultCollisionBox  bool
	IsSiegeWeapon           bool
	ProjectileCollision     bool
	HideNameplate           bool
	DontPlayMountedAnim     bool
	IsLinkAll               bool
	InteractOnlyWithCreator bool
	ForceGossip             bool

	SpeedWalk            float32 `xorm:"'speedWalk'"`
	SpeedRun             float32 `xorm:"'speedRun'"`
	UnitClass            uint32  `xorm:"'unitClass'"`
	Rank                 uint32  `xorm:"'rank'"`
	HealthMultiplier     float32 `xorm:"'healthMultiplier'"`
	PowerMultiplier      float32 `xorm:"'powerMultiplier'"`
	DamageMultiplier     float32 `xorm:"'damageMultiplier'"`
	DamageVariance       float32 `xorm:"'damageVariance'"`
	ArmorMultiplier      float32 `xorm:"'armorMultiplier'"`
	ExperienceMultiplier float32 `xorm:"'experienceMultiplier'"`
	MinLevelHealth       uint32  `xorm:"'minLevelHealth'"`
	MaxLevelHealth       uint32  `xorm:"'maxLevelHealth'"`
	MinLevelMana         uint32  `xorm:"'minLevelMana'"`
	MaxLevelMana         uint32  `xorm:"'maxLevelMana'"`
	MinMeleeDmg          float32 `xorm:"'minMeleeDmg'"`
	MaxMeleeDmg          float32 `xorm:"'maxMeleeDmg'"`
	MinRangedDmg         float32 `xorm:"'minRangedDmg'"`
	MaxRangedDmg         float32 `xorm:"'maxRangedDmg'"`
	Armor                uint32  `xorm:"'armor'"`
	MeleeAttackPower     uint32  `xorm:"'meleeAttackPower'"`
	RangedAttackPower    uint32  `xorm:"'rangedAttackPower'"`
	MeleeBaseAttackTime  uint32  `xorm:"'meleeBaseAttackTime'"`
	RangedBaseAttackTime uint32  `xorm:"'mangedBaseAttackTime'"`
	DamageSchool         int32   `xorm:"'damageSchool'"`
	MinLootGold          Money   `xorm:"'minLootGold'"`
	MaxLootGold          Money   `xorm:"'maxLootGold'"`
	LootId               uint32  `xorm:"'lootId'"`
	PickpocketLootId     uint32  `xorm:"'pickpocketLootId'"`
	SkinningLootId       uint32  `xorm:"'skinningLootId'"`
	KillCredit1          uint32  `xorm:"'killCredit1'"`
	KillCredit2          uint32  `xorm:"'killCredit2'"`
	MechanicImmuneMask   uint32  `xorm:"'mechanicImmuneMask'"`
	SchoolImmuneMask     uint32  `xorm:"'schoolImmuneMask'"`
	ResistanceHoly       int32   `xorm:"'resistanceHoly'"`
	ResistanceFire       int32   `xorm:"'resistanceFire'"`
	ResistanceNature     int32   `xorm:"'resistanceNature'"`
	ResistanceFrost      int32   `xorm:"'resistanceFrost'"`
	ResistanceShadow     int32   `xorm:"'resistanceShadow'"`
	ResistanceArcane     int32   `xorm:"'resistanceArcane'"`
	PetSpellDataId       uint32  `xorm:"'petSpellDataId'"`
	MovementType         uint32  `xorm:"'movementType'"`
	TrainerType          int32   `xorm:"'trainerType'"`
	TrainerSpell         uint32  `xorm:"'trainerSpell'"`
	TrainerClass         uint32  `xorm:"'trainerClass'"`
	TrainerRace          uint32  `xorm:"'trainerRace'"`
	TrainerTemplateId    uint32  `xorm:"'trainerTemplateId'"`
	VendorTemplateId     uint32  `xorm:"'vendorTemplateId'"`
	GossipMenuId         string  `xorm:"'gossipMenuId'"`
	EquipmentTemplateId  uint32  `xorm:"'equipmentTemplateId'"`
	DishonourableKill    bool    `xorm:"'dishonourable_kill'"`
	AIName               string  `xorm:"'aIName'"`
	ScriptName           string  `xorm:"'script_name'"`
}
