package grunt

import (
	"crypto/sha1"
	"crypto/subtle"
	"fmt"
)

func (session *Session) handle_auth_reconnect_proof(proof *AuthReconnectProof_Client) (err error) {
	if session.state != state_reconnect_challenging {
		err = fmt.Errorf("grunt: trying to prove reconnect incorrectly")
		return
	}

	var proof_response AuthReconnectProof_Server

	// verify session key
	hash := sha1.New()
	hash.Write([]byte(session.logon_info.AccountName))
	hash.Write(proof.R1[:])
	hash.Write(session.reconnect_proof[:])
	hash.Write(session.login_info.SessionKey)
	session_digest := hash.Sum(nil)

	// check if matches
	session_digest_match := subtle.ConstantTimeCompare(session_digest, proof.R2[:]) == 1

	if session_digest_match {
		session.set_state(state_authorized)

		proof_response.Error = 0
		proof_response.LogonProofFlags = session.login_info.LogonProofFlags
	} else {
		// TODO: log hack attempt
		proof_response.Error = 1
	}

	if err = WriteMessageType(session, ReconnectProof); err != nil {
		return
	}
	if err = WriteAuthReconnectProof_Server(session, &proof_response); err != nil {
		return
	}

	err = session.Send()

	return
}
