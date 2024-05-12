package serialization

// Read a string with a fixed capacity.
func (buffer *Buffer) ReadFixedString(n int) string {
	data := make([]byte, n)
	buffer.Read(data)
	return string(data)
}
