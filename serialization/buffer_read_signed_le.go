package serialization

// returns the 8-bit signed integer at the current offset
func (buffer *Buffer) ReadInt8() int8 {
	return int8(buffer.ReadUint8())
}

// Returns the 16-bit signed little endian integer positioned at the current offset
func (buffer *Buffer) ReadInt16() int16 {
	return int16(buffer.ReadUint16())
}

// Returns the 32-bit signed little endian integer positioned at the current offset
func (buffer *Buffer) ReadInt32() int32 {
	return int32(buffer.ReadUint32())
}

// Returns the 64-bit unsigned little endian integer positioned at the current offset
func (buffer *Buffer) ReadInt64() int64 {
	return int64(buffer.ReadUint64())
}
