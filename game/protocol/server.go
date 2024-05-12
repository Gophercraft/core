package protocol

import (
	"net"
)

// A Server listens on a TCP adress:port, awaiting incoming connections.

type Server struct {
	config *ServerConfiguration

	tcp_listener                  net.Listener
	message_handlers              []MessageHandlerFunc
	state_change_handlers         [NumConnectionStates][]NewConnectionStateHandlerFunc
	preConnectSecurityCheckpoints []PreConnectSecurityCheckpointFunc
}

func NewServer(config *ServerConfiguration) (s *Server, err error) {
	if config == nil {
		err = ErrNoServerConfig
		return
	}

	s = new(Server)
	s.config = config

	s.HandleMessage(serverMessageHandler)

	return
}

func (s *Server) HandleNewConnectionState(newstate ConnectionState, handlerFunc NewConnectionStateHandlerFunc) {
	i := int(newstate)
	s.state_change_handlers[i] = append(s.state_change_handlers[i], handlerFunc)
}

func (s *Server) listen() (err error) {
	s.tcp_listener, err = net.Listen("tcp", s.config.Bind)
	return
}

func (s *Server) Run() (err error) {
	if err = s.listen(); err != nil {
		return
	}

	for {
		err = s.dispatch_incoming_connection()
		if err != nil {
			return err
		}
	}
}
