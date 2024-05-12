package bnet_rest

import (
	"net/http"
)

func (service *Service) handle_get_login(rw http.ResponseWriter, r *http.Request) {
	service.get_user_info(rw, r)

	login_form := service.provider.GetLoginFormInputs()
	send_proto_json_response(http.StatusOK, rw, login_form)
}
