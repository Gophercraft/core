package connection

import (
	"context"
	"time"

	"github.com/Gophercraft/core/bnet"
	"github.com/Gophercraft/core/bnet/pb/bgs/protocol"
	connection_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/connection/v1"
	"github.com/Gophercraft/core/bnet/util"
)

type Service struct {
	connection_v1.UnimplementedConnectionServiceServer
}

func (service *Service) RequestDisconnect(ctx context.Context, request *connection_v1.DisconnectRequest) (connect_response *protocol.NO_RESPONSE, err error) {
	connection := bnet.GetConnection(ctx)

	if connection != nil {
		connection.Close()
	}

	return
}

func (service *Service) KeepAlive(ctx context.Context, nodata *protocol.NoData) (connect_response *protocol.NO_RESPONSE, err error) {
	return
}

func (service *Service) Connect(ctx context.Context, connect_request *connection_v1.ConnectRequest) (connect_response *connection_v1.ConnectResponse, err error) {
	connect_response = new(connection_v1.ConnectResponse)
	util.Set(&connect_response.UseBindlessRpc, true)

	connect_response.ServerId = new(protocol.ProcessId)
	util.Set(&connect_response.ServerId.Label, uint32(100))
	util.Set(&connect_response.ServerId.Epoch, uint32(time.Now().Unix()))

	util.Set(&connect_response.ServerTime, uint64(time.Now().UnixMilli()))

	return connect_response, nil
}

func New() (service *Service) {
	service = new(Service)
	return
}
