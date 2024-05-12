package message

import (
	"github.com/Gophercraft/core/serialization"
	"github.com/Gophercraft/core/version"
)

// Packet describes an arbitrary message packet sent or received through the network
type Packet struct {
	// The virtual (version-nonspecific) type of the packet. note that this is not the actual wire-format type (enum NETMESSAGE)
	Type Type
	// Arbitrary blob of binary data (version-specific)
	serialization.Buffer
}

// Encodable describes a structure that can be converted into a Packet.
type Encodable interface {
	// The wire format of Encode may vary from build to build
	Encode(build version.Build, p *Packet) error
}

// Encodable describes a structure that can be converted back from a Packet.
type Decodable interface {
	Decode(build version.Build, p *Packet) error
}

// A Codec is both Encodable and Decodable
type Codec interface {
	Encodable
	Decodable
}
