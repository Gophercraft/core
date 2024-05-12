package auth

import (
	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/game/network"
	"github.com/Gophercraft/core/game/protocol/message/packet/auth"
	"github.com/Gophercraft/core/version"
)

func start_auth_challenge(session *network.Session) error {
	// Create auth context with this session
	auth_context := new(SessionContext)
	if err := session.SetServiceContext(auth_service_id, auth_context); err != nil {
		return err
	}

	// Auth challenge requires different behavior depending on what protocol version is being emulated
	emulated_build := session.Build()

	// Create auth challenge structure
	auth_context.challenge = &auth.Challenge{
		DosZeroBits:  1,
		DosChallenge: crypto.RandBytes(32),
	}
	switch {
	case version.Range(0, 3368).Contains(emulated_build):
		// Alpha authenticates purely with client data (in plaintext 😬)
	case version.Range(5875, 18414).Contains(emulated_build):
		// after 1.12.1 challenge bytes are added
		auth_context.challenge.Challenge = crypto.RandBytes(4)
	case version.Range(19027, version.Max).Contains(emulated_build):
		// in 6.0.2 this is expanded to 16 bytes
		auth_context.challenge.Challenge = crypto.RandBytes(16)
	default:
		panic(emulated_build)
	}

	// Send challenge to client.
	return session.Encode(auth_context.challenge)
}
