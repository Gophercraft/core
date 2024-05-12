package protocol

import "net"

// A function that checks a raw TCP connection for anything suspicious or prohibited.
// Called before the net.Conn is used for anything.
// If an error is returned, the connection is terminated by the server.
type PreConnectSecurityCheckpointFunc func(s *Server, c net.Conn) error

func (s *Server) vetRawConnectionAgainstSecurityCheckpoints(c net.Conn) error {
	for _, checkpoint := range s.preConnectSecurityCheckpoints {
		if err := checkpoint(s, c); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) AddPreConnectSecurityCheckpointFunc(checkpoint PreConnectSecurityCheckpointFunc) {
	s.preConnectSecurityCheckpoints = append(s.preConnectSecurityCheckpoints, checkpoint)
}
