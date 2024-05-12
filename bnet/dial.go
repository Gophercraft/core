package bnet

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/Gophercraft/core/bnet/rpc"
)

type DialOption func(connection *Connection)

func DialWithUnaryClientInterceptor(interceptor rpc.UnaryClientInterceptor) DialOption {
	return func(connection *Connection) {
		connection.unary_client_interceptor = interceptor
	}
}

func DialWithUnaryServerInterceptor(interceptor rpc.UnaryServerInterceptor) DialOption {
	return func(connection *Connection) {
		connection.unary_server_interceptor = interceptor
	}
}

func DialWithTLSConfig(tls_config *tls.Config) DialOption {
	return func(connection *Connection) {
		connection.tls_config = tls_config
	}
}

func DialContext(ctx context.Context, target string, options ...DialOption) (connection *Connection, err error) {
	connection = new(Connection)
	connection.self_rpcs = make(map[uint32]*outbound_rpc)
	for _, option := range options {
		option(connection)
	}
	fmt.Println("Dialing...", connection.tls_config.InsecureSkipVerify)
	connection.ctx = ctx
	connection.link.tls_connection, err = tls.Dial("tcp", target, connection.tls_config)
	if err != nil {
		return
	}

	connection.done.Add(1)

	go func() {
		for {
			header, message, err := connection.read_message()
			if err != nil {
				fmt.Println("Error reading message", err)
				break
			}

			go func() {
				if err := connection.handle_message(header, message); err != nil {
					fmt.Println("bnet: error handling message:", err)
				}
			}()
		}

		connection.done.Done()
	}()

	return
}
