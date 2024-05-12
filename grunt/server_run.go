package grunt

import (
	"errors"
	"net"
)

func (server *Server) Run() error {
	for {
		connection, err := server.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			continue
		}

		var connection_event Event
		connection_event.RemoteAddress = connection.RemoteAddr()
		connection_event.Type = NewConnection
		if err := server.service_provider.Check(&connection_event); err != nil {
			connection.Close()
			continue
		}

		go server.handle_incoming_connection(connection)
	}
}
