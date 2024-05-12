package message

import "github.com/Gophercraft/core/version"

// TypeDescriptor describes how to convert between message Type codes and protocol-level opcodes.
// This allows the server to operate on an abstracted layer where the differences of NETMESSAGE enums between client versions no longer matter to the server.
type TypeDescriptor struct {
	opcodes map[Type]uint32
	types   map[uint32]Type
}

var TypeDescriptors = map[version.BuildRange]*TypeDescriptor{}

func QueryTypeDescriptor(build version.Build) (td *TypeDescriptor, err error) {
	err = version.QueryDescriptors(build, TypeDescriptors, &td)
	return
}

func RegisterOpcodes(buildRange version.BuildRange, opcodes map[Type]uint32) {
	types := make(map[uint32]Type, len(opcodes))
	for t, op := range opcodes {
		types[op] = t
	}

	TypeDescriptors[buildRange] = &TypeDescriptor{
		opcodes,
		types,
	}
}
