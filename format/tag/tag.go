package tag

import (
	"encoding/binary"
	"fmt"
	"io"
	"unicode/utf8"
)

type Tag uint32

// Return a short 4-character string encoded as a 32-bit integer value
func Make(text string) (tag Tag) {
	text_bytes := []byte(text)
	if len(text_bytes) != 4 {
		panic(fmt.Errorf("tag: cannot Make a tag with this size %d", len(text_bytes)))
	}
	tag = Tag(binary.BigEndian.Uint32(text_bytes[:]))
	return
}

func Read(reader io.Reader) (tag Tag, err error) {
	var tag_bytes [4]byte
	if _, err = io.ReadFull(reader, tag_bytes[:]); err != nil {
		return
	}
	tag = Tag(binary.LittleEndian.Uint32(tag_bytes[:]))
	return
}

func Write(writer io.Writer, tag Tag) (err error) {
	var tag_bytes [4]byte
	binary.LittleEndian.PutUint32(tag_bytes[:], uint32(tag))
	_, err = writer.Write(tag_bytes[:])
	return
}

func (tag Tag) String() string {
	var tag_bytes [4]byte
	binary.BigEndian.PutUint32(tag_bytes[:], uint32(tag))
	return string(tag_bytes[:])
}

func Clean(tag Tag) string {
	var tag_bytes [4]byte
	binary.BigEndian.PutUint32(tag_bytes[:], uint32(tag))
	cleaned := make([]byte, 0, 4)

	for _, tag_byte := range tag_bytes {
		if tag_byte != 0 {
			cleaned = append(cleaned, tag_byte)
		}
	}

	if !utf8.Valid(cleaned) {
		return ""
	}

	return string(cleaned)
}
