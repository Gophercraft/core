package character

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type LoginResult uint8

const (
	LoginNoWorld LoginResult = iota
	LoginSuccess
	LoginInProgress
	LoginDuplicateCharacter
	LoginNoInstances
	LoginDisabled
	LoginNoCharacter
	LoginLockedForTransfer
	LoginLockedByBilling
	LoginFailed
)

var LoginResultDescriptors = map[vsn.BuildRange]map[LoginResult]uint8{
	{0, 3368}: {
		LoginInProgress:         0x30,
		LoginSuccess:            0x31,
		LoginNoWorld:            0x32,
		LoginDuplicateCharacter: 0x33,
		LoginNoInstances:        0x34,
		LoginFailed:             0x35,
		LoginDisabled:           0x36,
	},

	{5875, vsn.V3_3_5a}: {
		LoginNoWorld:            0x01,
		LoginDuplicateCharacter: 0x02,
		LoginNoInstances:        0x03,
		LoginDisabled:           0x04,
		LoginNoCharacter:        0x05,
		LoginLockedForTransfer:  0x06,
		LoginLockedByBilling:    0x07,
		LoginFailed:             0x08,
	},
}

type LoginStatus struct {
	Type   packet.WorldType
	Result LoginResult
}

func (l *LoginStatus) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = l.Type
	out.WriteByte(uint8(l.Result))
	return nil
}
