package dbdefs

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"

	"github.com/Gophercraft/core/format/dbc/dbd"
	"github.com/cybriq/gotiny"
)

var (
	All map[string]*dbd.Definition
)

func init() {
	reader := bytes.NewReader(packedDefinitions)

	zr, err := zlib.NewReader(reader)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(zr)
	if err != nil {
		panic(err)
	}

	gotiny.Unmarshal(b, &All)
}
