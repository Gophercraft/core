package crypto

import (
	"bytes"
	"crypto/sha1"
	"hash"
)

func SHA1(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}

func Hash(hsh func() hash.Hash, input ...[]byte) []byte {
	h := hsh()
	for _, in := range input {
		h.Write(in)
	}
	return h.Sum(nil)
}

const (
	// mod is the largest prime that is less than 65536.
	mod = 65521
	// nmax is the largest n such that
	// 255 * n * (n+1) / 2 + (n+1) * (mod-1) <= 2^32-1.
	// It is mentioned in RFC 1950 (search for "5552").
	nmax = 5552
)

// Add p to the running checksum d.
func Adler32(d uint32, p []byte) uint32 {
	s1, s2 := uint32(d&0xffff), uint32(d>>16)
	for len(p) > 0 {
		var q []byte
		if len(p) > nmax {
			p, q = p[:nmax], p[nmax:]
		}
		for len(p) >= 4 {
			s1 += uint32(p[0])
			s2 += s1
			s1 += uint32(p[1])
			s2 += s1
			s1 += uint32(p[2])
			s2 += s1
			s1 += uint32(p[3])
			s2 += s1
			p = p[4:]
		}
		for _, x := range p {
			s1 += uint32(x)
			s2 += s1
		}
		s1 %= mod
		s2 %= mod
		p = q
	}
	return s2<<16 | s1
}
