package alphalist

import (
	"errors"
	"net"

	"github.com/Gophercraft/log"

	"github.com/Gophercraft/core/home/models"
)

type Server struct {
	Backend models.RealmList

	address  string
	listener net.Listener
}

func NewServer(address string, b models.RealmList) (s *Server) {
	s = new(Server)
	s.address = address
	s.Backend = b
	return
}

func (s *Server) Start() (err error) {
	s.listener, err = net.Listen("tcp", s.address)
	if err != nil {
		return
	}

	log.Println("Listening Alpha realmlist", s.address)

	go s.listen()
	return nil
}

func (s *Server) handle(c net.Conn) {
	log.Println("New Alpha list connection", c.RemoteAddr().String())
	sesh := &Session{
		Server: s,
		c:      c,
	}
	sesh.handle()
}

func (s *Server) listen() {
	for {
		c, err := s.listener.Accept()
		if err == nil {
			go s.handle(c)
		} else if errors.Is(err, net.ErrClosed) {
			return
		}
	}
}

func (s *Server) Shutdown() error {
	return s.listener.Close()
}
