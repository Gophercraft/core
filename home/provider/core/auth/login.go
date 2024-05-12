package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *auth_provider) Login(ctx context.Context, sign_in *auth.LoginRequest) (login_response *auth.LoginResponse, err error) {
	var address string
	address, err = protocol.GetPeerAddress(ctx)
	if err != nil {
		return
	}

	login_response = new(auth.LoginResponse)
	login_response.WebToken, err = login.LoginWeb(provider.home_db, models.Login_Core, sign_in.AccountName, sign_in.Password, address, "")

	return
}

func (provider *auth_provider) Logout(ctx context.Context, logout_request *auth.LogoutRequest) (logout_response *auth.LogoutResponse, err error) {
	var address string
	address, err = protocol.GetPeerAddress(ctx)
	if err != nil {
		return
	}

	logout_response = new(auth.LogoutResponse)
	logout_response.LoggedOut, err = login.LogoutWeb(provider.home_db, models.Login_Core, logout_request.Credential, address, "")

	return
}

func (provider *auth_provider) GetCredentialStatus(ctx context.Context, credential_status_request *auth.CredentialStatusRequest) (credential_status *auth.CredentialStatus, err error) {
	var (
		web_token models.WebToken
		account   models.Account
		found     bool
	)
	// get web token
	found, err = provider.home_db.Table("WebToken").Where(query.Eq("Token", credential_status_request.Credential)).Get(&web_token)
	if err != nil {
		err = fmt.Errorf("home/provider/core/auth: could not access database")
		return
	}
	if !found {
		err = fmt.Errorf("home/provider/core/auth: not logged in")
		return
	}

	// check existence of account
	found, err = provider.home_db.Table("Account").Where(query.Eq("ID", web_token.Account)).Get(&account)
	if err != nil {
		err = fmt.Errorf("home/provider/core/auth: could not access database")
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/core/auth: no account (%d) backing this WebToken", web_token.Account)
		return
	}

	credential_status = new(auth.CredentialStatus)
	credential_status.AccountId = account.ID

	// If account is banned, then credential is invalid
	if account.Banned {
		credential_status.WebTokenStatus = auth.WebTokenStatus_LOGGED_OUT
		return
	}

	credential_status.WebTokenStatus = web_token.Status
	if credential_status.WebTokenStatus == auth.WebTokenStatus_AUTH_NEEDED {
		if account.Authenticator {
			credential_status.TwoFactorAuthenticationMethod = auth.TwoFactorAuthenticationMethod_TOTP
		} else if account.EmailVerified {
			credential_status.TwoFactorAuthenticationMethod = auth.TwoFactorAuthenticationMethod_EMAIL
		} else {
			credential_status.TwoFactorAuthenticationMethod = auth.TwoFactorAuthenticationMethod_NONE
		}
	}

	return
}

func (provider *auth_provider) GetTwoFactorAuthenticationMethods(ctx context.Context, two_factor_methods_request *auth.TwoFactorAuthenticationMethodsRequest) (methods_list *auth.TwoFactorAuthenticationMethodsResponse, err error) {
	var (
		account *models.Account
	)
	account, _, err = provider.get_account(ctx, two_factor_methods_request.Credential)
	if err != nil {
		return
	}

	methods_list = new(auth.TwoFactorAuthenticationMethodsResponse)
	if account.Authenticator {
		methods_list.Methods = append(methods_list.Methods, auth.TwoFactorAuthenticationMethod_TOTP)
	}

	if account.EmailVerified {
		methods_list.Methods = append(methods_list.Methods, auth.TwoFactorAuthenticationMethod_EMAIL)
	}

	return
}

func (provider *auth_provider) SendEmailAuthenticationCode(ctx context.Context, send_email_request *auth.SendEmailAuthenticationCodeRequest) (send_email_response *auth.SendEmailAuthenticationCodeResponse, err error) {
	send_email_response = new(auth.SendEmailAuthenticationCodeResponse)
	send_email_response.Sent = false
	return
}

func (provider *auth_provider) AuthenticateCredential(ctx context.Context, credential_authentication_request *auth.CredentialAuthenticationRequest) (credential_authentication_response *auth.CredentialAuthenticationResponse, err error) {
	var (
		address   string
		web_token models.WebToken
		status    auth.WebTokenStatus
		found     bool
	)
	address, err = protocol.GetPeerAddress(ctx)
	if err != nil {
		return
	}

	credential_authentication_response = new(auth.CredentialAuthenticationResponse)

	_, status, err = provider.get_account(ctx, credential_authentication_request.Credential)
	if err != nil {
		return
	}

	if status == auth.WebTokenStatus_AUTHENTICATED {
		credential_authentication_response.Authenticated = true
		return
	}

	// get web token
	found, err = provider.home_db.Table("WebToken").Where(query.Eq("Token", credential_authentication_request.Credential)).Get(&web_token)
	if err != nil {
		err = fmt.Errorf("home/provider/core/auth: could not access database")
		return
	}
	if !found {
		err = fmt.Errorf("home/provider/core/auth: not logged in")
		return
	}

	credential_authentication_response.Authenticated, err = login.AuthenticateCredential(provider.home_db, models.Login_Core, credential_authentication_request.TwoFactorMethod, &web_token, credential_authentication_request.TwoFactorPassword, address, "")
	return
}

func (provider *auth_provider) GenerateLoginTicket(ctx context.Context, generate_login_ticket_request *auth.GenerateLoginTicketRequest) (generate_login_ticket_response *auth.GenerateLoginTicketResponse, err error) {
	var (
		account *models.Account
		status  auth.WebTokenStatus
	)
	account, status, err = provider.get_account(ctx, generate_login_ticket_request.Credential)
	if err != nil {
		return
	}

	if status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/core/auth: you must be fully authenticated to generate a login ticket")
		return
	}

	generate_login_ticket_response = new(auth.GenerateLoginTicketResponse)
	generate_login_ticket_response.Ticket, err = login.CreateTicket(provider.home_db, account.ID, time.Now().Add(7*24*time.Hour))
	return
}
