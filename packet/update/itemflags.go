package update

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/Gophercraft/core/vsn"
)

type ItemFlags struct {
	Bitmask Bitmask
}

type ItemFlag uint32

const (
	ItemFlagNoPickup ItemFlag = iota
	ItemFlagConjured
	ItemFlagHasLoot
	ItemFlagHeroicTooltip
	ItemFlagDeprecated
	ItemFlagNoUserDestroy
	ItemFlagPlayerCast
	ItemFlagNoEquipCooldown
	ItemFlagMultiLootQuest
	ItemFlagIsWrapper
	ItemFlagUsesResources
	ItemFlagMultiDrop
	ItemFlagItemPurchaseRecord
	ItemFlagPetition
	ItemFlagHasText
	ItemFlagNoDisenchant
	ItemFlagRealDuration
	ItemFlagNoCreator
	ItemFlagIsProspectable
	ItemFlagUniqueEquippable
	ItemFlagIgnoreForAuras
	ItemFlagIgnoreDefaultArenaRestrictions
	ItemFlagNoDurabilityLoss
	ItemFlagUseWhenShapeshifted
	ItemFlagHasQuestGlow
	ItemFlagHideUnusableRecipe
	ItemFlagNotUseableInArena
	ItemFlagBindsToAccount
	ItemFlagNoReagentCost
	ItemFlagIsMillable
	ItemFlagReportToGuildChat
	ItemFlagNoProgressiveLoot
	ItemFlagFactionHorde
	ItemFlagFactionAlliance
	ItemFlagDontIgnoreBuyPrice
	ItemFlagClassifyAsCaster
	ItemFlagClassifyAsPhysical
	ItemFlagEveryoneCanRollNeed
	ItemFlagNoTradeBindOnAcquire
	ItemFlagCanTradeBindOnAcquire
	ItemFlagCanOnlyRollGreed
	ItemFlagCasterWeapon
	ItemFlagDeleteOnLogin
	ItemFlagInternalItem
	ItemFlagNoVendorValue
	ItemFlagShowBeforeDiscovered
	ItemFlagOverrideGoldCost
	ItemFlagIgnoreDefaultRatedBgRestrictions
	ItemFlagNotUsableInRatedBg
	ItemFlagBnetAccountTradeOk
	ItemFlagConfirmBeforeUse
	ItemFlagReevaluateBondingOnTransform
	ItemFlagNoTransformOnChargeDepletion
	ItemFlagNoAlterItemVisual
	ItemFlagNoSourceForItemVisual
	ItemFlagIgnoreQualityForItemVisualSource
	ItemFlagNoDurability
	ItemFlagRoleTank
	ItemFlagRoleHealer
	ItemFlagRoleDamage
	ItemFlagCanDropInChallengeMode
	ItemFlagNeverStackInLootUi
	ItemFlagDisenchantToLootTable
	ItemFlagUsedInATradeskill
	ItemFlagDontDestroyOnQuestAccept
	ItemFlagItemCanBeUpgraded
	ItemFlagUpgradeFromItemOverridesDropUpgrade
	ItemFlagAlwaysFfaInLoot
	ItemFlagHideUpgradeLevelsIfNotUpgraded
	ItemFlagUpdateInteractions
	ItemFlagUpdateDoesntLeaveProgressiveWinHistory
	ItemFlagIgnoreItemHistoryTracker
	ItemFlagIgnoreItemLevelCapInPvp
	ItemFlagDisplayAsHeirloom
	ItemFlagSkipUseCheckOnPickup
	ItemFlagObsolete
	ItemFlagDontDisplayInGuildNews
	ItemFlagPvpTournamentGear
	ItemFlagRequiresStackChangeLog
	ItemFlagUnusedFlag
	ItemFlagHideNameSuffix
	ItemFlagPushLoot
	ItemFlagDontReportLootLogToParty
	ItemFlagAlwaysAllowDualWield
	ItemFlagObliteratable
	ItemFlagActsAsTransmogHiddenVisualOption
	ItemFlagExpireOnWeeklyReset
	ItemFlagDoesntShowUpInTransmogUntilCollected
	ItemFlagCanStoreEnchants
	ItemFlagHideQuestItemFromObjectTooltip
	ItemFlagDoNotToast
	ItemFlagIgnoreCreationContextForProgressiveWinHistory
	ItemFlagForceAllSpecsForItemHistory
	ItemFlagSaveOnConsume
	ItemFlagContainerSavesPlayerData
	ItemFlagNoVoidStorage
	ItemFlagHandleOnUseEffectImmediately
	ItemFlagAlwaysShowItemLevelInTooltip
	ItemFlagShowsGenerationWithRandomStats
	ItemFlagActivateOnEquipEffectsWhenTransmogrified
	ItemFlagEnforceTransmogWithChildItem
	ItemFlagScrapable
	ItemFlagBypassRepRequirementsForTransmog
	ItemFlagDisplayOnlyOnDefinedRaces
	ItemFlagRegulatedCommodity
	ItemFlagCreateLootImmediately
	ItemFlagGenerateLootSpecItem
	ItemFlagEnd
)

