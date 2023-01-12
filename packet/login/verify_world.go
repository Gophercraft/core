package login

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
)

type VerifyWorld struct {
	MapID    uint32
	Position tempest.C4Vector
}

func (vw *VerifyWorld) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_LOGIN_VERIFY_WORLD

	out.WriteUint32(vw.MapID)
	return vw.Position.Encode(out)
}
