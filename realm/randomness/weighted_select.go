package randomness

func WeightedSelect(weights []int) int {
	if len(weights) == 1 {
		return 0
	}

	var sum int

	for _, w := range weights {
		sum += w
	}

	selection := Int(1, sum)

	for i, weight := range weights {
		selection -= weight
		if selection <= 0 {
			return i
		}
	}

	return 0
}
