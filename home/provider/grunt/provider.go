package grunt

import (
	"crypto/subtle"
	"encoding/base32"
	"fmt"
	"strings"
	"time"

	"github.com/Gophercraft/core/format/tag"
	"github.com/Gophercraft/core/grunt"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type ServiceProviderConfig struct {
	Mandatory2FA bool
}

type service_provider struct {
	config *ServiceProviderConfig
	db     *database.Container
}

func (*service_provider) Check(evt *grunt.Event) error {
	return nil
}

func (provider *service_provider) get_account(account_name string, account_record *models.Account) (found bool, err error) {
	// query account from database by username
	found, err = provider.db.
		Table("Account").
		Where(
			query.Eq("Username", strings.ToLower(account_name))).
		Get(account_record)
	return
}

func (provider *service_provider) unsuspend_account(account_id uint64) (err error) {
	var updated uint64
	updated, err = provider.db.
		Table("Account").
		Where(
			query.Eq("ID", account_id),
		).
		Columns("Suspended", "UnsuspendAt").
		UpdateColumns(false, time.Time{})
	if err != nil {
		return
	}

	if updated == 0 {
		err = fmt.Errorf("failed to unsuspend account %d (Not found)", account_id)
	}

	return
}

func (provider *service_provider) unsuspend_game_account(game_account_id uint64) (err error) {
	var updated uint64
	updated, err = provider.db.
		Table("GameAccount").
		Where(
			query.Eq("ID", game_account_id),
		).
		Columns("Suspended", "UnsuspendAt").
		UpdateColumns(false, time.Time{})
	if err != nil {
		return
	}

	if updated == 0 {
		err = fmt.Errorf("failed to unsuspend game account %d (Not found)", game_account_id)
	}

	return
}

func (provider *service_provider) get_active_game_account(account_id uint64, game_account *models.GameAccount) (found bool, err error) {
	// query account from database by username
	found, err = provider.db.
		Table("GameAccount").
		Where(
			query.Eq("Owner", account_id),
			query.Eq("Active", true),
		).Get(game_account)
	return
}

func (provider *service_provider) GetAccountLoginInfo(account_name string, build version.Build, info *grunt.AccountLoginInfo) (err error) {
	var (
		account      models.Account
		game_account models.GameAccount
		found        bool
	)
	found, err = provider.get_account(account_name, &account)
	if err != nil {
		err = grunt.LoginErrorf(grunt.LoginDbbusy, "home/provider/grunt: service provider: query of account failed: %w", err)
		return
	}
	if !found {
		err = grunt.LoginErrorf(grunt.LoginUnknownAccount, "home/provider/grunt: service provider: account not found")
		return
	}

	if account.Banned {
		err = grunt.LoginErrorf(grunt.LoginBanned, "home/provider/grunt: account is banned")
		return
	}

	if account.Suspended {
		suspension_expired := account.UnsuspendAt.Compare(time.Now()) == -1
		if suspension_expired {
			provider.unsuspend_account(account.ID)
		} else {
			err = grunt.LoginErrorf(grunt.LoginSuspended, "home/provider/grunt: account is banned")
			return
		}
	}

	// TODO: send email code here
	any_2fa := account.Authenticator || account.EmailVerified

	if !any_2fa && provider.config.Mandatory2FA {
		err = grunt.LoginErrorf(grunt.LoginAccountIsLockedButUnlockable, "home/provider/grunt: use of authenticator is mandatory, and account %d has not configured their authenticator", account.ID)
		return
	}

	// Check if any game account is active. If not, we cannot proceed with login: the server doesn't have a profile to save game data.
	found, err = provider.get_active_game_account(account.ID, &game_account)
	if err != nil {
		err = grunt.LoginErrorf(grunt.LoginDbbusy, "home/provider/grunt: service provider: query of game account failed: %w", err)
		return
	}
	if !found {
		err = grunt.LoginErrorf(grunt.LoginAccountIsLockedButUnlockable, "home/provider/grunt: no active game account on record for %d", account.ID)
		return
	}
	// If game account is either banned or suspended, fail in the same way as if it were an account-wide ban
	if game_account.Banned {
		err = grunt.LoginErrorf(grunt.LoginBanned, "home/provider/grunt: account is banned")
		return
	}
	if game_account.Suspended {
		suspension_expired := game_account.UnsuspendAt.Compare(time.Now()) == -1
		if suspension_expired {
			provider.unsuspend_game_account(game_account.ID)
		} else {
			err = grunt.LoginErrorf(grunt.LoginSuspended, "home/provider/grunt: game_account is suspended")
			return
		}
	}

	if build < 8606 {
		info.PIN = account.Authenticator
	} else {
		info.Authenticator = account.Authenticator
	}

	info.IdentityHash = account.Identity_SHA_1
	info.SessionKey = account.SessionKey

	return nil
}

var totp_options = totp.ValidateOpts{
	Period:    30,
	Skew:      1,
	Digits:    otp.DigitsSix,
	Algorithm: otp.AlgorithmSHA1,
}

func (provider *service_provider) VerifyAuthenticatorCode(account_name, authenticator_code string) (valid bool) {
	t := time.Now()

	var (
		account models.Account
		found   bool
		err     error
	)
	found, err = provider.get_account(account_name, &account)
	if err != nil {
		return
	}

	if !found {
		return
	}

	any_2fa := account.EmailVerified || account.Authenticator

	if !any_2fa {
		return
	}

	// If account has authenticator, it takes precedence over email code
	if account.Authenticator {
		secret := base32.StdEncoding.EncodeToString(account.AuthenticatorSecret)

		valid, err = totp.ValidateCustom(authenticator_code, secret, t, totp_options)
		if err != nil {
			return
		}
	} else if account.EmailVerified {
		if time.Since(account.EmailVerificationCodeSentAt) > 5*time.Minute {
			return
		}
		valid = authenticator_code == account.EmailVerificationCode
	}

	return
}

func (provider *service_provider) VerifyPIN(account_name string, server_pin_info *grunt.ServerPINInfo, client_pin_info *grunt.ClientPINInfo) (valid bool) {
	t := time.Now()

	var (
		account models.Account
		found   bool
		err     error
	)
	found, err = provider.get_account(account_name, &account)
	if err != nil {
		return
	}

	if !found {
		return
	}

	any_2fa := account.EmailVerified || account.Authenticator

	if !any_2fa {
		return
	}

	// If account has authenticator, it takes precedence over email code
	if account.Authenticator {
		secret := base32.StdEncoding.EncodeToString(account.AuthenticatorSecret)

		valid, err = totp_validate_pin(server_pin_info, client_pin_info.Salt[:], client_pin_info.Proof[:], secret, t, totp_options)
		if err != nil {
			log.Warn(err)
			return
		}

		// scrambled_digits := grunt.ScramblePINNumber(server_pin_info.GridSeed, []byte("4444"))
		// correct_digits := grunt.GetPINProof(server_pin_info.Salt[:], scrambled_digits, client_pin_info.Salt[:])
		// valid = subtle.ConstantTimeCompare(client_pin_info.Proof[:], correct_digits) == 1
	} else if account.EmailVerified {
		if t.Sub(account.EmailVerificationCodeSentAt) > 5*time.Minute {
			return
		}
		scrambled_pin := grunt.ScramblePINNumber(server_pin_info.GridSeed, []byte(account.EmailVerificationCode))
		hashed_email_verification_code := grunt.GetPINProof(server_pin_info.Salt[:], scrambled_pin, client_pin_info.Salt[:])
		valid = subtle.ConstantTimeCompare(hashed_email_verification_code, client_pin_info.Proof[:]) == 1
	}

	return
}

func (provider *service_provider) GetRealmList(account_name string, build version.Build) (grunt_list []grunt.Realm, err error) {
	var account models.Account
	found, err := provider.get_account(account_name, &account)
	if err != nil {
		err = fmt.Errorf("home/provider/grunt: error getting account for realm list %w", err)
		return
	}
	if !found {
		err = fmt.Errorf("home/provider/grunt: account '%s' not found", account_name)
		return
	}

	var realms []models.Realm
	err = provider.db.
		Table("Realm").
		Where(query.Lte("RequiredTier", account.Tier)).
		Find(&realms)
	if err != nil {
		return
	}
	// Loop through the list of announced realms,
	// building the Grunt realmlist presentation from it
	for _, realm := range realms {
		var grunt_realm grunt.Realm
		grunt_realm.Population = 100.0
		grunt_realm.Name = realm.Name
		grunt_realm.Address = realm.Address
		// TODO:
		// incorrect category can cause the realm not not appear at all
		// change the value to a "known good"
		// to avoid confusion
		grunt_realm.Category = uint8(realm.Category)
		// grunt_realm.Flags = grunt.RealmHasBuildInfo
		grunt_list = append(grunt_list, grunt_realm)
	}
	return
}

func create_platform_id(os grunt.OS, arch grunt.Architecture) string {
	prefix := "??"
	switch os {
	case grunt.Linux:
		prefix = "Ln"
	case grunt.MacOS:
		prefix = "Mc"
	case grunt.Windows:
		prefix = "Wn"
	}
	suffix := "??"
	switch arch {
	case grunt.X86:
		suffix = "32"
	case grunt.X64:
		suffix = "64"
	case grunt.ARMv6:
		suffix = "A6"
	case grunt.ARMv7:
		suffix = "A7"
	case grunt.PowerPC:
		suffix = "PP"
	}
	return prefix + suffix
}

func (provider *service_provider) StoreAccountInfo(account_info *grunt.AccountInfo) error {
	// convert grunt format to consistent format

	// account must be registered
	var account models.Account
	found, err := provider.get_account(account_info.AccountName, &account)
	if err != nil {
		return fmt.Errorf("grunt: service provider: query failed: %w", err)
	}
	if !found {
		return fmt.Errorf("grunt: service provider: account not found")
	}

	account.Locale = tag.Clean(tag.Tag(account_info.Locale))
	account.OS = tag.Clean(tag.Tag(account_info.OS))
	account.Architecture = tag.Clean(tag.Tag(account_info.Architecture))
	account.SessionKey = account_info.SessionKey
	account.Platform = create_platform_id(account_info.OS, account_info.Architecture)

	_, err = provider.db.
		Table("Account").
		Where(
			query.Eq("ID", account.ID)).
		Columns(
			"SessionKey",
			"Locale",
			"OS",
			"Architecture",
			"Platform").
		Update(&account)

	if err != nil {
		return err
	}

	return nil
}

func New(config *ServiceProviderConfig, db *database.Container) (provider *service_provider) {
	provider = new(service_provider)
	provider.config = config
	provider.db = db
	return provider
}
