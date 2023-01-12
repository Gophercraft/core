package realm

import (
	"strconv"
	"strings"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/tempest"
)

func (s *Server) InitConfig(cfg *config.World) error {
	s.Config = cfg

	alreadySet := make(map[string]bool)

	// Load configs set by client
	for key := range cfg.WorldVars {
		alreadySet[key] = true
	}

	// load preset
	preset, ok := config.Presets[cfg.RealmType]
	if ok {
		for key, va := range preset {
			if !alreadySet[key] {
				s.Config.WorldVars[key] = va
			}
		}
	}

	return nil
}

// None of these var functions should ever error, but panic instead.
// Type checking of worldvars should be done at server startup.

func (s *Server) BoolVar(vn string) bool {
	b, ok := s.Config.WorldVars[vn]
	if !ok {
		return false
	}

	return b == "true"
}

func (s *Server) FloatVar(vn string) float32 {
	b, ok := s.Config.WorldVars[vn]
	if !ok {
		return 0.0
	}

	f, err := strconv.ParseFloat(b, 64)
	if err != nil {
		panic(err)
	}
	return float32(f)
}

func (s *Server) UintVar(vn string) uint64 {
	b, ok := s.Config.WorldVars[vn]
	if !ok {
		return 0.0
	}

	u, err := strconv.ParseUint(b, 10, 64)
	if err != nil {
		panic(err)
	}
	return u
}

func (s *Server) StringVar(vn string) string {
	b, ok := s.Config.WorldVars[vn]
	if !ok {
		return ""
	}

	return b
}

type Pos struct {
	tempest.C4Vector
	MapID uint32
}

func (s *Server) PosVar(vn string) (pos *Pos) {
	pos = new(Pos)

	b, ok := s.Config.WorldVars[vn]
	if !ok {
		return nil
	}

	strs := strings.SplitN(b, " ", 2)
	if len(strs) != 2 {
		panic("invalid pos string")
	}

	mapID, err := strconv.ParseUint(strs[0], 0, 32)
	if err != nil {
		panic(err)
	}

	vec4, err := tempest.ParseC4Vector(strs[1])
	if err != nil {
		panic(err)
	}

	pos.C4Vector = vec4
	pos.MapID = uint32(mapID)

	return pos
}
