package serialization

func (buffer *Buffer) WriteCString(s string) {
	WriteCString(buffer, s)
}
