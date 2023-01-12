package update

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Target struct {
	ID guid.GUID
}

func (t *Target) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	err = t.ID.EncodeUnpacked(build, out)
	return
}

func (t *Target) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	t.ID, err = guid.DecodeUnpacked(build, in)
	return
}
