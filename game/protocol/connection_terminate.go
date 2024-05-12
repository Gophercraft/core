package protocol

func (conn *Connection) Terminate(err error) {
	close(conn.send_queue)
	conn.logerr(err)
	conn.cancel_cause(err)
}
