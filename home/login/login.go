package login

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const TokenAge = 12 * time.Hour

// create a unique but
func create_random_token(account_id uint64) string {
	var random_bytes [48]byte
	if _, err := io.ReadFull(rand.Reader, random_bytes[:]); err != nil {
		panic(err)
	}

	now := uint64(time.Now().UnixNano())

	account_id_bytes := make([]byte, binary.MaxVarintLen64)
	account_id_len := binary.PutUvarint(account_id_bytes, account_id)
	account_id_bytes = account_id_bytes[:account_id_len]

	now_bytes := make([]byte, binary.MaxVarintLen64)
	now_len := binary.PutUvarint(now_bytes, now)
	now_bytes = now_bytes[:now_len]

	account_string := base64.RawURLEncoding.EncodeToString(account_id_bytes)
	now_string := base64.RawURLEncoding.EncodeToString(now_bytes)
	random_string := base64.RawURLEncoding.EncodeToString(random_bytes[:])

	return strings.Join([]string{
		account_string,
		now_string,
		random_string,
	}, "-")
}

func token_is_expired(web_token *models.WebToken) bool {
	now := time.Now()
	return web_token.Expiry.Before(now)
}

func expire_token(home_db *database.Container, web_token *models.WebToken) {
	home_db.Table("WebToken").Where(query.Eq("Token", web_token.Token)).Delete()
}

func LoginWeb(home_db *database.Container, method models.LoginMethod, username, password, address, user_agent string) (token string, err error) {
	var (
		account models.Account
		found   bool
	)

	found, err = home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/login: could not access database")
		return
	}

	if !found {
		err = fmt.Errorf("home/login: username/password is incorrect")
		return
	}

	if account.Banned {
		err = fmt.Errorf("home/login: account has been banned")
		return
	}

	if !Verify(username, password, account.Identity_Bcrypt) {
		err = fmt.Errorf("home/login: username/password is incorrect")
		return
	}

	if account.Locked {
		err = fmt.Errorf("home/login: account is locked")
		return
	}

	token = create_random_token(account.ID)

	var web_token models.WebToken
	web_token.Token = token
	web_token.Account = account.ID
	web_token.Method = method
	web_token.Expiry = time.Now().Add(TokenAge)
	web_token.Address = address
	web_token.UserAgent = user_agent

	if account.Authenticator {
		web_token.Status = auth.WebTokenStatus_AUTH_NEEDED
	} else {
		web_token.Status = auth.WebTokenStatus_AUTHENTICATED
	}

	if err = home_db.Table("WebToken").Insert(&web_token); err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/login: could not save web token")
		return
	}

	return
}

func LogoutWeb(home_db *database.Container, method models.LoginMethod, credential, address, user_agent string) (logged_out bool, err error) {
	var deleted uint64
	deleted, err = home_db.Table("WebToken").Where(query.Eq("Token", credential)).Delete()
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/login: could not access database")
		return
	}

	logged_out = deleted != 0
	return
}

func CheckCredential(home_db *database.Container, method models.LoginMethod, credential, address, user_agent string) (account *models.Account, status auth.WebTokenStatus, err error) {
	var (
		found bool
		token models.WebToken
	)

	found, err = home_db.Table("WebToken").Where(query.Eq("Token", credential)).Get(&token)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/login: cannot connect to database right now")
		return
	}

	if !found {
		err = fmt.Errorf("home/login: token is expired or invalid")
		return
	}

	if token_is_expired(&token) {
		expire_token(home_db, &token)
		err = fmt.Errorf("home/provider/core/account: token is expired or invalid")
		return
	}

	status = token.Status

	account = new(models.Account)

	found, err = home_db.Table("Account").Where(query.Eq("ID", token.Account)).Get(account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/account: error getting token account from database")
		return
	}
	if !found {
		expire_token(home_db, &token)
		err = fmt.Errorf("home/provider/web: token is expired or invalid")
		return
	}

	// Banned accounts are not allowed to use the account service
	if account.Banned {
		expire_token(home_db, &token)
		err = fmt.Errorf("home/provider/web: token is expired or invalid")
		return
	}

	// Suspended accounts may continue to use the web API
	if account.Suspended {
		// Check if suspension has expired and update account
		now := time.Now()
		suspension_expired := account.UnsuspendAt.Compare(now) == -1

		if suspension_expired {
			account.Suspended = false
			account.UnsuspendAt = time.Time{}
			if _, err = home_db.Table("Account").Where(query.Eq("ID", account.ID)).Columns("Suspended", "UnsuspendAt").Update(&account); err != nil {
				panic(err)
			}
		}
	}

	return
}

var totp_options = totp.ValidateOpts{
	Period:    30,
	Skew:      1,
	Digits:    otp.DigitsSix,
	Algorithm: otp.AlgorithmSHA1,
}

func AuthenticateCredential(home_db *database.Container, login_method models.LoginMethod, two_factor_method auth.TwoFactorAuthenticationMethod, token *models.WebToken, two_factor_password, address, user_agent string) (authenticated bool, err error) {
	var (
		account models.Account
		found   bool
	)

	t := time.Now()

	found, err = home_db.Table("Account").Where(query.Eq("ID", token.Account)).Get(&account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/login: cannot connect to database")
		return
	}

	if !found {
		err = fmt.Errorf("home/login: token does not correspond to an existing account")
		return
	}

	if account.Banned {
		err = fmt.Errorf("home/login: cannot authenticate, account is banned")
		return
	}

	switch two_factor_method {
	case auth.TwoFactorAuthenticationMethod_EMAIL:
		if !account.EmailVerified {
			err = fmt.Errorf("home/login: no verified email")
			return
		}

		if time.Since(account.EmailVerificationCodeSentAt) > (5 * time.Minute) {
			err = fmt.Errorf("home/login: email verification code is expired or was never sent")
			return
		}
		authenticated = account.EmailVerificationCode != two_factor_password
	case auth.TwoFactorAuthenticationMethod_TOTP:
		if !account.Authenticator {
			err = fmt.Errorf("home/login: you cannot authenticate using TOTP without being enrolled in TOTP")
			return
		}

		secret := base32.StdEncoding.EncodeToString(account.AuthenticatorSecret)

		authenticated, err = totp.ValidateCustom(two_factor_password, secret, t, totp_options)
		if err != nil {
			return
		}
	default:
		err = fmt.Errorf("home/login: cannot authenticate, method %s is unknown", two_factor_method)
		return
	}

	if authenticated {
		token.Status = auth.WebTokenStatus_AUTHENTICATED
		var updated uint64
		updated, err = home_db.Table("WebToken").Where(query.Eq("Token", token.Token)).Columns("Status").Update(token)
		if err != nil {
			log.Warn(err)
			err = fmt.Errorf("home/login: failed to connect to database")
			return
		}

		if updated == 0 {
			err = fmt.Errorf("home/login: could not find any token to be authenticated")
			return
		}
	}

	return
}

func ValidateEnrollmentSecret(secret, password string) bool {
	valid, err := totp.ValidateCustom(password, secret, time.Now(), totp_options)
	if err != nil {
		log.Warn(err)
		return false
	}
	return valid
}
