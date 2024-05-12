package web

import (
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_get_registration_challenge(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	registration_challenge, err := service.provider.NewRegistrationChallenge(user_info)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, registration_challenge)
}

func (service *Service) handle_post_registration_request(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var registration_request models.RegistrationRequest
	if err := read_request(r, &registration_request); err != nil {
		http.Error(rw, "malformed request", http.StatusBadRequest)
		return
	}

	registration_response, err := service.provider.Register(user_info, &registration_request)
	if err != nil {
		respond(rw, http.StatusOK, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, registration_response)
}
