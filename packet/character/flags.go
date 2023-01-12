package character

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Flags uint32

const (
	FlagLockedForTransfer Flags = 1 << iota
	FlagHideHelm
	FlagHideCloak
	FlagGhost
	FlagRename
	FlagLockedByBilling
	FlagDeclined
)

type FlagDescriptor map[Flags]uint32

var FlagDescriptors = map[vsn.BuildRange]FlagDescriptor{
	{0, vsn.Max}: {
		FlagLockedForTransfer: 0x00000004,
		FlagHideHelm:          0x00000400,
		FlagHideCloak:         0x00000800,
		FlagGhost:             0x00002000,
		FlagRename:            0x00004000,
		FlagLockedByBilling:   0x01000000,
		FlagDeclined:          0x02000000,
	},
}

func (f *Flags) Encode(build vsn.Build, out *packet.WorldPacket) error {
	flags := *f
	var desc FlagDescriptor
	if err := vsn.QueryDescriptors(build, FlagDescriptors, &desc); err != nil {
		return err
	}
	var uflags uint32
	for flag, code := range desc {
		if flags&flag != 0 {
			uflags |= code
		}
	}
	out.WriteUint32(uint32(flags))
	return nil
}

func (f *Flags) Decode(build vsn.Build, in *packet.WorldPacket) error {
	var desc FlagDescriptor
	if err := vsn.QueryDescriptors(build, FlagDescriptors, &desc); err != nil {
		return err
	}
	uflags := in.ReadUint32()
	var out Flags
	for flag, code := range desc {
		if uflags&code != 0 {
			out |= flag
		}
	}
	*f = out
	return nil
}
