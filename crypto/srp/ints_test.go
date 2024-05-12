package srp_test

import (
	"bytes"
	"testing"

	"github.com/Gophercraft/core/crypto/srp"
	"github.com/davecgh/go-spew/spew"
)

func TestInts(t *testing.T) {
	test := srp.NewInt(1)
	if !bytes.Equal(test.Array(8), []byte{1, 0, 0, 0, 0, 0, 0, 0}) {
		t.Fatal(spew.Sdump(test.Array(8)))
	}

	if !bytes.Equal(test.ArrayBigEndian(8), []byte{0, 0, 0, 0, 0, 0, 0, 1}) {
		t.Fatal(spew.Sdump(test.ArrayBigEndian(8)))
	}

	test = srp.NewInt(256)

	if !bytes.Equal(test.Array(8), []byte{0, 1, 0, 0, 0, 0, 0, 0}) {
		t.Fatal(spew.Sdump(test.Array(8)))
	}

	if !bytes.Equal(test.ArrayBigEndian(8), []byte{0, 0, 0, 0, 0, 0, 1, 0}) {
		t.Fatal(spew.Sdump(test.ArrayBigEndian(8)))
	}

	// test contamination
	value := srp.NewInt(5)
	value_2 := value.Add(srp.NewInt(2))
	if !value.Equals(srp.NewInt(5)) {
		t.Fatal(value.String())
	}
	if !value_2.Equals(srp.NewInt(7)) {
		t.Fatal(value_2.String())
	}
}
