package login

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type LogoutStatus uint32

const (
	LogoutOK LogoutStatus = iota
	LogoutInCombat
	LogoutInDuel
	LogoutJumpFall
)

type LogoutResponse struct {
	Status  LogoutStatus
	Instant bool
}

func (lr *LogoutResponse) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_LOGOUT_RESPONSE
	out.WriteUint32(uint32(lr.Status))
	out.WriteBool(lr.Instant)
	return nil
}