var ItemFlagNames = map[ItemFlag]string{
	ItemFlagNoPickup:                                      "NoPickup",
	ItemFlagConjured:                                      "Conjured",
	ItemFlagHasLoot:                                       "HasLoot",
	ItemFlagHeroicTooltip:                                 "HeroicTooltip",
	ItemFlagDeprecated:                                    "Deprecated",
	ItemFlagNoUserDestroy:                                 "NoUserDestroy",
	ItemFlagPlayerCast:                                    "PlayerCast",
	ItemFlagNoEquipCooldown:                               "NoEquipCooldown",
	ItemFlagMultiLootQuest:                                "MultiLootQuest",
	ItemFlagIsWrapper:                                     "IsWrapper",
	ItemFlagUsesResources:                                 "UsesResources",
	ItemFlagMultiDrop:                                     "MultiDrop",
	ItemFlagItemPurchaseRecord:                            "ItemPurchaseRecord",
	ItemFlagPetition:                                      "Petition",
	ItemFlagHasText:                                       "HasText",
	ItemFlagNoDisenchant:                                  "NoDisenchant",
	ItemFlagRealDuration:                                  "RealDuration",
	ItemFlagNoCreator:                                     "NoCreator",
	ItemFlagIsProspectable:                                "IsProspectable",
	ItemFlagUniqueEquippable:                              "UniqueEquippable",
	ItemFlagIgnoreForAuras:                                "IgnoreForAuras",
	ItemFlagIgnoreDefaultArenaRestrictions:                "IgnoreDefaultArenaRestrictions",
	ItemFlagNoDurabilityLoss:                              "NoDurabilityLoss",
	ItemFlagUseWhenShapeshifted:                           "UseWhenShapeshifted",
	ItemFlagHasQuestGlow:                                  "HasQuestGlow",
	ItemFlagHideUnusableRecipe:                            "HideUnusableRecipe",
	ItemFlagNotUseableInArena:                             "NotUseableInArena",
	ItemFlagBindsToAccount:                                "BindsToAccount",
	ItemFlagNoReagentCost:                                 "NoReagentCost",
	ItemFlagIsMillable:                                    "IsMillable",
	ItemFlagReportToGuildChat:                             "ReportToGuildChat",
	ItemFlagNoProgressiveLoot:                             "NoProgressiveLoot",
	ItemFlagFactionHorde:                                  "FactionHorde",
	ItemFlagFactionAlliance:                               "FactionAlliance",
	ItemFlagDontIgnoreBuyPrice:                            "DontIgnoreBuyPrice",
	ItemFlagClassifyAsCaster:                              "ClassifyAsCaster",
	ItemFlagClassifyAsPhysical:                            "ClassifyAsPhysical",
	ItemFlagEveryoneCanRollNeed:                           "EveryoneCanRollNeed",
	ItemFlagNoTradeBindOnAcquire:                          "NoTradeBindOnAcquire",
	ItemFlagCanTradeBindOnAcquire:                         "CanTradeBindOnAcquire",
	ItemFlagCanOnlyRollGreed:                              "CanOnlyRollGreed",
	ItemFlagCasterWeapon:                                  "CasterWeapon",
	ItemFlagDeleteOnLogin:                                 "DeleteOnLogin",
	ItemFlagInternalItem:                                  "InternalItem",
	ItemFlagNoVendorValue:                                 "NoVendorValue",
	ItemFlagShowBeforeDiscovered:                          "ShowBeforeDiscovered",
	ItemFlagOverrideGoldCost:                              "OverrideGoldCost",
	ItemFlagIgnoreDefaultRatedBgRestrictions:              "IgnoreDefaultRatedBgRestrictions",
	ItemFlagNotUsableInRatedBg:                            "NotUsableInRatedBg",
	ItemFlagBnetAccountTradeOk:                            "BnetAccountTradeOk",
	ItemFlagConfirmBeforeUse:                              "ConfirmBeforeUse",
	ItemFlagReevaluateBondingOnTransform:                  "ReevaluateBondingOnTransform",
	ItemFlagNoTransformOnChargeDepletion:                  "NoTransformOnChargeDepletion",
	ItemFlagNoAlterItemVisual:                             "NoAlterItemVisual",
	ItemFlagNoSourceForItemVisual:                         "NoSourceForItemVisual",
	ItemFlagIgnoreQualityForItemVisualSource:              "IgnoreQualityForItemVisualSource",
	ItemFlagNoDurability:                                  "NoDurability",
	ItemFlagRoleTank:                                      "RoleTank",
	ItemFlagRoleHealer:                                    "RoleHealer",
	ItemFlagRoleDamage:                                    "RoleDamage",
	ItemFlagCanDropInChallengeMode:                        "CanDropInChallengeMode",
	ItemFlagNeverStackInLootUi:                            "NeverStackInLootUi",
	ItemFlagDisenchantToLootTable:                         "DisenchantToLootTable",
	ItemFlagUsedInATradeskill:                             "UsedInATradeskill",
	ItemFlagDontDestroyOnQuestAccept:                      "DontDestroyOnQuestAccept",
	ItemFlagItemCanBeUpgraded:                             "ItemCanBeUpgraded",
	ItemFlagUpgradeFromItemOverridesDropUpgrade:           "UpgradeFromItemOverridesDropUpgrade",
	ItemFlagAlwaysFfaInLoot:                               "AlwaysFfaInLoot",
	ItemFlagHideUpgradeLevelsIfNotUpgraded:                "HideUpgradeLevelsIfNotUpgraded",
	ItemFlagUpdateInteractions:                            "UpdateInteractions",
	ItemFlagUpdateDoesntLeaveProgressiveWinHistory:        "UpdateDoesntLeaveProgressiveWinHistory",
	ItemFlagIgnoreItemHistoryTracker:                      "IgnoreItemHistoryTracker",
	ItemFlagIgnoreItemLevelCapInPvp:                       "IgnoreItemLevelCapInPvp",
	ItemFlagDisplayAsHeirloom:                             "DisplayAsHeirloom",
	ItemFlagSkipUseCheckOnPickup:                          "SkipUseCheckOnPickup",
	ItemFlagObsolete:                                      "Obsolete",
	ItemFlagDontDisplayInGuildNews:                        "DontDisplayInGuildNews",
	ItemFlagPvpTournamentGear:                             "PvpTournamentGear",
	ItemFlagRequiresStackChangeLog:                        "RequiresStackChangeLog",
	ItemFlagUnusedFlag:                                    "UnusedFlag",
	ItemFlagHideNameSuffix:                                "HideNameSuffix",
	ItemFlagPushLoot:                                      "PushLoot",
	ItemFlagDontReportLootLogToParty:                      "DontReportLootLogToParty",
	ItemFlagAlwaysAllowDualWield:                          "AlwaysAllowDualWield",
	ItemFlagObliteratable:                                 "Obliteratable",
	ItemFlagActsAsTransmogHiddenVisualOption:              "ActsAsTransmogHiddenVisualOption",
	ItemFlagExpireOnWeeklyReset:                           "ExpireOnWeeklyReset",
	ItemFlagDoesntShowUpInTransmogUntilCollected:          "DoesntShowUpInTransmogUntilCollected",
	ItemFlagCanStoreEnchants:                              "CanStoreEnchants",
	ItemFlagHideQuestItemFromObjectTooltip:                "HideQuestItemFromObjectTooltip",
	ItemFlagDoNotToast:                                    "DoNotToast",
	ItemFlagIgnoreCreationContextForProgressiveWinHistory: "IgnoreCreationContextForProgressiveWinHistory",
	ItemFlagForceAllSpecsForItemHistory:                   "ForceAllSpecsForItemHistory",
	ItemFlagSaveOnConsume:                                 "SaveOnConsume",
	ItemFlagContainerSavesPlayerData:                      "ContainerSavesPlayerData",
	ItemFlagNoVoidStorage:                                 "NoVoidStorage",
	ItemFlagHandleOnUseEffectImmediately:                  "HandleOnUseEffectImmediately",
	ItemFlagAlwaysShowItemLevelInTooltip:                  "AlwaysShowItemLevelInTooltip",
	ItemFlagShowsGenerationWithRandomStats:                "ShowsGenerationWithRandomStats",
	ItemFlagActivateOnEquipEffectsWhenTransmogrified:      "ActivateOnEquipEffectsWhenTransmogrified",
	ItemFlagEnforceTransmogWithChildItem:                  "EnforceTransmogWithChildItem",
	ItemFlagScrapable:                                     "Scrapable",
	ItemFlagBypassRepRequirementsForTransmog:              "BypassRepRequirementsForTransmog",
	ItemFlagDisplayOnlyOnDefinedRaces:                     "DisplayOnlyOnDefinedRaces",
	ItemFlagRegulatedCommodity:                            "RegulatedCommodity",
	ItemFlagCreateLootImmediately:                         "CreateLootImmediately",
	ItemFlagGenerateLootSpecItem:                          "GenerateLootSpecItem",
}

