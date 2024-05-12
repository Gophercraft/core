package serialization

import (
	"io"
	"strconv"
	"testing"
)

// Ensure 7 bits are written without moving to the next byte
func TestWriteBitsPartial(t *testing.T) {
	buffer := MakeBuffer(1)

	// Write seven bits
	buffer.WriteBit(true)  // 0
	buffer.WriteBit(false) // 1
	buffer.WriteBit(true)  // 2
	buffer.WriteBit(false) // 3
	buffer.WriteBit(true)  // 4
	buffer.WriteBit(false) // 5
	buffer.WriteBit(true)  // 6

	if buffer.bitOffset != 7 {
		t.Fatal(buffer.bitOffset)
	}

	cur, err := buffer.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatal(err)
	}

	if cur != 0 {
		t.Fatal(cur)
	}
}

// Ensure buffer moves to the next byte automatically after 8 bits
func TestWriteBitsFull(t *testing.T) {
	buffer := MakeBuffer(1)

	// Write seven bits
	buffer.WriteBit(true)  // 0
	buffer.WriteBit(false) // 1
	buffer.WriteBit(true)  // 2
	buffer.WriteBit(false) // 3
	buffer.WriteBit(true)  // 4
	buffer.WriteBit(false) // 5
	buffer.WriteBit(true)  // 6
	buffer.WriteBit(false) // 7

	if buffer.bitOffset != 0 {
		t.Fatal(buffer.bitOffset)
	}

	cur, err := buffer.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatal(err)
	}

	if cur != 1 {
		t.Fatal(cur)
	}

	buffer.Seek(0, io.SeekStart)

	binaryRepresentation := strconv.FormatInt(int64(buffer.ReadUint8()), 2)

	if binaryRepresentation != "10101010" {
		t.Fatal(binaryRepresentation)
	}
}

// Ensure
func TestWriteBitsFlushingCorrectly(t *testing.T) {
	buffer := MakeBuffer(1)

	// Write seven bits
	buffer.WriteBit(true)  // 0
	buffer.WriteBit(false) // 1
	buffer.WriteBit(true)  // 2
	buffer.WriteBit(false) // 3
	buffer.WriteBit(true)  // 4
	buffer.WriteBit(false) // 5
	buffer.WriteBit(true)  // 6

	buffer.FlushBits()

	if buffer.bitOffset != 0 {
		t.Fatal(buffer.bitOffset)
	}

	cur, err := buffer.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatal(err)
	}

	if cur != 1 {
		t.Fatal(cur)
	}
}
