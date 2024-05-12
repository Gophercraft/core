package update

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type SetActiveMover struct {
	ID guid.GUID
}

func (sam *SetActiveMover) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.CMSG_SET_ACTIVE_MOVER
	err = sam.ID.EncodeUnpacked(build, out)
	return
}

func (sam *SetActiveMover) Decode(build version.Build, in *message.Packet) (err error) {
	sam.ID, err = guid.DecodeUnpacked(build, in)
	return
}
