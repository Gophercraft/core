package network

type SessionState uint8

const (
	Terminated SessionState = iota
	NewConnection
	EndSessionStates
)

const NumSessionStates = int(EndSessionStates)

func (session *Session) Terminate(err error) {
	session.protocol_connection.Terminate(err)
}

type NewSessionStateHandlerFunc func(newstate SessionState, session *Session)

func (server *Server) HandleNewSessionState(newstate SessionState, handlerfunc NewSessionStateHandlerFunc) {
	server.session_state_change_handlers[newstate] = append(server.session_state_change_handlers[newstate], handlerfunc)
}

func (server *Server) dispatch_session_state_change(newstate SessionState, session *Session) {
	for _, handler := range server.session_state_change_handlers[newstate] {
		handler(newstate, session)
	}
}
