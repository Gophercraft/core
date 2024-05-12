package warden

import (
	"fmt"
	"io"

	"github.com/Gophercraft/core/version"
)

type MemoryCheck struct {
	MemType uint8
	Module  string
	Offset  uint32
	Data    []byte
}

type ServerRequestMemoryChecks struct {
	Checks []MemoryCheck
}

func (srmc *ServerRequestMemoryChecks) Command() Command {
	return CServerRequestMemoryChecks
}

func (srmc *ServerRequestMemoryChecks) Encode(build version.Build, out *Writer) error {
	for i, check := range srmc.Checks {
		sByte(out, check.MemType)
		sByte(out, uint8(i))
		sString(out, check.Module)
		sLe(out, check.Offset)
		sByte(out, uint8(len(check.Data)))
		out.Write(check.Data[:])
	}
	return nil
}

func (srmc *ServerRequestMemoryChecks) Decode(build version.Build, in *Reader) error {
	for i := 0; ; i++ {
		var err error
		var memCheck MemoryCheck
		memCheck.MemType, err = lByte(in)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		index, err := lByte(in)
		if err != nil {
			return err
		}
		if int(index) != i {
			return fmt.Errorf(("packet/warden: id-index mismatch"))
		}
		memCheck.Module, err = lString(in)
		if err != nil {
			return err
		}
		lLe(in, &memCheck.Offset)
		size, err := lByte(in)
		if err != nil {
			return err
		}
		memCheck.Data = make([]byte, size)
		if _, err := in.Read(memCheck.Data[:]); err != nil {
			return err
		}
	}
}
