package serialization

func (buffer *Buffer) ReadCString() string {
	result := make([]byte, 0)

	for {
		b := buffer.ReadUint8()

		if b == 0 {
			break
		}

		result = append(result, b)
	}

	return string(result)
}
