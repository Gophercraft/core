package network

import (
	"reflect"

	"github.com/Gophercraft/core/game/protocol/message"
)

type message_handler struct {
	//

	// if true, fn is called in its own goroutine
	spawn_goroutine bool
	// the callback to the message handler
	fn MessageHandlerFunc
}

// Describes all the message handlers in a function
type ServiceDescriptor struct {
	message_handlers []message_handler
}

type MessageHandlerFunc func(session *Session, message *message.Packet) error

// Returns a message handler that wraps a handler for a decoded type M.
// M must both be message.Decodable and a pointer to a struct.
func DecodeHandler[M message.Decodable](handler_func func(session *Session, decoded_message M) error) MessageHandlerFunc {
	message_type := reflect.TypeFor[M]().Elem()

	return func(session *Session, message_packet *message.Packet) (err error) {
		decodable_message := reflect.New(message_type).Interface().(M)

		err = decodable_message.Decode(session.Build(), message_packet)
		if err != nil {
			return
		}

		return handler_func(session, decodable_message)
	}
}

type MessageHandlerOption func(option *message_handler)

func WithGoroutine() MessageHandlerOption {
	return func(option *message_handler) {
		option.spawn_goroutine = true
	}
}

func NewServiceDescriptor() (desc *ServiceDescriptor) {
	desc = new(ServiceDescriptor)
	return
}

func (service_descriptor *ServiceDescriptor) Handle(message_type message.Type, handler_func MessageHandlerFunc, options ...MessageHandlerOption) {
	service_descriptor.message_handlers = append(service_descriptor.message_handlers)
}
