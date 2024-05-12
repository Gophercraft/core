package protocol

import "net"

func (connection *Connection) RemoteAddr() net.Addr {
	return connection.tcp_conn.RemoteAddr()
}
