package login

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type LogoutComplete struct {
}

func (l *LogoutComplete) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_LOGOUT_COMPLETE
	return nil
}
