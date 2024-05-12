package protocol

import (
	"context"
	"net"

	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/game/protocol/message"
)

func (server *Server) handle_tcp_conn(tcp_conn net.Conn) {
	emulatedProtocol := server.config.Build

	conn := new(Connection)
	// Connection has context to signal separate read/write goroutines to die when connection is terminated
	conn.ctx, conn.cancel_cause = context.WithCancelCause(context.WithValue(context.TODO(), connection_ctx_server, server))
	conn.build = emulatedProtocol
	conn.tcp_conn = tcp_conn
	conn.cipher = crypto.DummyCipher{}

	// Begin handshaking process by performing protocol identity check
	//  Short strings between client and server confirming use for MMO game and not some other protocol
	// ( This does nothing before protocol 20886 )
	if err := conn.confirm_initial_protocol_identity(); err != nil {
		conn.logerr(err)
		conn.tcp_conn.Close()
		return
	}

	// Start the connection according to the server's protocol
	//
	// NOTE: I would ideally like to support ANY protocol, but these protocols are not set up to handle this.
	// The client gets bricked immediately if the server does not send the first message of the connection in the correct format.

	// Start with the fundamentals (packet framing)
	if err := conn.initialize_io(true, emulatedProtocol); err != nil {
		conn.logerr(err)
		conn.tcp_conn.Close()
		return
	}

	// Receive messages and dispatch to server message handlers
	recv := make(chan *message.Packet)
	go conn.recv_worker(recv)

	// Server is required to send an authentication challenge to the client at this point
	// if err := server.start_auth_challenge(conn); err != nil {
	// 	conn.Terminate(err)
	// 	return
	// }
	server.dispatch_connection_state_change(New, conn, nil)

	// Receive messages from receiver worker
	for {
		select {
		case <-conn.ctx.Done():
			break
		case msg := <-recv:
			server.dispatch_message(conn, msg)
		}
	}

	conn.tcp_conn.Close()

	server.dispatch_connection_state_change(Disconnected, conn, nil)
}
