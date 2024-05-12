package auth

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/game/network"
	"github.com/Gophercraft/core/game/protocol/message/packet/addon"
	"github.com/Gophercraft/core/game/protocol/message/packet/auth"
	"github.com/Gophercraft/core/version"
)

func (service *Service) handle_client_auth_session(session *network.Session, auth_session *auth.SessionClient) error {
	// Get auth context from session
	auth_context := session.ServiceContext(auth_service_id).(*SessionContext)

	// if already authed, or have not sent auth challenge, terminate the session.
	if auth_context.HasState(Authed) {
		return fmt.Errorf("client context is already authenticated")
	}
	if auth_context.challenge == nil {
		return fmt.Errorf("client context doesn't have a challenge object")
	}

	// Reject session if client build does not match the server's.
	if auth_session.Build != session.Build() {
		return fmt.Errorf("client attempted to authenticate with build not matching the server's emulated build: %s", auth_session.Build)
	}

	// This connects back to the home server, checking to see if this client is actually a registered user.
	// The auth server will perform calculations, and if valid, return to us a session key.
	verify_world_login_request := &AuthenticateGameConnectionRequest{
		RealmID:        service.config.RealmID,
		Build:          session.Build(),
		Account:        auth_session.Account,
		IPAddress:      session.RemoteAddr().String(),
		Digest:         auth_session.Digest,
		Challenge:      auth_context.challenge.Challenge,
		LocalChallenge: auth_session.LocalChallenge,
	}

	// In Bnet, a list of game accounts is shown in-game
	if len(auth_session.RealmJoinTicket) > 0 {
		ticket := strings.SplitN(auth_session.RealmJoinTicket, ":", 2)
		verify_world_login_request.Account = ticket[0]
		verify_world_login_request.GameAccount = ticket[1]
	}

	// Make the request
	response, err := service.provider.AuthenticateGameConnection(verify_world_login_request)
	if err != nil {
		return err
	}

	// Todo: realm server admins should get their tier overridden
	// Server admin could be admin on their server, while having status as a normal player in the rest of the federation!
	auth_context.tier = response.Tier
	auth_context.locale = response.Locale
	auth_context.account = response.Account
	auth_context.game_account = response.GameAccount
	auth_context.session_key = response.SessionKey

	if err := enable_encryption(session, auth_context); err != nil {
		return err
	}

	// Pass all addons (necessary login step)
	if auth_session.AddonList != nil {
		addon_info := addon.SkipServerCheck(session.Build(), auth_session.AddonList)
		session.Encode(addon_info)
	}

	return nil
}

func enable_encryption(session *network.Session, auth_context *SessionContext) error {
	if session.Build().AddedIn(version.NewCryptSystem) {
		// In this version, encryption is deferred until the client responds with an ack
		signature, err := crypto.GenEnableEncryptionSignature(auth_context.session_key)
		if err != nil {
			return err
		}
		return session.Encode(&auth.EnterEncryptedMode{
			Signature: signature,
			Enabled:   true,
		})
	} else {
		// Before version.NewCryptSystem, encryption is enabled immediatley after a successful handshake
		if err := session.EnterEncryptedMode(auth_context.session_key); err != nil {
			return err
		}
		complete_handshake(session)
	}

	return nil
}

func complete_handshake(session *network.Session) {
	auth_context := session.ServiceContext(auth_service_id).(*SessionContext)

	if auth_context.HasState(PassedWaitQueue) {
		panic("handshake can't complete twice")
	}

	// if err := session.DB().Where("id = ?", session.Account).Find(session.Props); err != nil {
	// 	panic(err)
	// }
	session.SendSessionMetadata()
	session.EnterWaitQueue()
}