var ItemFlagLookup = map[string]ItemFlag{}

func init() {
	for k, v := range ItemFlagNames {
		ItemFlagLookup[v] = k
	}
}

func (it ItemFlags) HasFlag(flag ItemFlag) bool {
	return it.Bitmask.Enabled(uint32(flag))
}

func (it *ItemFlags) DecodeWord(str string) error {
	if str == "" {
		return nil
	}

	flags := strings.Split(str, "|")
	for _, flag := range flags {
		code, ok := ItemFlagLookup[flag]
		if !ok {
			return fmt.Errorf("update: couldn't find item flag named '%s'", flag)
		}
		it.Bitmask.Set(uint32(code), true)
	}
	return nil
}

func (it *ItemFlags) EncodeWord() (string, error) {
	names := []string{}

	for i := ItemFlag(0); i < ItemFlagEnd; i++ {
		if it.HasFlag(i) {
			names = append(names, i.String())
		}
	}

	return strings.Join(names, "|"), nil
}

func (x ItemFlag) String() string {
	return ItemFlagNames[x]
}

func (iflg *ItemFlags) Encode(build vsn.Build, out io.Writer) error {
	var codes []uint32
	codesLen := 1
	if build.AddedIn(10192) {
		codesLen = 2
	}
	if build.AddedIn(28153) {
		codesLen = 4
	}
	codes = make([]uint32, codesLen)
	if len(iflg.Bitmask) >= codesLen {
		copy(codes[:], iflg.Bitmask[0:codesLen])
	} else {
		copy(codes[:], iflg.Bitmask)
	}

	var bytes [4]byte
	for _, code := range codes {
		binary.LittleEndian.PutUint32(bytes[:], code)
		out.Write(bytes[:])
	}
	return nil
}

func DecodeItemFlags(build vsn.Build, in io.Reader) (ItemFlags, error) {
	var flags ItemFlags
	var codes []uint32
	codesLen := 1
	if build.AddedIn(10192) {
		codesLen = 2
	}
	if build.AddedIn(28153) {
		codesLen = 4
	}
	codes = make(Bitmask, codesLen)
	codeBytes := make([]byte, 4*codesLen)
	if _, err := in.Read(codeBytes); err != nil {
		return flags, err
	}
	for i := range codes {
		codes[i] = binary.LittleEndian.Uint32(codeBytes[i*4 : (i+1)*4])
	}
	flags.Bitmask = codes
	return flags, nil
}
