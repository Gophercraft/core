package srp

func CalculatePrivateB(N *Int) *Int {
	b := NewRandomInt(N.Bits())
	b = b.Mod(N.Subtract(NewInt(1)))
	return b
}

func CalculatePublicB(v, b, N, g, k *Int) *Int {
	return (g.ModExp(b, N).Add(v.Multiply(k))).Mod(N)
}
