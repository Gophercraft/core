package bnet

import (
	protocol "github.com/Gophercraft/core/bnet/bgs/protocol"
	v1 "github.com/Gophercraft/core/bnet/bgs/protocol/authentication/v1"
	chalv1 "github.com/Gophercraft/core/bnet/bgs/protocol/challenge/v1"
	"github.com/Gophercraft/log"
	"github.com/superp00t/etc"
)

// func (s *Listener) DoVerifyWebCredentials(conn *Conn, token uint32, args *v1.VerifyWebCredentialsRequest) {

// }

func (s *Listener) Logon(conn *Conn, token uint32, args *v1.LogonRequest) {
	if args.GetProgram() != "WoW" {
		conn.SendResponseCode(token, ERROR_BAD_PROGRAM)
		return
	}

	if len(args.GetCachedWebCredentials()) > 0 {
		s.VerifyWebCredentials(conn, token, &v1.VerifyWebCredentialsRequest{
			WebCredentials: args.GetCachedWebCredentials(),
		})
		return
	}

	conn.locale = args.GetLocale()
	conn.platform = args.GetPlatform()

	go func() {
		t := "web_auth_url"
		// payload := []byte("https://" + s.RESTAddress + "/bnetserver/login/")
		payload := []byte("https://" + s.HostExternal + ":1120/bnetserver/login/")

		var request chalv1.ChallengeExternalRequest
		request.PayloadType = &t
		request.Payload = payload

		log.Println("Sending request")
		log.Dump("request", request)
		err := conn.ChallengeListener_Request_OnExternalChallenge(&request)
		if err != nil {
			log.Warn(err)
		} else {
			log.Println("Sent request")
		}
	}()

	conn.SendResponseCode(token, ERROR_OK)
}

func (s *Listener) ModuleNotify(conn *Conn, token uint32, args *v1.ModuleNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) ModuleMessage(conn *Conn, token uint32, args *v1.ModuleMessageRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) SelectGameAccount_DEPRECATED(conn *Conn, token uint32, args *protocol.EntityId) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) GenerateSSOToken(conn *Conn, token uint32, args *v1.GenerateSSOTokenRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) SelectGameAccount(conn *Conn, token uint32, args *v1.SelectGameAccountRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) VerifyWebCredentials(conn *Conn, token uint32, args *v1.VerifyWebCredentialsRequest) {
	ticket := string(args.GetWebCredentials())
	us, _ := s.Backend.GetTicket(ticket)
	if us == "" {
		log.Println("sending invalid")

		invalid := uint32(ERROR_DENIED)
		conn.AuthenticationListener_Request_OnLogonComplete(&v1.LogonResult{
			ErrorCode: &invalid,
		})

		conn.SendResponseCode(token, ERROR_OK)
		conn.Close() // We want to show the invalid code, however NOT closing the connection causes the Login button to be greyed out.
		// "you have been disconnected"
		return
	}

	conn.user = us

	sessionKey := etc.NewBuffer().WriteRandom(64).Bytes()

	acc, gameAccounts, err := s.Backend.GetAccount(us)
	if err != nil {
		panic(err)
	}

	lr := &v1.LogonResult{
		ErrorCode: u32p(0),
		AccountId: &protocol.EntityId{
			Low:  u64p(uint64(acc.ID)),
			High: u64p(uint64(0x100000000000000)),
		},
		SessionKey: sessionKey,
	}

	highEntityID := uint64(0x200000200576F57)

	for _, v := range gameAccounts {
		lr.GameAccountId = append(lr.GameAccountId, &protocol.EntityId{
			Low:  &v.ID,
			High: &highEntityID,
		})
	}

	// Default to zero. TODO: figure out how to trigger the game account selection menu
	if len(gameAccounts) > 0 {
		conn.gameAccountName = gameAccounts[0].Name
	}

	conn.authed = true
	go conn.AuthenticationListener_Request_OnLogonComplete(lr)

	conn.SendResponseCode(token, ERROR_OK)
}

func (s *Listener) GenerateWebCredentials(conn *Conn, token uint32, args *v1.GenerateWebCredentialsRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
