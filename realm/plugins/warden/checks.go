package warden

import (
	"fmt"

	packetwarden "github.com/Gophercraft/core/packet/warden"
)

type FailType uint8

const (
	FailNone FailType = iota
	FailPacket
	FailHash
	FailVersion
	FailAnticheat
)

func (sd *SessionData) HandleClientCheatChecksResult(cccr *packetwarden.ClientCheatChecksResult) {
	for i, result := range cccr.CheckResults {
		serverCheck := cccr.CurrentChecks.Checks[i]
		switch serverCheck.Type {
		case packetwarden.CheckTiming:
		case packetwarden.CheckMem:
			sd.HandleMemoryCheck(cccr.)
		}
	}
}

func (sd *SessionData) Fail(ft FailType, data string, args ...interface{}) {
	var reason string
	switch ft {
	case FailNone:
		reason = "Unspecified reason"
	case FailPacket:
		reason = "Your client sent a badly formatted packet to Warden"
	case FailHash:
		reason = "Your client sent an incorrect hash value"
	case FailVersion:
		reason = "Your client is reporting a version different from the server's"
	case FailAnticheat:
		reason = "Cheats have been observed running in your client"
	}
	if reason == "" {
		return
	}

	str := fmt.Sprintf("%s: "+data, append([]interface{}{reason}, args...)...)
	sd.Wardenf(WarnFatal, "%s", str)
}


