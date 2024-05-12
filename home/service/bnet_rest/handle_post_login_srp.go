package bnet_rest

import (
	"net/http"

	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/log"
)

func (service *Service) handle_post_login_srp(rw http.ResponseWriter, r *http.Request) {
	var (
		login_form login.LoginForm
		err        error
		user_info  *UserInfo
	)
	if err = read_login_form(rw, r, &login_form); err != nil {
		log.Warn(err)
		return
	}

	user_info, err = service.get_user_info(rw, r)
	if err != nil {
		log.Warn("Error getting user info: ", err)
		send_error_as_login_result(rw, ErrorMessage(http.StatusBadRequest, "ERROR_INTERNAL", "Internal server error"))
		return
	}

	login_srp_challenge, err := service.provider.LoginSRP(user_info, &login_form)
	if err != nil {
		send_error_as_login_result(rw, err)
		return
	}

	if login_srp_challenge == nil {
		panic("home/service/bnet_rest: must return srp challenge")
	}

	send_proto_json_response(http.StatusOK, rw, login_srp_challenge)
}
