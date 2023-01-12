package auth

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

type Client struct {
	GameName   string
	Version    vsn.Build
	Address    string
	Username   string
	Password   string
	SessionKey []byte
	Realms     []RealmListing
	conn       *Connection
}

// connect using the legacy auth protocol.
func (cl *Client) connectLegacy() error {
	conn, err := Dial(cl.Address)
	if err != nil {
		return err
	}

	cl.conn = conn

	alcc := &AuthLogonChallenge_C{
		GameName: cl.GameName,
		Version:  Version(cl.Version),
		Build:    uint16(cl.Version),
		Platform: "x86",
		//
		OS:           "Win",
		Country:      "enUS",
		TimezoneBias: 0,
		IP:           16777343, //localhost
		I:            strings.ToUpper(cl.Username),
	}

	if err := cl.Send(alcc); err != nil {
		return err
	}

	for {
		if err := cl.Recv(); err != nil {
			return err
		}
	}

	return nil
}

func (cl *Client) handleLogonProof(alps *AuthLogonProof_S) {
	if alps.Error != GruntSuccess {
		log.Warn(fmt.Errorf("Server returned %s", alps.Error))
		return
	}
}

func (cl *Client) handleLogonChallenge(auth *AuthLogonChallenge_S) {
	if auth.Error != GruntSuccess {
		log.Warn(fmt.Errorf("auth: server returned %s", auth.Error))
		return
	}

	if len(auth.N) != 32 {
		log.Warn(fmt.Errorf("auth: server sent invalid prime number length %d", len(auth.N)))
		return
	}

	_, K, A, M1 := srp.SRPCalculate(cl.Username, cl.Password, auth.B, auth.N, auth.S)
	cl.SessionKey = K
	proof := &AuthLogonProof_C{
		A:            A,
		M1:           M1,
		NumberOfKeys: 0,
		SecFlags:     0,
	}

	if err := cl.Send(proof); err != nil {
		log.Warn(err)
		return
	}

}

func (cl *Client) handleRealmList(rlst *RealmList_S) {
	cl.Realms = rlst.Realms
}

func (cl *Client) Recv() error {
	var packet Packet
	at, err := cl.conn.readAuthType()
	if err != nil {
		return err
	}

	switch at {
	case LogonChallenge:
		packet = &AuthLogonChallenge_S{}
	case LogonProof:
		packet = &AuthLogonProof_S{}
	case RealmList:
		packet = &RealmList_S{}
	default:
		return fmt.Errorf("unhandled type %s", at)
	}

	if err := packet.Recv(cl.Version, cl.conn); err != nil {
		return err
	}

	switch at {
	case LogonChallenge:
		cl.handleLogonChallenge(packet.(*AuthLogonChallenge_S))
	case LogonProof:
		cl.handleLogonProof(packet.(*AuthLogonProof_S))
	case RealmList:
		cl.handleRealmList(packet.(*RealmList_S))
	default:
		return fmt.Errorf("create a handler for %s", at)
	}

	return nil
}

func (cl *Client) Send(p Packet) error {
	if _, err := cl.conn.Write([]byte{uint8(p.Type())}); err != nil {
		return err
	}
	if err := p.Send(cl.Version, cl.conn); err != nil {
		return err
	}
	return nil
}

func (cl *Client) RequestRealmListUpdates() error {
	return cl.Send(&RealmList_C{})
}

func Login(version vsn.Build, address, username, password string) (*Client, error) {
	cl := &Client{
		Version:  version,
		Address:  address,
		Username: strings.ToUpper(username),
		Password: strings.ToUpper(password),
	}

	switch {
	case version < vsn.V4_0_1:
		err := cl.connectLegacy()
		if err != nil {
			return nil, err
		}

		return cl, nil
	// todo: add bnet client support
	default:
		return nil, fmt.Errorf("auth: unsupported protocol version %d", version)
	}
}
