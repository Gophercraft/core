package tag_test

import (
	"bytes"
	"testing"

	"github.com/Gophercraft/core/format/tag"
)

func TestTag(t *testing.T) {
	stream := bytes.NewReader([]byte{
		0x57, 0x6f, 0x57, 0x00,
	})

	text := "\x00WoW"

	result, err := tag.Read(stream)
	if err != nil {
		t.Fatal(err)
	}

	if result != tag.Make(text) {
		t.Fatal("mismatch")
	}

	if text != result.String() {
		t.Fatal("string mismatch")
	}
}
