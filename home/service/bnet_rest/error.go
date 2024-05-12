package bnet_rest

import (
	"net/http"

	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
)

type login_error struct {
	status        int
	error_code    string
	error_message string
}

func (l login_error) Error() string {
	return l.error_message
}

func ErrorMessage(status int, error_code, error_message string) error {
	l := login_error{
		status:        status,
		error_code:    error_code,
		error_message: error_message,
	}

	return l
}

func send_error_as_login_result(rw http.ResponseWriter, err error) {
	var login_result login.LoginResult
	var status int
	util.Set(&login_result.AuthenticationState, login.AuthenticationState_LOGIN)
	if le, ok := err.(login_error); ok {
		status = le.status
		util.Set(&login_result.ErrorCode, le.error_code)
		util.Set(&login_result.ErrorMessage, le.error_message)
	} else {
		status = http.StatusOK
		util.Set(&login_result.ErrorCode, "ERROR_INTERNAL")
		util.Set(&login_result.ErrorMessage, "Internal server error")
	}

	send_proto_json_response(status, rw, &login_result)
}
