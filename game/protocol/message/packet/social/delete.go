package social

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/version"
)

type Delete struct {
	Type packet.WorldType
	ID   guid.GUID
}

func (del *Delete) Decode(build version.Build, in *message.Packet) (err error) {
	del.Type = in.Type
	del.ID, err = guid.DecodeUnpacked(build, in)
	return
}

func (del *Delete) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = del.Type
	err = del.ID.EncodeUnpacked(build, out)
	return
}
