package web

import (
	"net/http"

	api_models "github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_post_authenticate_credential(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var authentication_request api_models.CredentialAuthenticationRequest
	if err := read_request(r, &authentication_request); err != nil {
		http.Error(rw, "malformed request", http.StatusBadRequest)
		return
	}

	authentication_response, err := service.provider.AuthenticateCredential(user_info, &authentication_request)
	if err != nil {
		respond(rw, http.StatusOK, api_models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, authentication_response)
}

func (service *Service) handle_post_2fa_enroll(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var enrollment_request api_models.TwoFactorAuthenticationEnrollmentRequest
	if err := read_request(r, &enrollment_request); err != nil {
		http.Error(rw, "malformed request", http.StatusBadRequest)
		return
	}

	enrollment_response, err := service.provider.EnrollTwoFactorAuthentication(user_info, &enrollment_request)
	if err != nil {
		respond(rw, http.StatusOK, api_models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, enrollment_response)
}

func (service *Service) handle_get_2fa_methods(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	two_factor_methods_list, err := service.provider.GetTwoFactorAuthenticationMethods(user_info)
	if err != nil {
		respond(rw, http.StatusOK, api_models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, two_factor_methods_list)
}

func (service *Service) handle_post_send_2fa_email_code(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	send_email_response, err := service.provider.SendTwoFactorAuthenticationCodeEmail(user_info)
	if err != nil {
		respond(rw, http.StatusOK, api_models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, send_email_response)
}
