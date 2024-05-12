package serialization

func (buffer *Buffer) ReadBool() bool {
	boolField := buffer.ReadUint8()
	return boolField == 1
}

func (buffer *Buffer) ReadBool32() bool {
	boolField := buffer.ReadUint32()
	return boolField == 1
}
