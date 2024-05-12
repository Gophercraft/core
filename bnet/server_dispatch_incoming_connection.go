package bnet

import (
	"context"
	"net"

	"github.com/Gophercraft/log"
)

func (server *Server) dispatch_incoming_connection(tls_connection net.Conn) {
	var connection Connection
	connection.self_rpcs = make(map[uint32]*outbound_rpc)
	connection.ctx = context.WithValue(server.ctx, server_context_key, server)
	connection.tls_connection = tls_connection
	connection.unary_server_interceptor = server.unary_server_interceptor
	connection.unary_client_interceptor = server.unary_client_interceptor

	for {
		header, message, err := connection.read_message()
		if err != nil {
			log.Warn("bnet connection error", err)
			return
		}

		log.Println("Got message")

		go func() {
			if err := connection.handle_message(header, message); err != nil {
				log.Warn("bnet: error handling message:", err)
			}
		}()
	}
}
