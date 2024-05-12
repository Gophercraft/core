package login

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type Player struct {
	Character guid.GUID
}

func (j *Player) Decode(build version.Build, in *message.Packet) error {
	var err error
	j.Character, err = guid.DecodeUnpacked(build, in)
	return err
}

func (j *Player) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_PLAYER_LOGIN
	j.Character.EncodeUnpacked(build, out)
	return nil
}
