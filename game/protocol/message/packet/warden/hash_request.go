package warden

import "github.com/Gophercraft/core/version"

type ServerHashRequest struct {
	Seed []byte
}

func (shr *ServerHashRequest) Command() Command {
	return CServerHashRequest
}

func (shr *ServerHashRequest) Encode(build version.Build, out *Writer) error {
	out.Write(shr.Seed[:16])
	return nil
}

func (shr *ServerHashRequest) Decode(build version.Build, in *Reader) error {
	shr.Seed = make([]byte, 16)
	_, err := in.Read(shr.Seed[:])
	return err
}

type ClientHashResult struct {
	Response []byte
}

func (chr *ClientHashResult) Command() Command {
	return CClientHashResult
}

func (chr *ClientHashResult) Encode(build version.Build, out *Writer) error {
	out.Write(chr.Response[:20])
	return nil
}

func (chr *ClientHashResult) Decode(build version.Build, in *Reader) error {
	chr.Response = make([]byte, 20)
	_, err := in.Read(chr.Response)
	return err
}
