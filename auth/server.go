package auth

import (
	"crypto/rand"
	"errors"
	"net"

	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

var (
	VersionChallenge = [...]byte{0xBA, 0xA3, 0x1E, 0x99, 0xA0, 0x0B, 0x21, 0x57, 0xFC, 0x37, 0x3F, 0xB3, 0x69, 0xCD, 0xD2, 0xF1}
)

type Backend interface {
	GetAccount(user string) (*models.Account, []models.GameAccount, error)
	ListRealms() []models.Realm
	StoreKey(user, locale, platform string, K []byte)
}

type Server struct {
	Backend
}

const (
	stateUnauthorized = iota
	stateChallenging
	stateAuthorized
)

type Session struct {
	server *Server
	state  int // := stateUnauthorized
	K      []byte
	valid  bool
	M2     []byte
	// acc    *gcore.Account
	// g      *srp.BigNum
	// salt   *srp.BigNum
	// N      *srp.BigNum
	// v      *srp.BigNum
	// b      *srp.BigNum
	// B      *srp.BigNum
	// alc    *AuthLogonChallenge_C
	// locale string
	// // gameAccounts []gcore.GameAccount
	// build        vsn.Build
	// platformOS   string = "Wn"
	// platformArch string = "32"
	acc    *models.Account
	g      *srp.BigNum
	salt   *srp.BigNum
	N      *srp.BigNum
	v      *srp.BigNum
	b      *srp.BigNum
	B      *srp.BigNum
	alc    *AuthLogonChallenge_C
	locale string
	// gameAccounts []gcore.GameAccount
	build        vsn.Build
	platformOS   string
	platformArch string

	conn *Connection
}

func (session *Session) Send(packet Packet) error {
	log.Warn("Sending", packet.Type())
	if _, err := session.conn.Write([]byte{uint8(packet.Type())}); err != nil {
		return err
	}
	return packet.Send(session.build, session.conn)
}

func (session *Session) handleLogonProof(alpc *AuthLogonProof_C) {
	if session.state != stateChallenging {
		return
	}

	session.K, session.valid, session.M2 = srp.ServerLogonProof(session.acc.Username,
		srp.BigNumFromArray(alpc.A),
		srp.BigNumFromArray(alpc.M1),
		session.b,
		session.B,
		session.salt,
		session.N,
		session.v)

	if !session.valid {
		log.Println(session.acc.Username, "Invalid login")
		session.Send(&AuthLogonProof_S{
			Error: GruntFailUnknownAccount,
		})
		return
	}

	session.server.StoreKey(session.acc.Username, session.locale, session.platformOS+session.platformArch, session.K)

	proof := &AuthLogonProof_S{
		Error:    GruntSuccess,
		M2:       session.M2,
		SurveyID: 0,
		Unk3:     0,
	}

	if session.build.AddedIn(vsn.V2_0_1) {
		proof.AccountFlags = 0x00800000
	}

	err := session.Send(proof)
	if err != nil {
		log.Warn(err)
		return
	}

	session.state = stateAuthorized
	// Client requested a list of realms from the server.
}

func (session *Session) handleLogonChallenge(alc *AuthLogonChallenge_C) {
	if alc.Platform == "x86" {
		session.platformArch = "32"
	}
	if alc.Platform == "x64" {
		session.platformArch = "64"
	} else {
		session.platformArch = "??"
	}

	if alc.OS == "Win" {
		session.platformOS = "Wn"
	}
	if alc.OS == "OSX" {
		session.platformOS = "Mc"
	}

	session.locale = alc.Country

	session.build = vsn.Build(alc.Build)

	var err error
	session.acc, _, err = session.server.GetAccount(string(alc.I))
	if err != nil {
		// User could not be found.
		session.Send(&AuthLogonChallenge_S{
			Error: GruntFailUnknownAccount,
		})
		return
	}

	session.state = stateChallenging

	// Generate parameters
	session.salt = srp.BigNumFromRand(32)
	session.g = srp.Generator.Copy()
	session.N = srp.Prime.Copy()

	// Compute ephemeral (temporary) variables
	_, session.v = srp.CalculateVerifier(session.acc.IdentityHash, session.g, session.N, session.salt)
	session.b, session.B = srp.ServerGenerateEphemeralValues(session.g, session.N, session.v)

	challenge := &AuthLogonChallenge_S{
		Error:            GruntSuccess,
		B:                session.B.ToArray(32),
		G:                session.g.ToArray(1),
		N:                session.N.ToArray(32),
		S:                session.salt.ToArray(32),
		VersionChallenge: VersionChallenge[:],
	}

	session.Send(challenge)
}

func (session *Session) handleRealmList(rlst *RealmList_C) {
	if session.state != stateAuthorized {
		return
	}

	realms := []models.Realm{}

	realmList := session.server.ListRealms()
	for _, v := range realmList {
		if v.ClientVersion == session.build {
			realms = append(realms, v)
		}
	}

	realmListS := MakeRealmlist(realms)
	session.Send(realmListS)
}

func (server *Server) Handle(cn net.Conn) {
	session := &Session{}
	session.server = server
	session.conn = wrapConn(cn)

	log.Println("New authserver connection from", session.conn.RemoteAddr())

	// every iteration of this loop reads an opcode from the TCP socket, and associated data.
	for {
		at, err := session.conn.readAuthType()
		if err != nil {
			log.Warn(err)
			return
		}

		log.Println("Received", at)

		var packet Packet

		switch at {
		case LogonChallenge:
			packet = &AuthLogonChallenge_C{}
		case LogonProof:
			packet = &AuthLogonProof_C{}
		case RealmList:
			packet = &RealmList_C{}
		default:
			log.Warn(at)
			return
		}

		err = packet.Recv(session.build, session.conn)
		if err != nil {
			log.Warn(err)
			return
		}

		switch at {
		case LogonChallenge:
			session.handleLogonChallenge(packet.(*AuthLogonChallenge_C))
		case LogonProof:
			session.handleLogonProof(packet.(*AuthLogonProof_C))
		case RealmList:
			session.handleRealmList(packet.(*RealmList_C))
		default:
			panic(at)
			return
		}
	}
}

func RunServer(b Backend, l net.Listener) error {
	server := new(Server)
	server.Backend = b

	for {
		c, err := l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			continue
		}
		go server.Handle(c)
	}
}

func rnd(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}

func (h *Server) Close() error {
	panic("cannot close")
	return nil
}
