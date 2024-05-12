package serialization

// Appends signed 8-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteInt8(i8 int8) {
	buffer.WriteUint8(uint8(i8))
}

// Appends little-endian signed 16-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteInt16(i16 int16) {
	buffer.WriteUint16(uint16(i16))
}

// Appends little-endian signed 32-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteInt32(i32 int32) {
	buffer.WriteUint32(uint32(i32))
}

// Appends little-endian signed 64-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteInt64(i64 int64) {
	buffer.WriteUint64(uint64(i64))
}
