package update

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type StandState uint8

const (
	StateStand StandState = iota
	StateSit
	StateSitChair
	StateSleep
	StateSitLowChair
	StateSitMediumChair
	StateSitHighChair
	StateDead
	StateKneel
	StateCustom
)

type SetStandState struct {
	State StandState
}

func (sss *SetStandState) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.CMSG_STANDSTATECHANGE
	out.WriteUint32(uint32(sss.State))
	return nil
}

func (sss *SetStandState) Decode(build version.Build, in *message.Packet) (err error) {
	sss.State = StandState(in.ReadUint32())
	return nil
}
