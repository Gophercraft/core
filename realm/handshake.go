package realm

import (
	"context"
	"net"
	"strings"

	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/addon"
	"github.com/Gophercraft/core/packet/auth"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
	"github.com/superp00t/etc"
)

// End signals the termination of a session. End can be called anywhere, for any reason.
func (s *Session) End() {
	if s == nil {
		return
	}

	if s.ending == true {
		return
	}

	s.ending = true

	log.Warn("End called, state already ended")
	if s.State() == Ended {
		return
	}

	log.Warn("Cleaning up player")
	if !s.GuardSession.TryLock() {
		panic("who tf is locking")
	}

	inWorld := s.HasState(InWorld)

	if inWorld {
		log.Warn("Cleaning up player")
		s.CleanupPlayer()
	}

	s.GuardSession.Unlock()

	s.SetState(Ended)

	close(s.PacketPipe)

	if err := s.Connection.Close(); err != nil {
		log.Warn(err)
	}

	// s.Connection.Conn.Close()
}

func (s *Session) SendPacket(wPacket *packet.WorldPacket) {
	if s.State() != Ended {
		s.PacketPipe <- wPacket
	}
}

func (s *Session) Send(f packet.Encodable) {
	wPacket := &packet.WorldPacket{Buffer: etc.NewBuffer()}
	err := f.Encode(s.Build(), wPacket)
	if err != nil {
		panic(err)
	}

	if wPacket.Type == 0 {
		panic("cannot send packet with 0 type")
	}
	s.SendPacket(wPacket)
}

func (s *Session) StartAuthChallenge() {
	v := s.Build()
	s.AuthChallenge = &auth.Challenge{}
	s.AuthChallenge.DosZeroBits = 1
	s.AuthChallenge.DosChallenge = crypto.RandBytes(32)

	switch {
	case vsn.Range(0, 3368).Contains(v):
	// do nothing
	case vsn.Range(5875, 18414).Contains(v):
		// after 1.12.1 challenge bytes are added
		s.AuthChallenge.Challenge = crypto.RandBytes(4)
	case vsn.Range(19027, vsn.Max).Contains(v):
		// in 6.0.2 this is expanded to 16 bytes
		s.AuthChallenge.Challenge = crypto.RandBytes(16)
	default:
		panic(v)
	}

	// Send challenge to client.
	s.Send(s.AuthChallenge)
}

func (s *Session) HandleAuthSession(as *auth.SessionClient) {
	// if already authed, or have not sent auth challenge, terminate the session.
	if s.HasState(Authed) || s.AuthChallenge == nil {
		s.End()
		return
	}

	// Reject session if client build does not match the server's.
	if as.Build != s.Build() {
		log.Warn("Client attempted to join with build different than configured server-side", as.Build)
		s.End()
		return
	}

	// This connects back to the home server, checking to see if this client is actually a registered user.
	// The auth server will perform calculations, and if valid, return to us a session key.
	query := &rpcnet.VerifyWorldQuery{
		RealmID:     s.Server.Config.RealmID,
		Build:       uint32(as.Build),
		Account:     as.Account,
		GameAccount: "",
		IP:          s.Connection.Conn.RemoteAddr().String(),
		Digest:      as.Digest,
		Salt:        s.AuthChallenge.Challenge,
		Seed:        as.LocalChallenge,
	}

	// In Bnet, a list of game accounts is shown in-game
	if len(as.RealmJoinTicket) > 0 {
		ticket := strings.SplitN(as.RealmJoinTicket, ":", 2)
		query.Account = ticket[0]
		query.GameAccount = ticket[1]
	}

	resp, err := s.Server.HomeServiceClient.VerifyWorld(context.Background(), query)
	if err != nil {
		log.Warn(err)
		s.End()
		return
	}

	// Todo: realm server admins should get their tier overridden
	// Server admin could be admin on their server, while having status as a normal player in the rest of the federation!
	s.Tier = resp.Tier
	s.Locale = i18n.Locale(resp.Locale)
	s.Account = resp.Account
	s.GameAccount = resp.GameAccount
	s.SetState(Authed)

	switch resp.Status {
	case rpcnet.Status_OK:
	default:
		log.Warn("Login for user", query.Account, "failed")
		s.End()
		return
	}

	if err := s.EnableEncryption(resp.SessionKey); err != nil {
		log.Warn(err)
		s.End()
		return
	}

	if as.AddonList != nil {
		aInfoPacket := addon.SkipServerCheck(s.Build(), as.AddonList)
		s.Send(aInfoPacket)
		log.Println("Addon info sent")
	}
}

