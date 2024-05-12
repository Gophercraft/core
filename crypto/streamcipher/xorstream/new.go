package xorstream

import (
	"crypto/cipher"
	"fmt"
)

type cipher_state struct {
	key  []byte
	i, j uint8
}

type encrypt_cipher struct {
	cipher_state
}

type decrypt_cipher struct {
	cipher_state
}

// Creates an encryption-mode cipher
func NewEncrypt(key []byte) (stream cipher.Stream, err error) {
	if len(key) == 0 {
		err = fmt.Errorf("crypto/xorstream: key must have length")
		return
	}
	encryption_cipher := new(encrypt_cipher)
	encryption_cipher.key = key
	stream = encryption_cipher
	return
}

// Creates an decryption-mode cipher
func NewDecrypt(key []byte) (stream cipher.Stream, err error) {
	if len(key) == 0 {
		err = fmt.Errorf("crypto/xorstream: key must have length")
		return
	}
	decryption_cipher := new(decrypt_cipher)
	decryption_cipher.key = key
	stream = decryption_cipher
	return
}

func (cipher *encrypt_cipher) XORKeyStream(dst, src []byte) {
	for t := 0; t < len(src); t++ {
		cipher.i %= uint8(len(cipher.key))
		xor_byte := (src[t] ^ cipher.key[cipher.i]) + cipher.j
		cipher.i++
		cipher.j = xor_byte
		dst[t] = cipher.j
	}
}

func (cipher *decrypt_cipher) XORKeyStream(dst, src []byte) {
	for t := 0; t < len(src); t++ {
		cipher.i %= uint8(len(cipher.key))
		xor_byte := (src[t] - cipher.j) ^ cipher.key[cipher.i]
		cipher.i++
		cipher.j = src[t]
		dst[t] = xor_byte
	}
}
