package bnet_rest

import (
	"net/http"

	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/log"
)

func (service *Service) handle_post_login_authenticator(rw http.ResponseWriter, r *http.Request) {
	var (
		login_form login.LoginForm
		err        error
		user_info  *UserInfo
	)

	user_info, err = service.get_user_info(rw, r)
	if err != nil {
		log.Warn("Error getting user info: ", err)
		send_error_as_login_result(rw, ErrorMessage(http.StatusBadRequest, "ERROR_INTERNAL", "Internal server error"))
		return
	}

	log.Println("MY SESSION ID", user_info.SessionID)

	if err := read_login_form(rw, r, &login_form); err != nil {
		log.Warn(err)
		return
	}

	login_result, err := service.provider.LoginAuthenticator(user_info, &login_form)
	if err != nil {
		log.Warn("Login(): returned", err)
		send_error_as_login_result(rw, err)
		return
	}

	if login_result == nil {
		panic("home/service/bnet_rest: must return login result")
	}

	send_proto_json_response(http.StatusOK, rw, login_result)
}
