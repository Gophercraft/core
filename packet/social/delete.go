package social

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Delete struct {
	Type packet.WorldType
	ID   guid.GUID
}

func (del *Delete) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	del.Type = in.Type
	del.ID, err = guid.DecodeUnpacked(build, in)
	return
}

func (del *Delete) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = del.Type
	err = del.ID.EncodeUnpacked(build, out)
	return
}
