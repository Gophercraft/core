package grunt

// Write data to the buffer, without sending
func (session *Session) Write(b []byte) (n int, err error) {
	n, err = session.buffer.Write(b)
	return
}

// Send all data in the buffer, while resetting
func (session *Session) Send() (err error) {
	_, err = session.connection.Write(session.buffer.Bytes())
	session.buffer.Reset()
	return err
}
