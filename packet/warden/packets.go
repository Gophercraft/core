package warden

import (
	"io"

	"github.com/Gophercraft/core/vsn"
)

type Command uint8

const (
	CServerModuleUse Command = iota
	CServerModuleTransfer
	CServerRequestCheatChecks
	CServerModuleInitialize
	CServerRequestMemoryChecks
	CServerHashRequest
)

const (
	CClientModuleMissing Command = iota
	CClientModuleOK
	CClientCheatChecksResult
	CClientMemoryChecksResult
	CClientHashResult
	CClientModuleFailed
)

type Request interface {
	Command() Command
	Encode(build vsn.Build, out *Writer) error
	Decode(build vsn.Build, in *Reader) error
}

type ServerRequest Request

type ServerData struct {
	Requests []ServerRequest
}

func (sd *ServerData) Decode(build vsn.Build, in *Reader) error {
	for {
		command, err := readCommand(in)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		request, err := newServerRequest(command)
		if err != nil {
			return err
		}

		err = request.Decode(build, in)
		if err != nil {
			return err
		}

		sd.Requests = append(sd.Requests, request)
	}
}

func (sd *ServerData) Encode(build vsn.Build, out *Writer) error {
	for _, request := range sd.Requests {
		sByte(out, uint8(request.Command()))
		if err := request.Encode(build, out); err != nil {
			return err
		}
	}
	return nil
}

type ClientResult Request

type ClientData struct {
	Result ClientResult
}

type emptyRequest Command

func (er emptyRequest) Command() Command {
	return Command(er)
}

func (er emptyRequest) Decode(build vsn.Build, in *Reader) error {
	return nil
}

func (er emptyRequest) Encode(build vsn.Build, out *Writer) error {
	return nil
}

func (cd *ClientData) Decode(build vsn.Build, in *Reader) error {
	cmd, err := readCommand(in)
	if err != nil {
		return err
	}

	cr, err := newClientResult(cmd)
	if err != nil {
		return err
	}

	if err := cr.Decode(build, in); err != nil {
		return err
	}

	cd.Result = cr
	return nil
}

func (cd *ClientData) Encode(build vsn.Build, out *Writer) error {
	sLe(out, uint8(cd.Result.Command()))
	return cd.Result.Encode(build, out)
}
