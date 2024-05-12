package connection

import (
	"fmt"

	"github.com/Gophercraft/core/version"
)

// Cipher describes some method for setting up an encryption layer over a TCP socket.
type Cipher interface {
	Encrypt(data, tag []byte) error
	Decrypt(data, tag []byte) error
}

func NewCipher(build version.Build, session_key []byte, server bool) (cipher Cipher, err error) {
	switch {
	case build < 4062:
		// No encryption!
		cipher = NewDummyCipher()
	case build >= 4062 && build < version.V2_4_3:
		// Basic XOR+counter encryption was added to protocol 4062.
		cipher, err = NewXORStreamCipher(session_key)
	case build.AddedIn(version.V2_4_3) && build < 9614:
		// HMAC-SHA1 key agreement step added with pinned seed
		cipher, err = NewXORStreamHashCipher(session_key, XORStreamHashCipherSeed8606)
	case build.AddedIn(9056) && build.RemovedIn(version.NewCryptSystem):
		// The cipher now includes two HMAC-SHA1 key-agreement seeds.
		// And the send/receive stream ciphers are now ARC4, instead of basic XOR+counter
		cipher, err = NewARC4HashCipher(
			ARC4HashCipherServerDecryptionSeed9056,
			ARC4HashCipherServerEncryptionSeed9056,
			session_key,
			server)
	case build >= version.NewCryptSystem:
		// Blizzard eventually wised up and decided to use a proper encryption algorithm
		// ( Encrypting entire packets now, not just size + opcode fields! :D )
		cipher, err = NewAESGCMCipher(
			session_key,
			server)
	default:
		return nil, fmt.Errorf("crypto/connection: no cipher exists for protocol %s", build)
	}

	return
}
