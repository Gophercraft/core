package rpc

import (
	"context"

	"github.com/Gophercraft/core/bnet/util"
)

type Method uint32

type ClientConnectionInterface interface {
	Invoke(ctx context.Context, full_method string, service util.ServiceHash, method Method, wait_for_response bool, args, reply any, opts ...CallOption) error
}
