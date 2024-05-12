package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type Miss uint8

const (
	MissNone       = 0
	MissMiss       = 1
	MissResist     = 2
	MissDodge      = 3
	MissParry      = 4
	MissBlock      = 5
	MissEvade      = 6
	MissImmune     = 7
	MissTempImmune = 8 // one of these 2 is MISS_TEMPIMMUNE
	MissDeflect    = 9
	MissAbsorb     = 10
	MissReflect    = 11
)

type MissStatus struct {
	Target        guid.GUID
	Reason        Miss
	ReflectStatus uint8
}

func (ms *MissStatus) Encode(build version.Build, out *message.Packet) (err error) {
	ms.Target.EncodeUnpacked(build, out)
	out.WriteUint8(uint8(ms.Reason))
	if ms.Reason == MissReflect {
		out.WriteUint8(ms.ReflectStatus)
	}
	return nil
}

func (ms *MissStatus) Decode(build version.Build, in *message.Packet) (err error) {
	ms.Target, err = guid.DecodePacked(build, in)
	if err != nil {
		return
	}
	ms.Reason = Miss(in.ReadUint8())
	if ms.Reason == MissReflect {
		ms.ReflectStatus = in.ReadUint8()
	}
	return nil
}
