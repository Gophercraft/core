package game_utilities

import "github.com/Gophercraft/core/bnet/pb/bgs/protocol"

type ClientRequestParameters map[string]*protocol.Variant

func (parameters ClientRequestParameters) Get(name string) (variant *protocol.Variant) {
	return parameters[name]
}
