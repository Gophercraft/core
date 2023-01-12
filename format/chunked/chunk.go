// Package chunked implements RIFF-like encoding used in the WDT and ADT terrain formats
package chunked

import (
	"encoding/binary"
	"fmt"
	"io"
	"unicode/utf8"
)

type Tag uint32

var Nil Tag

func (tag Tag) String() string {
	if tag == Nil {
		return "Nil"
	}

	var tagdata [4]byte
	binary.BigEndian.PutUint32(tagdata[:], uint32(tag))

	var tagstring string = string(tagdata[:])

	if !utf8.ValidString(tagstring) {
		panic("chunked: invalid string")
	}

	return string(tagdata[:])
}

func ID(s string) Tag {
	if len(s) != 4 {
		panic(s)
	}

	var inputBytes [4]byte
	copy(inputBytes[:], []byte(s))

	return Tag(binary.BigEndian.Uint32(inputBytes[:]))
}

type Reader struct {
	Reader io.Reader
}

func (c *Reader) ReadChunk() (tag Tag, bytes []byte, err error) {
	err = binary.Read(c.Reader, binary.LittleEndian, &tag)
	if err != nil {
		return
	}
	if tag == Nil {
		return
	}

	var size uint32
	err = binary.Read(c.Reader, binary.LittleEndian, &size)
	if err != nil {
		return
	}

	if size > 0xFFFFF {
		err = fmt.Errorf("chunked: chunk %s is way too big: %d bytes", tag, size)
		return
	}

	bytes = make([]byte, size)

	var i int

	i, err = io.ReadFull(c.Reader, bytes)
	if err != nil {
		return
	}

	if uint32(i) != size {
		err = fmt.Errorf("chunked: stream did not return all %d bytes referenced in this chunk %s (only %d)", size, tag, i)
		return
	}

	return
}
