package protocol

import "github.com/Gophercraft/core/game/protocol/message"

func (server *Server) dispatch_message(conn *Connection, packet *message.Packet) {
	for _, handler := range conn.message_handlers {
		handler(conn, packet)
	}
}
