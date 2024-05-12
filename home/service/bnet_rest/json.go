package bnet_rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Send
func send_proto_json_response(status int, rw http.ResponseWriter, message proto.Message) {
	// Names from .proto file get used here
	json_response, err := protojson.MarshalOptions{
		UseProtoNames: true,
	}.Marshal(message)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json_response))
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	// rw.Header().Set("Connection", "close")
	rw.WriteHeader(status)
	rw.Write(json_response)
}

func read_login_form(rw http.ResponseWriter, request *http.Request, login_form *login.LoginForm) (err error) {
	var json_body []byte
	json_body, err = io.ReadAll(request.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("Received json body", string(json_body))

	// Attempt to read and decode login form
	var (
		login_result login.LoginResult
	)
	if err = protojson.Unmarshal(json_body, login_form); err != nil {
		util.Set(&login_result.AuthenticationState, login.AuthenticationState_LOGIN)
		util.Set(&login_result.ErrorCode, "UNABLE_TO_DECODE")
		util.Set(&login_result.ErrorMessage, "There was an internal error while connecting to Battle.net. Please try again later.")

		send_proto_json_response(http.StatusBadRequest, rw, &login_result)

		return
	}

	return
}
