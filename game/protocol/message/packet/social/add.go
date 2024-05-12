package social

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/version"
)

type Add struct {
	Type packet.WorldType
	Name string
}

func (add *Add) Decode(build version.Build, in *message.Packet) (err error) {
	add.Type = in.Type
	add.Name = in.ReadCString()
	return
}

func (add *Add) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = add.Type
	out.WriteCString(add.Name)
	return
}
