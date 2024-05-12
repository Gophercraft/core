package redirect

import "net"

func Run(listen string, redirect_handler func() string) (err error) {
	var server net.Listener
	server, err = net.Listen("tcp", listen)
	if err != nil {
		return
	}

	for {
		var connection net.Conn
		connection, err = server.Accept()
		if err != nil {
			return
		}

		go handle(connection, redirect_handler)
	}
}

func handle(connection net.Conn, redirect_handler func() string) {
	address := redirect_handler()
	connection.Write(append([]byte(address), 0))
	connection.Close()
}
