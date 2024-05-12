package protocol

import "github.com/Gophercraft/core/version"

type ServerConfiguration struct {
	// The address to listen on.
	Bind string

	// The protocol version to emulate.
	Build version.Build
}
