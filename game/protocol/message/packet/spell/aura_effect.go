package spell

type AuraEffect uint32

const (
	AuraEffectNone                                     AuraEffect = iota // 0
	AuraEffectBindSight                                                  // 1
	AuraEffectModPossess                                                 // 2
	AuraEffectPeriodicDamage                                             // 3
	AuraEffectDummy                                                      // 4
	AuraEffectModConfuse                                                 // 5
	AuraEffectModCharm                                                   // 6
	AuraEffectModFear                                                    // 7
	AuraEffectPeriodicHeal                                               // 8
	AuraEffectModAttackspeed                                             // 9
	AuraEffectModThreat                                                  // 10
	AuraEffectModTaunt                                                   // 11
	AuraEffectModStun                                                    // 12
	AuraEffectModDamageDone                                              // 13
	AuraEffectModDamageTaken                                             // 14
	AuraEffectDamageShield                                               // 15
	AuraEffectModStealth                                                 // 16
	AuraEffectModStealth_DETECT                                          // 17
	AuraEffectModInvisibility                                            // 18
	AuraEffectModInvisibilityDetect                                      // 19
	AuraEffectObsModHealth                                               // 20 // 20 21 unofficial
	AuraEffectObsModPower                                                // 21
	AuraEffectModResistance                                              // 22
	AuraEffectPeriodicTriggerSpell                                       // 23
	AuraEffectPeriodicEnergize                                           // 24
	AuraEffectModPacify                                                  // 25
	AuraEffectModRoot                                                    // 26
	AuraEffectModSilence                                                 // 27
	AuraEffectReflectSpells                                              // 28
	AuraEffectModStat                                                    // 29
	AuraEffectModSkill                                                   // 30
	AuraEffectModIncreaseSpeed                                           // 31
	AuraEffectModIncreaseMountedSpeed                                    // 32
	AuraEffectModDecreaseSpeed                                           // 33
	AuraEffectModIncreaseHealth                                          // 34
	AuraEffectModIncreaseEnergy                                          // 35
	AuraEffectModShapeshift                                              // 36
	AuraEffectEffectImmunity                                             // 37
	AuraEffectStateImmunity                                              // 38
	AuraEffectSchoolImmunity                                             // 39
	AuraEffectDamageImmunity                                             // 40
	AuraEffectDispelImmunity                                             // 41
	AuraEffectProcTriggerSpell                                           // 42
	AuraEffectProcTriggerDamage                                          // 43
	AuraEffectTrackCreatures                                             // 44
	AuraEffectTrackResources                                             // 45
	AuraEffect46                                                         // 46 // Ignore all Gear test spells
	AuraEffectModParryPercent                                            // 47
	AuraEffectPeriodicTriggerSpellFromClient                             // 48 // One periodic spell
	AuraEffectModDodgePercent                                            // 49
	AuraEffectModCriticalHealingAmount                                   // 50
	AuraEffectModBlockPercent                                            // 51
	AuraEffectModWeaponCritPercent                                       // 52
	AuraEffectPeriodicLeech                                              // 53
	AuraEffectModHitChance                                               // 54
	AuraEffectModSpellHitChance                                          // 55
	AuraEffectTransform                                                  // 56
	AuraEffectModSpellCritChance                                         // 57
	AuraEffectModIncreaseSwimSpeed                                       // 58
	AuraEffectModDamageDoneCreature                                      // 59
	AuraEffectModPacifySilence                                           // 60
	AuraEffectModScale                                                   // 61
	AuraEffectPeriodicHealthFunnel                                       // 62
	AuraEffectModAdditionalPowerCost                                     // 63
	AuraEffectPeriodicManaLeech                                          // 64
	AuraEffectModCastingSpeed_NOT_STACK                                  // 65
	AuraEffectFeignDeath                                                 // 66
	AuraEffectModDisarm                                                  // 67
	AuraEffectModStalked                                                 // 68
	AuraEffectSchoolAbsorb                                               // 69
	AuraEffectPeriodicWeaponPercentDamage                                // 70
	AuraEffectStoreTeleportReturnPoint                                   // 71
	AuraEffectModPowerCostSchoolPct                                      // 72
	AuraEffectModPowerCostSchool                                         // 73
	AuraEffectReflectSpells_SCHOOL                                       // 74
	AuraEffectModLanguage                                                // 75
	AuraEffectFarSight                                                   // 76
	AuraEffectMechanicImmunity                                           // 77
	AuraEffectMounted                                                    // 78
	AuraEffectModDamagePercentDone                                       // 79
	AuraEffectModPercentStat                                             // 80
	AuraEffectSplitDamagePct                                             // 81
	AuraEffectWaterBreathing                                             // 82
	AuraEffectModBaseResistance                                          // 83
	AuraEffectModRegen                                                   // 84
	AuraEffectModPowerRegen                                              // 85
	AuraEffectChannelDeathItem                                           // 86
	AuraEffectModDamagePercentTaken                                      // 87
	AuraEffectModHealthRegenPercent                                      // 88
	AuraEffectPeriodicDamagePercent                                      // 89
	AuraEffect90                                                         // 90 // old SPELL_AURA_MOD_RESIST_CHANCE
	AuraEffectModDetectRange                                             // 91
	AuraEffectPreventsFleeing                                            // 92
	AuraEffectModUnattackable                                            // 93
	AuraEffectInterruptRegen                                             // 94
	AuraEffectGhost                                                      // 95
	AuraEffectSpellMagnet                                                // 96
	AuraEffectManaShield                                                 // 97
	AuraEffectModSkillTalent                                             // 98
	AuraEffectModAttackPower                                             // 99
	AuraEffectAurasVisible                                               // 100
	AuraEffectModResistancePct                                           // 101
	AuraEffectModMeleeAttackPowerVersus                                  // 102
	AuraEffectModTotalThreat                                             // 103
	AuraEffectWaterWalk                                                  // 104
	AuraEffectFeatherFall                                                // 105
	AuraEffectHover                                                      // 106
	AuraEffectAddFlatModifier                                            // 107
	AuraEffectAddPctModifier                                             // 108
	AuraEffectAddTargetTrigger                                           // 109
	AuraEffectModPowerRegenPercent                                       // 110
	AuraEffectInterceptMeleeRangedAttacks                                // 111
	AuraEffectOverrideClassScripts                                       // 112
	AuraEffectModRangedDamageTaken                                       // 113
	AuraEffectModRangedDamageTaken_PCT                                   // 114
	AuraEffectModHealing                                                 // 115
	AuraEffectModRegen_DURING_COMBAT                                     // 116
	AuraEffectModMechanicResistance                                      // 117
	AuraEffectModHealing_PCT                                             // 118
	AuraEffectPvpTalents                                                 // 119
	AuraEffectUntrackable                                                // 120
	AuraEffectEmpathy                                                    // 121
	AuraEffectModOffhandDamagePct                                        // 122
	AuraEffectModTargetResistance                                        // 123
	AuraEffectModRangedAttackPower                                       // 124
	AuraEffectModMeleeDamageTaken                                        // 125
	AuraEffectModMeleeDamageTaken_PCT                                    // 126
	AuraEffectRangedAttackPowerAttackerBonus                             // 127
	AuraEffectModFixate                                                  // 128
	AuraEffectModSpeedAlways                                             // 129
	AuraEffectModMountedSpeedAlways                                      // 130
	AuraEffectModRangedAttackPowerVersus                                 // 131
	AuraEffectModIncreaseEnergy_PERCENT                                  // 132
	AuraEffectModIncreaseHealth_PERCENT                                  // 133
	AuraEffectModManaRegenInterrupt                                      // 134
	AuraEffectModHealingDone                                             // 135
	AuraEffectModHealingDone_PERCENT                                     // 136
	AuraEffectModTotalStatPercentage                                     // 137
	AuraEffectModMeleeHaste                                              // 138
	AuraEffectForceReaction                                              // 139
	AuraEffectModRangedHaste                                             // 140
	AuraEffect141                                                        // 141 // old SPELL_AURA_MOD_RANGED_AMMO_HASTE unused now
	AuraEffectModBaseResistance_PCT                                      // 142
	AuraEffectModRecoveryRate_BY_SPELL_LABEL                             // 143 // NYI
	AuraEffectSafeFall                                                   // 144
	AuraEffectModIncreaseHealthPercent2                                  // 145
	AuraEffectAllowTamePetType                                           // 146
	AuraEffectMechanicImmunity_MASK                                      // 147
	AuraEffectModChargeRecoveryRate                                      // 148 // NYI
	AuraEffectReducePushback                                             // 149 //    Reduce Pushback
	AuraEffectModShieldBlockvaluePct                                     // 150
	AuraEffectTrackStealthed                                             // 151 //    Track Stealthed
	AuraEffectModDetectedRange                                           // 152 //    Mod Detected Range
	AuraEffectModAutoattackRange                                         // 153
	AuraEffectModStealth_LEVEL                                           // 154 //    Stealth Level Modifier
	AuraEffectModWaterBreathing                                          // 155 //    Mod Water Breathing
	AuraEffectModReputationGain                                          // 156 //    Mod Reputation Gain
	AuraEffectPetDamageMulti                                             // 157 //    Mod Pet Damage
	AuraEffectAllowTalentSwapping                                        // 158
	AuraEffectNoPvpCredit                                                // 159
	AuraEffect160                                                        // 160 // old SPELL_AURA_MOD_AOE_AVOIDANCE. Unused 4.3.4
	AuraEffectModHealthRegenInCombat                                     // 161
	AuraEffectPowerBurn                                                  // 162
	AuraEffectModCritDamageBonus                                         // 163
	AuraEffectForceBreathBar                                             // 164
	AuraEffectMeleeAttackPowerAttackerBonus                              // 165
	AuraEffectModAttackPowerPct                                          // 166
	AuraEffectModRangedAttackPowerPct                                    // 167
	AuraEffectModDamageDoneVersus                                        // 168
	AuraEffectSetFfaPvp                                                  // 169
	AuraEffectDetectAmore                                                // 170
	AuraEffectModSpeedNotStack                                           // 171
	AuraEffectModMountedSpeedNotStack                                    // 172
	AuraEffectModRecoveryRate2                                           // 173 // NYI
	AuraEffectModSpellDamageOfStatPercent                                // 174 // by defeult intelect dependent from AuraEffectModSpellHealingOfStatPercent
	AuraEffectModSpellHealingOfStatPercent                               // 175
	AuraEffectSpiritOfRedemption                                         // 176
	AuraEffectAoeCharm                                                   // 177
	AuraEffectModMaxPowerPct                                             // 178
	AuraEffectModPowerDisplay                                            // 179
	AuraEffectModFlatSpellDamageVersus                                   // 180
	AuraEffectModSpellCurrencyReagentsCountPct                           // 181 // NYI
	AuraEffectSuppressItemPassiveEffectBySpellLabel                      // 182
	AuraEffectModCritChanceVersusTargetHealth                            // 183
	AuraEffectModAttackerMeleeHitChance                                  // 184
	AuraEffectModAttackerRangedHitChance                                 // 185
	AuraEffectModAttackerSpellHitChance                                  // 186
	AuraEffectModAttackerMeleeCritChance                                 // 187
	AuraEffectModAttackerRangedCritChance                                // 188
	AuraEffectModRating                                                  // 189
	AuraEffectModFactionReputationGain                                   // 190
	AuraEffectUseNormalMovementSpeed                                     // 191
	AuraEffectModMeleeRangedHaste                                        // 192
	AuraEffectMeleeSlow                                                  // 193
	AuraEffectModTargetAbsorbSchool                                      // 194
	AuraEffectLearnSpell                                                 // 195
	AuraEffectModCooldown                                                // 196 // only 24818 Noxious Breath
	AuraEffectModAttackerSpellAndWeaponCritChance                        // 197
	AuraEffectModCombatRatingFromCombatRating                            // 198
	AuraEffect199                                                        // 199 // old SPELL_AURA_MOD_INCREASES_SPELL_PCT_TO_HIT. unused 4.3.4
	AuraEffectModXpPct                                                   // 200
	AuraEffectFly                                                        // 201
	AuraEffectIgnoreCombatResult                                         // 202
	AuraEffectPreventInterrupt                                           // 203 // NYI
	AuraEffectPreventCorpseRelease                                       // 204 // NYI
	AuraEffectModChargeCooldown                                          // 205 // NYI
	AuraEffectModIncreaseVehicleFlightSpeed                              // 206
	AuraEffectModIncreaseMountedFlightSpeed                              // 207
	AuraEffectModIncreaseFlightSpeed                                     // 208
	AuraEffectModMountedFlightSpeedAlways                                // 209
	AuraEffectModVehicleSpeedAlways                                      // 210
	AuraEffectModFlightSpeedNotStack                                     // 211
	AuraEffectModHonorGainPct                                            // 212
	AuraEffectModRageFromDamageDealt                                     // 213
	AuraEffect214                                                        // 214
	AuraEffectArenaPreparation                                           // 215
	AuraEffectHasteSpells                                                // 216
	AuraEffectModMeleeHaste_2                                            // 217
	AuraEffectAddPctModifierBySpellLabel                                 // 218
	AuraEffectAddFlatModifier_BY_SPELL_LABEL                             // 219
	AuraEffectModAbilitySchoolMask                                       // 220 // NYI
	AuraEffectModDetaunt                                                 // 221
	AuraEffectRemoveTransmogCost                                         // 222
	AuraEffectRemoveBarberShopCost                                       // 223
	AuraEffectLearnTalent                                                // 224 // NYI
	AuraEffectModVisibilityRange                                         // 225
	AuraEffectPeriodicDummy                                              // 226
	AuraEffectPeriodicTriggerSpell_WITH_VALUE                            // 227
	AuraEffectDetectStealth                                              // 228
	AuraEffectModAoeDamageAvoidance                                      // 229
	AuraEffectModMaxHealth                                               // 230
	AuraEffectProcTriggerSpell_WITH_VALUE                                // 231
	AuraEffectMechanicDurationMod                                        // 232
	AuraEffectChangeModelForAllHumanoids                                 // 233 // client-side only
	AuraEffectMechanicDurationMod_NOT_STACK                              // 234
	AuraEffectModHoverNoHeightOffset                                     // 235
	AuraEffectControlVehicle                                             // 236
	AuraEffect237                                                        // 237
	AuraEffect238                                                        // 238
	AuraEffectModScale2                                                  // 239
	AuraEffectModExpertise                                               // 240
	AuraEffectForceMoveForward                                           // 241
	AuraEffectModSpellDamageFromHealing                                  // 242
	AuraEffectModFaction                                                 // 243
	AuraEffectComprehendLanguage                                         // 244
	AuraEffectModAuraDurationByDispel                                    // 245
	AuraEffectModAuraDurationByDispel_NOT_STACK                          // 246
	AuraEffectCloneCaster                                                // 247
	AuraEffectModCombatResultChance                                      // 248
	AuraEffectModDamagePercentDoneByTargetAuraMechanic                   // 249 // NYI
	AuraEffectModIncreaseHealth_2                                        // 250
	AuraEffectModEnemyDodge                                              // 251
	AuraEffectModSpeedSlowAll                                            // 252
	AuraEffectModBlockCritChance                                         // 253
	AuraEffectModDisarm_OFFHAND                                          // 254
	AuraEffectModMechanicDamageTakenPercent                              // 255
	AuraEffectNoReagentUse                                               // 256
	AuraEffectModTargetResistBySpellClass                                // 257
	AuraEffectOverrideSummonedObject                                     // 258
	AuraEffectModHotPct                                                  // 259
	AuraEffectScreenEffect                                               // 260
	AuraEffectPhase                                                      // 261
	AuraEffectAbilityIgnoreAurastate                                     // 262
	AuraEffectDisableCastingExceptAbilities                              // 263
	AuraEffectDisableAttackingExceptAbilities                            // 264
	AuraEffect265                                                        // 265
	AuraEffectSetVignette                                                // 266 // NYI
	AuraEffectModImmuneAuraApplySchool                                   // 267
	AuraEffectModArmorPctFromStat                                        // 268
	AuraEffectModIgnoreTargetResist                                      // 269
	AuraEffectModSchoolMaskDamageFromCaster                              // 270
	AuraEffectModSpellDamageFromCaster                                   // 271
	AuraEffectModBlockValuePct                                           // 272 // NYI
	AuraEffectXRay                                                       // 273
	AuraEffectModBlockValueFlat                                          // 274 // NYI
	AuraEffectModIgnoreShapeshift                                        // 275
	AuraEffectModDamageDoneForMechanic                                   // 276
	AuraEffect277                                                        // 277 // old SPELL_AURA_MOD_MAX_AFFECTED_TARGETS. unused 4.3.4
	AuraEffectModDisarmRanged                                            // 278
	AuraEffectInitializeImages                                           // 279
	AuraEffect280                                                        // 280 // old SPELL_AURA_MOD_ARMOR_PENETRATION_PCT unused 4.3.4
	AuraEffectProvideSpellFocus                                          // 281
	AuraEffectModBaseHealthPct                                           // 282
	AuraEffectModHealing_RECEIVED                                        // 283 // Possibly only for some spell family class spells
	AuraEffectLinked                                                     // 284
	AuraEffectLinked2                                                    // 285
	AuraEffectModRecoveryRate                                            // 286
	AuraEffectDeflectSpells                                              // 287
	AuraEffectIgnoreHitDirection                                         // 288
	AuraEffectPreventDurabilityLoss                                      // 289
	AuraEffectModCritPct                                                 // 290
	AuraEffectModXpQuestPct                                              // 291
	AuraEffectOpenStable                                                 // 292
	AuraEffectOverrideSpells                                             // 293
	AuraEffectPreventRegeneratePower                                     // 294
	AuraEffectModPeriodicDamageTaken                                     // 295
	AuraEffectSetVehicleId                                               // 296
	AuraEffectModRootDisableGravity                                      // 297 // NYI
	AuraEffectModStunDisableGravity                                      // 298 // NYI
	AuraEffect299                                                        // 299
	AuraEffectShareDamagePct                                             // 300
	AuraEffectSchoolHealAbsorb                                           // 301
	AuraEffect302                                                        // 302
	AuraEffectModDamageDoneVersusAurastate                               // 303
	AuraEffectModFakeInebriate                                           // 304
	AuraEffectModMinimumSpeed                                            // 305
	AuraEffectModCritChanceForCaster                                     // 306
	AuraEffectCastWhileWalkingBySpellLabel                               // 307
	AuraEffectModCritChanceForCasterWithAbilities                        // 308
	AuraEffectModResilience                                              // 309 // NYI
	AuraEffectModCreatureAoeDamageAvoidance                              // 310
	AuraEffectIgnoreCombat                                               // 311 // NYI
	AuraEffectAnimReplacementSet                                         // 312
	AuraEffectMountAnimReplacementSet                                    // 313
	AuraEffectPreventResurrection                                        // 314
	AuraEffectUnderwaterWalking                                          // 315
	AuraEffectSchoolAbsorbOverkill                                       // 316 // NYI - absorbs overkill damage
	AuraEffectModSpellPowerPct                                           // 317
	AuraEffectMastery                                                    // 318
	AuraEffectModMeleeHaste_3                                            // 319
	AuraEffect320                                                        // 320
	AuraEffectModNoActions                                               // 321
	AuraEffectInterfereTargetting                                        // 322
	AuraEffect323                                                        // 323 // Not used in 4.3.4
	AuraEffectOverrideUnlockedAzeriteEssenceRank                         // 324 // testing aura
	AuraEffectLearnPvpTalent                                             // 325 // NYI
	AuraEffectPhaseGroup                                                 //= 326             // Puts the player in all the phases that are in the group with id // miscB
	AuraEffectPhaseAlwaysVisible                                         // 327 // Sets PhaseShiftFlags::AlwaysVisible
	AuraEffectTriggerSpellOnPowerPct                                     //= 328             // Triggers spell when power goes above (MiscB = 0) or falls below (MiscB // 1) specified percent value (once not every time condition has meet)
	AuraEffectModPowerGainPct                                            // 329
	AuraEffectCastWhileWalking                                           // 330
	AuraEffectForceWeather                                               // 331
	AuraEffectOverrideActionbarSpells                                    // 332
	AuraEffectOverrideActionbarSpells_TRIGGERED                          // 333 // Spells cast with this override have no cast time or power cost
	AuraEffectModAutoattackCritChance                                    // 334
	AuraEffect335                                                        // 335
	AuraEffectMountRestrictions                                          // 336
	AuraEffectModVendorItemsPrices                                       // 337
	AuraEffectModDurabilityLoss                                          // 338
	AuraEffectModCritChanceForCaster_PET                                 // 339
	AuraEffectModResurrectedHealthByGuildMember                          // 340 // Increases health gained when resurrected by a guild member by X
	AuraEffectModSpellCategoryCooldown                                   // 341 // Modifies cooldown of all spells using affected category
	AuraEffectModMeleeRangedHaste_2                                      // 342
	AuraEffectModMeleeDamageFromCaster                                   // 343
	AuraEffectModAutoattackDamage                                        // 344
	AuraEffectBypassArmorForCaster                                       // 345
	AuraEffectEnableAltPower                                             // 346
	AuraEffectModSpellCooldownByHaste                                    // 347
	AuraEffectModMoneyGain                                               //  = 348             // Modifies gold gains from source: [Misc = 0 Quests][Misc // 1 Loot]
	AuraEffectModCurrencyGain                                            // 349
	AuraEffect350                                                        // 350
	AuraEffectModCurrencyCategoryGainPct                                 // 351 // NYI
	AuraEffect352                                                        // 352
	AuraEffectModCamouflage                                              // 353 // NYI
	AuraEffectModHealingDone_PCT_VERSUS_TARGET_HEALTH                    // = 354             // Restoration Shaman mastery - mod healing based on target's health (less // more healing)
	AuraEffectModCastingSpeed                                            // 355 // NYI
	AuraEffectProvideTotemCategory                                       // 356
	AuraEffectEnableBoss1UnitFrame                                       // 357
	AuraEffectWorgenAlteredForm                                          // 358
	AuraEffectModHealingDone_VERSUS_AURASTATE                            // 359
	AuraEffectProcTriggerSpell_COPY                                      // 360 // Procs the same spell that caused this proc (Dragonwrath Tarecgosa's Rest)
	AuraEffectOverrideAutoattackWithMeleeSpell                           // 361
	AuraEffect362                                                        // 362 // Not used in 4.3.4
	AuraEffectModNextSpell                                               // 363 // Used by 101601 Throw Totem - causes the client to initialize spell cast with specified spell
	AuraEffect364                                                        // 364 // Not used in 4.3.4
	AuraEffectMaxFarClipPlane                                            // 365 // Overrides client's View Distance setting to max("Fair" current_setting) and turns off terrain display
	AuraEffectOverrideSpellPowerByApPct                                  // 366 // NYI - Sets spellpower equal to % of attack power discarding all other bonuses (from gear and buffs)
	AuraEffectOverrideAutoattackWithRangedSpell                          // 367 // NYI
	AuraEffect368                                                        // 368 // Not used in 4.3.4
	AuraEffectEnablePowerBarTimer                                        // 369
	AuraEffectSpellOverrideNameGroup                                     // 370 // picks a random SpellOverrideName id from a group (group id in miscValue)
	AuraEffect371                                                        // 371
	AuraEffectOverrideMountFromSet                                       // 372 // NYI
	AuraEffectModSpeedNoControl                                          // 373 // NYI
	AuraEffectModifyFallDamagePct                                        // 374
	AuraEffectHideModelAndEquipementSlots                                // 375
	AuraEffectModCurrencyGainFromSource                                  // 376 // NYI
	AuraEffectCastWhileWalking_ALL                                       // 377 // Enables casting all spells while moving
	AuraEffectModPossess_PET                                             // 378
	AuraEffectModManaRegenPct                                            // 379
	AuraEffect380                                                        // 380
	AuraEffectModDamageTakenFromCasterPet                                // 381 // NYI
	AuraEffectModPetStatPct                                              // 382 // NYI
	AuraEffectIgnoreSpellCooldown                                        // 383 // NYI
	AuraEffect384                                                        // 384
	AuraEffect385                                                        // 385
	AuraEffect386                                                        // 386
	AuraEffect387                                                        // 387
	AuraEffectModTaxiFlightSpeed                                         // 388 // NYI
	AuraEffect389                                                        // 389
	AuraEffect390                                                        // 390
	AuraEffect391                                                        // 391
	AuraEffect392                                                        // 392
	AuraEffectBlockSpellsInFront                                         // 393 // NYI
	AuraEffectShowConfirmationPrompt                                     // 394
	AuraEffectAreaTrigger                                                // 395 // NYI
	AuraEffectTriggerSpellOnPowerAmount                                  //  = 396             // Triggers spell when power goes above (MiscA = 0) or falls below (MiscA // 1) specified percent value (once not every time condition has meet)
	AuraEffectBattlegroundPlayerPosition_FACTIONAL                       // 397
	AuraEffectBattlegroundPlayerPosition                                 // 398
	AuraEffectModTimeRate                                                // 399
	AuraEffectModSkill_2                                                 // 400
	AuraEffect401                                                        // 401
	AuraEffectModOverridePowerDisplay                                    // 402
	AuraEffectOverrideSpellVisual                                        // 403
	AuraEffectOverrideAttackPowerBySpPct                                 // 404
	AuraEffectModRatingPct                                               // 405
	AuraEffectKeyboundOverride                                           // 406 // NYI
	AuraEffectModFear2                                                   // 407 // NYI
	AuraEffectSetActionButtonSpellCount                                  // 408
	AuraEffectCanTurnWhileFalling                                        // 409
	AuraEffect410                                                        // 410
	AuraEffectModMaxCharges                                              // 411
	AuraEffect412                                                        // 412
	AuraEffectModRangedAttackDeflectChance                               // 413 // NYI
	AuraEffectModRangedAttackBlockChanceInFront                          // 414 // NYI
	AuraEffect415                                                        // 415
	AuraEffectModCooldownByHasteRegen                                    // 416
	AuraEffectModGlobalCooldownByHasteRegen                              // 417
	AuraEffectModMaxPower                                                // 418 // NYI
	AuraEffectModBaseManaPct                                             // 419
	AuraEffectModBattlePetXpPct                                          // 420
	AuraEffectModAbsorbEffectsDonePct                                    // 421 // NYI
	AuraEffectModAbsorbEffectsTakenPct                                   // 422 // NYI
	AuraEffectModManaCostPct                                             // 423
	AuraEffectCasterIgnoreLos                                            // 424 // NYI
	AuraEffect425                                                        // 425
	AuraEffect426                                                        // 426
	AuraEffectScalePlayerLevel                                           // 427 // NYI
	AuraEffectLinked_SUMMON                                              // 428
	AuraEffectModSummonDamage                                            // 429 // NYI - increases damage done by all summons not just controlled pets
	AuraEffectPlayScene                                                  // 430
	AuraEffectModOverrideZonePvpType                                     // 431 // NYI
	AuraEffect432                                                        // 432
	AuraEffect433                                                        // 433
	AuraEffect434                                                        // 434
	AuraEffect435                                                        // 435
	AuraEffectModEnvironmentalDamageTaken                                // 436
	AuraEffectModMinimumSpeed_RATE                                       // 437
	AuraEffectPreloadPhase                                               // 438 // NYI
	AuraEffect439                                                        // 439
	AuraEffectModMultistrikeDamage                                       // 440 // NYI
	AuraEffectModMultistrikeChance                                       // 441 // NYI
	AuraEffectModReadiness                                               // 442 // NYI
	AuraEffectModLeech                                                   // 443 // NYI
	AuraEffect444                                                        // 444
	AuraEffect445                                                        // 445
	AuraEffect446                                                        // 446
	AuraEffectModXpFromCreatureType                                      // 447
	AuraEffect448                                                        // 448
	AuraEffect449                                                        // 449
	AuraEffect450                                                        // 450
	AuraEffectOverridePetSpecs                                           // 451
	AuraEffect452                                                        // 452
	AuraEffectChargeRecoveryMod                                          // 453
	AuraEffectChargeRecoveryMultiplier                                   // 454
	AuraEffectModRoot_2                                                  // 455
	AuraEffectChargeRecoveryAffectedByHaste                              // 456
	AuraEffectChargeRecoveryAffectedByHasteRegen                         // 457
	AuraEffectIgnoreDualWieldHitPenalty                                  // 458
	AuraEffectIgnoreMovementForces                                       // 459
	AuraEffectResetCooldownsOnDuelStart                                  // 460 // NYI
	AuraEffect461                                                        // 461
	AuraEffectModHealing_AND_ABSORB_FROM_CASTER                          // 462 // NYI
	AuraEffectConvertCritRatingPctToParryRating                          // 463 // NYI
	AuraEffectModAttackPowerOfBonusArmor                                 // 464 // NYI
	AuraEffectModBonusArmor                                              // 465
	AuraEffectModBonusArmor_PCT                                          // 466 // Affects bonus armor gain from all sources except base stats
	AuraEffectModStat_BONUS_PCT                                          // 467 // Affects stat gain from all sources except base stats
	AuraEffectTriggerSpellOnHealthPct                                    // = 468 // Triggers spell when health goes above (MiscA = 0) or falls below (MiscA // 1) specified percent value (once not every time condition has meet)
	AuraEffectShowConfirmationPrompt_WITH_DIFFICULTY                     // 469
	AuraEffectModAuraTimeRateBySpellLabel                                // 470 // NYI
	AuraEffectModVersatility                                             // 471
	AuraEffect472                                                        // 472
	AuraEffectPreventDurabilityLoss_FROM_COMBAT                          // 473 // Prevents durability loss from dealing/taking damage
	AuraEffectReplaceItemBonusTree                                       // 474 // NYI
	AuraEffectAllowUsingGameobjectsWhileMounted                          // 475
	AuraEffectModCurrencyGainLooted                                      // 476
	AuraEffect477                                                        // 477
	AuraEffect478                                                        // 478
	AuraEffect479                                                        // 479
	AuraEffectModArtifactItemLevel                                       // 480
	AuraEffectConvertConsumedRune                                        // 481
	AuraEffect482                                                        // 482
	AuraEffectSuppressTransforms                                         // 483 // NYI
	AuraEffectAllowInterruptSpell                                        // 484 // NYI
	AuraEffectModMovementForceMagnitude                                  // 485
	AuraEffect486                                                        // 486
	AuraEffectCosmeticMounted                                            // 487
	AuraEffect488                                                        // 488
	AuraEffectModAlternativeDefaultLanguage                              // 489
	AuraEffect490                                                        // 490
	AuraEffect491                                                        // 491
	AuraEffect492                                                        // 492
	AuraEffect493                                                        // 493 // 1 spell 267116 - Animal Companion (modifies Call Pet)
	AuraEffectSetPowerPointCharge                                        // 494 // NYI
	AuraEffectTriggerSpellOnExpire                                       // 495
	AuraEffectAllowChangingEquipmentInTorghast                           // 496 // NYI
	AuraEffectModAnimaGain                                               // 497 // NYI
	AuraEffectCurrencyLossPctOnDeath                                     // 498 // NYI
	AuraEffectModRestedXpConsumption                                     // 499
	AuraEffectIgnoreSpellChargeCooldown                                  // 500 // NYI
	AuraEffectModCriticalDamageTakenFromCaster                           // 501
	AuraEffectModVersatility_DAMAGE_DONE_BENEFIT                         // 502 // NYI
	AuraEffectModVersatility_HEALING_DONE_BENEFIT                        // 503 // NYI
	AuraEffectModHealingTakenFromCaster                                  // 504
	AuraEffectModPlayerChoiceRerolls                                     // 505 // NYI
	AuraEffectDisableInertia                                             // 506
	AuraEffect507                                                        // 507
	AuraEffect508                                                        // 508
	AuraEffect509                                                        // 509
	AuraEffectModifiedRaidInstance                                       // 510 // Related to "Fated" raid affixes
	NumAuraEffects                                                       // 511
)
