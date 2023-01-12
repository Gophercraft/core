package packet

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/superp00t/etc"
)

func GetMSTime() uint32 {
	return uint32(time.Now().UnixNano() / int64(time.Millisecond))
}

func ReverseBuffer(input []byte) []byte {
	buf := make([]byte, len(input))
	inc := 0
	for x := len(input) - 1; x > -1; x-- {
		buf[inc] = input[x]
		inc++
	}
	return buf
}

func Uncompress(input []byte) ([]byte, error) {
	z, err := zlib.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(z)
}

func Compress(input []byte) []byte {
	b := etc.NewBuffer()
	z, err := zlib.NewWriterLevelDict(b, zlib.BestCompression, nil)
	if err != nil {
		panic(err)
	}
	// z := zlib.NewWriter(b)
	w, err := z.Write(input)
	if err != nil {
		panic(err)
	}
	if w != len(input) {
		panic(fmt.Errorf("%d/%d bytes compressed", w, len(input)))
	}
	if err := z.Close(); err != nil {
		panic(err)
	}
	return b.Bytes()
}
