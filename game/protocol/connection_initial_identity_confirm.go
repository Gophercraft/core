package protocol
s
import (
	"bytes"
	"fmt"
	"io"
)

func (conn *Connection) make_protocol_hello(serverMode bool) []byte {
	if !conn.build.AddedIn(20886) {
		return nil
	}

	hello := []byte("WORLD OF WARCRAFT CONNECTION - ")

	if serverMode {
		hello = append(hello, []byte("SERVER TO CLIENT")...)
	} else {
		hello = append(hello, []byte("CLIENT TO SERVER")...)
	}

	if conn.build.AddedIn(31478) {
		hello = append(hello, []byte(" - V2")...)
	}

	if conn.build.AddedIn(22248) {
		hello = append(hello, '\n')
	}

	return hello
}

func (conn *Connection) confirm_initial_protocol_identity() error {
	if !conn.build.AddedIn(20886) {
		return nil
	}

	// Server sends first hello
	if conn.serverMode {
		if _, err := conn.tcp_conn.Write(conn.make_protocol_hello(true)); err != nil {
			return err
		}
	} else {
		// Client expects server's first hello
		expectedServerHello := conn.make_protocol_hello(true)
		actualServerHello := make([]byte, len(expectedServerHello))
		if _, err := io.ReadFull(conn.tcp_conn, actualServerHello[:]); err != nil {
			return err
		}
		if !bytes.Equal(expectedServerHello, actualServerHello) {
			return fmt.Errorf("protocol: server sent us an incorrect hello")
		}
	}

	// Client sends second hello
	if conn.serverMode {
		// Server expects client's second hello
		expectedClientHello := conn.make_protocol_hello(false)
		actualClientHello := make([]byte, len(expectedClientHello))
		if _, err := io.ReadFull(conn.tcp_conn, actualClientHello[:]); err != nil {
			return err
		}
		if !bytes.Equal(expectedClientHello, actualClientHello) {
			return fmt.Errorf("protocol: client sent us an incorrect hello")
		}
	} else {
		if _, err := conn.tcp_conn.Write(conn.make_protocol_hello(false)); err != nil {
			return err
		}
	}

	return nil
}
