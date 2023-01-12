package bnet

import (
	"time"

	protocol "github.com/Gophercraft/core/bnet/bgs/protocol"
	v1 "github.com/Gophercraft/core/bnet/bgs/protocol/connection/v1"
	"github.com/Gophercraft/log"
)

func (s *Listener) Connect(conn *Conn, token uint32, args *v1.ConnectRequest) {
	resp := v1.ConnectResponse{}
	if args.ClientId != nil {
		resp.ClientId = args.ClientId
	}

	label := uint32(5)
	epoch := uint32(time.Now().Unix())

	resp.ServerId = &protocol.ProcessId{
		Label: &label,
		Epoch: &epoch,
	}

	resp.UseBindlessRpc = args.UseBindlessRpc

	conn.SendResponse(token, &resp)
}

func (s *Listener) Bind(conn *Conn, token uint32, args *v1.BindRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) Echo(conn *Conn, token uint32, args *v1.EchoRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) ForceDisconnect(conn *Conn, token uint32, args *v1.DisconnectNotification) {
	log.Println("disconnecting for reason", args.GetReason())
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (s *Listener) KeepAlive(conn *Conn, token uint32, args *protocol.NoData) {
	conn.SendResponseCode(token, ERROR_OK)
}
func (s *Listener) Encrypt(conn *Conn, token uint32, args *v1.EncryptRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (s *Listener) RequestDisconnect(conn *Conn, token uint32, args *v1.DisconnectRequest) {
	log.Println("disconnecting for reason", args.GetErrorCode())
	conn.SendResponseCode(token, ERROR_OK)

	go func() {
		time.Sleep(500 * time.Millisecond)
		conn.c.Close()
	}()
}
