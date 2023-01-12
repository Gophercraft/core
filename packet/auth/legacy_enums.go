package auth

type Legacy uint8

const (
	Ok                  Legacy = 0x0C
	Failed              Legacy = 0x0D
	Reject              Legacy = 0x0E
	BadServerProof      Legacy = 0x0F
	Unavailable         Legacy = 0x10
	SystemError         Legacy = 0x11
	BillingError        Legacy = 0x12
	BillingExpired      Legacy = 0x13
	VersionMismatch     Legacy = 0x14
	UnknownAccount      Legacy = 0x15
	IncorrectPassword   Legacy = 0x16
	SessionExpired      Legacy = 0x17
	ServerShuttingDown  Legacy = 0x18
	AlreadyLoggingIn    Legacy = 0x19
	LoginServerNotFound Legacy = 0x1A
	WaitQueue           Legacy = 0x1B
	Banned              Legacy = 0x1C
	AlreadyOnline       Legacy = 0x1D
	NoTime              Legacy = 0x1E
	DBBusy              Legacy = 0x1F
	Suspended           Legacy = 0x20
	ParentalControl     Legacy = 0x21
	LockedEnforced      Legacy = 0x22
)
