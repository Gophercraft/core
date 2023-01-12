package realm

import (
	"time"

	"github.com/Gophercraft/core/packet/update"
)

const MovementManagerTag = "Gophercraft/core/realm/movement_mgr"

type MovementManager struct {
	Moves        chan *update.MovementPacket
	LastMove     time.Time
	MoveSpeeds   update.Speeds
	MovementInfo *update.MovementInfo
}

func (s *Session) savePosition() {
	if _, err := s.DB().Where("id = ?", s.PlayerID()).Cols("x", "y", "z", "o", "map", "zone").Update(s.Char); err != nil {
		panic(err)
	}
}

func (s *Session) savePosInterval() time.Duration {
	return 5 * time.Second
}

func (s *Session) validateMove(move *update.MovementPacket) error {
	return nil
}

func (s *Session) handleMoveUpdate(move *update.MovementPacket) {
	if err := s.validateMove(move); err != nil {
		s.Log("Error validating movement", err)
		return
	}

	pos := move.Info.Position
	s.MovementInfo = move.Info
	s.Camera.Position = pos

	s.SetPosition(pos)

	s.Char.Map = s.CurrentMap.ID

	if time.Since(s.LastMove) > s.savePosInterval() {
		s.savePosition()
	}

	s.broadcastMovementLocal(move)

	s.Map().UpdateMapPosition(s.GUID())
}

func (s *Session) broadcastMovementLocal(move *update.MovementPacket) {
	units := s.Map().
		VisibleObjects(s).
		WithoutGUID(s.GUID()). // Client doesn't need to be reminded of their own movements
		Units()

	// log.Println("Relaying to ", len(units), "n units")

	units.NotifyMovement(move.Type, s)
}

func (s *Session) initMovementManager() {
	s.CreateProcess(MovementManagerTag, (*Session).movementManagerProcess)
}

func (s *Session) movementManagerProcess(cancel <-chan bool) {
	s.Moves = make(chan *update.MovementPacket, 16)
	ticker := time.NewTicker(s.savePosInterval())

	for {
		select {
		case <-cancel:
			close(s.Moves)
			s.Moves = nil
			ticker.Stop()
			return
		case move := <-s.Moves:
			s.handleMoveUpdate(move)
		case <-ticker.C:
			s.savePosition()
		}
	}
}
