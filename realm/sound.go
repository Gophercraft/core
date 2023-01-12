package realm

import (
	"fmt"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
)

func (s *Session) SendPlayObjectSound(id guid.GUID, soundID uint32) {
	fmt.Println("Sending direct sound play to", id, soundID)
	p := packet.NewWorldPacket(packet.SMSG_PLAY_OBJECT_SOUND)
	p.WriteUint32(soundID)
	id.EncodeUnpacked(s.Build(), p)
	s.SendPacket(p)
}

func (m *Map) PlaySound(id uint32) {
	m.AllObjects().Sessions().Iter(func(s *Session) {
		s.SendPlaySound(id)
	})
}

func (m *Map) PlayMusic(id uint32) {
	m.AllObjects().Sessions().Iter(func(s *Session) {
		s.SendPlayMusic(id)
	})
}

func (m *Map) PlayObjectSound(id guid.GUID, soundID uint32) {
	speaker := m.GetObject(id)
	if speaker == nil {
		fmt.Println("no speaker", id)
		return
	}

	fmt.Println("Playing", id, soundID)

	if player, ok := (speaker).(*Session); ok {
		player.SendPlayObjectSound(id, soundID)
	}

	for _, nearPlayer := range m.VisibleObjects(speaker).Sessions() {
		nearPlayer.SendPlayObjectSound(id, soundID)
	}
}

func (s *Session) SendPlaySound(id uint32) {
	pkt := packet.NewWorldPacket(packet.SMSG_PLAY_SOUND)
	pkt.WriteUint32(id)
	s.SendPacket(pkt)
}

func (s *Session) SendPlayMusic(id uint32) {
	pkt := packet.NewWorldPacket(packet.SMSG_PLAY_MUSIC)
	pkt.WriteUint32(id)
	s.SendPacket(pkt)
}
