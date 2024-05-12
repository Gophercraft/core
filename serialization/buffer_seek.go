package serialization

import (
	"errors"
	"io"
)

func (buffer *Buffer) Seek(offset int64, whence int) (int64, error) {
	buffer.ResetBits()

	var abs int64

	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = int64(buffer.offset) + offset
	case io.SeekEnd:
		abs = int64(len(buffer.data)) + offset
	default:
		return 0, errors.New("serialization.Buffer.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("serialization.Buffer.Seek: negative position")
	}
	buffer.offset = int(abs)
	return abs, nil
}
