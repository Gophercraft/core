package models

import (
	"github.com/Gophercraft/core/vsn"
)

type Resistance uint8

const (
	ResistPhysical Resistance = iota
	ResistHoly
	ResistFire
	ResistNature
	ResistFrost
	ResistShadow
	ResistArcane
	NumResists
)

type ResistanceDescriptor map[Resistance]int

var (
	ResistanceDescriptors = map[vsn.BuildRange]ResistanceDescriptor{
		vsn.Range(0, vsn.Alpha): {
			ResistPhysical: 0,
			ResistHoly:     1,
			ResistFire:     2,
			ResistNature:   3,
			ResistFrost:    4,
			ResistShadow:   5,
		},

		vsn.Range(vsn.V1_12_1, vsn.Max): {
			ResistPhysical: 0,
			ResistHoly:     1,
			ResistFire:     2,
			ResistNature:   3,
			ResistFrost:    4,
			ResistShadow:   5,
			ResistArcane:   6,
		},
	}
)

func ResistIndex(build vsn.Build, resist Resistance) int {
	var desc ResistanceDescriptor
	if err := vsn.QueryDescriptors(build, ResistanceDescriptors, &desc); err != nil {
		panic(err)
	}

	i, ok := desc[resist]
	if !ok {
		return -1
	}

	return i
}
