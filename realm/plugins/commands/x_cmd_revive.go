package commands

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/realm"
)

func cmdRevive(s *realm.Session) {
	target := s.GetTarget()

	if target == guid.Nil {
		target = s.GUID()
	}

	object := s.Map().GetObject(target)
	if object != nil {
		if unit, ok := (object).(realm.Unit); ok {
			if err := s.Map().Revive(unit); err != nil {
				s.Warnf("%s", err)
			}
		}
	}
}
