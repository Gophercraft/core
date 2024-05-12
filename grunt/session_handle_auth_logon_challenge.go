package grunt

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
	"github.com/davecgh/go-spew/spew"
)

const authenticator_added_build version.Build = 8606

var (
	version_challenge = [...]byte{0xBA, 0xA3, 0x1E, 0x99, 0xA0, 0x0B, 0x21, 0x57, 0xFC, 0x37, 0x3F, 0xB3, 0x69, 0xCD, 0xD2, 0xF1}
)

func (session *Session) handle_auth_logon_challenge(client_challenge *AuthLogonChallenge_Client) (err error) {
	// validate account name
	if client_challenge.Info.AccountName == "" || !utf8.Valid([]byte(client_challenge.Info.AccountName)) {
		err = fmt.Errorf("grunt: invalid account name")
		return
	}

	var login_attempt_event Event
	login_attempt_event.Type = LoginAttempt
	login_attempt_event.AccountName = client_challenge.Info.AccountName

	// stop mass login attempts
	if err = session.server.service_provider.Check(&login_attempt_event); err != nil {
		return
	}

	var server_challenge AuthLogonChallenge_Server

	if err = session.server.service_provider.GetAccountLoginInfo(
		client_challenge.Info.AccountName,
		session.logon_info.Build,
		&session.login_info); err != nil {
		log.Println("could not get login info", err)

		// if this is a grunt error, fail appropriately
		if grunt_err, ok := err.(LoginError); ok {
			server_challenge.LoginResult = grunt_err.Result()
		} else {
			server_challenge.LoginResult = LoginUnknownAccount
		}

		err = nil
	} else {
		// found login info

		session.logon_info = client_challenge.Info

		session.set_state(state_logon_challenging)

		// Generate parameters
		session.salt = srp.NewRandomInt(32)
		session.g = srp.Generator.Copy()
		session.N = srp.Prime.Copy()

		// Compute ephemeral (temporary) variables
		_, session.v = srp.CalculateVerifier(sha1.New, session.salt.Bytes(), session.login_info.IdentityHash, session.g, session.N)
		session.b, session.B = srp.ServerGenerateEphemeralValues(session.g, session.N, session.v)

		server_challenge.LoginResult = LoginOk

		// Enable authenticator if needed
		if session.login_info.Authenticator {
			// Check if the authenticator has
			// if session.logon_info.Build.AddedIn(authenticator_added_build)
			server_challenge.LogonFlags |= Token
			server_challenge.Token.Required = 1
		}

		if session.login_info.PIN {
			server_challenge.LogonFlags |= PIN
			server_challenge.PIN.GridSeed = random_pin_grid_seed()
			io.ReadFull(rand.Reader, server_challenge.PIN.Salt[:])
			session.server_pin_info = &server_challenge.PIN
		}

		copy(server_challenge.ServerPublicKey[:], session.B.Bytes())
		server_challenge.Generator = session.g.Bytes()
		server_challenge.LargeSafePrime = session.N.Bytes()
		copy(server_challenge.Salt[:], session.salt.Bytes())
		copy(server_challenge.VersionChallenge[:], version_challenge[:])
	}

	if err = WriteMessageType(session, LogonChallenge); err != nil {
		return
	}
	if err = WriteAuthLogonChallenge_Server(session, &server_challenge); err != nil {
		return
	}

	log.Println(spew.Sdump(&server_challenge))

	err = session.Send()

	return
}
