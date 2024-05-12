package serialization

import "math"

func (buffer *Buffer) WriteFloat32(f32 float32) {
	buffer.WriteUint32(math.Float32bits(f32))
}
