package serialization

import "io"

func ReadCString(reader io.Reader, cap int) (str string, err error) {
	data := make([]byte, 0, cap)

	for i := 0; i < cap; i++ {
		var character_byte [1]byte
		if _, err = io.ReadFull(reader, character_byte[:]); err != nil {
			return
		}

		if character_byte[0] == 0x00 {
			break
		}

		data = append(data, character_byte[0])
	}

	return string(data), err
}
