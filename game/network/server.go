package network

import (
	"github.com/Gophercraft/core/game/protocol"
)

type ServerConfig struct {
	Protocol protocol.ServerConfiguration
}

type Server struct {
	config                        *ServerConfig
	protocol_server               *protocol.Server
	message_handlers              message_handlers
	services                      map[ServiceID]Service
	session_state_change_handlers [NumSessionStates][]NewSessionStateHandlerFunc
	error                         chan error
}
