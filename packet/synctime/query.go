package synctime

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type QueryTime struct {
}

func (qt *QueryTime) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_QUERY_TIME
	return nil
}

func (qt *QueryTime) Decode(build vsn.Build, out *packet.WorldPacket) error {
	return nil
}

type QueryTimeResponse struct {
	Unix int32
}

func (tr *QueryTimeResponse) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_QUERY_TIME_RESPONSE
	out.WriteInt32(tr.Unix)
	return nil
}
