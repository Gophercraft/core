package grunt

import (
	"net"

	"github.com/Gophercraft/log"
	"github.com/davecgh/go-spew/spew"
)

type message_handler func(session *Session, message_type MessageType) error

var (
	session_message_handlers [MaxMessageType]message_handler
)

func init() {
	// setup message handlers
	session_message_handlers[LogonChallenge] = func(session *Session, message_type MessageType) (err error) {
		var message AuthLogonChallenge_Client
		if err = ReadAuthLogonChallenge_Client(session.connection, &message); err != nil {
			return
		}
		log.Println(spew.Sdump(message))
		if err = session.handle_auth_logon_challenge(&message); err != nil {
			return
		}
		return
	}

	session_message_handlers[LogonProof] = func(session *Session, message_type MessageType) (err error) {
		var message AuthLogonProof_Client
		if err = ReadAuthLogonProof_Client(session.connection, &message); err != nil {
			return
		}
		log.Println(spew.Sdump(message))
		if err = session.handle_auth_logon_proof(&message); err != nil {
			return
		}
		return
	}

	session_message_handlers[ReconnectChallenge] = func(session *Session, message_type MessageType) (err error) {
		var message AuthLogonChallenge_Client
		if err = ReadAuthLogonChallenge_Client(session.connection, &message); err != nil {
			return
		}
		if err = session.handle_auth_reconnect_challenge(&message); err != nil {
			return
		}
		return
	}

	session_message_handlers[ReconnectProof] = func(session *Session, message_type MessageType) (err error) {
		var message AuthReconnectProof_Client
		if err = ReadAuthReconnectProof_Client(session.connection, &message); err != nil {
			return
		}
		if err = session.handle_auth_reconnect_proof(&message); err != nil {
			return
		}
		return
	}

	session_message_handlers[RealmList] = func(session *Session, message_type MessageType) (err error) {
		var message RealmList_Client
		if err = ReadRealmList_Client(session.connection, &message); err != nil {
			return
		}
		if err = session.handle_realm_list(&message); err != nil {
			return
		}
		return
	}
}

func (server *Server) handle_incoming_connection(connection net.Conn) {
	session := new(Session)
	session.server = server
	session.connection = connection

	log.Println("New authserver connection from", session.connection.RemoteAddr())

	var (
		err          error
		message_type MessageType
	)

	// every iteration of this loop reads an opcode from the TCP socket, and associated data.
	for {
		message_type, err = ReadMessageType(session.connection)
		if err != nil {
			log.Warn("could not read message type: ", err)
			break
		}

		log.Println(session.connection.RemoteAddr(), message_type)

		if message_type >= MaxMessageType {
			log.Warn("remote client sent unknown type", message_type)
			break
		}

		message_handler := session_message_handlers[message_type]
		if message_handler == nil {
			log.Warn("remote client sent unhandled message type", message_type)
			continue
		}

		if err := message_handler(session, message_type); err != nil {
			log.Warn("fatal error in message handler: ", err)
			break
		}
	}

	connection.Close()
}
