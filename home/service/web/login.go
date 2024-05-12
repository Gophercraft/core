package web

import (
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_get_login_challenge(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	login_challenge, err := service.provider.NewLoginChallenge(user_info)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, login_challenge)
}

func (service *Service) handle_post_login_request(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var login_request models.LoginRequest
	if err := read_request(r, &login_request); err != nil {
		http.Error(rw, "malformed request", http.StatusBadRequest)
		return
	}

	login_response, err := service.provider.Login(user_info, &login_request)
	if err != nil {
		respond(rw, http.StatusOK, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, login_response)
}

func (service *Service) handle_get_logout_request(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	logout_response, err := service.provider.Logout(user_info)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, logout_response)
}
