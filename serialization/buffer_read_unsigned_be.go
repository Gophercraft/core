package serialization

import "encoding/binary"

// Returns the 16-bit unsigned little endian integer positioned at the current offset
func (buffer *Buffer) ReadUint16BE() uint16 {
	var u16 [2]byte
	buffer.Read(u16[:])
	return binary.BigEndian.Uint16(u16[:])
}

// Returns the 32-bit unsigned little endian integer positioned at the current offset
func (buffer *Buffer) ReadUint32BE() uint32 {
	var u32 [4]byte
	buffer.Read(u32[:])
	return binary.BigEndian.Uint32(u32[:])
}

// Returns the 64-bit unsigned little endian integer positioned at the current offset
func (buffer *Buffer) ReadUint64BE() uint64 {
	var u64 [8]byte
	buffer.Read(u64[:])
	return binary.BigEndian.Uint64(u64[:])
}
