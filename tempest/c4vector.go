package tempest

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type C4Vector struct {
	X float32
	Y float32
	Z float32
	W float32
}

func parseFloat(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	return float32(f), err
}

func ParseC4Vector(in string) (C4Vector, error) {
	strs := strings.Split(in, " ")

	var (
		pos C4Vector
		err error
	)

	if len(strs) < 3 {
		return pos, fmt.Errorf("update: invalid Position string: only has %d coordinates", len(strs))
	}

	pos.X, err = parseFloat(strs[0])
	if err != nil {
		return pos, err
	}

	pos.Y, err = parseFloat(strs[1])
	if err != nil {
		return pos, err
	}
	pos.Z, err = parseFloat(strs[2])
	if err != nil {
		return pos, err
	}

	if len(strs) > 3 {
		pos.W, err = parseFloat(strs[3])
		if err != nil {
			return pos, err
		}
	}

	return pos, nil
}

func (c4 C4Vector) Encode(wr io.Writer) error {
	var floats [4 * 4]byte
	binary.LittleEndian.PutUint32(floats[0:4], math.Float32bits(c4.X))
	binary.LittleEndian.PutUint32(floats[4:8], math.Float32bits(c4.Y))
	binary.LittleEndian.PutUint32(floats[8:12], math.Float32bits(c4.Z))
	binary.LittleEndian.PutUint32(floats[12:16], math.Float32bits(c4.W))
	_, err := wr.Write(floats[:])
	return err
}

func DecodeC4Vector(r io.Reader) (C4Vector, error) {
	var vector C4Vector
	var floats [4 * 4]byte
	i, err := r.Read(floats[:])
	if err != nil {
		return vector, err
	}
	if i != len(floats) {
		return vector, fmt.Errorf("tempest: incomplete read of C4Vector")
	}
	vector.X = math.Float32frombits(binary.LittleEndian.Uint32(floats[0:4]))
	vector.Y = math.Float32frombits(binary.LittleEndian.Uint32(floats[4:8]))
	vector.Z = math.Float32frombits(binary.LittleEndian.Uint32(floats[8:12]))
	vector.W = math.Float32frombits(binary.LittleEndian.Uint32(floats[12:16]))
	return vector, nil
}

func (c4 C4Vector) C2() C2Vector {
	return C2Vector{
		c4.X,
		c4.Y,
	}
}

func (c4 C4Vector) C3() C3Vector {
	return C3Vector{
		c4.X,
		c4.Y,
		c4.Z,
	}
}

func (c4 C4Vector) Distance(oc4 C4Vector) float32 {
	return c4.C3().Distance(oc4.C3())
}
