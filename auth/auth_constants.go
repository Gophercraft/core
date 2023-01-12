package auth

type LoginResult uint8
type AuthType uint8
type ErrorType uint8

//go:generate gcraft_stringer -type=AuthType
//go:generate gcraft_stringer -type=ErrorType
const (
	GruntSuccess                         ErrorType = 0x00
	GruntFailBanned                      ErrorType = 0x03
	GruntFailUnknownAccount              ErrorType = 0x04
	GruntFailIncorrectPassword           ErrorType = 0x05
	GruntFailAlreadyOnline               ErrorType = 0x06
	GruntFailNoTime                      ErrorType = 0x07
	GruntFailDbBusy                      ErrorType = 0x08
	GruntFailVersionInvalid              ErrorType = 0x09
	GruntFailVersionUpdate               ErrorType = 0x0A
	GruntFailInvalidServer               ErrorType = 0x0B
	GruntFailSuspended                   ErrorType = 0x0C
	GruntFailFailNoaccess                ErrorType = 0x0D
	GruntSuccessSurvey                   ErrorType = 0x0E
	GruntFailParentcontrol               ErrorType = 0x0F
	GruntFailLockedEnforced              ErrorType = 0x10
	GruntFailTrialEnded                  ErrorType = 0x11
	GruntFailUseBattlenet                ErrorType = 0x12
	GruntFailAntiIndulgence              ErrorType = 0x13
	GruntFailExpired                     ErrorType = 0x14
	GruntFailNoGameAccount               ErrorType = 0x15
	GruntFailChargeback                  ErrorType = 0x16
	GruntFailInternetGameRoomWithoutBnet ErrorType = 0x17
	GruntFailGameAccountLocked           ErrorType = 0x18
	GruntFailUnlockableLock              ErrorType = 0x19
	GruntFailConversionRequired          ErrorType = 0x20
	GruntFailDisconnected                ErrorType = 0xFF

	LoginOk              LoginResult = 0x00
	LoginFailed          LoginResult = 0x01
	LoginFailed2         LoginResult = 0x02
	LoginBanned          LoginResult = 0x03
	LoginUnknownAccount  LoginResult = 0x04
	LoginUnknownAccount3 LoginResult = 0x05
	LoginAlreadyonline   LoginResult = 0x06
	LoginNotime          LoginResult = 0x07
	LoginDbbusy          LoginResult = 0x08
	LoginBadversion      LoginResult = 0x09
	LoginDownloadFile    LoginResult = 0x0A
	LoginFailed3         LoginResult = 0x0B
	LoginSuspended       LoginResult = 0x0C
	LoginFailed4         LoginResult = 0x0D
	LoginConnected       LoginResult = 0x0E
	LoginParentalcontrol LoginResult = 0x0F
	LoginLockedEnforced  LoginResult = 0x10

	NoValidExpFlag uint8 = 0x0
	PreBcExpFlag   uint8 = 0x1
	PostBcExpFlag  uint8 = 0x2

	LogonChallenge     AuthType = 0x00
	LogonProof         AuthType = 0x01
	ReconnectChallenge AuthType = 0x02
	ReconnectProof     AuthType = 0x03
	RealmList          AuthType = 0x10
	XferInitiate       AuthType = 0x30
	XferData           AuthType = 0x31
	XferAccept         AuthType = 0x32
	XferResume         AuthType = 0x33
	XferCancel         AuthType = 0x34

	RealmGreen  uint8 = 0
	RealmYellow uint8 = 1
	RealmRed    uint8 = 2

	GruntLegacy         uint16 = 0x02
	GruntImprovedSystem uint16 = 0x08
)
