package update

import (
	"reflect"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type DescriptorOptions uint32

const (
	DescriptorOptionClassicGUIDs = 1 << iota
	DescriptorOptionHasHasTransport
	DescriptorOptionAlpha
)

// Descriptor describes all the info about a particular version's SMSG_UPDATE_OBJECT structure.
// Descriptor's underlying Types are backed by a Go structure
// However,
type Descriptor struct {
	DescriptorOptions
	ObjectDescriptors map[guid.TypeMask]reflect.Type
}

var (
	Descriptors = map[version.BuildRange]*Descriptor{}
)

func GetDescriptor(build version.Build) *Descriptor {
	var desc *Descriptor
	if err := version.QueryDescriptors(build, Descriptors, &desc); err != nil {
		panic(err)
	}

	return desc
}
