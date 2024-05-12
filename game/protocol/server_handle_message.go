package protocol

import "github.com/Gophercraft/core/game/protocol/message"

type MessageHandlerFunc func(conn *Connection, msg *message.Packet)

func (server *Server) HandleMessage(handler_func MessageHandlerFunc) {
	server.message_handlers = append(server.message_handlers, handler_func)
}
