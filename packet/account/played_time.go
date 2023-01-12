package account

import (
	"time"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type PlayedTimeRequest struct {
}

func (ptr *PlayedTimeRequest) Decode(build vsn.Build, in *packet.WorldPacket) error {
	return nil
}

type PlayedTime struct {
	TotalTime time.Duration
	LevelTime time.Duration
}

func (pt *PlayedTime) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_PLAYED_TIME
	out.WriteInt32(int32(pt.TotalTime / time.Millisecond))
	out.WriteInt32(int32(pt.LevelTime / time.Millisecond))
	return nil
}
