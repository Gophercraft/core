package web

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/home/login"
	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/home/service/web/models"
	api_models "github.com/Gophercraft/core/home/service/web/models"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *service_provider) NewRegistrationChallenge(user_info *api_models.UserInfo) (registration_challenge *api_models.RegistrationChallenge, err error) {
	// TODO: check IP

	if !provider.config.WebRegistration {
		err = fmt.Errorf("home/provider/web: server has disabled web registration")
		return
	}

	registration_challenge = new(models.RegistrationChallenge)

	if provider.config.UseCaptchas {
		registration_challenge.CaptchaID, err = provider.captcha_manager.new_challenge()
		if err != nil {
			return
		}
	}

	registration_challenge.MaxEmailLength = provider.config.MaxEmailLength
	registration_challenge.MaxUsernameLength = provider.config.MaxUsernameLength
	registration_challenge.MaxPasswordLength = provider.config.MaxPasswordLength

	return
}

func (provider *service_provider) Register(user_info *api_models.UserInfo, registration_request *api_models.RegistrationRequest) (registration_response *api_models.RegistrationResponse, err error) {
	if !provider.config.WebRegistration {
		err = fmt.Errorf("home/provider/web: server has disabled web registration")
		return
	}

	username := strings.ToLower(registration_request.Username)
	email := strings.ToLower(registration_request.Email)
	password := registration_request.Password
	if provider.config.UseCaptchas {
		valid := provider.captcha_manager.verify(registration_request.CaptchaID, registration_request.CaptchaSolution)
		if !valid {
			err = fmt.Errorf("invalid captcha")
			return
		}
	}

	if username == "" {
		err = fmt.Errorf("account name required")
		return
	}

	if password == "" {
		err = fmt.Errorf("password required")
		return
	}

	if len(email) > int(provider.config.MaxEmailLength) {
		err = fmt.Errorf("e-mail address is too long")
		return
	}

	if len(username) > int(provider.config.MaxUsernameLength) {
		err = fmt.Errorf("account name is too long")
		return
	}

	if len(password) > int(provider.config.MaxPasswordLength) {
		err = fmt.Errorf("password is too long")
		return
	}

	registration_response = new(models.RegistrationResponse)

	var (
		account home_models.Account
		found   bool
	)

	if provider.config.EmailRequired {
		if email == "" {
			err = fmt.Errorf("a valid email is required to register")
			return
		}
	}

	if email != "" {
		found, err = provider.home_db.Table("Account").Where(query.Eq("Email", email)).Get(&account)
		if err != nil {
			log.Warn("error accessing account table", err)
			err = fmt.Errorf("internal server error")
			return
		}

		if found {
			err = fmt.Errorf("e-mail address is already in use")
			return
		}
	}

	err = login.RegisterAccount(provider.home_db, false, auth.AccountTier_NORMAL, email, username, registration_request.Password, auth.AccountTier_NORMAL)
	if err != nil {
		return
	}

	registration_response.EmailVerificationNeeded = provider.config.EmailVerificationNeeded

	return
}
