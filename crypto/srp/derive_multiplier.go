package srp

import (
	"hash"

	"github.com/Gophercraft/core/crypto/hashutil"
)

// k = H(N, g)
func DeriveMultiplier(hash_func func() hash.Hash, N, g *Int) *Int {
	const material_length = 256

	digest := hashutil.H(hash_func, N.ArrayBigEndian(material_length), g.ArrayBigEndian(material_length))
	// digest := hashutil.H(hash_func, N.BytesBigEndian(), g.BytesBigEndian())

	return NewIntFromBytesBigEndian(digest)
}
