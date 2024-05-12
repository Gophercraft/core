package grunt

import (
	"bytes"
	"net"

	"github.com/Gophercraft/core/crypto/srp"
)

type Session struct {
	server          *Server
	state           int
	K               []byte
	valid           bool
	M2              []byte
	reconnect_proof [16]byte
	login_info      AccountLoginInfo
	server_pin_info *ServerPINInfo
	g               *srp.Int
	salt            *srp.Int
	N               *srp.Int
	v               *srp.Int
	b               *srp.Int
	B               *srp.Int
	logon_info      LogonInfo
	connection      net.Conn
	buffer          bytes.Buffer
}
