package item

import (
	"fmt"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type InvResult uint8

const (
	InvOk InvResult = iota
	InvCantEquipLevelI
	InvCantEquipSkill
	InvItemDoesntGoToSlot
	InvBagFull
	InvNonEmptyBagOverOtherBag
	InvCantTradeEquipBags
	InvOnlyAmmoCanGoHere
	InvNoRequiredProficiency
	InvNoEquipmentSlotAvailable
	InvYouCanNeverUseThatItem
	InvYouCanNeverUseThatItem2
	InvNoEquipmentSlotAvailable2
	InvCantEquipWithTwoHanded
	InvCantDualWield
	InvItemDoesntGoIntoBag
	InvItemDoesntGoIntoBag2
	InvCantCarryMoreOfThis
	InvNoEquipmentSlotAvailable3
	InvItemCantStack
	InvItemCantBeEquipped
	InvItemsCantBeSwapped
	InvSlotIsEmpty
	InvItemNotFound
	InvCantDropSoulbound
	InvOutOfRange
	InvTriedToSplitMoreThanCount
	InvCouldntSplitItems
	InvMissingReagent
	InvNotEnoughMoney
	InvNotABag
	InvCanOnlyDoWithEmptyBags
	InvDontOwnThatItem
	InvCanEquipOnly1Quiver
	InvMustPurchaseThatBagSlot
	InvTooFarAwayFromBank
	InvItemLocked
	InvYouAreStunned
	InvYouAreDead
	InvCantDoRightNow
	InvIntBagError
	InvCanEquipOnly1Bolt
	InvCanEquipOnly1Ammopouch
	InvStackableCantBeWrapped
	InvEquippedCantBeWrapped
	InvWrappedCantBeWrapped
	InvBoundCantBeWrapped
	InvUniqueCantBeWrapped
	InvBagsCantBeWrapped
	InvAlreadyLooted
	InvInventoryFull
	InvBankFull
	InvItemIsCurrentlySoldOut
	InvBagFull3
	InvItemNotFound2
	InvItemCantStack2
	InvBagFull4
	InvItemSoldOut
	InvObjectIsBusy
	InvNone
	InvNotInCombat
	InvNotWhileDisarmed
	InvBagFull6
	InvCantEquipRank
	InvCantEquipReputation
	InvTooManySpecialBags
	InvLootCantLootThatNow
	InvItemUniqueEquipable
	InvVendorMissingTurnins
	InvNotEnoughHonorPoints
	InvNotEnoughArenaPoints
	InvItemMaxCountSocketed
	InvMailBoundItem
	InvNoSplitWhileProspecting
	InvItemMaxCountEquippedSocketed
	InvItemUniqueEquippableSocketed
	InvTooMuchGold
	InvNotDuringArenaMatch
	InvCannotTradeThat
	InvPersonalArenaRatingTooLow
	InvEventAutoequipBindConfirm
	InvArtefactsOnlyForOwnCharacters
	InvItemMaxLimitCategoryCountExceeded
	InvItemMaxLimitCategorySocketedExceeded
	InvScalingStatItemLevelExceeded
	InvPurchaseLevelTooLow
	InvCantEquipNeedTalent
	InvItemMaxLimitCategoryEquippedExceeded
)

type InvResultDescriptor map[InvResult]uint8

var InvResultDescriptors = map[vsn.BuildRange]InvResultDescriptor{
	{0, vsn.V8_3_0}: {
		InvOk:                            0,
		InvCantEquipLevelI:               1,
		InvCantEquipSkill:                2,
		InvItemDoesntGoToSlot:            3,
		InvBagFull:                       4,
		InvNonEmptyBagOverOtherBag:       5,
		InvCantTradeEquipBags:            6,
		InvOnlyAmmoCanGoHere:             7,
		InvNoRequiredProficiency:         8,
		InvNoEquipmentSlotAvailable:      9,
		InvYouCanNeverUseThatItem:        10,
		InvYouCanNeverUseThatItem2:       11,
		InvNoEquipmentSlotAvailable2:     12,
		InvCantEquipWithTwoHanded:        13,
		InvCantDualWield:                 14,
		InvItemDoesntGoIntoBag:           15,
		InvItemDoesntGoIntoBag2:          16,
		InvCantCarryMoreOfThis:           17,
		InvNoEquipmentSlotAvailable3:     18,
		InvItemCantStack:                 19,
		InvItemCantBeEquipped:            20,
		InvItemsCantBeSwapped:            21,
		InvSlotIsEmpty:                   22,
		InvItemNotFound:                  23,
		InvCantDropSoulbound:             24,
		InvOutOfRange:                    25,
		InvTriedToSplitMoreThanCount:     26,
		InvCouldntSplitItems:             27,
		InvMissingReagent:                28,
		InvNotEnoughMoney:                29,
		InvNotABag:                       30,
		InvCanOnlyDoWithEmptyBags:        31,
		InvDontOwnThatItem:               32,
		InvCanEquipOnly1Quiver:           33,
		InvMustPurchaseThatBagSlot:       34,
		InvTooFarAwayFromBank:            35,
		InvItemLocked:                    36,
		InvYouAreStunned:                 37,
		InvYouAreDead:                    38,
		InvCantDoRightNow:                39,
		InvIntBagError:                   40,
		InvCanEquipOnly1Bolt:             41,
		InvCanEquipOnly1Ammopouch:        42,
		InvStackableCantBeWrapped:        43,
		InvEquippedCantBeWrapped:         44,
		InvWrappedCantBeWrapped:          45,
		InvBoundCantBeWrapped:            46,
		InvUniqueCantBeWrapped:           47,
		InvBagsCantBeWrapped:             48,
		InvAlreadyLooted:                 49,
		InvInventoryFull:                 50,
		InvBankFull:                      51,
		InvItemIsCurrentlySoldOut:        52,
		InvBagFull3:                      53,
		InvItemNotFound2:                 54,
		InvItemCantStack2:                55,
		InvBagFull4:                      56,
		InvItemSoldOut:                   57,
		InvObjectIsBusy:                  58,
		InvNone:                          59,
		InvNotInCombat:                   60,
		InvNotWhileDisarmed:              61,
		InvBagFull6:                      62,
		InvCantEquipRank:                 63,
		InvCantEquipReputation:           64,
		InvTooManySpecialBags:            65,
		InvLootCantLootThatNow:           66,
		InvItemUniqueEquipable:           67,
		InvVendorMissingTurnins:          68,
		InvNotEnoughHonorPoints:          69,
		InvNotEnoughArenaPoints:          70,
		InvItemMaxCountSocketed:          71,
		InvMailBoundItem:                 72,
		InvNoSplitWhileProspecting:       73,
		InvItemMaxCountEquippedSocketed:  75,
		InvItemUniqueEquippableSocketed:  76,
		InvTooMuchGold:                   77,
		InvNotDuringArenaMatch:           78,
		InvCannotTradeThat:               79,
		InvPersonalArenaRatingTooLow:     80,
		InvEventAutoequipBindConfirm:     81,
		InvArtefactsOnlyForOwnCharacters: 82,
		// no output                                 = 83,
		InvItemMaxLimitCategoryCountExceeded:    84,
		InvItemMaxLimitCategorySocketedExceeded: 85,
		InvScalingStatItemLevelExceeded:         86,
		InvPurchaseLevelTooLow:                  87,
		InvCantEquipNeedTalent:                  88,
		InvItemMaxLimitCategoryEquippedExceeded: 89,
	},
}

type InvChangeError struct {
	Result            InvResult
	SrcItem           guid.GUID
	DstItem           guid.GUID
	BagSubclass       uint8
	RequiredLevel     uint32
	ItemLimitCategory uint32
	AEItem            guid.GUID
	AESlot            models.ItemSlot
	AEContainer       guid.GUID
}

func (ice *InvChangeError) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_INVENTORY_CHANGE_FAILURE

	var desc InvResultDescriptor
	if err := vsn.QueryDescriptors(build, InvResultDescriptors, &desc); err != nil {
		return err
	}

	code, ok := desc[ice.Result]
	if !ok {
		return fmt.Errorf("packet: could not resolve InvResult %d", ice.Result)
	}

	out.WriteByte(code)
	switch ice.Result {
	case InvOk:
		return nil
	}

	ice.SrcItem.EncodeUnpacked(build, out)
	ice.DstItem.EncodeUnpacked(build, out)
	out.WriteByte(ice.BagSubclass) //bag type subclass

	switch ice.Result {
	case InvCantEquipLevelI, InvPurchaseLevelTooLow:
		out.WriteUint32(ice.RequiredLevel)
	case InvEventAutoequipBindConfirm:
		ice.AEItem.EncodeUnpacked(build, out)
		out.WriteUint32(uint32(ice.AESlot))
		ice.AEContainer.EncodeUnpacked(build, out)
	case InvItemMaxLimitCategoryCountExceeded, InvItemMaxLimitCategorySocketedExceeded, InvItemMaxLimitCategoryEquippedExceeded:
		out.WriteUint32(ice.ItemLimitCategory)
	default:
	}

	return nil
}

func (ice *InvChangeError) Decode(build vsn.Build, in *packet.WorldPacket) error {
	var desc InvResultDescriptor
	if err := vsn.QueryDescriptors(build, InvResultDescriptors, &desc); err != nil {
		return err
	}

	code := in.ReadByte()

	for k, v := range desc {
		if v == code {
			ice.Result = k
			break
		}
	}

	if ice.Result == InvOk {
		return nil
	}

	var err error
	ice.SrcItem, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return err
	}

	ice.DstItem, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return err
	}

	ice.BagSubclass = in.ReadByte()
	switch ice.Result {
	case InvCantEquipLevelI, InvPurchaseLevelTooLow:
		ice.RequiredLevel = in.ReadUint32()
	case InvEventAutoequipBindConfirm:
		ice.AEItem, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}
		ice.AESlot = models.ItemSlot(in.ReadUint32())
		ice.AEContainer, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}
	case InvItemMaxLimitCategoryCountExceeded, InvItemMaxLimitCategorySocketedExceeded, InvItemMaxLimitCategoryEquippedExceeded:
		ice.ItemLimitCategory = in.ReadUint32()
	default:
	}
	return nil
}
