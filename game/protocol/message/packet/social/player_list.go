package social

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/version"
)

type PlayerList struct {
	Type    packet.WorldType
	Players []guid.GUID
}

func (l *PlayerList) Encode(build version.Build, out *message.Packet) error {
	out.Type = l.Type
	out.WriteUint8(uint8(len(l.Players)))

	for _, player := range l.Players {
		player.EncodeUnpacked(build, out)
	}

	return nil
}

func (l *PlayerList) Decode(build version.Build, in *message.Packet) error {
	numPlayers := int(in.ReadUint8())
	l.Players = make([]guid.GUID, numPlayers)

	var err error

	for i := 0; i < numPlayers; i++ {
		l.Players[i], err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}
	}

	return nil
}
