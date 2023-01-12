package warden

import (
	"bytes"

	"github.com/superp00t/etc"
	"github.com/Gophercraft/core/crypto/warden"
	"github.com/Gophercraft/core/vsn"
)

type ServerModuleInitialize struct {
	SizeFlags     uint16
	Type          uint8
	StringLibrary uint8
	Function      []uint64
	FunctionSet   uint8
}

func (smu *ServerModuleInitialize) Command() Command {
	return CServerModuleInitialize
}

func (smi *ServerModuleInitialize) Decode(build vsn.Build, in *Reader) error {
	var size uint16
	var checksum uint32
	if err := lLe(in, &size); err != nil {
		return err
	}
	if err := lLe(in, &checksum); err != nil {
		return err
	}
	var data = make([]byte, size)
	if _, err := in.Read(data[:]); err != nil {
		return err
	}
	if checksum != warden.Checksum(data) {
		return ErrCheckBadChecksum
	}
	var reader = bytes.NewReader(data)
	lLe(reader, &smi.SizeFlags)

	var numAddress int

	// This isn't correct but info is scarce
	switch smi.SizeFlags {
	case 1:
		lLe(reader, &smi.Type)
		lLe(reader, &smi.StringLibrary)
		numAddress = (int(size) - 4) / 4
	case 4, 257:
		lLe(reader, &smi.StringLibrary)
		numAddress = (int(size) - 4) / 4
	default:
		panic(smi.SizeFlags)
	}

	smi.Function = make([]uint64, numAddress)
	for i := 0; i < numAddress; i++ {
		var address uint32
		lLe(reader, &address)
		smi.Function[i] = uint64(address)
	}

	switch smi.SizeFlags {
	case 4, 257:
		lLe(reader, &smi.FunctionSet)
	}

	return nil
}

func (smi *ServerModuleInitialize) Encode(build vsn.Build, out *Writer) error {
	data := etc.NewBuffer()
	data.WriteUint16(smi.SizeFlags)
	switch smi.SizeFlags {
	case 1:
		data.WriteByte(smi.Type)
		data.WriteByte(smi.StringLibrary)
	case 4, 257:
		data.WriteByte(smi.StringLibrary)
	}
	for _, fn := range smi.Function {
		sLe(out, uint32(fn))
	}
	switch smi.SizeFlags {
	case 4, 257:
		sLe(out, uint8(smi.FunctionSet))
	}
	sLe(out, uint16(data.Len()))
	sLe(out, warden.Checksum(data.Bytes()))
	out.Write(data.Bytes())
	return nil
}
