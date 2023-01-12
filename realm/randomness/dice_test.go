package randomness_test

import (
	"fmt"
	"testing"

	"github.com/Gophercraft/core/realm/randomness"
)

func TestDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(randomness.RollPercent(100))
	}
}
