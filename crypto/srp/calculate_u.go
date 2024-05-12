package srp

import (
	"hash"

	"github.com/Gophercraft/core/crypto/hashutil"
)

func CalculateU(hash_func func() hash.Hash, A, B *Int) (U *Int) {
	const material_length = 256

	digest := hashutil.H(hash_func, A.ArrayBigEndian(material_length), B.ArrayBigEndian(material_length))
	// digest := hashutil.H(hash_func, A.BytesBigEndian(), B.BytesBigEndian())
	U = NewIntFromBytesBigEndian(digest)

	return
}
