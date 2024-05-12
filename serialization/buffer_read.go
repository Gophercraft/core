package serialization

import "io"

func (buffer *Buffer) Read(b []byte) (n int, err error) {
	if buffer.Empty() {
		return 0, io.EOF
	}
	n = copy(b, buffer.data[buffer.offset:])
	buffer.offset += n
	return n, nil
}

func (buffer *Buffer) ReadBytes(n int) []byte {
	b := make([]byte, n)
	buffer.Read(b)
	return b
}
