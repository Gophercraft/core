package auth

import (
	"fmt"

	"github.com/Gophercraft/core/bnet"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/character"
	"github.com/Gophercraft/core/vsn"
)

const NewResponse = vsn.V4_0_1

type VirtualRealmNameInfo struct {
	IsLocal             bool
	IsInternalRealm     bool
	RealmNameActual     string
	RealmNameNormalized string
}

type VirtualRealmInfo struct {
	VirtualAddress uint32
	VirtualRealmNameInfo
}

type ClassAvailability struct {
	ClassID               uint8
	ActiveExpansionLevel  uint8
	AccountExpansionLevel uint8
}

type RaceClassAvailability struct {
	RaceID  uint8
	Classes []ClassAvailability
}

type GameTime struct {
	BillingPlan uint32
	TimeRemain  uint32
	Unknown735  uint32
	InGameRoom  bool
}

type SuccessInfo struct {
	ActiveExpansionLevel  uint8
	AccountExpansionLevel uint8

	TimeRested             uint32
	VirtualRealmAddress    uint32
	TimeSecondsUntilPCKick uint32
	CurrencyID             uint32
	Time                   int64

	GameTimeInfo GameTime

	VirtualRealms []VirtualRealmInfo
	Templates     []character.Template

	AvailableClasses []RaceClassAvailability

	IsExpansionTrial       bool
	ForceCharacterTemplate bool

	NumPlayersRedTeam  uint16
	NumPlayersBlueTeam uint16

	ExpansionTrialExpiration uint32
}

type WaitInfo struct {
	WaitCount uint32
	WaitTime  uint32
	HasFCM    bool
}

type Response struct {
	Result bnet.Status

	SuccessInfo *SuccessInfo
	WaitInfo    *WaitInfo
}

func (ar *Response) legacyResult() (r Legacy, err error) {
	switch ar.Result {
	case bnet.ERROR_OK:
		r = Ok
	case bnet.ERROR_DENIED:
		r = Reject
	case bnet.ERROR_LOGON_INVALID_SERVER_PROOF:
		r = Failed
	default:
		err = fmt.Errorf("unknown %d", ar.Result)
	}
	return
}

func (ar *Response) Encode(build vsn.Build, to *packet.WorldPacket) error {
	to.Type = packet.SMSG_AUTH_RESPONSE
	// Use old packet structure
	if build.RemovedIn(NewResponse) {
		res, err := ar.legacyResult()
		if err != nil {
			return err
		}

		if ar.WaitInfo != nil && res == Ok {
			to.WriteByte(uint8(WaitQueue))
		} else {
			to.WriteByte(uint8(res))
		}

		switch res {
		case Ok:
			if ar.SuccessInfo != nil {
				to.WriteUint32(ar.SuccessInfo.GameTimeInfo.TimeRemain)
				to.WriteByte(uint8(ar.SuccessInfo.GameTimeInfo.BillingPlan))
				to.WriteUint32(ar.SuccessInfo.TimeRested)
			} else {
				to.WriteUint32(0)
				to.WriteByte(0)
				to.WriteUint32(0)
			}

			if ar.WaitInfo != nil {
				to.WriteUint32(ar.WaitInfo.WaitCount)
			}

			if build.AddedIn(vsn.V2_0_1) {
				if ar.SuccessInfo != nil {
					to.WriteByte(ar.SuccessInfo.AccountExpansionLevel)
				} else {
					to.WriteByte(uint8(build.Exp()))
				}
			}
		default:
		}

		return nil
	}

	to.WriteUint32(uint32(ar.Result))
	to.WriteMSBit(ar.SuccessInfo != nil)
	to.WriteMSBit(ar.WaitInfo != nil)
	to.FlushBits()

	if ar.SuccessInfo != nil {
		to.WriteUint32(ar.SuccessInfo.VirtualRealmAddress)
		to.WriteUint32(uint32(len(ar.SuccessInfo.VirtualRealms)))
		to.WriteUint32(ar.SuccessInfo.TimeRested)
		to.WriteByte(ar.SuccessInfo.ActiveExpansionLevel)
		to.WriteByte(ar.SuccessInfo.AccountExpansionLevel)
		to.WriteUint32(ar.SuccessInfo.TimeSecondsUntilPCKick)
		to.WriteUint32(uint32(len(ar.SuccessInfo.AvailableClasses)))
		to.WriteUint32(uint32(len(ar.SuccessInfo.Templates)))
		to.WriteUint32(ar.SuccessInfo.CurrencyID)
		to.WriteInt64(ar.SuccessInfo.Time)

		for _, class := range ar.SuccessInfo.AvailableClasses {
			to.WriteByte(class.RaceID)
			to.WriteUint32(uint32(len(class.Classes)))

			for _, classAvailability := range class.Classes {
				to.WriteByte(classAvailability.ClassID)
				to.WriteByte(classAvailability.ActiveExpansionLevel)
				to.WriteByte(classAvailability.AccountExpansionLevel)
			}
		}

		includeNumPlayersR := ar.SuccessInfo.NumPlayersRedTeam != 0
		includeNumPlayersB := ar.SuccessInfo.NumPlayersBlueTeam != 0
		includeETE := ar.SuccessInfo.ExpansionTrialExpiration != 0

		to.WriteMSBit(ar.SuccessInfo.IsExpansionTrial)
		to.WriteMSBit(ar.SuccessInfo.ForceCharacterTemplate)
		to.WriteMSBit(includeNumPlayersR)
		to.WriteMSBit(includeNumPlayersB)
		to.WriteMSBit(includeETE)
		to.FlushBits()

		gti := ar.SuccessInfo.GameTimeInfo
		to.WriteUint32(gti.BillingPlan)
		to.WriteUint32(gti.TimeRemain)
		to.WriteUint32(gti.Unknown735)

		to.WriteMSBit(gti.InGameRoom)
		to.WriteMSBit(gti.InGameRoom)
		to.WriteMSBit(gti.InGameRoom)
		to.FlushBits()
	}

	if ar.WaitInfo != nil {

	}

	return nil
}
