package bnet

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"

	"github.com/Gophercraft/core/bnet/pb/bgs/protocol"
	"github.com/Gophercraft/core/bnet/rpc"
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/bnet/rpc/status"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/log"
	"google.golang.org/protobuf/proto"
)

type context_key uint8

const (
	server_context_key     context_key = 1
	connection_context_key context_key = 2
	client_context_key     context_key = 3
)

const (
	request_service  uint32 = 0
	response_service uint32 = 0xFE
)

type rpc_result struct {
	code codes.Code
	data []byte
}

type outbound_rpc struct {
	id     uint32
	result chan<- rpc_result
}

type Connection struct {
	done sync.WaitGroup

	guard_ctx sync.Mutex
	ctx       context.Context

	link

	tls_config *tls.Config

	unary_client_interceptor rpc.UnaryClientInterceptor
	unary_server_interceptor rpc.UnaryServerInterceptor

	guard_rpcs sync.Mutex
	token      uint32
	// rpcs initiated by this connection
	self_rpcs map[uint32]*outbound_rpc
}

// return the Server bound to this connection
func (connection *Connection) server() (server *Server) {
	value := connection.ctx.Value(server_context_key)
	if value == nil {
		return
	} else {
		server = value.(*Server)
	}
	return
}

// return the Client bound to this connection
func (connection *Connection) client() (client *Client) {
	value := connection.ctx.Value(client_context_key)
	if value == nil {
		return
	} else {
		client = value.(*Client)
	}
	return
}

// increment the token counter and return a unique token value
func (connection *Connection) next_token_sequence() (token uint32) {
	connection.guard_rpcs.Lock()
	connection.token++
	token = connection.token
	connection.guard_rpcs.Unlock()
	return
}

// safely set a value in this connection's Context, which propagates to service calls
func (connection *Connection) Set(key, value any) {
	connection.guard_ctx.Lock()
	connection.ctx = context.WithValue(connection.ctx, key, value)
	connection.guard_ctx.Unlock()
}

// returns the context of the connection safely
func (connection *Connection) get_context() (ctx context.Context) {
	connection.guard_ctx.Lock()
	ctx = connection.ctx
	connection.guard_ctx.Unlock()
	return
}

func (connection *Connection) get_registered_service(service_hash util.ServiceHash) (service *registered_service, err error) {
	server := connection.server()
	if server != nil {
		service = server.get_registered_service(service_hash)
	} else {
		client := connection.client()
		if client != nil {
			service = client.get_registered_service(service_hash)
		}
	}

	// todo: look for client-registered services (aka Listeners)

	if service == nil {
		err = fmt.Errorf("bnet: could not find service for 0x%08X", service_hash)
	}

	return
}

// handle a message sent in response to an RPC we made
func (connection *Connection) handle_response(header *protocol.Header, message []byte) (err error) {
	connection.guard_rpcs.Lock()
	outbound_rpc, found := connection.self_rpcs[header.GetToken()]
	if !found {
		err = fmt.Errorf("bnet: peer sent response to an RPC we didn't make (%d)", header.GetToken())
		return
	}
	// cleanup rpc
	delete(connection.self_rpcs, header.GetToken())
	connection.guard_rpcs.Unlock()

	// send result of RPC back to invoker
	var result rpc_result
	result.code = codes.Code(header.GetStatus())
	result.data = message

	outbound_rpc.result <- result
	close(outbound_rpc.result)

	return nil
}

func (connection *Connection) handle_request(header *protocol.Header, message []byte) (err error) {
	token := header.GetToken()
	service_hash := util.ServiceHash(header.GetServiceHash())
	method := rpc.Method(header.GetMethodId())

	// Attempt to find service with this hash
	var service *registered_service
	service, err = connection.get_registered_service(service_hash)
	if err != nil {
		return
	}

	// Attempt to find the particular method within this service
	for _, service_method := range service.descriptor.Methods {
		if service_method.MethodId == method {
			// Method was found

			// arguments decoder
			dec := func(m any) error {
				return proto.Unmarshal(message, m.(proto.Message))
			}

			// invoke the service/method handler
			var response_status codes.Code = codes.ERROR_OK
			var reply any
			ctx := context.WithValue(connection.get_context(), connection_context_key, connection)
			reply, err = service_method.Handler(service.service, ctx, dec, connection.unary_server_interceptor)
			if err != nil {
				status_error, is_status_error := err.(*status.CodedError)
				if is_status_error {
					response_status = status_error.Code
				} else {
					// generic error for unknown error types
					response_status = codes.ERROR_INTERNAL
				}
			}
			// encode the
			var response_data []byte
			if reply != nil {
				response_data, err = proto.Marshal(reply.(proto.Message))
				if err != nil {
					return
				}
			}

			return connection.send_response(token, response_status, response_data)
		}
	}

	err = fmt.Errorf("bnet: peer wants a method (%d) that this service (%s) doesn't have", method, service.descriptor.ServiceName)
	return
}

