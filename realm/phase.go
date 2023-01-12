package realm

import (
	"fmt"
	"sync"
	"time"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/log"
)

// Phase describes a plane of existence which contains multiple maps.
// When the list of maps is empty after too long, Phase itself becomes unloaded.
type Phase struct {
	sync.Mutex
	ID     string
	Server *Server

	maps map[uint32]*Map
}

func (phase *Phase) Map(id uint32) *Map {
	log.Warn("getting map", id)
	m, err := phase.loadMap(id)
	if err != nil {
		log.Warn(err)
		return nil
	}
	log.Warn("got map", id)
	return m
}

func (phase *Phase) loadMap(id uint32) (*Map, error) {
	phase.Lock()
	defer phase.Unlock()
	m, ok := phase.maps[id]
	if ok {
		return m, nil
	}

	log.Println(phase.ID, "Map not loaded. Time to load.", id)

	m = new(Map)
	m.Phase = phase
	m.ID = id
	m.objectrequest = make(chan *objectrequest)
	param := m.terrainMapParam()
	m.objects = make(map[guid.GUID]*mapobject)
	m.playerlist = make(map[guid.GUID]*mapobject)
	m.blocks = make([]*MapBlock, param.BlockSize.X*param.BlockSize.Y)

	def := m.Definition()

	if def == nil {
		return nil, fmt.Errorf("realm: Phase.LoadMap: map does not exist")
	}

	go m.run()

	phase.maps[id] = m

	return m, nil
}

func (phase *Phase) removeMap(mapID uint32) error {
	phase.Lock()
	delete(phase.maps, mapID)
	phase.Unlock()
	return nil
}

func (phase *Phase) removeFromList() {
	phase.Server.GuardPhases.Lock()
	delete(phase.Server.Phases, phase.ID)
	phase.Server.GuardPhases.Unlock()
}

func (phase *Phase) sweep() {
	sweepTick := time.NewTicker(2 * time.Minute)

	for {
		select {
		case <-sweepTick.C:
			if !phase.TryLock() {
				continue
			}
			shouldQuit := false
			if len(phase.maps) == 0 {
				phase.removeFromList()
				shouldQuit = true
			}
			phase.Unlock()
			if shouldQuit {
				sweepTick.Stop()
				return
			}
		}
	}
}

func (s *Server) Phase(id string) *Phase {
	s.GuardPhases.Lock()
	ph, ok := s.Phases[id]
	if !ok {
		log.Println("Allocating new phase.")
		ph = new(Phase)
		ph.Server = s
		ph.ID = id
		ph.maps = make(map[uint32]*Map)
		go ph.sweep()
		s.Phases[id] = ph
	}
	s.GuardPhases.Unlock()
	return ph
}
