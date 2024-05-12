package connection

import (
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"

	"github.com/Gophercraft/core/crypto/streamcipher/xorstream"
)

// addBuild(8606, 2, 4, 3, "", "", "", "", "319AFAA3F2559682F9FF658BE01456255F456FB1", "D8B0ECFE534BC1131E19BAD1D4C0E813EEE4994F")

var XORStreamHashCipherSeed8606 = []byte{0x38, 0xA7, 0x83, 0x15, 0xF8, 0x92, 0x25, 0x30, 0x71, 0x98, 0x67, 0xB1, 0x8C, 0x4, 0xE2, 0xAA}

type xor_stream_hash_cipher struct {
	encrypt cipher.Stream
	decrypt cipher.Stream
}

func (xor_stream_hash_cipher *xor_stream_hash_cipher) Encrypt(data, tag []byte) (err error) {
	xor_stream_hash_cipher.encrypt.XORKeyStream(data, data)
	return
}

func (xor_stream_hash_cipher *xor_stream_hash_cipher) Decrypt(data, tag []byte) (err error) {
	xor_stream_hash_cipher.decrypt.XORKeyStream(data, data)
	return
}

func NewXORStreamHashCipher(key, seed []byte) (connection_cipher Cipher, err error) {
	hash := hmac.New(sha1.New, seed)
	hash.Write(key)
	derived_key := hash.Sum(nil)

	c := new(xor_stream_hash_cipher)
	c.encrypt, err = xorstream.NewEncrypt(derived_key)
	if err != nil {
		return
	}
	c.decrypt, err = xorstream.NewDecrypt(derived_key)
	if err != nil {
		return
	}
	connection_cipher = c
	return
}
