package protocol

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

func init() {
	funcs := io_functions{}

	funcs[client_send] = io_client_send_0_12340
	funcs[client_recv] = io_client_recv_0_12340

	funcs[server_send] = io_server_send_0_12340
	funcs[server_recv] = io_server_recv_0_12340

	io_protocols[version.Range(0, 12340)] = funcs
}

/*******************
* Client -> server *
********************/

func io_client_send_0_12340(conn *Connection, packet *message.Packet) error {
	// Convert high-level type to the low-level protocol level representation.
	resolvedOpcode, err := conn.type_descriptor.LookupOpcode(packet.Type)
	if err != nil {
		return err
	}

	// Start building message envelope
	message := make([]byte, packet.Len()+2+4)
	// Opcodes are included within the packet size. (network byte order)
	binary.BigEndian.PutUint16(message[0:2], uint16(packet.Len()+4))
	binary.LittleEndian.PutUint32(message[2:6], resolvedOpcode)

	copy(message[6:], packet.Bytes())

	if err := conn.cipher.Encrypt(message[:6], nil); err != nil {
		return err
	}

	_, err = conn.tcp_conn.Write(message)
	return err
}

func io_client_recv_0_12340(conn *Connection, packet *message.Packet) error {
	// Receive size prefix. May or may not be encrypted
	var sizePrefix [2]byte
	if _, err := io.ReadFull(conn.tcp_conn, sizePrefix[:]); err != nil {
		return err
	}
	if err := conn.cipher.Decrypt(sizePrefix[:], nil); err != nil {
		return err
	}
	size := int(binary.BigEndian.Uint16(sizePrefix[:]))
	if size < 2 {
		return fmt.Errorf("game/protocol: client received server message without opcode")
	}

	// Receive message data. The opcode prefix may be encrypted,
	// while the content of the message is always in plaintext.
	message := make([]byte, size)
	if _, err := io.ReadFull(conn.tcp_conn, message); err != nil {
		return err
	}
	// Decrypt opcode
	if err := conn.cipher.Decrypt(message[0:2], nil); err != nil {
		return err
	}
	opcode := uint32(binary.LittleEndian.Uint16(message[0:2]))
	// Convert opcode to higher-level type
	messageType, err := conn.type_descriptor.LookupType(opcode)
	if err != nil {
		return err
	}
	packet.Type = messageType
	// Set packet data
	packet.Buffer.SetBytes(message[2:])

	return nil
}

/*******************
* Server -> client *
********************/

func io_server_send_0_12340(conn *Connection, packet *message.Packet) error {
	// Convert high-level type to the low-level protocol level representation.
	resolvedOpcode, err := conn.type_descriptor.LookupOpcode(packet.Type)
	if err != nil {
		return err
	}

	// Start building message envelope
	message := make([]byte, packet.Len()+2+2)
	// Opcodes are included within the packet size. (network byte order)
	binary.BigEndian.PutUint16(message[0:2], uint16(packet.Len()+4))
	binary.LittleEndian.PutUint16(message[2:4], uint16(resolvedOpcode))

	copy(message[4:], packet.Bytes())

	if err := conn.cipher.Encrypt(message[:4], nil); err != nil {
		return err
	}

	_, err = conn.tcp_conn.Write(message)
	return err
}

func io_server_recv_0_12340(conn *Connection, packet *message.Packet) error {
	// Receive size prefix. May or may not be encrypted
	var sizePrefix [2]byte
	if _, err := io.ReadFull(conn.tcp_conn, sizePrefix[:]); err != nil {
		return err
	}
	if err := conn.cipher.Decrypt(sizePrefix[:], nil); err != nil {
		return err
	}
	size := int(binary.BigEndian.Uint16(sizePrefix[:]))
	if size < 4 {
		return fmt.Errorf("game/protocol: server received client message without opcode")
	}

	// Receive message data. The opcode prefix may be encrypted,
	// while the content of the message is always in plaintext.
	message := make([]byte, size)
	if _, err := io.ReadFull(conn.tcp_conn, message); err != nil {
		return err
	}
	// Decrypt opcode
	if err := conn.cipher.Decrypt(message[0:4], nil); err != nil {
		return err
	}
	opcode := binary.LittleEndian.Uint32(message[0:4])
	// Convert opcode to higher-level type
	messageType, err := conn.type_descriptor.LookupType(opcode)
	if err != nil {
		return err
	}
	packet.Type = messageType
	// Set packet data
	packet.Buffer.SetBytes(message[4:])
	return nil
}
