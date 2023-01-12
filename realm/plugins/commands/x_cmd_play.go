package commands

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/realm"
)

func cmdPlayVisual(s *realm.Session, visualKit uint32) {
	id := s.GetTarget()

	if id == guid.Nil {
		id = s.GUID()
	}

	if id == guid.Nil {
		return
	}

	object := s.Map().GetObject(id)

	if object != nil {
		unit, ok := object.(realm.Unit)
		if ok {
			s.Map().PlaySpellVisual(unit, visualKit)
		}
	}
}

func cmdPlaySound(s *realm.Session, sound uint32) {
	s.Map().PlaySound(sound)
}

func cmdPlayMusic(s *realm.Session, music uint32) {
	s.Map().PlayMusic(music)
}
