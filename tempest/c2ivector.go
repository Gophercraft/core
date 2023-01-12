package tempest

import (
	"encoding/binary"
	"fmt"
	"io"
)

type C2iVector struct {
	X int32
	Y int32
}

func DecodeC2iVector(r io.Reader) (C2iVector, error) {
	var vector C2iVector
	var ints [2 * 4]byte
	i, err := r.Read(ints[:])
	if err != nil {
		return vector, err
	}
	if i != len(ints) {
		return vector, fmt.Errorf("tempest: incomplete read of C2Vector")
	}
	vector.X = int32(binary.LittleEndian.Uint32(ints[0:4]))
	vector.Y = int32(binary.LittleEndian.Uint32(ints[4:8]))
	return vector, nil
}
