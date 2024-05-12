package connection

import (
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"

	"github.com/Gophercraft/core/crypto/streamcipher/arc4"
)

var (
	ARC4HashCipherServerDecryptionSeed9056 = []byte{
		0xC2, 0xB3, 0x72, 0x3C, 0xC6, 0xAE, 0xD9, 0xB5,
		0x34, 0x3C, 0x53, 0xEE, 0x2F, 0x43, 0x67, 0xCE,
	}

	ARC4HashCipherServerEncryptionSeed9056 = []byte{
		0xCC, 0x98, 0xAE, 0x04, 0xE8, 0x97, 0xEA, 0xCA,
		0x12, 0xDD, 0xC0, 0x93, 0x42, 0x91, 0x53, 0x57,
	}
)

// AuthCipher creates two ARC4 states for sending and recieving.
// They are created from pre-generated seeds, computed along with the session key.
type arc4_hash_cipher struct {
	ServerEncryptionKey []byte
	ServerDecryptionKey []byte

	encrypt cipher.Stream
	decrypt cipher.Stream
}

func NewARC4HashCipher(server_decryption_seed, server_encryption_seed, session_key []byte, server_mode bool) (connection_cipher Cipher, err error) {
	var (
		// The derived key used by the server to encrypt messages
		// (and therefore also the key used by the client to decrypt messages)
		derived_server_encryption_key []byte
		// The derived key used by the server to decrypt messages
		// (and therefore also the key used by the client to enccrypt messages)
		derived_server_decryption_key []byte
	)

	decryption_hash := hmac.New(sha1.New, server_decryption_seed)
	decryption_hash.Write(session_key)
	derived_server_decryption_key = decryption_hash.Sum(nil)

	encryption_hash := hmac.New(sha1.New, server_encryption_seed)
	encryption_hash.Write(session_key)
	derived_server_encryption_key = encryption_hash.Sum(nil)

	c := new(arc4_hash_cipher)

	if server_mode {
		c.encrypt = arc4.New(derived_server_encryption_key)
		c.decrypt = arc4.New(derived_server_decryption_key)
	} else {
		c.encrypt = arc4.New(derived_server_decryption_key)
		c.decrypt = arc4.New(derived_server_encryption_key)
	}

	connection_cipher = c
	return
}

func (c *arc4_hash_cipher) Encrypt(data, tag []byte) error {
	c.encrypt.XORKeyStream(data, data)
	return nil
}

func (c *arc4_hash_cipher) Decrypt(data, tag []byte) error {
	c.decrypt.XORKeyStream(data, data)
	return nil
}
