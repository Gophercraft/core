package grunt

import (
	"net"

	"github.com/Gophercraft/log"
)

func (server *Server) start_run_server() {
	if err := server.Run(); err != nil {
		log.Warn("grunt server stopped running:", err)
	}
}

func (server *Server) SetListener(listener net.Listener) {
	server.listener = listener
}

func (server *Server) Start() (err error) {
	// bind on address:port
	var listener net.Listener
	listener, err = net.Listen("tcp", server.config.Address)
	if err != nil {
		return
	}

	server.SetListener(listener)

	// run server in a separate goroutine
	go server.start_run_server()

	return
}
