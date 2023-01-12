// randomness provides pseudo-random utilities for in-game use.
package randomness

func Float32(min, max float32) float32 {
	rng := max - min
	return min + float32(randRangeFloat(float64(rng)))
}

func Float64(min, max float64) float64 {
	rng := max - min
	return min + randRangeFloat(float64(rng))
}
