// package multiproto allows multiple protocols to be interpreted on a single socket, by receiving a small segment of bytes
package multiproto

import (
	"bytes"
	"errors"
	"net"

	"github.com/Gophercraft/log"
)

type Router struct {
	addr      *net.TCPAddr
	l         *net.TCPListener
	listeners []*Listener
}

func NewRouter(address string) (r *Router, err error) {
	r = new(Router)
	r.addr, err = net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return
	}

	r.l, err = net.ListenTCP("tcp", r.addr)
	if err != nil {
		return
	}

	return
}

func (r *Router) SelectListener(hello []byte) *Listener {
	l := r.newListener()
	l.ProtocolHello = hello
	return l
}

func (r *Router) FallbackListener() *Listener {
	l := r.newListener()
	l.Fallback = true
	return l
}

func (r *Router) handle(c net.Conn) {
	bc := NewBufferedConn(c)
	var maxmagic int
	for _, s := range r.listeners {
		if len(s.ProtocolHello) > maxmagic {
			maxmagic = len(s.ProtocolHello)
		}
	}

	pHello, err := bc.Peek(maxmagic)
	if err != nil {
		log.Warn("Peeking error", err)
		return
	}

	// decide what protocol it looks like
	var selector *Listener

	// look for a protocol handler with a hello like what we want
	for i, s := range r.listeners {
		if !s.Fallback && bytes.Equal(pHello[:len(s.ProtocolHello)], s.ProtocolHello) {
			selector = r.listeners[i]
			break
		}
	}

	// If not, choose a fallback protocol handler
	if selector == nil {
		for i := range r.listeners {
			s := r.listeners[i]
			if s.Fallback {
				selector = s
				break
			}
		}
	}

	if selector != nil {
		selector.sendConn(bc)
	}
}

func (r *Router) Close() error {
	for _, l := range r.listeners {
		l.Close()
	}
	return nil
}

func (r *Router) Serve() error {
	for {
		conn, err := r.l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			continue
		}

		go r.handle(conn)
	}
}
