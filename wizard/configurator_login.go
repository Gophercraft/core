package wizard

import (
	"context"
	"fmt"

	"github.com/Gophercraft/core/home/protocol/pb/account"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
)

func (c *Configurator) Login(login_request *auth.LoginRequest) (err error) {
	var login_response *auth.LoginResponse
	auth_client := auth.NewAuthServiceClient(c.home_connection)
	login_response, err = auth_client.Login(context.Background(), login_request)
	if err != nil {
		return
	}

	c.login_info.WebToken = login_response.GetWebToken()

	var credential_status *auth.CredentialStatus
	credential_status, err = c.GetCredentialStatus()
	if err != nil {
		return
	}

	if credential_status.WebTokenStatus == auth.WebTokenStatus_AUTHENTICATED {
		c.login_info.LoggedIn = true
		if err = c.write_login_info(); err != nil {
			return
		}
	}

	return
}

func (c *Configurator) Logout() (err error) {
	if c.Connected() {
		if c.LoggedIn() {
			var logout_request auth.LogoutRequest
			logout_request.Credential = c.login_info.WebToken
			auth_client := auth.NewAuthServiceClient(c.home_connection)
			auth_client.Logout(context.Background(), &logout_request)
		}
	}

	c.login_info = login_info{}

	if err = c.write_login_info(); err != nil {
		return
	}

	return
}

func (c *Configurator) LoggedIn() bool {
	return c.login_info.LoggedIn
}

func (c *Configurator) GetCredentialStatus() (credential_status *auth.CredentialStatus, err error) {
	auth_client := auth.NewAuthServiceClient(c.home_connection)

	var credential_status_request auth.CredentialStatusRequest
	credential_status_request.Credential = c.login_info.WebToken

	credential_status, err = auth_client.GetCredentialStatus(context.Background(), &credential_status_request)
	return
}

func (c *Configurator) GetTwoFactorAuthenticationMethods() (methods *auth.TwoFactorAuthenticationMethodsResponse, err error) {
	auth_client := auth.NewAuthServiceClient(c.home_connection)

	if c.login_info.WebToken == "" {
		err = fmt.Errorf("wizard: cannot get two factor methods without credential")
		return
	}

	var two_factor_auth_methods_request auth.TwoFactorAuthenticationMethodsRequest
	two_factor_auth_methods_request.Credential = c.login_info.WebToken

	methods, err = auth_client.GetTwoFactorAuthenticationMethods(context.Background(), &two_factor_auth_methods_request)
	return
}

func (c *Configurator) AuthenticateCredential(method auth.TwoFactorAuthenticationMethod, code string) (err error) {
	auth_client := auth.NewAuthServiceClient(c.home_connection)

	if c.login_info.WebToken == "" {
		err = fmt.Errorf("wizard: cannot authenticate credential without a token")
		return
	}

	var (
		credential_authentication_request  auth.CredentialAuthenticationRequest
		credential_authentication_response *auth.CredentialAuthenticationResponse
	)

	credential_authentication_request.Credential = c.login_info.WebToken
	credential_authentication_request.TwoFactorPassword = code
	credential_authentication_request.TwoFactorMethod = method

	credential_authentication_response, err = auth_client.AuthenticateCredential(context.Background(), &credential_authentication_request)
	if err != nil {
		return
	}

	if !credential_authentication_response.Authenticated {
		err = fmt.Errorf("incorrect authenticator code")
		return
	}

	c.login_info.LoggedIn = true
	if err = c.write_login_info(); err != nil {
		return
	}

	return
}

func (c *Configurator) GetLoggedInAccountStatus() (account_status *account.AccountStatus, err error) {
	if !c.LoggedIn() {
		err = fmt.Errorf("wizard: cannot get account status if not logged in")
		return
	}

	var credential_status *auth.CredentialStatus
	credential_status, err = c.GetCredentialStatus()
	if err != nil {
		return
	}

	return c.GetAccountStatus(credential_status.AccountId)
}

func (c *Configurator) GetAccountStatus(id uint64) (account_status *account.AccountStatus, err error) {
	account_client := account.NewAccountServiceClient(c.home_connection)

	if !c.LoggedIn() {
		err = fmt.Errorf("wizard: cannot retrieve account status without being logged in")
		return
	}

	var account_status_request account.AccountStatusRequest
	account_status_request.Id = id
	account_status_request.Credential = c.login_info.WebToken

	return account_client.GetAccountStatus(context.Background(), &account_status_request)
}

func (c *Configurator) GenerateLoginTicket() (ticket string, err error) {
	auth_client := auth.NewAuthServiceClient(c.home_connection)

	if !c.LoggedIn() {
		err = fmt.Errorf("wizard: cannot authenticate logged in")
		return
	}

	var (
		generate_login_ticket_request  auth.GenerateLoginTicketRequest
		generate_login_ticket_response *auth.GenerateLoginTicketResponse
	)

	generate_login_ticket_request.Credential = c.login_info.WebToken

	generate_login_ticket_response, err = auth_client.GenerateLoginTicket(context.Background(), &generate_login_ticket_request)
	if err != nil {
		return
	}

	ticket = generate_login_ticket_response.Ticket
	return
}
