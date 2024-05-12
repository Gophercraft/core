package connection

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
)

var (
	srvr = []byte("SRVR")
	clnt = []byte("CLNT")
)

type aes_gcm_cipher struct {
	encrypt_nonce_suffix []byte
	decrypt_nonce_suffix []byte
	encrypt              cipher.AEAD
	decrypt              cipher.AEAD
	encrypt_counter      uint64
	decrypt_counter      uint64
}

func NewAESGCMCipher(key []byte, server_mode bool) (connection_cipher Cipher, err error) {
	c := new(aes_gcm_cipher)

	if server_mode {
		c.encrypt_nonce_suffix = srvr
		c.decrypt_nonce_suffix = clnt
	} else {
		c.encrypt_nonce_suffix = clnt
		c.decrypt_nonce_suffix = srvr
	}

	var (
		encrypt cipher.Block
		decrypt cipher.Block
	)

	encrypt, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	decrypt, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	c.encrypt, err = cipher.NewGCMWithTagSize(encrypt, 12)
	if err != nil {
		return
	}
	c.decrypt, err = cipher.NewGCMWithTagSize(decrypt, 12)
	if err != nil {
		return
	}

	c.encrypt_counter = 1
	c.decrypt_counter = 1

	connection_cipher = c
	return
}

func (aec *aes_gcm_cipher) Encrypt(data, tag []byte) error {
	var nonce [12]byte
	binary.LittleEndian.PutUint64(nonce[0:8], aec.encrypt_counter)
	copy(nonce[8:], aec.encrypt_nonce_suffix)
	result := aec.encrypt.Seal(nil, nonce[:], data, nil)
	// crypto/aes makes the decision to append the authentication tag to the end of our ciphertext
	// We don't want this. Split them up into the slices that we want.
	copy(data[:], result[:len(result)-12])
	copy(tag[:], result[len(result)-12:])
	aec.encrypt_counter++
	return nil
}

func (aec *aes_gcm_cipher) Decrypt(data, tag []byte) error {
	var nonce [12]byte
	binary.LittleEndian.PutUint64(nonce[0:8], aec.decrypt_counter)
	copy(nonce[8:], aec.decrypt_nonce_suffix)
	// crypto/aes needs to have the tag at the end
	envelope := append(data, tag...)
	result, err := aec.decrypt.Open(nil, nonce[:], envelope, nil)
	if err != nil {
		return err
	}
	if len(data) != len(result) {
		return fmt.Errorf("crypto/connection: (*aes_gcm_cipher).Decrypt: encrypted data length is not the same as the decrypted size %d != %d", len(data), len(result))
	}
	copy(data[:], result[:])
	aec.decrypt_counter++
	return nil
}
