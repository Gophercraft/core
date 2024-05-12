package network

import (
	"github.com/Gophercraft/core/game/protocol"
)

func (server *Server) Run() (err error) {
	// Create protocol server (TCP/WebSocket)
	server.protocol_server, err = protocol.NewServer(&server.config.Protocol)
	if err != nil {
		return err
	}

	if err = server.set_protocol_handlers(); err != nil {
		return err
	}

	return server.protocol_server.Run()
}
