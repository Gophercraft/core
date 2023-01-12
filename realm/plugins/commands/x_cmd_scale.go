package commands

import "github.com/Gophercraft/core/realm"

func cmdScale(s *realm.Session, scale float32) {
	if scale < .1 || scale > 1000 {
		s.Warnf("scale must be [0.1 - 1000.0]")
		return
	}

	if scale == 0 {
		scale = 1
	}

	s.SetFloat32("ScaleX", scale)
	s.UpdatePlayer()
}
