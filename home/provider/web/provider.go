package web

import (
	"time"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/phylactery/database"
)

type service_provider struct {
	config                *ProviderConfig
	home_db               *database.Container
	captcha_manager       *captcha_manager
	last_credential_sweep time.Time
}

// func (provider *service_provider) SignIn(user_info *models.UserInfo, sign_in_request *models.SignInRequest) (sign_in_response *models.SignInResponse, err error) {
// 	return
// }

type ProviderConfig struct {
	WebRegistration          bool
	ServiceAddresses         map[config.HomeServiceID]string
	Brand                    string
	UseCaptchas              bool
	EmailRequired            bool
	EmailVerificationNeeded  bool
	MaxEmailLength           int32
	MaxUsernameLength        int32
	MaxPasswordLength        int32
	MaxGameAccounts          int32
	MaxGameAccountNameLength int32
}

func New(config *ProviderConfig, home_db *database.Container) (provider *service_provider) {
	provider = new(service_provider)
	provider.home_db = home_db
	provider.config = config

	if provider.config.MaxEmailLength == 0 {
		provider.config.MaxEmailLength = 128
	}
	if provider.config.MaxUsernameLength == 0 {
		provider.config.MaxUsernameLength = 32
	}
	if provider.config.MaxPasswordLength == 0 {
		provider.config.MaxPasswordLength = 128
	}
	if provider.config.MaxGameAccounts == 0 {
		provider.config.MaxGameAccounts = 16
	}
	if provider.config.MaxGameAccountNameLength == 0 {
		provider.config.MaxGameAccountNameLength = 128
	}

	if provider.config.UseCaptchas {
		provider.captcha_manager = new_captcha_manager()
	}
	return provider
}
