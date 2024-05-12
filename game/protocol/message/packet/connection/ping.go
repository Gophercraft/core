package connection

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type Ping struct {
	Ping    uint32
	Latency uint32
}

func (pi *Ping) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_PING
	out.WriteUint32(pi.Ping)
	out.WriteUint32(pi.Latency)
	return nil
}

func (pi *Ping) Decode(build version.Build, in *message.Packet) error {
	pi.Ping = in.ReadUint32()
	pi.Latency = in.ReadUint32()
	return nil
}

type Pong struct {
	Ping uint32
}

func (po *Pong) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_PONG
	out.WriteUint32(po.Ping)
	return nil
}

func (po *Pong) Decode(build version.Build, out *message.Packet) error {
	out.WriteUint32(po.Ping)
	return nil
}
