package update

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type Target struct {
	ID guid.GUID
}

func (t *Target) Encode(build version.Build, out *message.Packet) (err error) {
	err = t.ID.EncodeUnpacked(build, out)
	return
}

func (t *Target) Decode(build version.Build, in *message.Packet) (err error) {
	t.ID, err = guid.DecodeUnpacked(build, in)
	return
}
