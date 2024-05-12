package social

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/version"
)

type WhoRequest struct {
	LevelMin, LevelMax    uint32
	PlayerName, GuildName string
	RaceMask, ClassMask   uint32
	ZonesCount            uint32
	Strings               []string
}

func (wr *WhoRequest) Decode(build version.Build, e *message.Packet) error {
	wr.LevelMin = e.ReadUint32()
	wr.LevelMax = e.ReadUint32()
	wr.PlayerName = e.ReadCString()
	wr.GuildName = e.ReadCString()
	wr.RaceMask = e.ReadUint32()
	wr.ClassMask = e.ReadUint32()
	wr.ZonesCount = e.ReadUint32()
	strCount := e.ReadUint32()

	if wr.ZonesCount > 10 {
		return fmt.Errorf("packet: too many zones")
	}

	if strCount > 4 {
		return fmt.Errorf("packet: too many strings")
	}

	strs := []string{}

	for x := uint32(0); x < strCount; x++ {
		strs = append(strs, strings.ToLower(e.ReadCString()))
	}

	wr.Strings = strs
	return nil
}

type WhoMatch struct {
	PlayerName string
	GuildName  string
	Level      uint32
	Class      uint32
	Race       uint32
	ZoneID     uint32
}

type Who struct {
	DisplayCount uint32
	WhoMatches   []WhoMatch
}

func (w *Who) Encode(build version.Build, p *message.Packet) error {
	p.Type = message.SMSG_WHO

	displayCt := uint32(len(w.WhoMatches))

	p.WriteUint32(displayCt) // placeholder value
	p.WriteUint32(uint32(len(w.WhoMatches)))

	for i, m := range w.WhoMatches {
		if p.Size() > packet.MaxLength-6 {
			break
		}

		p.WriteCString(m.PlayerName)
		p.WriteCString(m.GuildName)
		p.WriteUint32(m.Level)
		p.WriteUint32(m.Class)
		p.WriteUint32(m.Race)
		p.WriteUint32(m.ZoneID)

		displayCt = uint32(i) + 1
	}

	p.Start()

	p.WriteUint32(displayCt)

	return nil
}

func (w *Who) Decode(build version.Build, in *message.Packet) error {
	return nil
}
