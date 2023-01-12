package randomness_test

import (
	"fmt"
	"testing"

	"github.com/Gophercraft/core/realm/randomness"
)

func TestWeightedSelect(t *testing.T) {
	weights := []int{
		1000,
		3,
		500,
		90,
		20,
	}

	for i := 0; i < 20; i++ {
		sel := randomness.WeightedSelect(weights)
		fmt.Println(weights[sel])
	}
}
