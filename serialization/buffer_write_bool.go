package serialization

// For more efficient storage of multiple boolean fields, see buffer_write_bits.go

// Write an 8-bit aligned boolean field to the buffer.
func (buffer *Buffer) WriteBool(b bool) {
	var data uint8 = 0

	if b {
		data = 1
	}

	buffer.WriteUint8(data)
}

// Write a 32-bit aligned boolean field to the buffer
// For some reason, in older versions of the protocol a full 32-bits was used to communicate a boolean value.
func (buffer *Buffer) WriteBool32(b bool) {
	var data uint32 = 0

	if b {
		data = 1
	}

	buffer.WriteUint32(data)
}
