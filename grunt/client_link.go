package grunt

import (
	"bytes"
	"net"
	"time"
)

type ClientLink struct {
	dialer     net.Dialer
	connection net.Conn
	buffer     bytes.Buffer
}

func (link *ClientLink) SetTimeout(timeout time.Duration) {
	link.dialer.Timeout = timeout
}

func (link *ClientLink) Establish(address string) (err error) {
	link.connection, err = link.dialer.Dial("tcp", address)
	if err != nil {
		return
	}
	link.connection.(*net.TCPConn).SetNoDelay(true)
	return
}

func (link *ClientLink) Read(b []byte) (n int, err error) {
	n, err = link.connection.Read(b)
	return
}

func (link *ClientLink) Write(b []byte) (n int, err error) {
	n, err = link.buffer.Write(b)
	return
}

func (link *ClientLink) Send() (err error) {
	_, err = link.connection.Write(link.buffer.Bytes())
	link.buffer.Reset()
	return
}

func (link *ClientLink) ReadMessageType() (message_type MessageType, err error) {
	message_type, err = ReadMessageType(link)
	return
}

func (link *ClientLink) WriteMessageType(message_type MessageType) (err error) {
	err = WriteMessageType(link, message_type)
	return
}
