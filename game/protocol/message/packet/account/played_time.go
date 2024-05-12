package account

import (
	"time"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type PlayedTimeRequest struct {
}

func (ptr *PlayedTimeRequest) Decode(build version.Build, in *message.Packet) error {
	return nil
}

type PlayedTime struct {
	TotalTimeMilliseconds int32
	LevelTimeMilliseconds int32
}

func MakePlayedTime(total, level time.Duration) (pt PlayedTime) {
	pt.TotalTimeMilliseconds = int32(total / time.Millisecond)
	pt.LevelTimeMilliseconds = int32(level / time.Millisecond)
	return
}

func (pt *PlayedTime) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_PLAYED_TIME
	out.WriteInt32(pt.TotalTimeMilliseconds)
	out.WriteInt32(pt.LevelTimeMilliseconds)
	return nil
}
