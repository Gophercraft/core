package crypto

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

var (
	srvr = []byte("SRVR")
	clnt = []byte("CLNT")
)

type AesCipher struct {
	server                   bool
	send, recv               cipher.AEAD
	sendCounter, recvCounter uint64
}

func (aec *AesCipher) Init(server bool, key []byte) error {
	aec.server = server
	send, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	recv, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	aec.send, err = cipher.NewGCMWithTagSize(send, 12)
	if err != nil {
		return err
	}
	aec.recv, err = cipher.NewGCMWithTagSize(recv, 12)
	if err != nil {
		return err
	}
	aec.sendCounter = 1
	aec.recvCounter = 1
	return nil
}

func (aec *AesCipher) Encrypt(data, tag []byte) error {
	var nonce [12]byte
	binary.LittleEndian.PutUint64(nonce[0:8], aec.sendCounter)
	if aec.server {
		copy(nonce[8:], srvr)
	} else {
		copy(nonce[8:], clnt)
	}
	result := aec.send.Seal(nil, nonce[:], data, nil)
	copy(data[:], result[:len(result)-12])
	copy(tag[:], result[len(result)-12:])
	aec.sendCounter++
	return nil
}

func (aec *AesCipher) Decrypt(data, tag []byte) error {
	var nonce [12]byte
	binary.LittleEndian.PutUint64(nonce[0:8], aec.recvCounter)
	if aec.server {
		copy(nonce[8:], clnt)
	} else {
		copy(nonce[8:], srvr)
	}
	envelope := append(data, tag...)
	result, err := aec.recv.Open(nil, nonce[:], envelope, nil)
	if err != nil {
		return err
	}
	if len(data) != len(result) {
		return fmt.Errorf("encrypted data length is not the same as the decrypted size %d != %d", len(data), len(result))
	}
	copy(data[:], result[:])
	aec.recvCounter++
	return nil
}

func GenEnableEncryptionSignature(encryptKey []byte) ([]byte, error) {
	if len(encryptKey) != 16 {
		return nil, fmt.Errorf("crypto: invalid encryption key, needs to be 16 long")
	}

	connectToRsa := GetConnectionRSAKey()

	enableEncryptionSeed := [16]byte{0x90, 0x9C, 0xD0, 0x50, 0x5A, 0x2C, 0x14, 0xDD, 0x5C, 0x2C, 0xC0, 0x64, 0x14, 0xF3, 0xFE, 0xC9}

	hmc := hmac.New(sha256.New, encryptKey)
	hmc.Write([]byte{1}) // enabled
	hmc.Write(enableEncryptionSeed[:])
	enableEncryptionResult := hmc.Sum(nil)

	signature, err := rsa.SignPKCS1v15(nil, connectToRsa, crypto.SHA256, enableEncryptionResult)
	if err != nil {
		return nil, err
	}

	if len(signature) != 256 {
		panic(len(signature))
	}

	if err := rsa.VerifyPKCS1v15(connectToRsa.Public().(*rsa.PublicKey), crypto.SHA256, enableEncryptionResult, signature); err != nil {
		return nil, err
	}

	ReverseBytes(signature)
	return signature, nil
}
