package login

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
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

func (lr *LogoutResponse) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_LOGOUT_RESPONSE
	out.WriteUint32(uint32(lr.Status))
	out.WriteBool(lr.Instant)
	return nil
}
