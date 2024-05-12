package bls

import "bytes"

func Read(b []byte) (*Shader, error) {
	shader := new(Shader)
	bytesReader := bytes.NewReader(b)
	shaderReader := newReader(bytesReader, shader)
	err := shaderReader.read()
	if err != nil {
		return nil, err
	}
	return shader, nil
}
