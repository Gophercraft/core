package spell

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type CastStatus uint8

const (
	Success                    CastStatus = 0
	AffectingCombat            CastStatus = 1
	AlreadyAtFullHealth        CastStatus = 2
	AlreadyAtFullMana          CastStatus = 3
	AlreadyAtFullPower         CastStatus = 4
	AlreadyBeingTamed          CastStatus = 5
	AlreadyHaveCharm           CastStatus = 6
	AlreadyHaveSummon          CastStatus = 7
	AlreadyOpen                CastStatus = 8
	AuraBounced                CastStatus = 9
	AutotrackInterrupted       CastStatus = 10
	BadImplicitTargets         CastStatus = 11
	BadTargets                 CastStatus = 12
	CantBeCharmed              CastStatus = 13
	CantBeDisenchanted         CastStatus = 14
	CantBeDisenchantedSkill    CastStatus = 15
	CantBeMilled               CastStatus = 16
	CantBeProspected           CastStatus = 17
	CantCastOnTapped           CastStatus = 18
	CantDuelWhileInvisible     CastStatus = 19
	CantDuelWhileStealthed     CastStatus = 20
	CantStealth                CastStatus = 21
	CasterAurastate            CastStatus = 22
	CasterDead                 CastStatus = 23
	Charmed                    CastStatus = 24
	ChestInUse                 CastStatus = 25
	Confused                   CastStatus = 26
	DontReport                 CastStatus = 27
	EquippedItem               CastStatus = 28
	EquippedItemClass          CastStatus = 29
	EquippedItemClassMainhand  CastStatus = 30
	EquippedItemClassOffhand   CastStatus = 31
	Error                      CastStatus = 32
	Fizzle                     CastStatus = 33
	Fleeing                    CastStatus = 34
	FoodLowlevel               CastStatus = 35
	Highlevel                  CastStatus = 36
	HungerSatiated             CastStatus = 37
	Immune                     CastStatus = 38
	IncorrectArea              CastStatus = 39
	Interrupted                CastStatus = 40
	InterruptedCombat          CastStatus = 41
	ItemAlreadyEnchanted       CastStatus = 42
	ItemGone                   CastStatus = 43
	ItemNotFound               CastStatus = 44
	ItemNotReady               CastStatus = 45
	LevelRequirement           CastStatus = 46
	LineOfSight                CastStatus = 47
	Lowlevel                   CastStatus = 48
	LowCastlevel               CastStatus = 49
	MainhandEmpty              CastStatus = 50
	Moving                     CastStatus = 51
	NeedAmmo                   CastStatus = 52
	NeedAmmoPouch              CastStatus = 53
	NeedExoticAmmo             CastStatus = 54
	NeedMoreItems              CastStatus = 55
	Nopath                     CastStatus = 56
	NotBehind                  CastStatus = 57
	NotFishable                CastStatus = 58
	NotFlying                  CastStatus = 59
	NotHere                    CastStatus = 60
	NotInfront                 CastStatus = 61
	NotInControl               CastStatus = 62
	NotKnown                   CastStatus = 63
	NotMounted                 CastStatus = 64
	NotOnTaxi                  CastStatus = 65
	NotOnTransport             CastStatus = 66
	NotReady                   CastStatus = 67
	NotShapeshift              CastStatus = 68
	NotStanding                CastStatus = 69
	NotTradeable               CastStatus = 70
	NotTrading                 CastStatus = 71
	NotUnsheathed              CastStatus = 72
	NotWhileGhost              CastStatus = 73
	NotWhileLooting            CastStatus = 74
	NoAmmo                     CastStatus = 75
	NoChargesRemain            CastStatus = 76
	NoChampion                 CastStatus = 77
	NoComboPoints              CastStatus = 78
	NoDueling                  CastStatus = 79
	NoEndurance                CastStatus = 80
	NoFish                     CastStatus = 81
	NoItemsWhileShapeshifted   CastStatus = 82
	NoMountsAllowed            CastStatus = 83
	NoPet                      CastStatus = 84
	NoPower                    CastStatus = 85
	NothingToDispel            CastStatus = 86
	NothingToSteal             CastStatus = 87
	OnlyAbovewater             CastStatus = 88
	OnlyDaytime                CastStatus = 89
	OnlyIndoors                CastStatus = 90
	OnlyMounted                CastStatus = 91
	OnlyNighttime              CastStatus = 92
	OnlyOutdoors               CastStatus = 93
	OnlyShapeshift             CastStatus = 94
	OnlyStealthed              CastStatus = 95
	OnlyUnderwater             CastStatus = 96
	OutOfRange                 CastStatus = 97
	Pacified                   CastStatus = 98
	Possessed                  CastStatus = 99
	Reagents                   CastStatus = 100
	RequiresArea               CastStatus = 101
	RequiresSpellFocus         CastStatus = 102
	Rooted                     CastStatus = 103
	Silenced                   CastStatus = 104
	SpellInProgress            CastStatus = 105
	SpellLearned               CastStatus = 106
	SpellUnavailable           CastStatus = 107
	Stunned                    CastStatus = 108
	TargetsDead                CastStatus = 109
	TargetAffectingCombat      CastStatus = 110
	TargetAurastate            CastStatus = 111
	TargetDueling              CastStatus = 112
	TargetEnemy                CastStatus = 113
	TargetEnraged              CastStatus = 114
	TargetFriendly             CastStatus = 115
	TargetInCombat             CastStatus = 116
	TargetIsPlayer             CastStatus = 117
	TargetIsPlayerControlled   CastStatus = 118
	TargetNotDead              CastStatus = 119
	TargetNotInParty           CastStatus = 120
	TargetNotLooted            CastStatus = 121
	TargetNotPlayer            CastStatus = 122
	TargetNoPockets            CastStatus = 123
	TargetNoWeapons            CastStatus = 124
	TargetNoRangedWeapons      CastStatus = 125
	TargetUnskinnable          CastStatus = 126
	ThirstSatiated             CastStatus = 127
	TooClose                   CastStatus = 128
	TooManyOfItem              CastStatus = 129
	TotemCategory              CastStatus = 130
	Totems                     CastStatus = 131
	TryAgain                   CastStatus = 132
	UnitNotBehind              CastStatus = 133
	UnitNotInfront             CastStatus = 134
	WrongPetFood               CastStatus = 135
	NotWhileFatigued           CastStatus = 136
	TargetNotInInstance        CastStatus = 137
	NotWhileTrading            CastStatus = 138
	TargetNotInRaid            CastStatus = 139
	TargetFreeforall           CastStatus = 140
	NoEdibleCorpses            CastStatus = 141
	OnlyBattlegrounds          CastStatus = 142
	TargetNotGhost             CastStatus = 143
	TransformUnusable          CastStatus = 144
	WrongWeather               CastStatus = 145
	DamageImmune               CastStatus = 146
	PreventedByMechanic        CastStatus = 147
	PlayTime                   CastStatus = 148
	Reputation                 CastStatus = 149
	MinSkill                   CastStatus = 150
	NotInArena                 CastStatus = 151
	NotOnShapeshift            CastStatus = 152
	NotOnStealthed             CastStatus = 153
	NotOnDamageImmune          CastStatus = 154
	NotOnMounted               CastStatus = 155
	TooShallow                 CastStatus = 156
	TargetNotInSanctuary       CastStatus = 157
	TargetIsTrivial            CastStatus = 158
	BmOrInvisgod               CastStatus = 159
	ExpertRidingRequirement    CastStatus = 160
	ArtisanRidingRequirement   CastStatus = 161
	NotIdle                    CastStatus = 162
	NotInactive                CastStatus = 163
	PartialPlaytime            CastStatus = 164
	NoPlaytime                 CastStatus = 165
	NotInBattleground          CastStatus = 166
	NotInRaidInstance          CastStatus = 167
	OnlyInArena                CastStatus = 168
	TargetLockedToRaidInstance CastStatus = 169
	OnUseEnchant               CastStatus = 170
	NotOnGround                CastStatus = 171
	CustomError                CastStatus = 172
	CantDoThatRightNow         CastStatus = 173
	TooManySockets             CastStatus = 174
	InvalidGlyph               CastStatus = 175
	UniqueGlyph                CastStatus = 176
	GlyphSocketLocked          CastStatus = 177
	NoValidTargets             CastStatus = 178
	ItemAtMaxCharges           CastStatus = 179
	NotInBarbershop            CastStatus = 180
	FishingTooLow              CastStatus = 181
	ItemEnchantTradeWindow     CastStatus = 182
	SummonPending              CastStatus = 183
	MaxSockets                 CastStatus = 184
	PetCanRename               CastStatus = 185
	TargetCannotBeResurrected  CastStatus = 186
	Unknown                    CastStatus = 187
)

type CastResult struct {
	CastID    byte
	SpellID   uint32
	Status    CastStatus
	CastCount byte
}

func (cr *CastResult) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_CAST_RESULT
	if build.AddedIn(vsn.V3_0_2) {
		out.WriteByte(cr.CastID)
	}
	out.WriteUint32(cr.SpellID)
	out.WriteByte(uint8(cr.Status))
	if build.AddedIn(vsn.V3_0_2) {
		out.WriteByte(cr.CastCount)
	}
	return nil
}
