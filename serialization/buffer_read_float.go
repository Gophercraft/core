package serialization

import "math"

func (buffer *Buffer) ReadFloat32() float32 {
	return math.Float32frombits(buffer.ReadUint32())
}
