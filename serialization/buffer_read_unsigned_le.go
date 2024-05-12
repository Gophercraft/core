package serialization

import "encoding/binary"

// returns the 8-bit unsigned integer at the current offset
func (buffer *Buffer) ReadUint8() uint8 {
	var b [1]byte
	buffer.Read(b[:])
	return b[0]
}

// Returns the 16-bit unsigned little endian integer positioned at the current offset
func (buffer *Buffer) ReadUint16() uint16 {
	var u16 [2]byte
	buffer.Read(u16[:])
	return binary.LittleEndian.Uint16(u16[:])
}

// Returns the 32-bit unsigned little endian integer positioned at the current offset
func (buffer *Buffer) ReadUint32() uint32 {
	var u32 [4]byte
	buffer.Read(u32[:])
	return binary.LittleEndian.Uint32(u32[:])
}

// Returns the 64-bit unsigned little endian integer positioned at the current offset
func (buffer *Buffer) ReadUint64() uint64 {
	var u64 [8]byte
	buffer.Read(u64[:])
	return binary.LittleEndian.Uint64(u64[:])
}
