package serialization

// Finalize the writing of an 8-bit field.
func (buffer *Buffer) FlushBits() {
	buffer.bitOffset = 0
	buffer.offset += 1
}

// Write a bit at the current bit and byte offset.
// Will start writing at the most-significant bit (1 << 7) (0x80), with each successive call setting a less significant bit.
// Note that this will not increment the byte-offset until you reach your 8th bit (1 << 0) or (0x1)
func (buffer *Buffer) WriteBit(bit bool) {
	if buffer.bitOffset == 0 && buffer.offset == len(buffer.data) {
		// Reserve bit field space
		buffer.data = append(buffer.data, 0)
	}

	bitFlag := uint8(1 << (7 - buffer.bitOffset))

	if bit {
		buffer.data[buffer.offset] |= bitFlag
	} else {
		buffer.data[buffer.offset] &= ^(bitFlag)
	}

	buffer.bitOffset++

	if buffer.bitOffset == 8 {
		buffer.bitOffset = 0
		buffer.offset++
	}
}
