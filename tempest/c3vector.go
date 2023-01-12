package tempest

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/arl/math32"
)

type C3Vector struct {
	X float32
	Y float32
	Z float32
}

func (c3 C3Vector) Encode(wr io.Writer) error {
	var floats [3 * 4]byte
	binary.LittleEndian.PutUint32(floats[0:4], math.Float32bits(c3.X))
	binary.LittleEndian.PutUint32(floats[4:8], math.Float32bits(c3.Y))
	binary.LittleEndian.PutUint32(floats[8:12], math.Float32bits(c3.Z))
	_, err := wr.Write(floats[:])
	return err
}

func DecodeC3Vector(r io.Reader) (C3Vector, error) {
	var vector C3Vector
	var floats [3 * 4]byte
	i, err := r.Read(floats[:])
	if err != nil {
		return vector, err
	}
	if i != len(floats) {
		return vector, fmt.Errorf("tempest: incomplete read of C3Vector")
	}
	vector.X = math.Float32frombits(binary.LittleEndian.Uint32(floats[0:4]))
	vector.Y = math.Float32frombits(binary.LittleEndian.Uint32(floats[4:8]))
	vector.Z = math.Float32frombits(binary.LittleEndian.Uint32(floats[8:12]))
	return vector, nil
}

func (c3 *C3Vector) Multiply(a float32) {
	c3.X *= a
	c3.Y *= a
	c3.Z *= a
}

func (c3 *C3Vector) Sub(oc3 C3Vector) C3Vector {
	return C3Vector{
		c3.X - oc3.X,
		c3.Y - oc3.Y,
		c3.Z - oc3.Z,
	}
}

func (c3 *C3Vector) SquaredMag() float32 {
	return c3.X*c3.X + c3.Y*c3.Y + c3.Z*c3.Z
}

func (c3 C3Vector) Mag() float32 {
	return math32.Sqrt(c3.SquaredMag())
}

func (c3 *C3Vector) Normalize() {
	c3.Multiply(1.0 / c3.Mag())
}

func (c3 *C3Vector) Dot(oc3 C3Vector) float32 {
	return c3.X*oc3.X + c3.Y*oc3.Y + c3.Z*oc3.Z
}

func (c3 C3Vector) Distance(oc3 C3Vector) float32 {
	return c3.Sub(oc3).Mag()
}

func (c3 C3Vector) C4() C4Vector {
	return C4Vector{
		c3.X,
		c3.Y,
		c3.Z,
		0,
	}
}

func (c3 C3Vector) C2() C2Vector {
	return C2Vector{
		c3.X,
		c3.Y,
	}
}
