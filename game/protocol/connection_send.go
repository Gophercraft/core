package protocol

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
)

func (conn *Connection) send_queue_worker() {
	for {
		select {
		case <-conn.ctx.Done():
			return
		case msg := <-conn.send_queue:
			if err := conn.send(msg); err != nil {
				conn.cancel_cause(err)
			}
		}
	}
}

// Enqueues a message packet for transmission to the peer.
func (conn *Connection) Send(msg *message.Packet) error {
	if conn.ctx.Err() != nil {
		return fmt.Errorf("protocol: send on terminated connection")
	}
	conn.send_queue <- msg
	return nil
}
