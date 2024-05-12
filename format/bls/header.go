package bls

import (
	"fmt"
	"io"
)

const (
	Version1_3 = 0x10003
)

type Header struct {
	Magic            [4]byte
	Version          uint32
	PermutationCount uint32
	// ShaderCount            uint32
	// OffsetCompressedChunks uint32
	// CompressedChunkCount   uint32
	// OffsetCompressedData   uint32
}

func (r *reader) versionError() error {
	switch r.shader.Header.Version {
	case Version1_3:
		return nil
	default:
		return fmt.Errorf("bls: invalid shader format 0x%016X", r.shader.Header.Version)
	}
}

func (r *reader) readHeader1_3() (err error) {
	return r.readU32(&r.shader.Header.PermutationCount)
}

func (r *reader) validationError() (err error) {
	if r.shader.PermutationCount > 1024 {
		err = fmt.Errorf("bls: unreasonably large permutation count (%d)", r.shader.PermutationCount)
	}

	return
}

func (r *reader) readHeader(header *Header) (err error) {
	r.file.Seek(0, io.SeekStart)
	_, err = r.file.Read(r.shader.Header.Magic[:])
	if err != nil {
		return
	}

	if err = r.readU32(&r.shader.Header.Version); err != nil {
		return
	}

	if err = r.versionError(); err != nil {
		return
	}

	switch r.shader.Header.Version {
	case Version1_3:
		err = r.readHeader1_3()
	default:
		panic(r.shader.Header.Version)
	}

	if err != nil {
		return err
	}

	return r.validationError()
}
