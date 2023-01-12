package auth

import (
	"github.com/Gophercraft/core/vsn"
)

type Packet interface {
	Type() AuthType
	Send(build vsn.Build, out *Connection) error
	Recv(build vsn.Build, in *Connection) error
}
