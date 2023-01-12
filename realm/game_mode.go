package realm

import (
	"fmt"

	"github.com/Gophercraft/core/realm/wdb/models"
)

type GameMode models.GameMode

func (s *Session) GameMode() models.GameMode {
	return s.gameMode
}

func (s *Session) God() bool {
	return s.gameMode == models.GameMode_God
}

func (s *Session) SetGameMode(gm models.GameMode) error {
	s.GuardSession.Lock()
	prev := s.gameMode
	s.gameMode = gm
	s.GuardSession.Unlock()

	if prev == gm {
		return fmt.Errorf("realm: gamemode is already %s", gm)
	}

	s.Warnf("Game mode set from %s -> %s", prev, gm)
	return nil
}
