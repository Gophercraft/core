package network

import "net"

func (session *Session) RemoteAddr() net.Addr {
	if session.protocol_connection == nil {
		return nil
	}
	return session.protocol_connection.RemoteAddr()
}
