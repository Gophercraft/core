package randomness

func Int(min, max int) int {
	return min + int(randRangeInt(int64(max-min)))
}
