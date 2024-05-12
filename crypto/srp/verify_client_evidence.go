package srp

import (
	"hash"

	"github.com/Gophercraft/log"

	"github.com/davecgh/go-spew/spew"
)

func GetBrokenEvidenceSlice(number *Int) []byte {
	bytes := (int32(number.Bits()) + int32(8)) >> int32(3)
	return number.SliceBigEndian(int(bytes))
}

func CalculateEvidence(hash_func func() hash.Hash, numbers ...*Int) *Int {
	h := hash_func()

	for _, number := range numbers {
		h.Write(GetBrokenEvidenceSlice(number))
	}

	return NewIntFromBytesBigEndian(h.Sum(nil))
}

func VerifyClientEvidence(hash_func func() hash.Hash, A, b, B, N, v, client_M1 *Int) (K *Int, valid bool) {
	if A.Mod(N).Equals(Zero) {
		return
	}

	u := CalculateU(hash_func, A, B)

	if u.Mod(N).Equals(Zero) {
		return
	}

	S := (A.Multiply(v.ModExp(u, N)).ModExp(b, N))

	server_M1 := CalculateEvidence(hash_func, A, B, S)
	if !server_M1.Equals(client_M1) {
		log.Println("==================================================")
		log.Println("Does not match.")
		log.Println("Server M1 = ", spew.Sdump(server_M1.Bytes()))
		log.Println("Client M1 = ", spew.Sdump(client_M1.Bytes()))
		log.Println("==================================================")
		return
	}

	valid = true
	K = S
	return
}

func CalculateServerEvidence(hash_func func() hash.Hash, A, client_M1, K *Int) (server_M2 *Int) {
	return CalculateEvidence(hash_func, A, client_M1, K)
}
