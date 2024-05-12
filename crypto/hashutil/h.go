package hashutil

import "hash"

func H(hash_function func() hash.Hash, material ...[]byte) []byte {
	h := hash_function()
	for _, m := range material {
		h.Write(m)
	}
	return h.Sum(nil)
}
