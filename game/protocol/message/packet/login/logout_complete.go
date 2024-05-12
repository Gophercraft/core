package login

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type LogoutComplete struct {
}

func (l *LogoutComplete) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_LOGOUT_COMPLETE
	return nil
}
