package network

import "github.com/Gophercraft/core/game/protocol/message"

func (session *Session) dispatch_message(msg *message.Packet) {
	// Look up handler for type
}

func (session *Session) Encode(encodable message.Encodable) error {
	// Encode packet data
	packet := new(message.Packet)
	if err := encodable.Encode(session.Build(), packet); err != nil {
		return err
	}

	// Send data to connection
	return session.Send(packet)
}

func (session *Session) Send(packet *message.Packet) error {
	return session.protocol_connection.Send(packet)
}
