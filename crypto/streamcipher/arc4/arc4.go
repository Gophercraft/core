// Package arc4 provides an implementation of the RC4-drop-1024 stream cipher
// WARNING: RC4 is widely deemed as UNSAFE by cryptography specialists.
// We only are providing this package for backward compatibility with older game versions
package arc4

import "crypto/cipher"

type arc4_cipher struct {
	S    [256]byte
	i, j byte
}

func New(key []byte) (stream_cipher cipher.Stream) {
	c := new(arc4_cipher)
	// Identity permutation
	for i := 0; i < 256; i++ {
		c.S[i] = byte(i)
	}

	// Key scheduling
	var j uint8 = 0
	var t uint8 = 0
	for i := 0; i < 256; i++ {
		j = (j + c.S[i] + key[i%len(key)]) & 255
		t = c.S[i]
		c.S[i] = c.S[j]
		c.S[j] = t
	}

	c.i = 0
	c.j = 0

	// Drop-1024 variant
	for range 1024 {
		c.next_byte()
	}

	return c
}

func (c *arc4_cipher) next_byte() uint8 {
	var t uint8
	c.i = (c.i + 1) & 255
	c.j = (c.j + c.S[c.i]) & 255
	t = c.S[c.i]
	c.S[c.i] = c.S[c.j]
	c.S[c.j] = t
	return c.S[(t+c.S[c.i])&255]
}

func (c *arc4_cipher) XORKeyStream(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ c.next_byte()
	}
}
