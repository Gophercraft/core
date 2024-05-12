package connection

import (
	"crypto/cipher"

	"github.com/Gophercraft/core/crypto/streamcipher/xorstream"
)

type xor_stream_cipher struct {
	encrypt cipher.Stream
	decrypt cipher.Stream
}

func (xor_stream_cipher *xor_stream_cipher) Encrypt(data, tag []byte) (err error) {
	xor_stream_cipher.encrypt.XORKeyStream(data, data)
	return
}

func (xor_stream_cipher *xor_stream_cipher) Decrypt(data, tag []byte) (err error) {
	xor_stream_cipher.decrypt.XORKeyStream(data, data)
	return
}

// The first encryption scheme, added in version 4062
func NewXORStreamCipher(key []byte) (connection_cipher Cipher, err error) {
	c := new(xor_stream_cipher)
	c.encrypt, err = xorstream.NewEncrypt(key)
	if err != nil {
		return
	}
	c.decrypt, err = xorstream.NewDecrypt(key)
	if err != nil {
		return
	}
	connection_cipher = c
	return
}
