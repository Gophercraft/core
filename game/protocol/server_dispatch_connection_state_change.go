package protocol

func (server *Server) dispatch_connection_state_change(newState ConnectionState, conn *Connection, err error) {
	handlers := server.state_change_handlers[newState]

	for _, handler := range handlers {
		handler(newState, conn, err)
	}
}
