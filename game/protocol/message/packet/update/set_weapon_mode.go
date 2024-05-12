package update

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type SetWeaponMode struct {
	Mode uint32
}

func (swm *SetWeaponMode) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_SETWEAPONMODE
	out.WriteUint32(swm.Mode)
	return nil
}

func (swm *SetWeaponMode) Decode(build version.Build, in *message.Packet) error {
	swm.Mode = in.ReadUint32()
	return nil
}

type SetSheathe struct {
	Mode uint32
}

func (swm *SetSheathe) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_SETSHEATHED
	out.WriteUint32(swm.Mode)
	return nil
}

func (swm *SetSheathe) Decode(build version.Build, in *message.Packet) error {
	swm.Mode = in.ReadUint32()
	return nil
}
