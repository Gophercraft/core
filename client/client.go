package client

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"net"

	"github.com/Gophercraft/core/auth"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Config struct {
	vsn.Build
	Username, Password string
	Player             string
	Realmlist          string
}

type Realm struct {
	ID      uint64
	Name    string
	Address string
}

type Client struct {
	Auth       *auth.Client
	Player     string
	PlayerGUID guid.GUID
	Config     Config
	RealmList  []Realm
	// Warden     *warden.Warden
	SessionKey []byte
}

func New(cfg Config) (c *Client, err error) {
	c = &Client{}
	c.Player = cfg.Player
	c.Config = cfg
	err = c.Login(cfg.Username, cfg.Password)
	if err != nil {
		return
	}

	c.RealmList, err = c.Auth.GetRealmlist()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cl *Client) Connect(ip string) error {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return err
	}

	buf := make([]byte, 512)
	if _, err := conn.Read(buf); err != nil {
		return err
	}

	gp, err := packet.UnmarshalSMSGAuthPacket(cl.Cfg.Build, buf)
	if err != nil {
		return err
	}

	seed := randomBuffer(4)
	h := hash(
		[]byte(cl.Cfg.Username),
		[]byte{0, 0, 0, 0},
		seed,
		gp.Salt,
		cl.Auth.SessionKey,
	)

	app := &packet.CMSGAuthSession{
		Build:     uint32(cl.Cfg.Build),
		Account:   cl.Cfg.Username,
		Seed:      seed,
		Digest:    h,
		AddonData: packet.ClientAddonData,
	}

	if _, err = wc.Write(app.Encode()); err != nil {
		return nil
	}

	cl.Handlers = make(map[packet.WorldType]*ClientHandler)
	cl.World = wc
	cl.Crypter = packet.NewCrypter(cl.Cfg.Build, wc, cl.SessionKey, false)

	return cl.Handle()
}

func hash(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}

func randomBuffer(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}
