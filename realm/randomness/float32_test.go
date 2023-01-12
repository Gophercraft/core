package randomness_test

import (
	"fmt"
	"testing"

	"github.com/Gophercraft/core/realm/randomness"
)

func TestFloat32(t *testing.T) {
	for x := 0; x < 10; x++ {
		value := randomness.Float32(0, 100)

		if value < 0 {
			t.Fatal(value)
		}

		if value > 100 {
			t.Fatal(value)
		}

		fmt.Println(value)
	}
}
