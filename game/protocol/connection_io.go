package protocol

import (
	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type io_mode uint8

const (
	client_send io_mode = iota
	client_recv
	server_send
	server_recv
	io_modes
)

type io_func func(conn *Connection, packet *message.Packet) error

type io_functions [int(io_modes)]io_func

var (
	io_protocols = map[version.BuildRange]io_functions{}
)

func (conn *Connection) initialize_io(server_mode bool, emulated version.Build) error {
	// Save protocol emulation settings
	conn.server_mode = server_mode
	conn.build = emulated

	// Connections begin unencrypted
	conn.cipher = crypto.DummyCipher{}

	// Set up abstract opcode mapping
	var err error
	conn.type_descriptor, err = message.QueryTypeDescriptor(conn.build)
	if err != nil {
		return err
	}

	// Set up packet framing/encryption/IO functions
	var desired_protocol io_functions
	if err := version.QueryDescriptors(conn.build, io_protocols, &desired_protocol); err != nil {
		return err
	}
	if server_mode {
		conn.send_io = desired_protocol[server_send]
		conn.recv_io = desired_protocol[server_recv]
	} else {
		conn.send_io = desired_protocol[client_send]
		conn.recv_io = desired_protocol[client_recv]
	}

	// Set up message queue
	conn.send_queue = make(chan *message.Packet)
	go conn.send_queue_worker()

	return nil
}

func (conn *Connection) recv() (packet *message.Packet, err error) {
	packet = new(message.Packet)
	err = conn.recv_io(conn, packet)
	if err != nil {
		packet = nil
	}
	return
}

func (conn *Connection) send(packet *message.Packet) (err error) {
	return conn.send_io(conn, packet)
}
