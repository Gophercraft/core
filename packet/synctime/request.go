package synctime

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Request struct {
	ServerTimeMs uint32
}

func (r *Request) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.SMSG_TIME_SYNC_REQ
	out.WriteUint32(r.ServerTimeMs)
	return
}

func (r *Request) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	r.ServerTimeMs = in.ReadUint32()
	return nil
}
