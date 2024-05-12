package auth

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

func (response *Response) encode_13164_(build version.Build, to *message.Packet) (err error) {
	to.WriteUint32(uint32(response.Result))
	to.WriteBit(response.SuccessInfo != nil)
	to.WriteBit(response.WaitInfo != nil)
	to.FlushBits()

	if response.SuccessInfo != nil {
		to.WriteUint32(response.SuccessInfo.VirtualRealmAddress)
		to.WriteUint32(uint32(len(response.SuccessInfo.VirtualRealms)))
		to.WriteUint32(response.SuccessInfo.TimeRested)
		to.WriteUint8(response.SuccessInfo.ActiveExpansionLevel)
		to.WriteUint8(response.SuccessInfo.AccountExpansionLevel)
		to.WriteUint32(response.SuccessInfo.TimeSecondsUntilPCKick)
		to.WriteUint32(uint32(len(response.SuccessInfo.AvailableClasses)))
		to.WriteUint32(uint32(len(response.SuccessInfo.Templates)))
		to.WriteUint32(response.SuccessInfo.CurrencyID)
		to.WriteInt64(response.SuccessInfo.Time)

		for _, class := range response.SuccessInfo.AvailableClasses {
			to.WriteUint8(class.RaceID)
			to.WriteUint32(uint32(len(class.Classes)))

			for _, classAvailability := range class.Classes {
				to.WriteUint8(classAvailability.ClassID)
				to.WriteUint8(classAvailability.ActiveExpansionLevel)
				to.WriteUint8(classAvailability.AccountExpansionLevel)
			}
		}

		includeNumPlayersR := response.SuccessInfo.NumPlayersRedTeam != 0
		includeNumPlayersB := response.SuccessInfo.NumPlayersBlueTeam != 0
		includeETE := response.SuccessInfo.ExpansionTrialExpiration != 0

		to.WriteBit(response.SuccessInfo.IsExpansionTrial)
		to.WriteBit(response.SuccessInfo.ForceCharacterTemplate)
		to.WriteBit(includeNumPlayersR)
		to.WriteBit(includeNumPlayersB)
		to.WriteBit(includeETE)
		to.FlushBits()

		gti := response.SuccessInfo.GameTimeInfo
		to.WriteUint32(gti.BillingPlan)
		to.WriteUint32(gti.TimeRemain)
		to.WriteUint32(gti.Unknown735)

		to.WriteBit(gti.InGameRoom)
		to.WriteBit(gti.InGameRoom)
		to.WriteBit(gti.InGameRoom)
		to.FlushBits()
	}

	if response.WaitInfo != nil {
		err = fmt.Errorf("Wait info nyi")
	}

	return
}
