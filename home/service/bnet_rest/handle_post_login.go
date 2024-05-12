package bnet_rest

import (
	"net/http"
	"net/http/httputil"

	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/log"
)

const fake_ticket = "TC-0000000000000000000000000000000000000000"

const (
	login_unknown = iota
	login_plaintext
	login_srp
)

func (service *Service) handle_post_login(rw http.ResponseWriter, r *http.Request) {

	dat, _ := httputil.DumpRequest(r, false)
	log.Dump("login post", dat)
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

	if err := read_login_form(rw, r, &login_form); err != nil {
		log.Warn(err)
		return
	}

	login_result, err := service.provider.Login(user_info, &login_form)
	if err != nil {
		log.Warn("Login(): returned", err)
		send_error_as_login_result(rw, err)
		return
	}

	if login_result == nil {
		panic("home/service/bnet_rest: must return login result")
	}

	send_proto_json_response(http.StatusOK, rw, login_result)

	// // Attempt to parse login form inputs
	// var (
	// 	login_type               int = login_unknown
	// 	login_username           string
	// 	login_password           string
	// 	login_public_A           *srp.Int
	// 	login_client_evidence_M1 *srp.Int
	// 	login_use_srp            bool
	// )
	// for _, form_input := range login_form.GetInputs() {
	// 	switch form_input.GetInputId() {
	// 	case "use_srp":
	// 		login_use_srp = form_input.GetValue() == "true"
	// 	case "account_name":
	// 		login_username = strings.ToLower(form_input.GetValue())
	// 	case "password":
	// 		login_password = form_input.GetValue()
	// 	case "public_A":
	// 		login_public_A = srp.NewIntFromHexString(form_input.GetValue())
	// 	case "client_evidence_M1":
	// 		login_client_evidence_M1 = srp.NewIntFromHexString(form_input.GetValue())
	// 	}
	// }

	// if len(login_username) > 0 && len(login_password) > 0 {
	// 	login_type = login_plaintext
	// } else {
	// 	if login_use_srp && len(login_username) > 0 {
	// 		if login_public_A != nil && login_client_evidence_M1 != nil {
	// 			if !login_public_A.Equals(srp.Zero) && !login_client_evidence_M1.Equals(srp.Zero) {
	// 				login_type = login_srp
	// 			}
	// 		}
	// 	}
	// }

	// log.Println("login type", login_type)

	// var login_result login.LoginResult
	// var status = http.StatusOK

	// switch login_type {
	// case login_unknown:
	// 	log.Warn("unknown login type")
	// 	status = http.StatusBadRequest
	// 	util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 	util.Set(&login_result.LoginTicket, fake_ticket)
	// 	util.Set(&login_result.ErrorCode, "ERROR_INTERNAL")
	// 	util.Set(&login_result.ErrorMessage, "Internal server error")
	// case login_plaintext:
	// 	// Check login details
	// 	login_ticket, valid, err := service.provider.Login(r.RemoteAddr, login_username, login_password)
	// 	if err != nil {
	// 		log.Warn("provider.Login() returned", err)
	// 		status = http.StatusBadRequest
	// 		util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 		util.Set(&login_result.LoginTicket, fake_ticket)
	// 		util.Set(&login_result.ErrorCode, "ERROR_INTERNAL")
	// 		util.Set(&login_result.ErrorMessage, "Internal server error")
	// 	} else {
	// 		if valid {
	// 			// login is valid
	// 			util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 			util.Set(&login_result.LoginTicket, login_ticket)
	// 		} else {
	// 			log.Warn("provider.Login() returned no error, but was invalid")
	// 			// login is invalid
	// 			status = http.StatusBadRequest
	// 			util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 			util.Set(&login_result.LoginTicket, fake_ticket)
	// 			util.Set(&login_result.ErrorCode, "ERROR_LOGON_INVALID_SERVER_PROOF")
	// 			util.Set(&login_result.ErrorMessage, "Invalid username or password")
	// 		}
	// 	}
	// case login_srp:
	// 	var server_M2 *srp.Int
	// 	srp_session, platform_id, err := service.provider.CheckoutSRPSession(r.RemoteAddr, login_username)
	// 	if err != nil {
	// 		log.Warn("No srp session was found for", r.RemoteAddr, login_username)
	// 		// No SRP session found
	// 		status = http.StatusBadRequest
	// 		util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 		util.Set(&login_result.LoginTicket, fake_ticket)
	// 		util.Set(&login_result.ErrorCode, "ERROR_INTERNAL")
	// 		util.Set(&login_result.ErrorMessage, "Internal server error")
	// 	} else {
	// 		log.Println("verifier", spew.Sdump(srp_session.Verifier.Bytes()))

	// 		session_key, valid := srp.VerifyClientEvidence(
	// 			srp_session.Hash,
	// 			login_public_A,
	// 			srp_session.PrivateB,
	// 			srp_session.PublicB,
	// 			srp_session.LargeSafePrime,
	// 			srp_session.Verifier,
	// 			login_client_evidence_M1,
	// 		)

	// 		if !valid {
	// 			log.Warn("Client evidence was wrong")
	// 			status = http.StatusBadRequest
	// 			util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 			// util.Set(&login_result.LoginTicket, "TC-0000000000000000000000000000000000000000")
	// 			// util.Set(&login_result.ErrorCode, "ERROR_LOGON_INVALID_SERVER_PROOF")
	// 			// util.Set(&login_result.ErrorMessage, "Invalid username or password")
	// 		} else {
	// 			server_M2 = srp.CalculateServerEvidence(srp_session.Hash, login_public_A, login_client_evidence_M1, session_key)
	// 			if err := service.provider.UpdateAccountData(login_username, platform_id, session_key.Bytes()); err != nil {
	// 				log.Warn("Couldn't update account data", err)
	// 				util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 				util.Set(&login_result.LoginTicket, fake_ticket)
	// 				// util.Set(&login_result.ErrorCode, "ERROR_INTERNAL")
	// 				// util.Set(&login_result.ErrorMessage, "Internal server error")
	// 			} else {
	// 				new_ticket, err := service.provider.CreateTicket(login_username)
	// 				if err != nil {
	// 					log.Warn("Couldn't create ticket", err)
	// 					util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 					// util.Set(&login_result.LoginTicket, fake_ticket)
	// 					// util.Set(&login_result.ErrorCode, "ERROR_INTERNAL")
	// 					// util.Set(&login_result.ErrorMessage, "Internal server error")
	// 				} else {
	// 					util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	// 					util.Set(&login_result.LoginTicket, new_ticket)
	// 					util.Set(&login_result.ServerEvidence_M2, server_M2.HexString())
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// send_proto_json_response(status, rw, &login_result)
}
