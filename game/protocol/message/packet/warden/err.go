package warden

type Error struct {
	Msg string
}

var (
	ErrCheckBadChecksum  = Error{"Bad checksum in client cheat checks"}
	ErrNeedCurrentChecks = Error{"CClientCheatChecksResult contains no intrinsic structure. It requires the last set of warden checks to be parsed"}
)

func (e Error) Error() string {
	return "packet/warden: " + e.Msg
}
