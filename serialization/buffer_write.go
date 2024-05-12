package serialization

func (buffer *Buffer) Write(b []byte) (n int, err error) {
	availableSpace := len(buffer.data) - buffer.offset

	if availableSpace >= len(b) {
		n = copy(buffer.data[buffer.offset:], b)
		buffer.offset += len(b)
		buffer.bitOffset = 0
		return
	}

	// How far we need to grow the buffer
	neededSpace := len(b) - availableSpace

	// Ensure space is available in the buffer to hold bytes by reallocating
	buffer.data = append(buffer.data, make([]byte, neededSpace)...)

	// Copy data into expanded buffer
	n = copy(buffer.data[buffer.offset:], b)
	buffer.offset += len(b)
	buffer.bitOffset = 0
	return
}
