package rest

import (
	"crypto/subtle"
	"encoding/base32"
	"fmt"
	"time"

	pb_login "github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/service/bnet_rest"
	"github.com/Gophercraft/phylactery/database/query"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var totp_options = totp.ValidateOpts{
	Period:    30,
	Skew:      1,
	Digits:    otp.DigitsSix,
	Algorithm: otp.AlgorithmSHA1,
}

func (provider *Provider) LoginAuthenticator(user_info *bnet_rest.UserInfo, login_form *pb_login.LoginForm) (login_result *pb_login.LoginResult, err error) {
	begin := time.Now()

	authenticator_input := get_login_form_input(login_form, "authenticator_input")
	if authenticator_input == "" {
		err = fmt.Errorf("no authenticator input")
		return
	}
	var (
		account models.Account
		session *identity_session
		found   bool
	)

	session = provider.checkout_identity(user_info.SessionID)
	if session == nil {
		err = fmt.Errorf("no session")
		return
	}

	session.guard.Lock()
	defer session.guard.Unlock()

	if session.authenticated {
		err = fmt.Errorf("already authenticated")
		return
	}

	if !session.identified {
		err = fmt.Errorf("not identified")
	}

	found, err = provider.home_db.Table("Account").Where(query.Eq("ID", session.account_ID)).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = fmt.Errorf("account not found")
		return
	}

	if account.Authenticator {
		secret := base32.StdEncoding.EncodeToString(account.AuthenticatorSecret)

		session.authenticated, err = totp.ValidateCustom(authenticator_input, secret, begin, totp_options)
		if err != nil {
			return
		}
	} else if account.EmailVerified {
		if begin.Sub(account.EmailVerificationCodeSentAt) > 5*time.Minute {
			err = fmt.Errorf("email verification code expired")
			return
		}

		session.authenticated = subtle.ConstantTimeCompare([]byte(account.EmailVerificationCode), []byte(authenticator_input)) == 1
	}

	var new_ticket string

	if session.authenticated {
		login_result = session.pending_login_result
		session.pending_login_result = nil
		new_ticket, err = provider.create_ticket(account.ID)
		if err != nil {
			return
		}
		util.Set(&login_result.LoginTicket, new_ticket)
	} else {
		err = fmt.Errorf("authentication failed")
	}

	return
}
