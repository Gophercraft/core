//Package srp implements a backward-compatible version of SRP-6.
//Warning: this package is ONLY suitable for Gophercraft.
//Do not use it in any other case: it provides little in the way of security.

package srp

import (
	"bytes"
	"crypto/sha1"
	"strings"

	"github.com/Gophercraft/core/crypto/hashutil"
)

var (
	Zero       = NewInt(0)
	Generator  = NewInt(7)
	Multiplier = NewInt(3)
	Prime      = NewIntFromString("62100066509156017342069496140902949863249758336000796928566441170293728648119")
)

// Ngh = XOR[20](H(N), H(g))
func HashPrimeAndGenerator(N, g *Int) []byte {
	Nh := hashutil.H(sha1.New, N.Bytes())
	gh := hashutil.H(sha1.New, g.Bytes())

	Ngh := make([]byte, 20)
	for i := 0; i < 20; i++ {
		Ngh[i] = Nh[i] ^ gh[i]
	}

	return Ngh
}

func ServerGenerateEphemeralValues(g, N, v *Int) (b *Int, B *Int) {
	b = NewRandomInt(19)
	gMod := g.ModExp(b, N)
	B = ((v.Multiply(Multiplier.Copy())).Add(gMod)).Mod(N)
	return
}

func HashCalculate(username string, identity_hash, server_public_key, large_safe_prime, generator, salt []byte) (verifier *Int, session_key []byte, public_key []byte, proof []byte) {
	g := NewIntFromBytes(generator)

	k := Multiplier.Copy()

	N := NewIntFromBytes(large_safe_prime)

	x, v := CalculateVerifier(sha1.New, salt, identity_hash, g, N)

	a := NewRandomInt(19)
	A := g.ModExp(a, N)

	B := NewIntFromBytes(server_public_key)

	uh := hashutil.H(sha1.New, A.Bytes(), B.Bytes())
	u := NewIntFromBytes(uh)

	kgx := k.Multiply(g.ModExp(x, N))
	aux := a.Add(u.Multiply(x))

	_S := B.Subtract(kgx).ModExp(aux, N)
	S := _S.Bytes()

	if len(S) > 32 {
		S = S[:32]
	} else {
		S = append(S, make([]byte, 32-len(S))...)
	}

	S1, S2 := make([]byte, 16), make([]byte, 16)

	for i := 0; i < 16; i++ {
		S1[i] = S[i*2]
		S2[i] = S[i*2+1]
	}

	S1h := hashutil.H(sha1.New, S1)
	S2h := hashutil.H(sha1.New, S2)

	K := make([]byte, 40)

	for i := 0; i < 20; i++ {
		K[i*2] = S1h[i]
		K[i*2+1] = S2h[i]
	}

	userh := hashutil.H(sha1.New, []byte(strings.ToUpper(username)))

	Ngh := HashPrimeAndGenerator(N, g)

	M1 := hashutil.H(sha1.New,
		Ngh,
		userh,
		salt,
		A.Bytes(),
		B.Bytes(),
		K,
	)

	return v, K, A.Bytes(), M1
}

func ServerLogonProof(username string, A, M1, b, B, s, N, v *Int) ([]byte, bool, []byte) {
	g := Generator.Copy()

	u := NewIntFromBytes(hashutil.H(sha1.New, A.Bytes(), B.Bytes()))
	if A.Mod(N).Equals(Zero) {
		return nil, false, nil
	}

	_S := (A.Multiply(v.ModExp(u, N))).ModExp(b, N)

	S := _S.Bytes()

	S1, S2 := make([]byte, 16), make([]byte, 16)

	for i := 0; i < 16; i++ {
		if len(S) < 32 {
			return nil, false, nil
		}
		S1[i] = S[i*2]
		S2[i] = S[i*2+1]
	}

	S1h := hashutil.H(sha1.New, S1)
	S2h := hashutil.H(sha1.New, S2)

	vK := make([]byte, 40)

	for i := 0; i < 20; i++ {
		vK[i*2] = S1h[i]
		vK[i*2+1] = S2h[i]
	}

	K := NewIntFromBytes(vK)

	Nh := hashutil.H(sha1.New, N.Bytes())
	gh := hashutil.H(sha1.New, g.Bytes())

	for i := 0; i < 20; i++ {
		Nh[i] ^= gh[i]
	}

	t3 := NewIntFromBytes(Nh)
	t4 := hashutil.H(sha1.New, []byte(strings.ToUpper(username)))

	final := hashutil.H(sha1.New,
		t3.Bytes(),
		t4,
		s.Bytes(),
		A.Bytes(),
		B.Bytes(),
		K.Bytes(),
	)

	M2 := hashutil.H(sha1.New, A.Bytes(), final, K.Bytes())
	return vK, bytes.Equal(final, M1.Bytes()), M2
}
