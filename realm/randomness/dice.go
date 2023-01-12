package randomness

// Return true %percent of the times RollPercent is called.
func RollPercent(percent float32) bool {
	roll := Float32(0, 100)
	return roll > (100 - percent)
}
