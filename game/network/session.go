package network

import (
	"sync"

	"github.com/Gophercraft/core/game/protocol"
	"github.com/Gophercraft/core/version"
)

type Session struct {
	// the server that the session was created with
	server *Server
	// the connection currently associated with this connection
	protocol_connection *protocol.Connection

	//
	guard_service_contexts sync.Mutex
	service_contexts       map[ServiceID]any
}

func (session *Session) Build() version.Build {
	return session.protocol_connection.Build()
}

func (session *Session) EnterEncryptedMode(session_key []byte) (err error) {
	return session.protocol_connection.EnterEncryptedMode(session_key)
}
