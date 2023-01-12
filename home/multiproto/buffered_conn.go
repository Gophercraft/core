package multiproto

import (
	"bufio"
	"net"
)

type BufferedConn struct {
	net.Conn

	Reader *bufio.Reader
}

func NewBufferedConn(c net.Conn) *BufferedConn {
	return &BufferedConn{
		Conn:   c,
		Reader: bufio.NewReader(c),
	}
}

func (bc *BufferedConn) Peek(n int) ([]byte, error) {
	return bc.Reader.Peek(n)
}

func (bc *BufferedConn) Read(p []byte) (int, error) {
	return bc.Reader.Read(p)
}
