package protocol

import "github.com/Gophercraft/core/game/protocol/message"

func (conn *Connection) recv_worker(recv chan<- *message.Packet) {
	for {
		packet, err := conn.recv()
		if err != nil {
			conn.cancel_cause(err)
			// Sender always closes channel
			close(recv)
			return
		}

		recv <- packet
	}
}
