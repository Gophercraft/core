package commands

import (
	"log"

	"github.com/Gophercraft/core/realm"
)

func cmdMorph(s *realm.Session, displayID uint32) {
	log.Println("Morphing to ", displayID)

	s.SetUint32("DisplayID", displayID)
	s.UpdatePlayer()
}
