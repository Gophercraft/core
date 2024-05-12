package network

import (
	"github.com/Gophercraft/core/game/protocol"
	"github.com/Gophercraft/core/game/protocol/message"
)

func (server *Server) set_protocol_handlers() error {
	// Handle messages from remote connection
	server.protocol_server.HandleMessage(func(conn *protocol.Connection, msg *message.Packet) {
		token := conn.TokenObject()
		if token != nil {
			session := token.(*Session)
			session.dispatch_message(msg)
		}
	})

	// Handle new remote connection
	server.protocol_server.HandleNewConnectionState(protocol.New, func(st protocol.ConnectionState, conn *protocol.Connection, err error) {
		server.create_new_session(conn)
	})

	// Handle remote connection reset (fail to correct state)
	server.protocol_server.HandleNewConnectionState(protocol.Disconnected, func(st protocol.ConnectionState, conn *protocol.Connection, err error) {
		token := conn.TokenObject()
		if token != nil {
			session := token.(*Session)
			session.Terminate(err)
		}
	})
}
