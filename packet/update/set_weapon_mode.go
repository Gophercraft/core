package update

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type SetWeaponMode struct {
	Mode uint32
}

func (swm *SetWeaponMode) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_SETWEAPONMODE
	out.WriteUint32(swm.Mode)
	return nil
}

func (swm *SetWeaponMode) Decode(build vsn.Build, in *packet.WorldPacket) error {
	swm.Mode = in.ReadUint32()
	return nil
}

type SetSheathe struct {
	Mode uint32
}

func (swm *SetSheathe) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_SETSHEATHED
	out.WriteUint32(swm.Mode)
	return nil
}

func (swm *SetSheathe) Decode(build vsn.Build, in *packet.WorldPacket) error {
	swm.Mode = in.ReadUint32()
	return nil
}
