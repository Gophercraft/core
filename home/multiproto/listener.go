package multiproto

import (
	"net"
)

type Listener struct {
	Fallback      bool
	ProtocolHello []byte

	addr net.Addr

	conns  chan net.Conn
	close  chan bool
	closed bool
}

func (l *Listener) sendConn(c net.Conn) {
	l.conns <- c
}

func (l *Listener) setupConn(c net.Conn) (net.Conn, error) {
	return c, nil
}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.conns:
		return l.setupConn(conn)
	case <-l.close:
		l.closed = true
		return nil, net.ErrClosed
	}
}

func (l *Listener) Addr() net.Addr {
	return l.addr
}

func (l *Listener) Close() error {
	l.close <- true
	return nil
}

func (r *Router) newListener() *Listener {
	l := new(Listener)
	l.conns = make(chan net.Conn, 16)
	l.close = make(chan bool)
	l.addr = r.addr
	r.listeners = append(r.listeners, l)
	return l
}
