package grunt

import (
	"fmt"

	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
)

func (session *Session) handle_auth_logon_proof(client_proof *AuthLogonProof_Client) (err error) {
	if session.state != state_logon_challenging {
		err = fmt.Errorf("grunt: auth logon proof sent without challenge")
		return
	}

	session.K, session.valid, session.M2 = srp.ServerLogonProof(session.logon_info.AccountName,
		srp.NewIntFromBytes(client_proof.ClientPublicKey[:]),
		srp.NewIntFromBytes(client_proof.ClientProof[:]),
		session.b,
		session.B,
		session.salt,
		session.N,
		session.v)

	var server_proof AuthLogonProof_Server

	if session.valid {
		if session.login_info.Authenticator {
			session.valid = session.server.service_provider.VerifyAuthenticatorCode(session.logon_info.AccountName, client_proof.Token.Token)
		} else if session.login_info.PIN {
			log.Warn("Verifiying PIN...")
			session.valid = session.server.service_provider.VerifyPIN(session.logon_info.AccountName, session.server_pin_info, &client_proof.PIN)
			log.Warn("Verified PIN", session.valid)
		}
	}

	// log.Println("client proof should be ", spew.Sdump(client_proof.PIN.Proof[:]))
	// log.Println("client proof should be ", spew.Sdump(GetPINProof(make([]byte, 16), "1234", client_proof.PIN.Salt[:])))

	if !session.valid {
		server_proof.LoginResult = LoginUnknownAccount
	} else {
		session.set_state(state_authorized)

		// Store account info
		var stored_account_info AccountInfo
		stored_account_info.AccountName = session.logon_info.AccountName
		stored_account_info.Architecture = session.logon_info.Architecture
		stored_account_info.OS = session.logon_info.OS
		stored_account_info.Locale = session.logon_info.Locale
		stored_account_info.SessionKey = session.K

		if err = session.server.service_provider.StoreAccountInfo(&stored_account_info); err != nil {
			// if grunt_err, ok := err.(Error); ok {
			// 	server_proof
			// }
			log.Warn("Could not store account info")
			return
		}

		server_proof.LoginResult = LoginOk
		copy(server_proof.ServerProof[:], session.M2)
		server_proof.SurveyID = 0

		if session.logon_info.Build.AddedIn(version.V2_0_1) {
			server_proof.AccountFlags = 0x00800000
		}
	}

	if err = WriteMessageType(session, LogonProof); err != nil {
		log.Warn("Failed to write message type", err)
		return
	}
	if err = WriteAuthLogonProof_Server(session, session.logon_info.Build, &server_proof); err != nil {
		log.Warn("Failed to write auth logon proof", err)
		return
	}

	err = session.Send()

	return
}
