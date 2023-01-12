package realm

import (
	"github.com/Gophercraft/core/tempest"
)

type Sight struct {
	Position tempest.C4Vector
}

func (s *Session) GetSight() *Sight {
	return &s.Camera
}

func (m *Map) CanSee(eye WorldObject, object WorldObject, vis float32) bool {
	if eye.GUID() == object.GUID() {
		return true
	}

	if vis == -1 {
		vis = m.VisibilityDistance()
	}
	if seer, ok := eye.(Seer); ok {
		sight := seer.GetSight()
		if sight != nil {
			eyePosition := sight.Position
			objectPosition := object.Movement().Position
			distance := eyePosition.Distance(objectPosition)
			return distance <= vis
		}
	}

	return eye.Movement().Position.Distance(object.Movement().Position) <= vis
}

// func (s *Session) UpdateCameraPosition(pos tempest.C3Vector) {

// }
