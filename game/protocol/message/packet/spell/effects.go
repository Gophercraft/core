package spell

type Effect int

const (
	EffectNone                                 Effect = iota
	EffectInstaKill                                   //  1 SPELL_EFFECT_INSTAKILL
	EffectSchoolDMG                                   //  2 SPELL_EFFECT_SCHOOL_DAMAGE
	EffectDummy                                       //  3 SPELL_EFFECT_DUMMY
	EffectPortalTeleport                              //  4 SPELL_EFFECT_PORTAL_TELEPORT          unused
	Effect5                                           //  5 SPELL_EFFECT_5
	EffectApplyAura                                   //  6 SPELL_EFFECT_APPLY_AURA
	EffectEnvironmentalDMG                            //  7 SPELL_EFFECT_ENVIRONMENTAL_DAMAGE
	EffectPowerDrain                                  //  8 SPELL_EFFECT_POWER_DRAIN
	EffectHealthLeech                                 //  9 SPELL_EFFECT_HEALTH_LEECH
	EffectHeal                                        // 10 SPELL_EFFECT_HEAL
	EffectBind                                        // 11 SPELL_EFFECT_BIND
	EffectPortal                                      // 12 SPELL_EFFECT_PORTAL
	EffectTeleportToReturnPoint                       // 13 SPELL_EFFECT_TELEPORT_TO_RETURN_POINT
	EffectIncreaseCurrencyCap                         // 14 SPELL_EFFECT_INCREASE_CURRENCY_CAP
	EffectTeleportUnitsWithVisualLoadingScreen        // 15 SPELL_EFFECT_TELEPORT_WITH_SPELL_VISUAL_KIT_LOADING_SCREEN
	EffectQuestComplete                               // 16 SPELL_EFFECT_QUEST_COMPLETE
	EffectWeaponDamageNoSchool                        // 17 SPELL_EFFECT_WEAPON_DAMAGE_NOSCHOOL
	EffectResurrect                                   // 18 SPELL_EFFECT_RESURRECT
	EffectAddExtraAttacks                             // 19 SPELL_EFFECT_ADD_EXTRA_ATTACKS
	EffectDodge                                       // 20 SPELL_EFFECT_DODGE                    one spell: Dodge
	EffectEvade                                       // 21 SPELL_EFFECT_EVADE                    one spell: Evade (DND)
	EffectParry                                       // 22 SPELL_EFFECT_PARRY
	EffectBlock                                       // 23 SPELL_EFFECT_BLOCK                    one spell: Block
	EffectCreateItem                                  // 24 SPELL_EFFECT_CREATE_ITEM
	EffectWeapon                                      // 25 SPELL_EFFECT_WEAPON
	EffectDefense                                     // 26 SPELL_EFFECT_DEFENSE                  one spell: Defense
	EffectPersistentAA                                // 27 SPELL_EFFECT_PERSISTENT_AREA_AURA
	EffectSummonType                                  // 28 SPELL_EFFECT_SUMMON
	EffectLeap                                        // 29 SPELL_EFFECT_LEAP
	EffectEnergize                                    // 30 SPELL_EFFECT_ENERGIZE
	EffectWeaponPercentDamage                         // 31 SPELL_EFFECT_WEAPON_PERCENT_DAMAGE
	EffectTriggerMissile                              // 32 SPELL_EFFECT_TRIGGER_MISSILE
	EffectOpenLock                                    // 33 SPELL_EFFECT_OPEN_LOCK
	EffectSummonChangeItem                            // 34 SPELL_EFFECT_SUMMON_CHANGE_ITEM
	EffectApplyAreaAuraParty                          // 35 SPELL_EFFECT_APPLY_AREA_AURA_PARTY
	EffectLearnSpell                                  // 36 SPELL_EFFECT_LEARN_SPELL
	EffectSpellDefense                                // 37 SPELL_EFFECT_SPELL_DEFENSE            one spell: SPELLDEFENSE (DND)
	EffectDispel                                      // 38 SPELL_EFFECT_DISPEL
	EffectLanguage                                    // 39 SPELL_EFFECT_LANGUAGE
	EffectDualWield                                   // 40 SPELL_EFFECT_DUAL_WIELD
	EffectJump                                        // 41 SPELL_EFFECT_JUMP
	EffectJumpDest                                    // 42 SPELL_EFFECT_JUMP_DEST
	EffectTeleUnitsFaceCaster                         // 43 SPELL_EFFECT_TELEPORT_UNITS_FACE_CASTER
	EffectLearnSkill                                  // 44 SPELL_EFFECT_SKILL_STEP
	EffectPlayMovie                                   // 45 SPELL_EFFECT_PLAY_MOVIE
	EffectSpawn                                       // 46 SPELL_EFFECT_SPAWN clientside unit appears as if it was just spawned
	EffectTradeSkill                                  // 47 SPELL_EFFECT_TRADE_SKILL
	EffectStealth                                     // 48 SPELL_EFFECT_STEALTH                  one spell: Base Stealth
	EffectDetect                                      // 49 SPELL_EFFECT_DETECT                   one spell: Detect
	EffectTransmitted                                 // 50 SPELL_EFFECT_TRANS_DOOR
	EffectForceCriticalHit                            // 51 SPELL_EFFECT_FORCE_CRITICAL_HIT       unused
	EffectSetMaxBattlePetCount                        // 52 SPELL_EFFECT_SET_MAX_BATTLE_PET_COUNT
	EffectEnchantItemPerm                             // 53 SPELL_EFFECT_ENCHANT_ITEM
	EffectEnchantItemTmp                              // 54 SPELL_EFFECT_ENCHANT_ITEM_TEMPORARY
	EffectTameCreature                                // 55 SPELL_EFFECT_TAMECREATURE
	EffectSummonPet                                   // 56 SPELL_EFFECT_SUMMON_PET
	EffectLearnPetSpell                               // 57 SPELL_EFFECT_LEARN_PET_SPELL
	EffectWeaponDamage                                // 58 SPELL_EFFECT_WEAPON_DAMAGE
	EffectCreateRandomItem                            // 59 SPELL_EFFECT_CREATE_RANDOM_ITEM       create item base at spell specific loot
	EffectProficiency                                 // 60 SPELL_EFFECT_PROFICIENCY
	EffectSendEvent                                   // 61 SPELL_EFFECT_SEND_EVENT
	EffectPowerBurn                                   // 62 SPELL_EFFECT_POWER_BURN
	EffectThreat                                      // 63 SPELL_EFFECT_THREAT
	EffectTriggerSpell                                // 64 SPELL_EFFECT_TRIGGER_SPELL
	EffectApplyAreaAuraRaid                           // 65 SPELL_EFFECT_APPLY_AREA_AURA_RAID
	EffectRechargeItem                                // 66 SPELL_EFFECT_RECHARGE_ITEM
	EffectHealMaxHealth                               // 67 SPELL_EFFECT_HEAL_MAX_HEALTH
	EffectInterruptCast                               // 68 SPELL_EFFECT_INTERRUPT_CAST
	EffectDistract                                    // 69 SPELL_EFFECT_DISTRACT
	EffectCompleteAndRewardWorldQuest                 // 70 SPELL_EFFECT_COMPLETE_AND_REWARD_WORLD_QUEST
	EffectPickPocket                                  // 71 SPELL_EFFECT_PICKPOCKET
	EffectAddFarsight                                 // 72 SPELL_EFFECT_ADD_FARSIGHT
	EffectUntrainTalents                              // 73 SPELL_EFFECT_UNTRAIN_TALENTS
	EffectApplyGlyph                                  // 74 SPELL_EFFECT_APPLY_GLYPH
	EffectHealMechanical                              // 75 SPELL_EFFECT_HEAL_MECHANICAL          one spell: Mechanical Patch Kit
	EffectSummonObjectWild                            // 76 SPELL_EFFECT_SUMMON_OBJECT_WILD
	EffectScriptEffect                                // 77 SPELL_EFFECT_SCRIPT_EFFECT
	EffectAttack                                      // 78 SPELL_EFFECT_ATTACK
	EffectSanctuary                                   // 79 SPELL_EFFECT_SANCTUARY
	EffectModifyFollowerItemLevel                     // 80 SPELL_EFFECT_MODIFY_FOLLOWER_ITEM_LEVEL
	EffectPushAbilityToActionBar                      // 81 SPELL_EFFECT_PUSH_ABILITY_TO_ACTION_BAR
	EffectBindSight                                   // 82 SPELL_EFFECT_BIND_SIGHT
	EffectDuel                                        // 83 SPELL_EFFECT_DUEL
	EffectStuck                                       // 84 SPELL_EFFECT_STUCK
	EffectSummonPlayer                                // 85 SPELL_EFFECT_SUMMON_PLAYER
	EffectActivateObject                              // 86 SPELL_EFFECT_ACTIVATE_OBJECT
	EffectGameObjectDamage                            // 87 SPELL_EFFECT_GAMEOBJECT_DAMAGE
	EffectGameObjectRepair                            // 88 SPELL_EFFECT_GAMEOBJECT_REPAIR
	EffectGameObjectSetDestructionState               // 89 SPELL_EFFECT_GAMEOBJECT_SET_DESTRUCTION_STATE
	EffectKillCreditPersonal                          // 90 SPELL_EFFECT_KILL_CREDIT              Kill credit but only for single person
	EffectThreatAll                                   // 91 SPELL_EFFECT_THREAT_ALL
	EffectEnchantHeldItem                             // 92 SPELL_EFFECT_ENCHANT_HELD_ITEM
	EffectForceDeselect                               // 93 SPELL_EFFECT_FORCE_DESELECT
	EffectSelfResurrect                               // 94 SPELL_EFFECT_SELF_RESURRECT
	EffectSkinning                                    // 95 SPELL_EFFECT_SKINNING
	EffectCharge                                      // 96 SPELL_EFFECT_CHARGE
	EffectCastButtons                                 // 97 SPELL_EFFECT_CAST_BUTTON (totem bar since 3.2.2a)
	EffectKnockBack                                   // 98 SPELL_EFFECT_KNOCK_BACK
	EffectDisEnchant                                  // 99 SPELL_EFFECT_DISENCHANT
	EffectInebriate                                   //100 SPELL_EFFECT_INEBRIATE
	EffectFeedPet                                     //101 SPELL_EFFECT_FEED_PET
	EffectDismissPet                                  //102 SPELL_EFFECT_DISMISS_PET
	EffectReputation                                  //103 SPELL_EFFECT_REPUTATION
	EffectSummonObject                                //104 SPELL_EFFECT_SUMMON_OBJECT_SLOT1
	EffectSurvey                                      //105 SPELL_EFFECT_SURVEY
	EffectChangeRaidMarker                            //106 SPELL_EFFECT_CHANGE_RAID_MARKER
	EffectShowCorpseLoot                              //107 SPELL_EFFECT_SHOW_CORPSE_LOOT
	EffectDispelMechanic                              //108 SPELL_EFFECT_DISPEL_MECHANIC
	EffectResurrectPet                                //109 SPELL_EFFECT_RESURRECT_PET
	EffectDestroyAllTotems                            //110 SPELL_EFFECT_DESTROY_ALL_TOTEMS
	EffectDurabilityDamage                            //111 SPELL_EFFECT_DURABILITY_DAMAGE
	Effect112                                         //112 SPELL_EFFECT_112
	EffectCancelConversation                          //113 SPELL_EFFECT_CANCEL_CONVERSATION
	EffectTaunt                                       //114 SPELL_EFFECT_ATTACK_ME
	EffectDurabilityDamagePCT                         //115 SPELL_EFFECT_DURABILITY_DAMAGE_PCT
	EffectSkinPlayerCorpse                            //116 SPELL_EFFECT_SKIN_PLAYER_CORPSE       one spell: Remove Insignia bg usage required special corpse flags...
	EffectSpiritHeal                                  //117 SPELL_EFFECT_SPIRIT_HEAL              one spell: Spirit Heal
	EffectSkill                                       //118 SPELL_EFFECT_SKILL                    professions and more
	EffectApplyAreaAuraPet                            //119 SPELL_EFFECT_APPLY_AREA_AURA_PET
	EffectTeleportGraveyard                           //120 SPELL_EFFECT_TELEPORT_GRAVEYARD
	EffectWeaponDmg                                   //121 SPELL_EFFECT_NORMALIZED_WEAPON_DMG
	Effect122                                         //122 SPELL_EFFECT_122                      unused
	EffectSendTaxi                                    //123 SPELL_EFFECT_SEND_TAXI                taxi/flight related (misc value is taxi path id)
	EffectPullTowards                                 //124 SPELL_EFFECT_PULL_TOWARDS
	EffectModifyThreatPercent                         //125 SPELL_EFFECT_MODIFY_THREAT_PERCENT
	EffectStealBeneficialBuff                         //126 SPELL_EFFECT_STEAL_BENEFICIAL_BUFF    spell steal effect?
	EffectProspecting                                 //127 SPELL_EFFECT_PROSPECTING              Prospecting spell
	EffectApplyAreaAuraFriend                         //128 SPELL_EFFECT_APPLY_AREA_AURA_FRIEND
	EffectApplyAreaAuraEnemy                          //129 SPELL_EFFECT_APPLY_AREA_AURA_ENEMY
	EffectRedirectThreat                              //130 SPELL_EFFECT_REDIRECT_THREAT
	EffectPlaySound                                   //131 SPELL_EFFECT_PLAY_SOUND               sound id in misc value (SoundEntries.dbc)
	EffectPlayMusic                                   //132 SPELL_EFFECT_PLAY_MUSIC               sound id in misc value (SoundEntries.dbc)
	EffectUnlearnSpecialization                       //133 SPELL_EFFECT_UNLEARN_SPECIALIZATION   unlearn profession specialization
	EffectKillCredit                                  //134 SPELL_EFFECT_KILL_CREDIT              misc value is creature entry
	EffectCallPet                                     //135 SPELL_EFFECT_CALL_PET
	EffectHealPct                                     //136 SPELL_EFFECT_HEAL_PCT
	EffectEnergizePct                                 //137 SPELL_EFFECT_ENERGIZE_PCT
	EffectLeapBack                                    //138 SPELL_EFFECT_LEAP_BACK                Leap back
	EffectQuestClear                                  //139 SPELL_EFFECT_CLEAR_QUEST              Reset quest status (miscValue - quest ID)
	EffectForceCast                                   //140 SPELL_EFFECT_FORCE_CAST
	EffectForceCastWithValue                          //141 SPELL_EFFECT_FORCE_CAST_WITH_VALUE
	EffectTriggerSpellWithValue                       //142 SPELL_EFFECT_TRIGGER_SPELL_WITH_VALUE
	EffectApplyAreaAuraOwner                          //143 SPELL_EFFECT_APPLY_AREA_AURA_OWNER
	EffectKnockBackDest                               //144 SPELL_EFFECT_KNOCK_BACK_DEST
	EffectPullTowardsDest                             //145 SPELL_EFFECT_PULL_TOWARDS_DEST        Black Hole Effect
	EffectRestoreGarrisonTroopVitality                //146 SPELL_EFFECT_RESTORE_GARRISON_TROOP_VITALITY
	EffectQuestFail                                   //147 SPELL_EFFECT_QUEST_FAIL               quest fail
	EffectTriggerMissileSpell                         //148 SPELL_EFFECT_TRIGGER_MISSILE_SPELL_WITH_VALUE
	EffectChargeDest                                  //149 SPELL_EFFECT_CHARGE_DEST
	EffectQuestStart                                  //150 SPELL_EFFECT_QUEST_START
	EffectTriggerRitualOfSummoning                    //151 SPELL_EFFECT_TRIGGER_SPELL_2
	EffectSummonRaFFriend                             //152 SPELL_EFFECT_SUMMON_RAF_FRIEND        summon Refer-a-Friend
	EffectCreateTamedPet                              //153 SPELL_EFFECT_CREATE_TAMED_PET         misc value is creature entry
	EffectDiscoverTaxi                                //154 SPELL_EFFECT_DISCOVER_TAXI
	EffectTitanGrip                                   //155 SPELL_EFFECT_TITAN_GRIP Allows you to equip two-handed axes maces and swords in one hand but you attack $49152s1% slower than normal.
	EffectEnchantItemPrismatic                        //156 SPELL_EFFECT_ENCHANT_ITEM_PRISMATIC
	EffectCreateItem2                                 //157 SPELL_EFFECT_CREATE_ITEM_2            create item or create item template and replace by some randon spell loot item
	EffectMilling                                     //158 SPELL_EFFECT_MILLING                  milling
	EffectRenamePet                                   //159 SPELL_EFFECT_ALLOW_RENAME_PET         allow rename pet once again
	EffectForceCast2                                  //160 SPELL_EFFECT_FORCE_CAST_2
	EffectTalentSpecCount                             //161 SPELL_EFFECT_TALENT_SPEC_COUNT        second talent spec (learn/revert)
	EffectActivateSpec                                //162 SPELL_EFFECT_TALENT_SPEC_SELECT       activate primary/secondary spec
	EffectObliterateItem                              //163 SPELL_EFFECT_OBLITERATE_ITEM
	EffectRemoveAura                                  //164 SPELL_EFFECT_REMOVE_AURA
	EffectDamageFromMaxHealthPCT                      //165 SPELL_EFFECT_DAMAGE_FROM_MAX_HEALTH_PCT
	EffectGiveCurrency                                //166 SPELL_EFFECT_GIVE_CURRENCY
	EffectUpdatePlayerPhase                           //167 SPELL_EFFECT_UPDATE_PLAYER_PHASE
	EffectAllowControlPet                             //168 SPELL_EFFECT_ALLOW_CONTROL_PET
	EffectDestroyItem                                 //169 SPELL_EFFECT_DESTROY_ITEM
	EffectUpdateZoneAurasAndPhases                    //170 SPELL_EFFECT_UPDATE_ZONE_AURAS_AND_PHASES
	EffectSummonPersonalGameObject                    //171 SPELL_EFFECT_SUMMON_PERSONAL_GAMEOBJECT
	EffectResurrectWithAura                           //172 SPELL_EFFECT_RESURRECT_WITH_AURA
	EffectUnlockGuildVaultTab                         //173 SPELL_EFFECT_UNLOCK_GUILD_VAULT_TAB
	EffectApplyAuraOnPet                              //174 SPELL_EFFECT_APPLY_AURA_ON_PET
	Effect175                                         //175 SPELL_EFFECT_175
	EffectSanctuary2                                  //176 SPELL_EFFECT_SANCTUARY_2
	EffectDespawnPersistentAreaAura                   //177 SPELL_EFFECT_DESPAWN_PERSISTENT_AREA_AURA
	Effect178                                         //178 SPELL_EFFECT_178 unused
	EffectCreateAreaTrigger                           //179 SPELL_EFFECT_CREATE_AREATRIGGER
	EffectUpdateAreaTrigger                           //180 SPELL_EFFECT_UPDATE_AREATRIGGER
	EffectRemoveTalent                                //181 SPELL_EFFECT_REMOVE_TALENT
	EffectDespawnAreaTrigger                          //182 SPELL_EFFECT_DESPAWN_AREATRIGGER
	Effect183                                         //183 SPELL_EFFECT_183
	EffectReputation2                                 //184 SPELL_EFFECT_REPUTATION
	Effect185                                         //185 SPELL_EFFECT_185
	Effect186                                         //186 SPELL_EFFECT_186
	EffectRandomizeArchaeologyDigsites                //187 SPELL_EFFECT_RANDOMIZE_ARCHAEOLOGY_DIGSITES
	EffectSummonStabledPetAsGuardian                  //188 SPELL_EFFECT_SUMMON_STABLED_PET_AS_GUARDIAN
	EffectLoot                                        //189 SPELL_EFFECT_LOOT
	EffectChangePartyMembers                          //190 SPELL_EFFECT_CHANGE_PARTY_MEMBERS
	EffectTeleportToDigsite                           //191 SPELL_EFFECT_TELEPORT_TO_DIGSITE
	EffectUncageBattlePet                             //192 SPELL_EFFECT_UNCAGE_BATTLEPET
	EffectStartPetBattle                              //193 SPELL_EFFECT_START_PET_BATTLE
	Effect194                                         //194 SPELL_EFFECT_194
	EffectPlaySceneScriptPackage                      //195 SPELL_EFFECT_PLAY_SCENE_SCRIPT_PACKAGE
	EffectCreateSceneObject                           //196 SPELL_EFFECT_CREATE_SCENE_OBJECT
	EffectCreatePrivateSceneObject                    //197 SPELL_EFFECT_CREATE_PERSONAL_SCENE_OBJECT
	EffectPlayScene                                   //198 SPELL_EFFECT_PLAY_SCENE
	EffectDespawnSummon                               //199 SPELL_EFFECT_DESPAWN_SUMMON
	EffectHealBattlePetPct                            //200 SPELL_EFFECT_HEAL_BATTLEPET_PCT
	EffectEnableBattlePets                            //201 SPELL_EFFECT_ENABLE_BATTLE_PETS
	EffectApplyAreaAuraSummons                        //202 SPELL_EFFECT_APPLY_AREA_AURA_SUMMONS
	EffectRemoveAura2                                 //203 SPELL_EFFECT_REMOVE_AURA_2
	EffectChangeBattlePetQuality                      //204 SPELL_EFFECT_CHANGE_BATTLEPET_QUALITY
	EffectLaunchQuestChoice                           //205 SPELL_EFFECT_LAUNCH_QUEST_CHOICE
	EffectAlterItem                                   //206 SPELL_EFFECT_ALTER_ITEM
	EffectLaunchQuestTask                             //207 SPELL_EFFECT_LAUNCH_QUEST_TASK
	EffectSetReputation                               //208 SPELL_EFFECT_SET_REPUTATION
	Effect209                                         //209 SPELL_EFFECT_209
	EffectLearnGarrisonBuilding                       //210 SPELL_EFFECT_LEARN_GARRISON_BUILDING
	EffectLearnGarrisonSpecialization                 //211 SPELL_EFFECT_LEARN_GARRISON_SPECIALIZATION
	EffectRemoveAuraBySpellLabel                      //212 SPELL_EFFECT_REMOVE_AURA_BY_SPELL_LABEL
	EffectJumpDest2                                   //213 SPELL_EFFECT_JUMP_DEST_2
	EffectCreateGarrison                              //214 SPELL_EFFECT_CREATE_GARRISON
	EffectUpgradeCharacterSpells                      //215 SPELL_EFFECT_UPGRADE_CHARACTER_SPELLS
	EffectCreateShipment                              //216 SPELL_EFFECT_CREATE_SHIPMENT
	EffectUpgradeGarrison                             //217 SPELL_EFFECT_UPGRADE_GARRISON
	Effect218                                         //218 SPELL_EFFECT_218
	EffectCreateConversation                          //219 SPELL_EFFECT_CREATE_CONVERSATION
	EffectAddGarrisonFollower                         //220 SPELL_EFFECT_ADD_GARRISON_FOLLOWER
	EffectAddGarrisonMission                          //221 SPELL_EFFECT_ADD_GARRISON_MISSION
	EffectCreateHeirloomItem                          //222 SPELL_EFFECT_CREATE_HEIRLOOM_ITEM
	EffectChangeItemBonuses                           //223 SPELL_EFFECT_CHANGE_ITEM_BONUSES
	EffectActivateGarrisonBuilding                    //224 SPELL_EFFECT_ACTIVATE_GARRISON_BUILDING
	EffectGrantBattlePetLevel                         //225 SPELL_EFFECT_GRANT_BATTLEPET_LEVEL
	EffectTriggerActionSet                            //226 SPELL_EFFECT_TRIGGER_ACTION_SET
	EffectTeleportToLFGDungeon                        //227 SPELL_EFFECT_TELEPORT_TO_LFG_DUNGEON
	Effect228                                         //228 SPELL_EFFECT_228
	EffectSetFollowerQuality                          //229 SPELL_EFFECT_SET_FOLLOWER_QUALITY
	Effect230                                         //230 SPELL_EFFECT_230
	EffectIncreaseFollowerExperience                  //231 SPELL_EFFECT_INCREASE_FOLLOWER_EXPERIENCE
	EffectRemovePhase                                 //232 SPELL_EFFECT_REMOVE_PHASE
	EffectRandomizeFollowerAbilities                  //233 SPELL_EFFECT_RANDOMIZE_FOLLOWER_ABILITIES
	Effect234                                         //234 SPELL_EFFECT_234
	Effect235                                         //235 SPELL_EFFECT_235
	EffectGiveExperience                              //236 SPELL_EFFECT_GIVE_EXPERIENCE
	EffectGiveRestedExperienceBonus                   //237 SPELL_EFFECT_GIVE_RESTED_EXPERIENCE_BONUS
	EffectIncreaseSkill                               //238 SPELL_EFFECT_INCREASE_SKILL
	EffectEndGarrisonBuildingConstruction             //239 SPELL_EFFECT_END_GARRISON_BUILDING_CONSTRUCTION
	EffectGiveArtifactPower                           //240 SPELL_EFFECT_GIVE_ARTIFACT_POWER
	Effect241                                         //241 SPELL_EFFECT_241
	EffectGiveArtifactPowerNoBonus                    //242 SPELL_EFFECT_GIVE_ARTIFACT_POWER_NO_BONUS
	EffectApplyEnchantIllusion                        //243 SPELL_EFFECT_APPLY_ENCHANT_ILLUSION
	EffectLearnFollowerAbility                        //244 SPELL_EFFECT_LEARN_FOLLOWER_ABILITY
	EffectUpgradeHeirloom                             //245 SPELL_EFFECT_UPGRADE_HEIRLOOM
	EffectFinishGarrisonMission                       //246 SPELL_EFFECT_FINISH_GARRISON_MISSION
	EffectAddGarrisonMissionSet                       //247 SPELL_EFFECT_ADD_GARRISON_MISSION_SET
	EffectFinishShipment                              //248 SPELL_EFFECT_FINISH_SHIPMENT
	EffectForceEquipItem                              //249 SPELL_EFFECT_FORCE_EQUIP_ITEM
	EffectTakeScreenshot                              //250 SPELL_EFFECT_TAKE_SCREENSHOT
	EffectSetGarrisonCacheSize                        //251 SPELL_EFFECT_SET_GARRISON_CACHE_SIZE
	EffectTeleportUnits                               //252 SPELL_EFFECT_TELEPORT_UNITS
	EffectGiveHonor                                   //253 SPELL_EFFECT_GIVE_HONOR
	EffectJumpCharge                                  //254 SPELL_EFFECT_JUMP_CHARGE
	EffectLearnTransmogSet                            //255 SPELL_EFFECT_LEARN_TRANSMOG_SET
	Effect256                                         //256 SPELL_EFFECT_256
	Effect257                                         //257 SPELL_EFFECT_257
	EffectModifyKeystone                              //258 SPELL_EFFECT_MODIFY_KEYSTONE
	EffectRespecAzeriteEmpoweredItem                  //259 SPELL_EFFECT_RESPEC_AZERITE_EMPOWERED_ITEM
	EffectSummonStabledPet                            //260 SPELL_EFFECT_SUMMON_STABLED_PET
	EffectScrapItem                                   //261 SPELL_EFFECT_SCRAP_ITEM
	Effect262                                         //262 SPELL_EFFECT_262
	EffectRepairItem                                  //263 SPELL_EFFECT_REPAIR_ITEM
	EffectRemoveGem                                   //264 SPELL_EFFECT_REMOVE_GEM
	EffectLearnAzeriteEssencePower                    //265 SPELL_EFFECT_LEARN_AZERITE_ESSENCE_POWER
	EffectSetItemBonusListGroupEntry                  //266 SPELL_EFFECT_SET_ITEM_BONUS_LIST_GROUP_ENTRY
	EffectCreatePrivateConversation                   //267 SPELL_EFFECT_CREATE_PRIVATE_CONVERSATION
	EffectApplyMountEquipment                         //268 SPELL_EFFECT_APPLY_MOUNT_EQUIPMENT
	EffectIncreaseItemBonusListGroupStep              //269 SPELL_EFFECT_INCREASE_ITEM_BONUS_LIST_GROUP_STEP
	Effect270                                         //270 SPELL_EFFECT_270
	EffectApplyAreaAuraPartyNonRandom                 //271 SPELL_EFFECT_APPLY_AREA_AURA_PARTY_NONRANDOM
	EffectSetCovenant                                 //272 SPELL_EFFECT_SET_COVENANT
	EffectCraftRuneforgeLegendary                     //273 SPELL_EFFECT_CRAFT_RUNEFORGE_LEGENDARY
	Effect274                                         //274 SPELL_EFFECT_274
	Effect275                                         //275 SPELL_EFFECT_275
	EffectLearnTransmogIllusion                       //276 SPELL_EFFECT_LEARN_TRANSMOG_ILLUSION
	EffectSetChromieTime                              //277 SPELL_EFFECT_SET_CHROMIE_TIME
	Effect278                                         //278 SPELL_EFFECT_278
	EffectLearnGarrTalent                             //279 SPELL_EFFECT_LEARN_GARR_TALENT
	Effect280                                         //280 SPELL_EFFECT_280
	EffectLearnSoulbindConduit                        //281 SPELL_EFFECT_LEARN_SOULBIND_CONDUIT
	EffectConvertItemsToCurrency                      //282 SPELL_EFFECT_CONVERT_ITEMS_TO_CURRENCY
	EffectCompleteCampaign                            //283 SPELL_EFFECT_COMPLETE_CAMPAIGN
	EffectSendChatMessage                             //284 SPELL_EFFECT_SEND_CHAT_MESSAGE
	EffectModifyKeystone2                             //285 SPELL_EFFECT_MODIFY_KEYSTONE_2
	EffectGrantBattlePetExperience                    //286 SPELL_EFFECT_GRANT_BATTLEPET_EXPERIENCE
	EffectSetGarrisonFollowerLevel                    //287 SPELL_EFFECT_SET_GARRISON_FOLLOWER_LEVEL
	Effect288                                         //288 SPELL_EFFECT_288
	Effect289                                         //289 SPELL_EFFECT_289
	NumEffects
)
