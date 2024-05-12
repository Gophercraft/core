package web

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/home/login"
	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/service/web/models"
	api_models "github.com/Gophercraft/core/home/service/web/models"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *service_provider) NewLoginChallenge(user_info *api_models.UserInfo) (login_challenge *api_models.LoginChallenge, err error) {
	// TODO: check IP

	login_challenge = new(models.LoginChallenge)

	if provider.config.UseCaptchas {
		login_challenge.CaptchaID, err = provider.captcha_manager.new_challenge()
		if err != nil {
			return
		}
	}

	return
}

func (provider *service_provider) Login(user_info *api_models.UserInfo, login_request *api_models.LoginRequest) (login_response *api_models.LoginResponse, err error) {
	if provider.config.UseCaptchas {
		valid := provider.captcha_manager.verify(login_request.CaptchaID, login_request.CaptchaSolution)
		if !valid {
			err = fmt.Errorf("invalid captcha")
			return
		}
	}

	username := strings.ToLower(login_request.Username)
	password := login_request.Password

	login_response = new(api_models.LoginResponse)
	login_response.WebToken, err = login.LoginWeb(provider.home_db, home_models.Login_Web, username, password, user_info.Address, user_info.UserAgent)
	if err != nil {
		return
	}

	return
}

func (provider *service_provider) Logout(user_info *api_models.UserInfo) (logout_response *api_models.LogoutResponse, err error) {
	logout_response = new(api_models.LogoutResponse)

	provider.home_db.Table("WebToken").Where(query.Eq("Token", user_info.Token)).Delete()
	return
}
