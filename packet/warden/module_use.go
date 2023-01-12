package warden

import (
	"encoding/binary"

	"github.com/Gophercraft/core/vsn"
)

type ServerModuleUse struct {
	ModuleID  [16]byte
	ModuleKey [16]byte
	Size      uint32
}

func (smu *ServerModuleUse) Command() Command {
	return CServerModuleUse
}

func (smu *ServerModuleUse) Decode(build vsn.Build, in *Reader) error {
	return binary.Read(in, binary.LittleEndian, smu)
}

func (smu *ServerModuleUse) Encode(build vsn.Build, out *Writer) error {
	return binary.Write(out, binary.LittleEndian, smu)
}
