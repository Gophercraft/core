package commands

import "github.com/Gophercraft/core/realm"

func cmdModLevel(s *realm.Session, level int) {
	s.LevelUp(uint32(level))
}
