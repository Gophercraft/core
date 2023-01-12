package auth

import (
	"io"
	"net"

	"github.com/superp00t/etc"
)

type Connection struct {
	net.Conn
}

func wrapConn(c net.Conn) *Connection {
	return &Connection{c}
}

func (conn *Connection) readAuthType() (a AuthType, err error) {
	var at [1]byte
	_, err = conn.Read(at[:])
	a = AuthType(at[0])
	return
}

func (conn *Connection) RecvBuffer(size int) (*etc.Buffer, error) {
	data := make([]byte, size)
	_, err := io.ReadFull(conn, data)
	if err != nil {
		return nil, err
	}
	return etc.OfBytes(data), nil
}

func (conn *Connection) SendBuffer(e *etc.Buffer) error {
	_, err := conn.Write(e.Bytes())
	return err
}

func Dial(address string) (*Connection, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return wrapConn(conn), nil
}
