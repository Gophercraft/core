package character

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
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

var FlagDescriptors = map[version.BuildRange]FlagDescriptor{
	{0, version.Max}: {
		FlagLockedForTransfer: 0x00000004,
		FlagHideHelm:          0x00000400,
		FlagHideCloak:         0x00000800,
		FlagGhost:             0x00002000,
		FlagRename:            0x00004000,
		FlagLockedByBilling:   0x01000000,
		FlagDeclined:          0x02000000,
	},
}

func (f *Flags) Encode(build version.Build, out *message.Packet) error {
	flags := *f
	var desc FlagDescriptor
	if err := version.QueryDescriptors(build, FlagDescriptors, &desc); err != nil {
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

func (f *Flags) Decode(build version.Build, in *message.Packet) error {
	var desc FlagDescriptor
	if err := version.QueryDescriptors(build, FlagDescriptors, &desc); err != nil {
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
