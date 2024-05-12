package bls

import (
	"encoding/binary"
	"io"
)

type reader struct {
	file   io.ReadSeeker
	shader *Shader
}

func newReader(file io.ReadSeeker, shader *Shader) *reader {
	r := new(reader)
	r.file = file
	r.shader = shader
	return r
}

func (r *reader) readU32(u32 *uint32) error {
	var u32data [4]byte
	_, err := r.file.Read(u32data[:])

	if err == nil {
		*u32 = binary.LittleEndian.Uint32(u32data[:])
	}

	return err
}

func (r *reader) readU16(u16 *uint16) error {
	var u16data [2]byte
	_, err := r.file.Read(u16data[:])

	if err == nil {
		*u16 = binary.LittleEndian.Uint16(u16data[:])
	}

	return err
}

func (r *reader) read() error {
	if err := r.readHeader(&r.shader.Header); err != nil {
		return err
	}

	if err := r.readPermutations(); err != nil {
		return err
	}

	return nil
}
