package tempest

import (
	"encoding/binary"
	"io"
	"math"
)

type C4QuaternionCompressed uint64

type C4Quaternion C4Vector

func (cq *C4Quaternion) DecodePacked(reader io.Reader) error {
	var packedq C4QuaternionCompressed

	if err := binary.Read(reader, binary.LittleEndian, &packedq); err != nil {
		return err
	}

	*cq = packedq.Unpack()
	return nil
}

func (cq *C4Quaternion) Decode(reader io.Reader) (err error) {
	var c4v C4Vector
	c4v, err = DecodeC4Vector(reader)
	*cq = C4Quaternion(c4v)
	return
}

const (
	packCoeffYZ = 1 << 20
	packCoeffX  = 1 << 21
)

func (c4 *C4QuaternionCompressed) Unpack() C4Quaternion {
	uraw := uint64(*c4)

	raw := int64(uraw)

	x := float64(raw>>4) / packCoeffX
	y := float64(raw<<22>>43) / packCoeffYZ
	z := float64(raw<<43>>42) / packCoeffYZ
	w := 1 - (x*x + y*y + z*z)

	w = math.Sqrt(w)

	return C4Quaternion{
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	}
}
func (cq *C4Quaternion) Pack() (raw C4QuaternionCompressed) {
	var wSign int32
	if cq.W >= 0 {
		wSign = 1
	} else {
		wSign = -1
	}

	x := C4QuaternionCompressed(int32(int32(cq.X)*packCoeffX*wSign) & ((1 << 22) - 1))
	y := C4QuaternionCompressed(int32(int32(cq.Y)*packCoeffYZ*wSign) & ((1 << 21) - 1))
	z := C4QuaternionCompressed(int32(int32(cq.Z)*packCoeffYZ*wSign) & ((1 << 21) - 1))

	raw = z | (y << 21) | (x << 42)
	return
}

func (cq *C4Quaternion) EncodePacked(writer io.Writer) (err error) {
	var uints [8]byte
	binary.LittleEndian.PutUint64(uints[:], uint64(cq.Pack()))
	_, err = writer.Write(uints[:])
	return
}
