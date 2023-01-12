package social

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type PlayerList struct {
	Type    packet.WorldType
	Players []guid.GUID
}

func (l *PlayerList) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = l.Type
	out.WriteByte(uint8(len(l.Players)))

	for _, player := range l.Players {
		player.EncodeUnpacked(build, out)
	}

	return nil
}

func (l *PlayerList) Decode(build vsn.Build, in *packet.WorldPacket) error {
	numPlayers := int(in.ReadByte())
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
