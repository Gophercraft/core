package network

import (
	"github.com/Gophercraft/core/game/protocol"
)

func (server *Server) create_new_session(conn *protocol.Connection) {
	// New connection, construct session
	session := new(Session)
	session.protocol_connection = conn
	// associate this session with the connection.
	session.protocol_connection.SetTokenObject(session)
	session.server = server

	// Begin auth process
	server.dispatch_session_state_change(NewConnection, session)
}
