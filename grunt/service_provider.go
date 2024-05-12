package grunt

import (
	"net"

	"github.com/Gophercraft/core/version"
)

type NetEventType uint8

const (
	NewConnection NetEventType = iota
	LoginAttempt
)

type Event struct {
	Type          NetEventType
	RemoteAddress net.Addr
	AccountName   string
}

type AccountInfo struct {
	AccountName  string
	OS           OS
	Architecture Architecture
	Locale       Locale
	SessionKey   []byte
}

type AccountLoginInfo struct {
	Authenticator   bool
	PIN             bool
	LogonProofFlags LogonProofFlags
	IdentityHash    []byte
	SessionKey      []byte
}

// The
// account_name is in upper-case
type ServiceProvider interface {
	// Trace Network event (antispam)
	// you can define antispam with this method, such as rejecting connections from certain addresses
	Check(event *Event) (err error)
	GetAccountLoginInfo(account_name string, build version.Build, info *AccountLoginInfo) (err error)
	VerifyAuthenticatorCode(account_name string, code string) bool
	VerifyPIN(account string, server_pin_info *ServerPINInfo, client_pin_info *ClientPINInfo) bool
	GetRealmList(account_name string, build version.Build) (list []Realm, err error)
	StoreAccountInfo(info *AccountInfo) (err error)
}
