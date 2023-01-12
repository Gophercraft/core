package terrain_test

import (
	"fmt"
	"testing"

	"github.com/Gophercraft/core/format/terrain"
	"github.com/Gophercraft/core/tempest"
	"github.com/davecgh/go-spew/spew"
)

func TestIndices(t *testing.T) {
	var param = terrain.DefaultMap

	pos := tempest.C2Vector{
		634, 1233,
	}

	bi, err := terrain.CalcBlockIndex(&param, pos)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(spew.Sdump(bi))
	ci, err := terrain.CalcChunkIndex(&param, pos)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(spew.Sdump(ci))
}

func TestIndexQueryInverse(t *testing.T) {
	m := &terrain.DefaultMap

	for _, check := range []struct {
		Input          terrain.BlockIndex
		ExpectedOutput tempest.C2Vector
	}{
		{
			terrain.BlockIndex{
				32, 32,
			},
			tempest.C2Vector{
				0,
				0,
			},
		},
		{
			terrain.BlockIndex{
				0, 0,
			},
			tempest.C2Vector{
				-17066.666016,
				-17066.666016,
			},
		},
	} {
		output, err := terrain.CalcBlockCornerPos(m, check.Input)
		if err != nil {
			t.Fatal(err)
		}

		if output != check.ExpectedOutput {
			t.Fatalf("CalcBlockCornerPos(..., %d:%d) should have returned %f:%f, but returned %f:%f", check.Input.X, check.Input.Y, check.ExpectedOutput.X, check.ExpectedOutput.Y, output.X, output.Y)
		}
	}
}
