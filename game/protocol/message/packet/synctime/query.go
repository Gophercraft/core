package synctime

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type QueryTime struct {
}

func (qt *QueryTime) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_QUERY_TIME
	return nil
}

func (qt *QueryTime) Decode(build version.Build, out *message.Packet) error {
	return nil
}

type QueryTimeResponse struct {
	Unix int32
}

func (tr *QueryTimeResponse) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_QUERY_TIME_RESPONSE
	out.WriteInt32(tr.Unix)
	return nil
}
