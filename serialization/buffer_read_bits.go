package serialization

func (buffer *Buffer) ResetBits() {
	buffer.bitOffset = 0
}

// The first bit is the most significant bit (equivalent to a flag of 0x80)
// the second bit returned is less significant (eq. to a flag of 0x40)
// and so on, until the bit flag is 1 (least significant bit)
// After which, the byte offset will move on to the next byte (or 8-bit field).
func (buffer *Buffer) ReadBit() bool {
	// Move
	if buffer.bitOffset == 7 {
		buffer.bitOffset = 0
		buffer.offset = buffer.offset + 1
	}

	bitField := buffer.data[buffer.offset]

	flag := uint8(1 << (7 - buffer.bitOffset))

	buffer.bitOffset++

	return bitField&flag != 0
}
