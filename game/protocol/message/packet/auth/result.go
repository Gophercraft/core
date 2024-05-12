package auth

import (
	"fmt"

	"github.com/Gophercraft/core/bnet/rpc/codes"
)

type Result uint8

const (
	Ok                  Result = 0x0C
	Failed              Result = 0x0D
	Reject              Result = 0x0E
	BadServerProof      Result = 0x0F
	Unavailable         Result = 0x10
	SystemError         Result = 0x11
	BillingError        Result = 0x12
	BillingExpired      Result = 0x13
	VersionMismatch     Result = 0x14
	UnknownAccount      Result = 0x15
	IncorrectPassword   Result = 0x16
	SessionExpired      Result = 0x17
	ServerShuttingDown  Result = 0x18
	AlreadyLoggingIn    Result = 0x19
	LoginServerNotFound Result = 0x1A
	WaitQueue           Result = 0x1B
	Banned              Result = 0x1C
	AlreadyOnline       Result = 0x1D
	NoTime              Result = 0x1E
	DBBusy              Result = 0x1F
	Suspended           Result = 0x20
	ParentalControl     Result = 0x21
	LockedEnforced      Result = 0x22
)

func (response *Response) legacy_result() (r Result, err error) {
	switch response.Result {
	case codes.ERROR_OK:
		r = Ok
	case codes.ERROR_DENIED:
		r = Reject
	case codes.ERROR_LOGON_INVALID_SERVER_PROOF:
		r = BadServerProof
	default:
		err = fmt.Errorf("unknown %d", response.Result)
	}
	return
}
