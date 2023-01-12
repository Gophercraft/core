package packet

import "github.com/Gophercraft/core/vsn"

type Ping struct {
	Ping    uint32
	Latency uint32
}

func (pi *Ping) Encode(build vsn.Build, out *WorldPacket) error {
	out.Type = CMSG_PING
	out.WriteUint32(pi.Ping)
	out.WriteUint32(pi.Latency)
	return nil
}

func (pi *Ping) Decode(build vsn.Build, in *WorldPacket) error {
	pi.Ping = in.ReadUint32()
	pi.Latency = in.ReadUint32()
	return nil
}

type Pong struct {
	Ping uint32
}

func (po *Pong) Encode(build vsn.Build, out *WorldPacket) error {
	out.Type = SMSG_PONG
	out.WriteUint32(po.Ping)
	return nil
}

func (po *Pong) Decode(build vsn.Build, out *WorldPacket) error {
	out.WriteUint32(po.Ping)
	return nil
}
