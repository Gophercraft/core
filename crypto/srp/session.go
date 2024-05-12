package srp

import "hash"

type Session struct {
	Version    uint32
	Iterations uint32
	Hash       func() hash.Hash
	// N
	LargeSafePrime *Int
	// g
	Generator *Int
	// k
	Multiplier *Int
	// v
	Verifier *Int
	// b
	PrivateB *Int
	// B
	PublicB *Int
	// s
	Salt []byte
}
