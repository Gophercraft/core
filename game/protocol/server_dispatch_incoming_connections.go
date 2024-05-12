package protocol

import (
	"errors"
	"net"
)

func (s *Server) dispatch_incoming_connection() error {
	// Accept incoming raw TCP connection.
	// May be anything, including spam/malware.
	tcpConn, err := s.tcp_listener.Accept()
	if err != nil {
		if errors.Is(err, net.ErrClosed) {
			return nil
		}
		return err
	}

	// We don't trust this connection yet. run tcpConn through all pre-connection security checkpoints.
	// In practice, we should reject the tcpConn based on a fast lookup table of compromised IP addresses.
	if err := s.vetRawConnectionAgainstSecurityCheckpoints(tcpConn); err != nil {
		s.logerr(err)
		tcpConn.Close()
		// Not an error with dispatch, only the client is in error.
		return nil
	}

	// Spawn new goroutine to handle TCP connection.
	go s.handle_tcp_conn(tcpConn)

	return nil
}
