package home

import (
	"crypto/tls"
	"net"

	"github.com/Gophercraft/core/auth"
	"github.com/Gophercraft/core/home/multiproto"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (s *Server) tlsConfig() *tls.Config {
	cfg := &tls.Config{
		Certificates: []tls.Certificate{s.Config.Certificate},
		MinVersion:   tls.VersionTLS12,
		ClientAuth:   tls.RequireAnyClientCert,
	}
	return cfg
}

func (s *Server) listenRpcNet(l net.Listener) {
	config := s.tlsConfig()
	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(config)))

	rpcnet.RegisterHomeServiceServer(grpcServer, &rpcServer{Server: s})

	log.Fatal(grpcServer.Serve(l))
}

func (s *Server) listenAuthServer(l net.Listener) {
	auth.RunServer(s, l)
}

func (s *Server) multiprotoListen() {
	log.Printf("Listening auth/rpc server at %s", s.Config.AuthListen)

	r, err := multiproto.NewRouter(s.Config.AuthListen)
	if err != nil {
		log.Fatal(err)
		return
	}

	// if a TLS hello [x16] is sent to the home server, it will be routed
	// to the GRPC handler
	tlsListener := r.SelectListener([]byte{0x16})

	// if it's anything else, it gets routed to the legacy auth server
	authListener := r.FallbackListener()

	go s.listenRpcNet(tlsListener)
	go s.listenAuthServer(authListener)

	log.Fatal(r.Serve())
}
