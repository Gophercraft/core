package synctime

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type Request struct {
	ServerTimeMs uint32
}

func (r *Request) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.SMSG_TIME_SYNC_REQ
	out.WriteUint32(r.ServerTimeMs)
	return
}

func (r *Request) Decode(build version.Build, in *message.Packet) (err error) {
	r.ServerTimeMs = in.ReadUint32()
	return nil
}
