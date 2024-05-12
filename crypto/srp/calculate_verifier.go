package srp

import (
	"hash"
)

// Compute auth := H('username' + ':' + 'pass')
// g := 7
// ....
// x := H(salt, auth)
// v := (g^x) % N
func CalculateVerifier(hash_func func() hash.Hash, salt, auth []byte, g, N *Int) (x *Int, v *Int) {
	h := hash_func()
	h.Write(salt)
	h.Write(auth)
	digest := h.Sum(nil)

	x = NewIntFromBytes(digest)
	v = g.ModExp(x, N)
	return x, v
}

// BNet protocol equivalent
// Compute auth := PBKDF2_SHA256('username' + ':' + 'pass', salt)
func CalculateVerifier2(hash_func func() hash.Hash, salt, auth []byte, g, N *Int) (x *Int, v *Int) {
	var x_bytes [64]byte
	copy(x_bytes[:], auth)
	x = NewIntFromBytesBigEndian(x_bytes[:])
	if x_bytes[0]&0x80 != 0 {
		var fix [65]byte
		fix[64] = 1
		x = x.Subtract(NewIntFromBytes(fix[:]))
	}

	x = x.Mod(N.Subtract(NewInt(1)))
	v = g.ModExp(x, N)

	return
}
