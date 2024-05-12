package game_utilities

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"

	"github.com/Gophercraft/core/bnet/pb/bgs/protocol"
	"github.com/Gophercraft/core/bnet/util"
)

func split_blob_parameter(blob []byte) (key []byte, value []byte, err error) {
	var found bool
	key, value, found = bytes.Cut(blob, []byte(":"))
	if !found {
		err = fmt.Errorf("home/provider/bnet/rpc/game_utilities: split_blob_parameter: could not find separator :")
		return
	}

	return
}

func append_blob_value_to_response(response *[]*protocol.Attribute, name string, value []byte) {
	attribute := new(protocol.Attribute)
	attribute.Value = new(protocol.Variant)
	util.Set(&attribute.Name, name)
	attribute.Value.BlobValue = value

	*response = append(*response, attribute)
}

func append_compressed_blob_value_to_response(response *[]*protocol.Attribute, name string, value []byte) {
	buffer := new(bytes.Buffer)
	var length_prefix [4]byte
	binary.LittleEndian.PutUint32(length_prefix[:], uint32(len(value)))
	buffer.Write(length_prefix[:])
	writer := zlib.NewWriter(buffer)
	writer.Write(value)
	writer.Close()

	append_blob_value_to_response(response, name, buffer.Bytes())
}

func remove_null_byte(b []byte) []byte {
	b, _, _ = bytes.Cut(b, []byte{0})
	return b
}
