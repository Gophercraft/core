package models

import (
	"github.com/Gophercraft/core/version"
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
	ResistanceDescriptors = map[version.BuildRange]ResistanceDescriptor{
		version.Range(0, version.Alpha): {
			ResistPhysical: 0,
			ResistHoly:     1,
			ResistFire:     2,
			ResistNature:   3,
			ResistFrost:    4,
			ResistShadow:   5,
		},

		version.Range(version.V1_12_1, version.Max): {
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

func ResistIndex(build version.Build, resist Resistance) int {
	var desc ResistanceDescriptor
	if err := version.QueryDescriptors(build, ResistanceDescriptors, &desc); err != nil {
		panic(err)
	}

	i, ok := desc[resist]
	if !ok {
		return -1
	}

	return i
}
