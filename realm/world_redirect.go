package realm

import (
	"net"

	"github.com/Gophercraft/log"
)

// Needed for the Alpha protocol.
func (ws *Server) serveRedirect() {
	var redirectAddress = ws.Config.Redirect

	if redirectAddress == "" {
		return
	}

	var publicRedirect = ws.Config.PublicRedirect
	if publicRedirect == "" {
		publicRedirect = redirectAddress
	}

	srv, err := net.Listen("tcp", redirectAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening redirect server at", redirectAddress)

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go ws.sendRedirectAddress(conn)
	}
}

func (ws *Server) sendRedirectAddress(conn net.Conn) {
	redirectAddress := ws.Config.PublicAddress
	conn.Write(append([]byte(redirectAddress), 0))
	conn.Close()
}
