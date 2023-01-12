package spell

import (
	"fmt"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/vsn"
)

type Attribute uint32

const (
	// Attributes
	Attr_ProcFailureBurnsCharge      Attribute = iota //  0
	Attr_UsesRangedSlot                               //  1 All ranged abilites have this flag
	Attr_OnNextSwing_NO_DAMAGE                        //  2 on next swing
	Attr_DoNotLog_IMMUNE_MISSES                       //  3 not set in 3.0.3
	Attr_IsAbility                                    //  4 Displays ability instead of spell clientside
	Attr_IsTradeskill                                 //  5 trade spells, will be added by client to a sublist of profession spell
	Attr_Passive                                      //  6 Passive spell
	Attr_DoNotDisplay                                 //  7 Hidden in Spellbook, Aura Icon, Combat Log
	Attr_DoNotLog                                     //  8
	Attr_HeldItemOnly                                 //  9 Client automatically selects item from mainhand slot as a cast target
	Attr_OnNextSwing                                  //  10 on next swing 2
	Attr_WearerCastsProcTrigger                       //  11
	Attr_DaytimeOnly                                  //  12 only useable at daytime, not set in 2.4.2
	Attr_NightOnly                                    //  13 only useable at night, not set in 2.4.2
	Attr_OnlyIndoors                                  //  14 only useable indoors, not set in 2.4.2
	Attr_OnlyOutdoors                                 //  15 Only useable outdoors.
	Attr_NotShapeshift                                //  16 Not while shapeshifted
	Attr_OnlyStealthed                                //  17 Must be in stealth
	Attr_DoNotSheath                                  //  18 client won't hide unit weapons in sheath on cast/channel TODO: Implement
	Attr_ScalesWithCreatureLevel                      //  19 spelldamage depends on caster level
	Attr_CancelsAutoAttackCombat                      //  20 Stop attack after use this spell (and not begin attack if use)
	Attr_NoActiveDefense                              //  21 Cannot be dodged/parried/blocked
	Attr_TrackTargetInCastPlayerOnly                  // 22 SetTrackingTarget
	Attr_AllowCastWhileDead                           //  23 castable while dead
	Attr_AllowWhileMounted                            //  24 castable while mounted
	Attr_CooldownOnEvent                              //  25 Activate and start cooldown after aura fade or remove summoned creature or go
	Attr_AuraIsDebuff                                 //  26
	Attr_AllowWhileSitting                            //  27 castable while sitting
	Attr_NotInCombatOnlyPeaceful                      //  28 Cannot be used in combat
	Attr_NoImmunities                                 //  29 unaffected by invulnerability
	Attr_HeartbeatResist                              //  30 Chance for spell effects to break early (heartbeat resist)
	Attr_NoAuraCancel                                 //  31 positive aura can't be canceled
	// AttributesEx
	AttrEx_DismissPet                  //  0
	AttrEx_DrainAllPower               //  1 use all power (Only paladin Lay of Hands and Bunyanize)
	AttrEx_Channeled1                  //  2 channeled 1
	AttrEx_CantBeRedirected            //  3
	AttrEx_Unk4                        //  4
	AttrEx_NotBreakStealth             //  5 Not break stealth
	AttrEx_Channeled2                  //  6 channeled 2
	AttrEx_CantBeReflected             //  7
	AttrEx_NotInCombatTarget           //  8 Spell req target not to be in combat state
	AttrEx_FacingTarget                //  9 TODO: CONFIRM!
	AttrEx_NoThreat                    //  10 no generates threat on cast 100%
	AttrEx_DontRefreshDurationOnRecast //  11 Aura will not refresh its duration when recast
	AttrEx_FailureBreaksStealth        //  12
	AttrEx_ToggleFarsight              //  13
	AttrEx_ChannelTrackTarget          //  14
	AttrEx_DispelAurasOnImmunity       //  15 remove auras on immunity
	AttrEx_UnaffectedBySchoolImmune    //  16 unaffected by school immunity
	AttrEx_UnautocastableByCharmed     //  17 TODO: Investigate more Chero version: SPELL_ATTR_EX_PLAYER_CANT_CAST_CHARMED, likely related to MC
	AttrEx_PreventsAnim                //  18
	AttrEx_CantTargetSelf              //  19 spells with area effect or friendly targets that exclude the caster
	AttrEx_ReqTargetComboPoints        //  20 Req combo points on target
	AttrEx_ThreatOnlyOnMiss            //  21
	AttrEx_ReqComboPoints              //  22 Use combo points (in 4.x not required combo point target selected)
	AttrEx_Unk23                       //  23
	AttrEx_Unk24                       //  24 Req fishing pole?? SPELL_ATTR_EX_FISHING
	AttrEx_Unk25                       //  25 not set in 2.4.2
	AttrEx_RequireAllTargets           //  26
	AttrEx_RefundPower                 //  27 All these spells refund power on parry or deflect
	AttrEx_DontDisplayInAuraBar        //  28
	AttrEx_ChannelDisplaySpellName     //  29
	AttrEx_EnableAtDodge               //  30 overpower
	AttrEx_Unk31                       //  31
	//AttributesEx2
	AttrEx2_CanTargetDead                //  0 can target dead unit or corpse
	AttrEx2_Unk1                         //  1
	AttrEx2_IgnoreLos                    //  2 do not need LOS (e.g. 18220 since 3.3.3)
	AttrEx2_Unk3                         //  3 auto targeting? (e.g. fishing skill enhancement items since 3.3.3)
	AttrEx2_DisplayInStanceBar           //  4 client displays icon in stance bar when learned, even if not shapeshift
	AttrEx2_AutorepeatFlag               //  5
	AttrEx2_CantTargetTapped             //  6 only usable on tabbed by yourself
	AttrEx2_Unk7                         //  7
	AttrEx2_Unk8                         //  8 not set in 2.4.2
	AttrEx2_Unk9                         //  9
	AttrEx2_Unk10                        //  10 SPELL_ATTR_EX2_TAME_SPELLS
	AttrEx2_HealthFunnel                 //  11
	AttrEx2_Unk12                        //  12 SPELL_ATTR_EX2_CLASS_CLEAVE
	AttrEx2_Unk13                        //  13 TODO: Implement from TC SPELL_ATTR_EX2_CASTABLE_ON_ITEMS
	AttrEx2_Unk14                        //  14
	AttrEx2_Unk15                        //  15 not set in 2.4.2
	AttrEx2_TameBeast                    //  16
	AttrEx2_NotResetAutoActions          //  17 suspend weapon timer instead of resetting it, (?Hunters Shot and Stings only have this flag?)
	AttrEx2_ReqDeadPet                   //  18 Only Revive pet - possible req dead pet
	AttrEx2_NotNeedShapeshift            //  19 does not necessary need shapeshift (pre-3.x not have passive spells with this attribute)
	AttrEx2_FacingTargetsBack            //  20 TODO: CONFIRM!
	AttrEx2_DamageReducedShield          //  21 for ice blocks, pala immunity buffs, priest absorb shields, but used also for other spells -> not sure!
	AttrEx2_NoInitialThreat              //  22
	AttrEx2_IsArcaneConcentration        //  23 Only mage Arcane Concentration have this flag
	AttrEx2_Unk24                        //  24
	AttrEx2_Unk25                        //  25
	AttrEx2_UnaffectedByAuraSchoolImmune //  26
	AttrEx2_Unk27                        //  27
	AttrEx2_Unk28                        //  28 no breaks stealth if it fails??
	AttrEx2_CantCrit                     //  29 Spell can't crit
	AttrEx2_TriggeredCanTriggerProc      //  30 Chero hint: SPELL_ATTR_EX2_CAN_TRIGGER_VICTIM
	AttrEx2_FoodBuff                     //  31 Food or Drink Buff (like Well Fed)
	// AttributesEx3
	AttrEx3_OutOfCombatAttack          //  0 Spell landed counts as hostile action against enemy even if it doesn't trigger combat state, propagates PvP flags
	AttrEx3_Unk1                       //  1
	AttrEx3_Unk2                       //  2
	AttrEx3_BlockableSpell             //  3 TODO: Investigate more
	AttrEx3_IgnoreResurrectionTimer    //  4 Druid Rebirth only this spell have this flag
	AttrEx3_Unk5                       //  5
	AttrEx3_Unk6                       //  6
	AttrEx3_StackForDiffCasters        //  7 create a separate (de)buff stack for each caster
	AttrEx3_TargetOnlyPlayer           //  8 Can target only player
	AttrEx3_TriggeredCanTriggerSpecial //  9 Can only proc auras
	AttrEx3_MainHand                   //  10 Main hand weapon required
	AttrEx3_Battleground               //  11 Can casted only on battleground
	AttrEx3_CastOnDead                 //  12 target is a dead player (not every spell has this flag)
	AttrEx3_DontDisplayChannelBar      //  13
	AttrEx3_IsHonorlessTarget          //  14 "Honorless Target" only this spells have this flag
	AttrEx3_RangedAttack               //  15 Spells with this attribute are processed as ranged attacks in client
	AttrEx3_SuppressCasterProcs        //  16
	AttrEx3_SuppressTargetProcs        //  17
	AttrEx3_AlwaysHit                  //  18 Spell should always hit its target
	AttrEx3_Unk19                      //  19 TODO: Implement from TC
	AttrEx3_DeathPersistent            //  20 Death persistent spells
	AttrEx3_Unk21                      //  21
	AttrEx3_ReqWand                    //  22 Req wand
	AttrEx3_Unk23                      //  23
	AttrEx3_ReqOffhand                 //  24 Req offhand weapon
	AttrEx3_TreatAsPeriodic            //  25 Treated as periodic spell
	AttrEx3_CanProcFromTriggered       //  26 Auras with this attribute can proc off AttrEx3_TriggeredCanTriggerSpecial
	AttrEx3_Unk27                      //  27
	AttrEx3_Unk28                      //  28 always cast ok ? (requires more research)
	AttrEx3_IgnoreCasterModifiers      //  29 Resistances should still affect damage
	AttrEx3_DontDisplayRange           //  30
	AttrEx3_Unk31                      //  31
	// AttributesEx3
	AttrEx4_IgnoreResistances          //  0
	AttrEx4_ProcOnlyOnCaster           //  1 Only proc on self-cast
	AttrEx4_AuraExpiresOffline         //  2
	AttrEx4_Unk3                       //  3
	AttrEx4_Unk4                       //  4 This will no longer cause guards to attack on use??
	AttrEx4_Unk5                       //  5
	AttrEx4_NotStealable               //  6 although such auras might be dispellable, they cannot be stolen
	AttrEx4_CanCastWhileCasting        //  7 In theory, can use this spell while another is channeled/cast/autocast
	AttrEx4_IgnoreDamageTakenModifiers //  8
	AttrEx4_TriggerActivate            //  9 initially disabled / trigger activate from event (Execute, Riposte, Deep Freeze end other)
	AttrEx4_SpellVsExtendCost          //  10 Rogue Shiv have this flag
	AttrEx4_Unk11                      //  11
	AttrEx4_Unk12                      //  12
	AttrEx4_Unk13                      //  13
	AttrEx4_DamageDoesntBreakAuras     //  14
	AttrEx4_Unk15                      //  15 Dont add to spellbook
	AttrEx4_NotUsableInArena           //  16 not usable in arena
	AttrEx4_UsableInArena              //  17 usable in arena
	AttrEx4_Unk18                      //  18 TODO: Investigate from TC
	AttrEx4_Unk19                      //  19
	AttrEx4_NotCheckSelfcastPower      //  20 do not give "more powerful spell" error message
	AttrEx4_Unk21                      //  21
	AttrEx4_Unk22                      //  22
	AttrEx4_Unk23                      //  23
	AttrEx4_AutoRangedCombatSpell      //  24
	AttrEx4_IsPetScaling               //  25 pet scaling auras
	AttrEx4_CastOnlyInOutland          //  26 Can only be used in Outland.
	AttrEx4_Unk27                      //  27 Always shown in combat log
	AttrEx4_Unk28                      //  28
	AttrEx4_Unk29                      //  29 Related to client check about dispel, CC seems to have them - dispel effect 0
	AttrEx4_Unk30                      //  30 - dispel effect 1
	AttrEx4_Unk31                      //  31 - dispel effect 2
)

func LoadAttributes(build vsn.Build, sp *dbdefs.Ent_Spell) (*update.Bitmask, error) {
	if build.RemovedIn(vsn.V2_0_1) {
		b := update.Bitmask{
			uint32(sp.Attributes),
			uint32(sp.AttributesEx),
			uint32(sp.AttributesExB),
			uint32(sp.AttributesExC),
			uint32(sp.AttributesExD),
		}
		return &b, nil
	}

	return nil, fmt.Errorf("spell: fail")
}
