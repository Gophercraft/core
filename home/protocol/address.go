package protocol

import (
	context "context"
	"fmt"

	"google.golang.org/grpc/peer"
)

func GetPeerAddress(ctx context.Context) (peer_address string, err error) {
	var peer_context *peer.Peer
	var ok bool
	peer_context, ok = peer.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("cannot get address from RPC context")
		return
	}
	peer_address = peer_context.Addr.String()
	return
}
