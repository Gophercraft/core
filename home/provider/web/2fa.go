package web

import (
	"encoding/base32"
	"fmt"

	"github.com/Gophercraft/core/home/login"
	home_models "github.com/Gophercraft/core/home/models"
	api_models "github.com/Gophercraft/core/home/service/web/models"

	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *service_provider) AuthenticateCredential(user_info *api_models.UserInfo, authenticate_credential_request *api_models.CredentialAuthenticationRequest) (authenticate_credential_response *api_models.CredentialAuthenticationResponse, err error) {
	// Find the token
	var (
		web_token home_models.WebToken
		found     bool
	)

	found, err = provider.home_db.Table("WebToken").Where(query.Eq("Token", user_info.Token)).Get(&web_token)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: cannot connect to database")
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/web: cannot find credential")
		return
	}

	authenticate_credential_response = new(api_models.CredentialAuthenticationResponse)

	if web_token.Status == auth.WebTokenStatus_AUTHENTICATED {
		// already authenticated!
		authenticate_credential_response.Authenticated = true
		return
	}

	// Find the 2fa method
	twofa_enum_value, ok := auth.TwoFactorAuthenticationMethod_value[authenticate_credential_request.TwoFactorAuthenticationMethod]
	if !ok {
		err = fmt.Errorf("home/provider/web: invalid 2fa method")
		return
	}

	two_factor_method := auth.TwoFactorAuthenticationMethod(twofa_enum_value)

	authenticate_credential_response.Authenticated, err = login.AuthenticateCredential(provider.home_db, home_models.Login_Web, two_factor_method, &web_token, authenticate_credential_request.AuthenticatorPassword, user_info.Address, user_info.UserAgent)
	return
}

func (provider *service_provider) EnrollTwoFactorAuthentication(auth_info *api_models.UserInfo, enrollment_request *api_models.TwoFactorAuthenticationEnrollmentRequest) (enrollment_response *api_models.TwoFactorAuthenticationEnrollmentResponse, err error) {
	var (
		account *home_models.Account
		status  auth.WebTokenStatus
	)

	account, status, err = provider.get_account(auth_info)
	if err != nil {
		return
	}

	if status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/web: not authenticated")
		return
	}

	var decoded_secret []byte

	decoded_secret, err = base32.StdEncoding.DecodeString(enrollment_request.Secret)
	if err != nil {
		return
	}

	if len(decoded_secret) != 24 {
		err = fmt.Errorf("home/provider/web: secret must be 24 bytes long")
		return
	}

	if len(enrollment_request.Password) != 6 {
		err = fmt.Errorf("home/provider/web: password must be 6 digits")
		return
	}

	valid := login.ValidateEnrollmentSecret(enrollment_request.Secret, enrollment_request.Password)

	enrollment_response = new(api_models.TwoFactorAuthenticationEnrollmentResponse)

	if valid {
		account.Authenticator = true
		account.AuthenticatorSecret = decoded_secret
		var updated uint64
		updated, err = provider.home_db.Table("Account").Where(query.Eq("ID", account.ID)).Columns("Authenticator", "AuthenticatorSecret").Update(account)
		if err != nil {
			log.Warn(err)
			err = fmt.Errorf("home/provider/web: cannot connect to database")
			return
		}

		if updated == 0 {
			err = fmt.Errorf("home/provider/web: cannot account to be updated")
			return
		}
		enrollment_response.Enrolled = true
	} else {
		enrollment_response.Enrolled = false
	}

	return
}

func (provider *service_provider) GetTwoFactorAuthenticationMethods(auth_info *api_models.UserInfo) (two_factor_methods_list *api_models.TwoFactorAuthMethods, err error) {
	var (
		account *home_models.Account
		// status  auth.WebTokenStatus
	)

	account, _, err = provider.get_account(auth_info)
	if err != nil {
		return
	}

	two_factor_methods_list = new(api_models.TwoFactorAuthMethods)
	if account.Authenticator {
		two_factor_methods_list.Methods = append(two_factor_methods_list.Methods, auth.TwoFactorAuthenticationMethod_TOTP.String())
	}

	if account.EmailVerified {
		two_factor_methods_list.Methods = append(two_factor_methods_list.Methods, auth.TwoFactorAuthenticationMethod_EMAIL.String())
	}

	return
}
func (provider *service_provider) SendTwoFactorAuthenticationCodeEmail(user_info *api_models.UserInfo) (send_2fa_code_email_response *api_models.SentTwoFactorAuthenticationCodeEmailResponse, err error) {
	send_2fa_code_email_response = new(api_models.SentTwoFactorAuthenticationCodeEmailResponse)
	send_2fa_code_email_response.Sent = false
	return
}