func (s *Session) drainPacketPipe(cancel <-chan bool) {
	for {
		select {
		// graceful shutdown
		case <-cancel:
			return
		case msg, ok := <-s.PacketPipe:
			if !ok {
				return
			}
			s.Log("Sending", msg.Type, msg.Len(), "bytes")
			if err := s.Connection.Send(msg); err != nil {
				log.Warn("Error sending:", err)
				s.End()
			}
		}
	}
}

func (s *Session) Build() vsn.Build {
	return s.Server.Build()
}

const PacketPipePid = "Gophercraft/core/realm/handshake.drainPacketPipe"

func (s *Session) initPacketPipe() {
	// ensures Send returns immediately by queueing up to 64 packets
	s.PacketPipe = make(chan *packet.WorldPacket, 64)
	s.CreateProcess(PacketPipePid, (*Session).drainPacketPipe)
}

func (server *Server) HandleConn(conn net.Conn) {
	// important todo: prevent memory exhaustion from attacker opening loads of connections
	// Perhaps using LRU map which tracks high numbers of connections from a particular IP

	c, err := packet.NewConnection(server.Build(), conn, true)
	if err != nil {
		panic(err)
	}

	if err := c.ConfirmProtocol(); err != nil {
		log.Warn(err)
		c.Close()
		return
	}

	session := &Session{
		Server:     server,
		Connection: c,
	}

	session.initPacketPipe()
	log.Println("started packet pipe")

	session.StartAuthChallenge()
	log.Println("started auth challenge")

	for {
		if session.ending {
			break
		}

		log.Println("Receiving")
		wPacket, err := session.Connection.Recv()
		if err != nil {
			log.Warn("recv error", err)
			session.End()
			return
		}

		if session.ending {
			break
		}

		log.Println(session.String(), wPacket.Type, "received, size", wPacket.Len())

		h, ok := session.Server.ProtocolHandlers[wPacket.Type]
		if !ok {
			continue
		}

		var opts Options
		for _, opt := range h.Options {
			if err := opt(wPacket, &opts); err != nil {
				// todo: close connection upon fatal error
				log.Warn(err)
				session.End()
				return
			}
		}

		if h.RequiredState <= session.State() {
			// Most opcodes are better left in this goroutine, but spawn a new one if specified
			if opts.OptionFlags&OptionFlagAsync != 0 {
				go session.dispatchHandler(h, wPacket)
			} else {
				session.dispatchHandler(h, wPacket)
			}
		}
	}
}

func (s *Session) EnableEncryption(sessionKey []byte) error {
	s.SessionKey = sessionKey
	if s.Build().AddedIn(vsn.NewCryptSystem) {
		// In this version, encryption is deferred until the client responds with an ack
		sig, err := crypto.GenEnableEncryptionSignature(sessionKey)
		if err != nil {
			return err
		}
		s.Send(&auth.EnterEncryptedMode{
			Signature: sig,
			Enabled:   true,
		})
	} else {
		// Before vsn.NewCryptSystem, encryption is assumed to be enabled
		if err := s.Connection.InitEncryption(sessionKey); err != nil {
			return err
		}
		s.CompleteHandshake()
	}

	return nil
}

func (s *Session) HandleEnterEncryptedModeAck() {
	if err := s.Connection.InitEncryption(s.SessionKey); err != nil {
		panic(err)
	}
	s.CompleteHandshake()
}

func (s *Session) HandleEnableNaglesAlgorithm() {
	log.Println("Client has requested to enable Nagle's algorithm (delay)")
	s.Connection.SetNagle(true)
}

func (s *Session) HandleDisconnect(dc *packet.ClientDisconnect) {
	log.Println(dc)
	s.End()
}

// This function is called after initial encryption checks are found to be correct.
func (session *Session) CompleteHandshake() {
	if session.state >= Waiting {
		panic("handshake can't complete twice")
	}

	// if err := session.DB().Where("id = ?", session.Account).Find(session.Props); err != nil {
	// 	panic(err)
	// }
	session.SendSessionMetadata()
	session.EnterWaitQueue()
}

// This function gets called after wait queue is completed or skipped.
func (s *Session) CompleteLogin() {
	s.SetState(CharacterSelectMenu)
	s.SendAuthSuccess()

	s.Server.Call(SessionCreateEvent, nil, s)
}
