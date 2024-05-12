package alphalist

import (
	"errors"
	"net"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/log"
)

var log_category = log.Category("service/alphalist")

type Service struct {
	list_provider ServiceProvider
	address       string
	listener      net.Listener
}

func New(address string, provider ServiceProvider) (s *Service) {
	s = new(Service)
	s.address = address
	s.list_provider = provider
	return
}

func (s *Service) Start() (err error) {
	s.listener, err = net.Listen("tcp", s.address)
	if err != nil {
		return
	}

	log.Println("Listening Alpha realmlist", s.address)

	go s.run()
	return nil
}

func (s *Service) handle(connection net.Conn) {
	log.Println("New Alpha list connection", connection.RemoteAddr().String())
	var session Session
	session.connection = connection
	session.handle()
	session.connection.Close()
}

func (s *Service) run() {
	for {
		c, err := s.listener.Accept()
		if err == nil {
			go s.handle(c)
		} else if errors.Is(err, net.ErrClosed) {
			return
		}
	}
}

func (s *Service) Stop() (err error) {
	return s.listener.Close()
}

func (s *Service) ID() config.HomeServiceID {
	return config.OldRealmlistService
}
