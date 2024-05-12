package auth

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

func (response *Response) encode_0_12340(build version.Build, to *message.Packet) (err error) {
	var (
		result_code Result
	)
	result_code, err = response.legacy_result()
	if err != nil {
		return err
	}

	if response.WaitInfo != nil {
		to.WriteUint8(uint8(WaitQueue))
		to.WriteUint32(response.WaitInfo.WaitCount)
		return
	}
	to.WriteUint8(uint8(result_code))

	if result_code != Ok {
		return
	}

	if response.SuccessInfo != nil {
		to.WriteUint32(response.SuccessInfo.GameTimeInfo.TimeRemain)
		to.WriteUint8(uint8(response.SuccessInfo.GameTimeInfo.BillingPlan))
		to.WriteUint32(response.SuccessInfo.TimeRested)
	} else {
		to.WriteUint32(0)
		to.WriteUint8(0)
		to.WriteUint32(0)
	}

	if build.AddedIn(version.V2_0_1) {
		if response.SuccessInfo != nil {
			to.WriteUint8(response.SuccessInfo.AccountExpansionLevel)
		} else {
			to.WriteUint8(uint8(build.Exp()))
		}
	}

	return nil
}
