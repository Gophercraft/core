package update

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
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

func (sss *SetStandState) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.CMSG_STANDSTATECHANGE
	out.WriteUint32(uint32(sss.State))
	return nil
}

func (sss *SetStandState) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	sss.State = StandState(in.ReadUint32())
	return nil
}
