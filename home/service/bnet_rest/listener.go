package bnet_rest

import (
	"crypto/tls"
	"net"

	"github.com/Gophercraft/log"
	"github.com/davecgh/go-spew/spew"
)

const intercept_debug = true

type listener struct {
	net.Listener
}

type conn struct {
	net.Conn
}

func listen(address string, tls_config *tls.Config) (l net.Listener, err error) {
	new_listener := new(listener)
	new_listener.Listener, err = tls.Listen("tcp", address, tls_config)
	if err != nil {
		return
	}
	l = new_listener
	return
}

func (listener *listener) Accept() (c net.Conn, err error) {
	conn := new(conn)
	conn.Conn, err = listener.Listener.Accept()
	if err != nil {
		return
	}

	if !intercept_debug {
		return conn.Conn, nil
	}

	c = conn
	return
}

func (c *conn) Write(b []byte) (n int, err error) {
	if intercept_debug {
		log.Println("[bnet_rest] http output", spew.Sdump(b))
	}
	n, err = c.Conn.Write(b)
	return
}
