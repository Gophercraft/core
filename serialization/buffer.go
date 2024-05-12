package serialization

type Buffer struct {
	data      []byte
	offset    int
	bitOffset uint8
}

func NewBuffer(from []byte) *Buffer {
	buffer := new(Buffer)
	buffer.data = from
	return buffer
}

func MakeBuffer(cap int) *Buffer {
	buffer := new(Buffer)
	buffer.data = make([]byte, 0, cap)
	return buffer
}

func (buffer *Buffer) SetBytes(bytes []byte) {
	buffer.data = bytes
}

func (buffer *Buffer) Bytes() []byte {
	return buffer.data
}

func (buffer *Buffer) Len() int {
	return len(buffer.data)
}

func (buffer *Buffer) Cap() int {
	return cap(buffer.data)
}

func (buffer *Buffer) Reset() {
	buffer.data = buffer.data[:0]
	buffer.offset = 0
	buffer.bitOffset = 0
}

// Empty reports whether the unread portion of the buffer is empty.
func (b *Buffer) Empty() bool {
	return len(b.data) <= b.offset
}

func (b *Buffer) Remaining() int {
	if b.Empty() {
		return 0
	}

	return len(b.data) - b.offset
}
