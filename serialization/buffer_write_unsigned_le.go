package serialization

import "encoding/binary"

// Appends unsigned 8-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteUint8(u8 uint8) {
	buffer.Write([]byte{u8})
}

// Appends little-endian unsigned 16-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteUint16(u16 uint16) {
	var b [2]byte
	binary.LittleEndian.PutUint16(b[:], u16)
	buffer.Write(b[:])
}

// Appends little-endian unsigned 32-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteUint32(u32 uint32) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], u32)
	buffer.Write(b[:])
}

// Appends little-endian unsigned 64-bit integer to the buffer at the current offset
func (buffer *Buffer) WriteUint64(u64 uint64) {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], u64)
	buffer.Write(b[:])
}
