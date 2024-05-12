package network

import "github.com/Gophercraft/core/game/protocol"

func NewServer(config *ServerConfig) (server *Server, err error) {
	// Make Server object
	server = new(Server)
	server.config = config
	server.services = make(map[ServiceID]Service)

	// Create protocol server
	server.protocol_server, err = protocol.NewServer(&server.config.Protocol)
	if err != nil {
		return
	}

	// Set up message handlers
	server.message_handlers.setup()
	return
}
