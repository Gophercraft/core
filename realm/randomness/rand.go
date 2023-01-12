package randomness

import (
	"crypto/rand"
	"math/big"
)

func randRangeInt(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

func randRangeFloat(r float64) float64 {
	scale := float64(randRangeInt(1<<53)) / (1 << 53)
	return scale * r
}
