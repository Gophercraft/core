package warden

import (
	"encoding/binary"
	"io"
)

type CryptoData struct {
	ClientKey []byte // _inputKey in TrinityCore
	ServerKey []byte // _outputKey in TrinityCore
}

type Reader struct {
	CryptoData
	io.Reader
}

type Writer struct {
	CryptoData
	io.Writer
}

func lLe(in io.Reader, value interface{}) error {
	return binary.Read(in, binary.LittleEndian, value)
}

func sLe(out io.Writer, value interface{}) error {
	return binary.Write(out, binary.LittleEndian, value)
}

func lString(in io.Reader) (string, error) {
	var data []byte
	for {
		b, err := lByte(in)
		if err != nil {
			return "", err
		}
		if b == 0x00 {
			break
		}
		data = append(data, b)
	}
	return string(data), nil
}

func lByte(i io.Reader) (c uint8, err error) {
	var cmd [1]byte
	_, err = i.Read(cmd[:])
	c = cmd[0]
	return
}

func sByte(out io.Writer, value byte) error {
	_, err := out.Write([]byte{value})
	return err
}

func readCommand(i io.Reader) (c Command, err error) {
	var b uint8
	b, err = lByte(i)
	c = Command(b)
	return
}

func sString(out io.Writer, value string) error {
	_, err := out.Write(append([]byte(value), 0))
	return err
}
