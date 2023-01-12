package commands

import (
	"github.com/Gophercraft/core/realm"
	"github.com/Gophercraft/core/realm/wdb/models"
)

func cmdSetGamemode(s *realm.Session, gm models.GameMode) {
	s.SetGameMode(gm)
}
