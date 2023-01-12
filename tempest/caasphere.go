package tempest

import "github.com/arl/math32"

var (
	inf = math32.Inf(1)
)

type CAaSphere struct {
	Position C3Vector
	Radius   float32
}

func (sphere *CAaSphere) Contains(point C3Vector) bool {
	if sphere.Radius == inf {
		return true
	}

	return sphere.Position.Distance(point) <= sphere.Radius
}
