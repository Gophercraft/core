package chunked

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestChunkedID(t *testing.T) {
	tg := ID("LMAO")

	var d bytes.Buffer
	binary.Write(&d, binary.LittleEndian, &tg)

	fmt.Println(spew.Sdump(d.Bytes()))

	var tg2 Tag
	binary.Read(&d, binary.LittleEndian, &tg2)

	fmt.Println(tg2.String())
}

func TestChunkedReader(t *testing.T) {
	file, err := os.Open("Work\\World\\Maps\\Azeroth\\Azeroth.wdt")
	if err != nil {
		panic(err)
	}

	reader := &Reader{file}

	for {
		id, data, err := reader.ReadChunk()
		fmt.Println(id)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(len(data))
	}
}
