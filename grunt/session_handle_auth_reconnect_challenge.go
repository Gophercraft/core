package grunt

import (
	"crypto/rand"
	"io"
)

func (session *Session) handle_auth_reconnect_challenge(challenge *AuthLogonChallenge_Client) (err error) {
	session.logon_info = challenge.Info

	var server_challenge AuthReconnectChallenge_Server

	// Lookup client username in database
	if err = session.server.service_provider.GetAccountLoginInfo(session.logon_info.AccountName, session.logon_info.Build, &session.login_info); err != nil {
		server_challenge.Error = 1
	} else if len(session.login_info.SessionKey) == 0 {
		// cannot reconnect without session key
		server_challenge.Error = 1
	} else {
		// generate random server proof
		if _, err = io.ReadFull(rand.Reader, session.reconnect_proof[:]); err != nil {
			return
		}
		// copy
		copy(server_challenge.ServerProof[:], session.reconnect_proof[:])
		// copy version challenge seed
		copy(server_challenge.VersionChallenge[:], version_challenge[:])

		session.set_state(state_reconnect_challenging)
	}

	// write reconnect challenge
	if err = WriteMessageType(session, ReconnectChallenge); err != nil {
		return
	}
	if err = WriteAuthReconnectChallenge_Server(session, &server_challenge); err != nil {
		return
	}
	err = session.Send()

	return
}
