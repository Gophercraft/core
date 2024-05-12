package auth

import (
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/packet/character"
	"github.com/Gophercraft/core/version"
)

const NewResponse = version.V4_0_1

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
	Result codes.Code

	SuccessInfo *SuccessInfo
	WaitInfo    *WaitInfo
}
