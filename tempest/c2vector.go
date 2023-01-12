package tempest

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type C2Vector struct {
	X float32
	Y float32
}

func (c2 C2Vector) String() string {
	return fmt.Sprintf("%f,%f", c2.X, c2.Y)
}

func (c2 C2Vector) Encode(wr io.Writer) error {
	var floats [2 * 4]byte
	binary.LittleEndian.PutUint32(floats[0:4], math.Float32bits(c2.X))
	binary.LittleEndian.PutUint32(floats[4:8], math.Float32bits(c2.Y))
	_, err := wr.Write(floats[:])
	return err
}

func DecodeC2Vector(r io.Reader) (C2Vector, error) {
	var vector C2Vector
	var floats [2 * 4]byte
	i, err := r.Read(floats[:])
	if err != nil {
		return vector, err
	}
	if i != len(floats) {
		return vector, fmt.Errorf("tempest: incomplete read of C2Vector")
	}
	vector.X = math.Float32frombits(binary.LittleEndian.Uint32(floats[0:4]))
	vector.Y = math.Float32frombits(binary.LittleEndian.Uint32(floats[4:8]))
	return vector, nil
}
