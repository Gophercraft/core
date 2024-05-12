package protocol

import (
	"context"
	"net"

	"github.com/Gophercraft/core/crypto/connection"
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type Connection struct {
	ctx          context.Context
	cancel_cause context.CancelCauseFunc

	tcp_conn    net.Conn // TCP connection
	build       version.Build
	server_mode bool

	cipher           connection.Cipher
	type_descriptor  *message.TypeDescriptor
	send_io          io_func
	recv_io          io_func
	send_queue       chan *message.Packet
	message_handlers []MessageHandlerFunc

	token_object any
}

// If the connection was established by a server, returns the server object.
func (conn *Connection) Server() *Server {
	value := conn.ctx.Value(connection_ctx_server)
	if value == nil {
		return nil
	}
	return value.(*Server)
}

func (conn *Connection) Build() version.Build {
	return conn.build
}

func (conn *Connection) SetTokenObject(object any) {
	conn.token_object = object
}

func (conn *Connection) TokenObject() any {
	return conn.token_object
}

func (conn *Connection) EnterEncryptedMode(session_key []byte) (err error) {
	conn.cipher, err = connection.NewCipher(conn.build, session_key, conn.server_mode)
	if err != nil {
		return
	}
	return
}
