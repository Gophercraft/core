package movemap

import (
	"encoding/binary"
	"fmt"
	"io"
)

var Magic = [4]byte{'P', 'A', 'M', 'M'} // MMAP

const (
	MmapVersion = 4
	SizeOfGrids = 533.33333
)

type TileHeader struct {
	Magic          [4]byte
	DtVersion      uint32
	MoveMapVersion uint32
	Size           uint32
	UsesLiquids    bool
	Padding        [3]byte
}

func ReadTileHeader(stream io.Reader) (*TileHeader, error) {
	th := new(TileHeader)

	var err error
	err = binary.Read(stream, binary.LittleEndian, th)
	if err != nil {
		return nil, err
	}

	if th.Magic != Magic {
		return nil, fmt.Errorf("invalid magic %s", th.Magic)
	}

	return th, nil
}
