package grunt

import "fmt"

type LoginResult uint8

const (
	LoginOk                           LoginResult = 0x00
	LoginFailed                       LoginResult = 0x01
	LoginFailed2                      LoginResult = 0x02
	LoginBanned                       LoginResult = 0x03
	LoginUnknownAccount               LoginResult = 0x04
	LoginUnknownAccount3              LoginResult = 0x05
	LoginAlreadyonline                LoginResult = 0x06
	LoginNotime                       LoginResult = 0x07
	LoginDbbusy                       LoginResult = 0x08
	LoginBadversion                   LoginResult = 0x09
	LoginDownloadFile                 LoginResult = 0x0A
	LoginFailed3                      LoginResult = 0x0B
	LoginSuspended                    LoginResult = 0x0C
	LoginFailed4                      LoginResult = 0x0D
	LoginConnected                    LoginResult = 0x0E
	LoginParentalcontrol              LoginResult = 0x0F
	LoginLockedEnforced               LoginResult = 0x10
	LoginTrialExpired                 LoginResult = 0x11
	LoginMigratedToBnetAccount        LoginResult = 0x12
	LoginTemporaryClosureChargeback   LoginResult = 0x16
	LoginMergeWithIGRTime             LoginResult = 0x17
	LoginAccountAccessDisabled        LoginResult = 0x18
	LoginAccountIsLockedButUnlockable LoginResult = 0x19
	LoginYouMustMigrateToBnetAccount  LoginResult = 0x20
)

type LoginError struct {
	result LoginResult
	err    error
}

func (err LoginError) Result() LoginResult {
	return err.result
}

func (err LoginError) Error() string {
	return fmt.Errorf("grunt: LoginResult(%d): %w", err.result, err.err).Error()
}

func LoginErrorf(code LoginResult, format string, args ...any) LoginError {
	return LoginError{
		result: code,
		err:    fmt.Errorf(format, args...),
	}
}
