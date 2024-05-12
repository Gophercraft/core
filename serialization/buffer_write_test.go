package serialization

import (
	"io"
	"testing"
)

func TestBufferWrite(t *testing.T) {
	testBufferCommon(NewBuffer([]byte{}), t)
	testBufferCommon(MakeBuffer(100), t)
}

func testBufferCommon(buffer *Buffer, t *testing.T) {
	buffer.WriteUint32(1)
	buffer.WriteUint32(2)
	buffer.WriteUint32(3)

	if buffer.Len() != 12 {
		t.Fatal(buffer.Len())
	}

	buffer.WriteFloat32(4.0)

	if _, err := buffer.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}

	var u uint32

	u = buffer.ReadUint32()
	if u != 1 {
		t.Fatal(u)
	}

	u = buffer.ReadUint32()
	if u != 2 {
		t.Fatal(u)
	}

	u = buffer.ReadUint32()
	if u != 3 {
		t.Fatal(u)
	}

	buffer.Seek(0, io.SeekStart)
	buffer.WriteUint32()
}
