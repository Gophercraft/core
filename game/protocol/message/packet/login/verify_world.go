package login

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/version"
)

type VerifyWorld struct {
	MapID    uint32
	Position tempest.C4Vector
}

func (vw *VerifyWorld) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_LOGIN_VERIFY_WORLD

	out.WriteUint32(vw.MapID)
	return vw.Position.Encode(out)
}
