package network

// import (
// 	"fmt"
// 	"reflect"

// 	"github.com/Gophercraft/core/game/protocol/message"
// )

// // types for discerning reflected message handlers
// var (
// 	message_command_type   = reflect.TypeOf(message.Type(0))
// 	decodable_message_type = reflect.TypeOf((*message.Decodable)(nil)).Elem()
// 	message_packet_type    = reflect.TypeOf((*message.Packet)(nil))
// 	session_type           = reflect.TypeOf((*Session)(nil))
// 	error_type             = reflect.TypeOf((*error)(nil)).Elem()
// )

// // this is the basic form of both types of handler
// type packet_handler_func func(session *Session, msg *message.Packet) error

// // container for packet lookup table
// type message_handlers struct {
// 	table map[message.Type]packet_handler_func
// }

// // allocate lookup table
// func (handlers *message_handlers) setup() {
// 	handlers.table = make(map[message.Type]packet_handler_func)
// }

// // set any kind of acceptable message handler
// func (handlers *message_handlers) set_message_handler(message_type message.Type, message_handler_func any) {
// 	// reflect on supplied handler value
// 	message_handler_func_value := reflect.ValueOf(message_handler_func)
// 	message_handler_func_type := message_handler_func_value.Type()

// 	// must be `func(*Session, [message type]) error`
// 	if message_handler_func_type.NumIn() != 2 {
// 		panic(fmt.Errorf("invalid message handler function argument count: %s", message_handler_func_type))
// 	}
// 	if message_handler_func_type.NumOut() != 1 {
// 		panic(fmt.Errorf("invalid message handler function return count: %s", message_handler_func_type))
// 	}
// 	if message_handler_func_type.In(0) != session_type {
// 		panic(fmt.Errorf("message handler function first argument must be *network.Session: %s", message_handler_func_type))
// 	}
// 	if message_handler_func_type.Out(0) != error_type {
// 		panic(fmt.Errorf("message handler function needs to return an error type: %s", message_handler_func_type))
// 	}

// 	// Is this `func(*Session, *message.Packet) error` ?
// 	if message_handler_func_type.In(1) == message_packet_type {
// 		handlers.set_packet_handler(message_type, message_handler_func.(packet_handler_func))
// 		return
// 	}

// 	// Is this `func(*Session, [type that satisfies message.Decodable interface]) error` ?
// 	if message_handler_func_type.In(1).Implements(decodable_message_type) {
// 		handlers.set_decodable_handler(message_type, message_handler_func)
// 		return
// 	}

// 	panic(fmt.Errorf("message handler function has unusable message type: %s", message_handler_func_type.In(1)))
// }

// // Set a packet handler only
// func (handlers *message_handlers) set_packet_handler(message_type message.Type, handler_func packet_handler_func) {
// 	handlers.table[message_type] = handler_func
// }

// // Set a decodable handler only
// func (handlers *message_handlers) set_decodable_handler(message_type message.Type, decoded_message_handler any) {
// 	handler_func := reflect.ValueOf(decoded_message_handler)
// 	handler_type := handler_func.Type()

// 	// Handler must be function
// 	if handler_type.Kind() != reflect.Func {
// 		panic("decodable handler must be function")
// 	}

// 	num_handler_arguments := handler_type.NumIn()
// 	// Handler must be two arguments
// 	if num_handler_arguments != 2 {
// 		panic("decodable handler must have two arguments (func(message.Type, message.Decodable))")
// 	}

// 	if handler_type.NumOut() != 1 {
// 		panic("decodable handler should return error")
// 	}

// 	// Ensure correct types
// 	command_parameter := handler_type.In(0)
// 	decodable_parameter := handler_type.In(1)
// 	if command_parameter != message_command_type {
// 		panic("decodable handler must handle the message type")
// 	}
// 	if !decodable_parameter.Implements(decodable_message_type) {
// 		panic(fmt.Sprintf("type %s does not implement %s", decodable_parameter, decodable_message_type))
// 	}
// 	error_parameter := handler_type.Out(0)
// 	if error_parameter != error_type {
// 		panic("function should return error")
// 	}

// 	// Create function to decode the message and call the handler
// 	handlers.set_packet_handler(message_type, func(session *Session, msg *message.Packet) error {
// 		decodable_message_value := reflect.New(decodable_parameter)
// 		decodable_message := decodable_message_value.Interface().(message.Decodable)

// 		if err := decodable_message.Decode(session.Build(), msg); err != nil {
// 			return err
// 		}

// 		err_value := handler_func.Call([]reflect.Value{
// 			reflect.ValueOf(message_type),
// 			reflect.ValueOf(msg),
// 		})

// 		return err_value[0].Interface().(error)
// 	})
// }
