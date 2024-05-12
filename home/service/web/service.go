package web

import (
	"net"
	"net/http"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/home/service/web/models"
	"github.com/Gophercraft/log"
)

type Service struct {
	config        *ServiceConfig
	provider      ServiceProvider
	web_serve_mux *http.ServeMux
	listener      net.Listener
}

type ServiceProvider interface {
	GetServiceAddresses(user_info *models.UserInfo) (endpoint_addresses *models.ServiceAddresses, err error)
	GetVersionInfo() (version_info *models.VersionInfo, err error)
	// NewCaptcha() (id string, err error)
	GetCredentialStatus(user_info *models.UserInfo) (status *models.CredentialStatus, err error)
	GetCaptchaImage(user_info *models.UserInfo, captcha_id string) (content_type string, data []byte, err error)
	// VerifyCaptcha(id string, solution_data string) (ok bool, err error)
	NewRegistrationChallenge(user_info *models.UserInfo) (challenge *models.RegistrationChallenge, err error)
	Register(auth_info *models.UserInfo, registration_request *models.RegistrationRequest) (registration_response *models.RegistrationResponse, err error)
	NewLoginChallenge(user_info *models.UserInfo) (challenge *models.LoginChallenge, err error)
	Login(auth_info *models.UserInfo, login_request *models.LoginRequest) (login_response *models.LoginResponse, err error)
	Logout(auth_info *models.UserInfo) (logout_response *models.LogoutResponse, err error)
	AuthenticateCredential(auth_info *models.UserInfo, authenticate_request *models.CredentialAuthenticationRequest) (authenticate_response *models.CredentialAuthenticationResponse, err error)
	EnrollTwoFactorAuthentication(auth_info *models.UserInfo, enrollment_request *models.TwoFactorAuthenticationEnrollmentRequest) (enrollment_response *models.TwoFactorAuthenticationEnrollmentResponse, err error)
	GetTwoFactorAuthenticationMethods(auth_info *models.UserInfo) (two_factor_methods_list *models.TwoFactorAuthMethods, err error)
	SendTwoFactorAuthenticationCodeEmail(user_info *models.UserInfo) (send_2fa_code_email_response *models.SentTwoFactorAuthenticationCodeEmailResponse, err error)
	GetAccountStatus(auth_info *models.UserInfo) (status *models.AccountStatus, err error)
	NewGameAccount(auth_info *models.UserInfo, new_game_account_request *models.NewGameAccountRequest) (new_game_account_response *models.NewGameAccountResponse, err error)
	ActivateGameAccount(auth_info *models.UserInfo, id string) (err error)
	DeleteGameAccount(auth_info *models.UserInfo, id string) (err error)
	RenameGameAccount(auth_info *models.UserInfo, id string, rename_request *models.RenameGameAccountRequest) (err error)
	GetRealmStatusList(auth_info *models.UserInfo) (realm_status_list *models.RealmStatusList, err error)
}

func New(config *ServiceConfig, provider ServiceProvider) (service *Service) {
	service = new(Service)
	service.config = config
	service.provider = provider
	service.web_serve_mux = http.NewServeMux()
	service.mount_api()

	if config.WebApp != nil {
		web_app_fileserver := http.FileServerFS(config.WebApp)
		service.web_serve_mux.Handle("/{app_file...}", web_app_fileserver)
	}

	return
}

func (service *Service) ID() config.HomeServiceID {
	return config.WebService
}

func (service *Service) Start() (err error) {
	service.listener, err = net.Listen("tcp", service.config.Address)
	if err != nil {
		return
	}
	go func() {
		if err := http.Serve(service.listener, service.web_serve_mux); err != nil {
			log.Warn(err)
		}
	}()
	return
}

func (service *Service) Stop() (err error) {
	err = service.listener.Close()
	return
}
