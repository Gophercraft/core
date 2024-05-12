package serialization

import "io"

var null = [1]byte{0}

func WriteCString(writer io.Writer, str string) (err error) {
	_, err = io.WriteString(writer, str)
	if err != nil {
		return
	}
	if _, err = writer.Write(null[:]); err != nil {
		return
	}
	return
}