func (connection *Connection) handle_message(header *protocol.Header, message []byte) (err error) {
	switch header.GetServiceId() {
	case response_service:
		return connection.handle_response(header, message)
	case request_service:
		return connection.handle_request(header, message)
	default:
		return fmt.Errorf("bnet: unknown service id %d", header.GetServiceId())
	}
}

func (connection *Connection) Invoke(ctx context.Context, full_method string, service util.ServiceHash, method rpc.Method, listen_for_response bool, args, reply any, opts ...rpc.CallOption) error {
	if connection.unary_client_interceptor != nil {
		return connection.unary_client_interceptor(ctx, full_method, service, method, listen_for_response, args, reply, connection, invoke, opts...)
	}

	return invoke(ctx, service, method, listen_for_response, args, reply, connection, opts...)
}

func invoke(ctx context.Context, service util.ServiceHash, method rpc.Method, listen_for_response bool, args, reply any, cc rpc.ClientConnectionInterface, opts ...rpc.CallOption) error {
	connection := cc.(*Connection)
	// marshal rpc arguments
	message, err := proto.Marshal(args.(proto.Message))
	if err != nil {
		return err
	}

	// get new RPC sequence
	token := connection.next_token_sequence()

	// Many RPCs have a response
	if listen_for_response {
		// create rpc
		result_channel := make(chan rpc_result)
		self_rpc := new(outbound_rpc)
		self_rpc.id = token
		self_rpc.result = result_channel

		// register rpc
		connection.guard_rpcs.Lock()
		connection.self_rpcs[token] = self_rpc
		connection.guard_rpcs.Unlock()
		// send request
		if err = connection.send_request(token, service, method, message); err != nil {
			// error happened, clean up
			close(self_rpc.result)
			connection.guard_rpcs.Lock()
			delete(connection.self_rpcs, token)
			connection.guard_rpcs.Unlock()
			return err
		}

		log.Println("waiting for reply token", token)

		// receive reply from message dispatcher goroutine
		result := <-result_channel

		log.Println("got reply token", token)

		// decode reply
		if err = proto.Unmarshal(result.data, reply.(proto.Message)); err != nil {
			return err
		}

		// success! (message handler cleans up RPC for us)
	} else {
		// If this RPC has no response, just send the RPC and return immediately
		// send request with no reply
		if err = connection.send_request(token, service, method, message); err != nil {
			// error happened, clean up
			return err
		}
	}

	return nil
}

func (connection *Connection) send_response(token uint32, code codes.Code, data []byte) (err error) {
	var header protocol.Header
	// Set serviceID to response service (0xFE)
	util.Set(&header.ServiceId, response_service)
	util.Set(&header.Status, uint32(code))
	util.Set(&header.Token, token)

	if data != nil {
		util.Set(&header.Size, uint32(len(data)))
	}

	err = connection.write_message(&header, data)

	return
}

// send a rpc call to a connected peer. whether data is nil or
func (connection *Connection) send_request(token uint32, service util.ServiceHash, method rpc.Method, data []byte) (err error) {
	var header protocol.Header
	// Set serviceID to request service (0)
	util.Set(&header.ServiceId, request_service)
	util.Set(&header.ServiceHash, uint32(service))
	util.Set(&header.MethodId, uint32(method))
	util.Set(&header.Token, token)

	if data != nil {
		util.Set(&header.Size, uint32(len(data)))
	}

	err = connection.write_message(&header, data)

	return
}

// Returns the the Connection associated with an RPC
func GetConnection(ctx context.Context) *Connection {
	value := ctx.Value(connection_context_key)
	if value == nil {
		return nil
	}
	return value.(*Connection)
}

func (connection *Connection) Close() (err error) {
	connection.request_close_after_next_write()
	return
}

func (connection *Connection) Wait() {
	connection.done.Wait()
}
