package warden

import (
	"encoding/binary"

	"github.com/Gophercraft/core/version"
)

type ServerModuleTransfer struct {
	Data []byte
}

func (smt *ServerModuleTransfer) Command() Command {
	return CServerModuleTransfer
}

func (smt *ServerModuleTransfer) Decode(build version.Build, in *Reader) (err error) {
	var size uint16
	err = binary.Read(in, binary.LittleEndian, &size)
	if err != nil {
		return
	}
	smt.Data = make([]byte, size)
	if _, err = in.Read(smt.Data); err != nil {
		return
	}
	return
}

func (smt *ServerModuleTransfer) Encode(build version.Build, out *Writer) (err error) {
	var size uint16 = uint16(len(smt.Data))
	err = binary.Write(out, binary.LittleEndian, &size)
	if err != nil {
		return
	}
	_, err = out.Write(smt.Data)
	return
}
