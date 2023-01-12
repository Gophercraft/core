package realm

import (
	"fmt"
	"sync"
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/format/terrain"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/log"
	"github.com/arl/math32"
)

type orop uint16

const (
	o_modify orop = 1 << iota
	o_add
	o_get
	o_byid
	o_byrange
	o_position
)

type objectrequest struct {
	Op    orop
	Value any
	// True is sent when
	Confirm chan error
}

// Map describes a world map.
// Map is meant to be used asynchronously.
// Map has its a goroutine with which it can emplace/remove world objects, and query them repeatedly
// this way we can avoid race conditions.
type Map struct {
	sync.Mutex

	ID uint32

	// If true, no part of the map is cut off from visibility.
	ShowEverywhere bool
	Phase          *Phase

	objects       map[guid.GUID]*mapobject
	objectrequest chan *objectrequest

	blocks []*MapBlock

	playerlist map[guid.GUID]*mapobject
}

func (m *Map) Definition() *dbdefs.Ent_Map {
	var def *dbdefs.Ent_Map
	m.Phase.Server.DB.Lookup(wdb.BucketKeyUint32ID, m.ID, &def)
	return def
}

type mapobject struct {
	lastpos tempest.C3Vector
	object  WorldObject
}

func (m *Map) sendrequest(or *objectrequest) bool {
	m.Lock()
	if m.objectrequest == nil {
		m.Unlock()
		return false
	}
	m.objectrequest <- or
	m.Unlock()
	return true
}

func (m *Map) terrainMapParam() *terrain.MapParam {
	return &terrain.DefaultMap
}

func (m *Map) calcBlockRange(indexCenter terrain.BlockIndex, extent int32) (min terrain.BlockIndex, max terrain.BlockIndex, err error) {
	param := m.terrainMapParam()

	min.X = 0
	min.Y = 0

	min.X = indexCenter.X - extent
	min.Y = indexCenter.Y - extent

	max.X = indexCenter.X + extent
	max.Y = indexCenter.Y + extent

	if min.X < 0 {
		min.X = 0
	}
	if min.Y < 0 {
		min.Y = 0
	}

	if max.X >= param.BlockSize.X {
		max.X = param.BlockSize.X - 1
	}
	if max.Y >= param.BlockSize.Y {
		max.Y = param.BlockSize.Y - 1
	}
	return
}

// 6 ┌─────┐
// 5 │     │
// 4 │  _  │
// 3 │ |x| │ blockIndex = C2Vector{3,3}
// 2 │  ͞   │ min = C2Vector{2,2}
// 1 │     │ max = C2Vector{4,4}
// 0 └─────┘
// . 0123456
//
// Rather than querying the full width and length of the Map, only check positions from a subset of loaded chunks
// TODO: benchmark to see if this is indeed faster than simply querying all of the individual position vectors of all WorldObjects.
func (m *Map) getObjectSetByMapBlocks(sphere *tempest.CAaSphere, filters []FilterFunc) (WorldObjectSet, error) {
	param := m.terrainMapParam()

	punt := sphere.Position
	blockIndex, err := terrain.CalcBlockIndex(param, punt.C2())
	if err != nil {
		return nil, err
	}

	blockExtent := int32(sphere.Radius/terrain.BlockSize) + 1

	// Build a rectangle based on this extent for use in a grid-search
	min, max, err := m.calcBlockRange(blockIndex, blockExtent)
	if err != nil {
		return nil, err
	}

	var objects WorldObjectSet

	// Perform the grid search.
	for x := int32(min.X); x <= int32(max.X); x++ {
		for y := int32(min.Y); y <= int32(max.Y); y++ {
			// Find index
			index1d := (x * param.BlockSize.X) + y
			currentLength := len(m.blocks)
			// if currentLength <= int(singleIndex) {
			// 	m.blocks = append(m.blocks, make([]*MapBlock, (singleIndex+1)-currentLength))
			// }

			if currentLength > int(index1d) {
				mapBlock := m.blocks[index1d]
				// If map block is nil, that means nothing is using it. Which is fine.
				if mapBlock != nil {
					for _, object := range mapBlock.objects {
						if sphere.Contains(object.Movement().Position.C3()) {
							pass := true
							for _, filter := range filters {
								if !filter(object) {
									pass = false
									break
								}
							}

							if pass {
								objects = append(objects, object)
							}
						}
					}
				}
			}
		}
	}

	return objects, nil
}

func (m *Map) getObjectsByRange(sphere *tempest.CAaSphere, filters ...FilterFunc) (WorldObjectSet, error) {
	// Since everything is synced, we do not care what mapblock the objects occupy.
	if m.ShowEverywhere {
		var out WorldObjectSet

		for _, mobject := range m.objects {
			object := mobject.object

			if sphere.Contains(mobject.lastpos) {
				pass := true
				for _, filter := range filters {
					if !filter(object) {
						pass = false
						break
					}
				}

				if pass {
					out = append(out, object)
				}
			}
		}

		return out, nil
	}

	return m.getObjectSetByMapBlocks(sphere, filters)
}

type rangeObjectRequest struct {
	Range   tempest.CAaSphere
	Filters []FilterFunc
}

// Handle a query of map objects within a particular range.
func (m *Map) handleObjectGetByRange(or *objectrequest) error {
	rng := or.Value.(*rangeObjectRequest)
	objects, err := m.getObjectsByRange(&rng.Range, rng.Filters...)
	or.Value = objects
	return err
}

func (m *Map) handleObjectGet(or *objectrequest) error {
	switch {
	case or.Op&o_byid != 0:
		id := or.Value.(guid.GUID)
		object, ok := m.objects[id]
		if !ok {
			return fmt.Errorf("realm: no object by id %s", id)
		}
		or.Value = object.object
		return nil
	case or.Op&o_byrange != 0:
		return m.handleObjectGetByRange(or)
	default:
		return fmt.Errorf("invalid object get")
	}
}

func (m *Map) handleRemoveObjects(or *objectrequest) error {
	ids := or.Value.([]guid.GUID)

	for _, id := range ids {
		object, ok := m.objects[id]
		if ok {
			m.removeObject(object)
		}
	}

	return nil
}

func (m *Map) handleAddObject(or *objectrequest) error {
	return m.addObjects(or.Value.(WorldObjectSet))
}

func (m *Map) handleModifyPosition(or *objectrequest) error {
	id := or.Value.(guid.GUID)
	mobject, ok := m.objects[id]
	if !ok {
		return fmt.Errorf("realm: could not find referred guid %s", id)
	}

	return m.updateObjectPosition(mobject, false)
}

func (m *Map) handleObjectrequest(or *objectrequest) error {
	if or.Op&o_get != 0 {
		return m.handleObjectGet(or)
	}

	switch or.Op {
	case o_modify:
		return m.handleRemoveObjects(or)
	case o_modify | o_position:
		return m.handleModifyPosition(or)
	case o_modify | o_add:
		return m.handleAddObject(or)
	}
	return nil
}

func (m *Map) remove() {
	for _, wo := range m.objects {
		if unit, ok := wo.object.(Unit); ok {
			go unit.Notify(UnitNotifyDelete, wo.object)
		}
	}

	m.blocks = nil
	m.objects = nil
	m.playerlist = nil

	m.Phase.removeMap(m.ID)
}

func (m *Map) unloadBlock(index terrain.BlockIndex) {
	param := m.terrainMapParam()

	linearIndex := int((index.X * param.BlockSize.X) + index.Y)

	if linearIndex >= len(m.blocks) {
		return
	}

	block := m.blocks[linearIndex]
	if block == nil {
		// no point unloading a block that wasn't loaded
		return
	}

	for _, object := range block.objects {
		if unit, ok := object.(Unit); ok {
			unit.Notify(UnitNotifyDelete, unit)
		}
	}

	block.objects = nil

	m.blocks[linearIndex] = nil
}

// Cleans up unused sections of the map.
func (m *Map) sweep() {
	mapRange := m.VisibilityDistance()

	// If sync is universal, disregard sweeping.
	if math32.IsInf(mapRange, 1) {
		return
	}

	param := m.terrainMapParam()
	// if bit is false, it will be unloaded (if it is loaded)
	mask := make([]uint32, param.BlockSize.X*param.BlockSize.Y)

	blockExtent := int32(mapRange/terrain.BlockSize) + 1

	// Save the blocks that players are near from destruction.
	for _, player := range m.playerlist {
		pos := player.lastpos.C2()

		blockIndex, err := terrain.CalcBlockIndex(param, pos)
		if err != nil {
			panic(err)
		}

		min, max, err := m.calcBlockRange(blockIndex, blockExtent)
		if err != nil {
			panic(err)
		}

		for x := min.X; x <= max.X; x++ {
			for y := min.Y; y <= max.Y; y++ {
				linearIndex := (x * param.BlockSize.X) + y
				maskIndex := linearIndex / 32
				maskFlag := uint32(1) << (uint32(linearIndex) % 32)
				mask[int(maskIndex)] |= maskFlag
			}
		}
	}

	for x := int32(0); x < param.BlockSize.X; x++ {
		for y := int32(0); y < param.BlockSize.Y; y++ {
			linearIndex := (x * param.BlockSize.X) + y
			maskIndex := linearIndex / 32
			maskFlag := uint32(1) << (uint32(linearIndex) % 32)

			if mask[int(maskIndex)]&maskFlag == 0 {
				m.unloadBlock(terrain.BlockIndex{
					x, y,
				})
			}
		}
	}

}

func (m *Map) shouldRemove() bool {
	return len(m.playerlist) == 0
}

func (m *Map) run() {
	sweepTick := time.NewTicker(2 * time.Minute)

	for {
		select {
		case or := <-m.objectrequest:
			err := m.handleObjectrequest(or)
			or.Confirm <- err
			close(or.Confirm)
			//
		case <-sweepTick.C:
			if m.shouldRemove() {
				// Someone is in the middle of talking to the map.
				// Hold off on cleanup
				if !m.TryLock() {
					continue
				}
				m.remove()
				close(m.objectrequest)
				m.objectrequest = nil
				m.Unlock()
				sweepTick.Stop()
				return
			} else {
				m.sweep()
			}
		}
	}
}

func (m *Map) updateObjectPosition(mobject *mapobject, first bool) error {
	nextpos := mobject.object.Movement().Position.C3()
	param := m.terrainMapParam()
	updatedUnit, updatedIsUnit := mobject.object.(Unit)
	vis := m.VisibilityDistance()

	c2 := nextpos.C2()

	block := m.getBlockByPosition(c2)

	blockIndex, err := terrain.CalcBlockIndex(param, c2)
	if err != nil {
		return err
	}

	// TODO: if position delta <= n, skip these costly calculations.

	if m.ObjectIsPlayer(mobject.object) {
		m.loadContentEnvirons(blockIndex)
	}

	// If this is not the first time an object's position has been updated,
	// then it needs to be disassociated from its last position.
	// if not, bad things can happen like a mapobject being considered in two places at once. Yikes!
	if !first {
		lastpos := mobject.lastpos

		lastBlock := m.getBlockByPosition(lastpos.C2())

		// Block may have been unloaded due to not being used. Which is okay. Because that means our object's position is already removed.
		if lastBlock != nil {
			if !lastBlock.removeobject(mobject.object) {
				panic(fmt.Errorf("object %s was not in block to begin with?", mobject.object.GUID()))
			}
		}

		// Next, we need to check nearby mapblocks for Units that can see us.
		nearLastPos := &tempest.CAaSphere{
			Position: lastpos,
			Radius:   vis,
		}
		nearNextPos := &tempest.CAaSphere{
			Position: nextpos,
			Radius:   vis,
		}
		nearOldPosObjects, err := m.getObjectsByRange(nearLastPos)
		if err != nil {
			panic(err)
		}
		nearNewPosObjects, err := m.getObjectsByRange(nearNextPos)
		if err != nil {
			panic(err)
		}

		// We Go here to avoid deadlock on the Map goroutine
		// Notify may call Map methods
		go func() {
			// if len(nearOldPosObjects) > 0 {
			// 	unitsNearOldPos := nearOldPosObjects.Units()
			// 	for _, unit := range unitsNearOldPos {
			// 		unitPos := unit.Movement().Position.C3()
			// 		// If this Unit was considered "near" to where the object was, but not where the unit is now
			// 		//  let's Notify this Unit that our object is now out of range.
			// 		if nearLastPos.Contains(unitPos) && !nearNextPos.Contains(unitPos) {
			// 			unit.Notify(UnitNotifyOutOfRange, mobject.object)
			// 			if updatedIsUnit {
			// 				updatedUnit.Notify(UnitNotifyOutOfRange, unit)
			// 			}
			// 		}
			// 	}
			// }
			if len(nearOldPosObjects) > 0 {
				for _, object := range nearOldPosObjects {
					objectPos := object.Movement().Position.C3()
					unit, isUnit := object.(Unit)
					// If this Unit was considered "near" to where the object was, but not where the unit is now
					//  let's Notify this Unit that our object is now out of range.
					if nearLastPos.Contains(objectPos) && !nearNextPos.Contains(objectPos) {
						// Tell this object that our updated object is out of range.
						if isUnit {
							unit.Notify(UnitNotifyOutOfRange, mobject.object)
						}
						// Tell our updated object that this object is out of range
						if updatedIsUnit {
							updatedUnit.Notify(UnitNotifyOutOfRange, object)
						}
					}
				}
			}
			if len(nearNewPosObjects) > 0 {
				for _, object := range nearNewPosObjects {
					objectPos := object.Movement().Position.C3()
					unit, isUnit := object.(Unit)
					if object != mobject.object {
						// But if a Unit is out of range of where the object was, and is no longer out of range
						// Then as far as that unit is concerned, the object has been created.
						if !nearLastPos.Contains(objectPos) && nearNextPos.Contains(objectPos) {
							if isUnit {
								unit.Notify(UnitNotifyCreate, mobject.object)
							}
							if updatedIsUnit {
								updatedUnit.Notify(UnitNotifyCreate, object)
							}
						}
					}
				}
			}
			// if len(nearNewPosObjects) > 0 {
			// 	unitsNearNewPos := nearNewPosObjects.Units()
			// 	for _, unit := range unitsNearNewPos {
			// 		if unit.GUID() != mobject.object.GUID() {
			// 			unitPos := unit.Movement().Position.C3()
			// 			// But if a Unit is out of range of where the object was, and is no longer out of range
			// 			// Then as far as that unit is concerned, the object has been created.
			// 			if !nearLastPos.Contains(unitPos) && nearNextPos.Contains(unitPos) {
			// 				unit.Notify(UnitNotifyCreate, mobject.object)
			// 				if updatedIsUnit {
			// 					updatedUnit.Notify(UnitNotifyCreate, unit)
			// 				}
			// 			}
			// 		}
			// 	}
			// }

		}()
	}

	// else {
	// 	// Just show me the objects that are near this new position (first)
	// 	nearNextPos := &tempest.CAaSphere{
	// 		Position: nextpos,
	// 		Radius:   vis,
	// 	}
	// 	nearNewPosObjects, err := m.getObjectsByRange(nearNextPos)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	go func() {
	// 		for _, object := range nearNewPosObjects {
	// 			unit, isUnit := object.(Unit)
	// 			if object != mobject.object {
	// 				// But if a Unit is out of range of where the object was, and is no longer out of range
	// 				// Then as far as that unit is concerned, the object has been created.
	// 				if isUnit {
	// 					unit.Notify(UnitNotifyCreate, mobject.object)
	// 				}
	// 				if updatedIsUnit {
	// 					updatedUnit.Notify(UnitNotifyCreate, unit)
	// 				}
	// 			}
	// 		}
	// 	}()
	// }

	mobject.lastpos = nextpos

	if block == nil {
		block = m.initMapBlock(blockIndex)
	}

	block.emplaceobject(mobject.object)

	return nil
}

func (m *Map) addObjects(objects WorldObjectSet) error {
	for _, obj := range objects {
		if err := m.addObject(obj); err != nil {
			return err
		}
	}
	return nil
}

func (m *Map) addObject(newObject WorldObject) error {
	if newObject == nil {
		return fmt.Errorf("realm: tried to add nil object")
	}

	if newObject.GUID() == guid.Nil {
		return fmt.Errorf("realm: tried to add object with nil guid")
	}

	if _, ok := m.objects[newObject.GUID()]; ok {
		return fmt.Errorf("realm: tried to add object %s, an object which is already in the map", newObject.GUID())
	}

	mobject := &mapobject{
		object: newObject,
	}

	m.objects[newObject.GUID()] = mobject

	pos := newObject.Movement().Position.C3()

	m.updateObjectPosition(mobject, true)

	objectIsPlayer := m.ObjectIsPlayer(newObject)
	if objectIsPlayer {
		m.playerlist[newObject.GUID()] = mobject
	}

	prange := m.getPositionRange(pos)
	// Here, load all objects near the position of the new object
	inRangeObjects, err := m.getObjectsByRange(prange)
	if err != nil {
		return err
	}

	go func() {
		// Broadcast create to all units here (including the new object itself if it is a Unit)
		inRangeObjects.Units().Notify(UnitNotifyCreate, newObject)

		log.Println("In range objects", len(inRangeObjects))

		// If this object is a Unit, notify the unit of all nearby world objects
		if unit, ok := newObject.(Unit); ok {
			for _, wo := range inRangeObjects {
				log.Println("AddObject, notifying unit of nearby object", wo.GUID())
				// Don't self-notify twice
				if wo.GUID() != newObject.GUID() {
					log.Println("Adding object which is different")
					unit.Notify(UnitNotifyCreate, wo)
				}
			}
		}
	}()

	return nil
}

func (m *Map) ObjectIsPlayer(player WorldObject) bool {
	if _, ok := player.(*Session); ok {
		return true
	}
	return false
}

func (m *Map) getPositionRange(position tempest.C3Vector) *tempest.CAaSphere {
	visDistance := m.VisibilityDistance()
	var prange tempest.CAaSphere
	prange.Position = position
	prange.Radius = visDistance
	return &prange
}

func (m *Map) removeObject(destroyed *mapobject) {
	block := m.getBlockByPosition(destroyed.lastpos.C2())
	block.removeobject(destroyed.object)

	wasPlayer := m.ObjectIsPlayer(destroyed.object)

	mobject, ok := m.objects[destroyed.object.GUID()]
	if !ok {
		panic(fmt.Errorf("tried to destroy map object (%s) that does not exist?", mobject.object.GUID().String()))
	}

	delete(m.objects, destroyed.object.GUID())

	deleteRange := m.getPositionRange(destroyed.lastpos)

	objects, err := m.getObjectsByRange(deleteRange, IsUnit)

	if err == nil {
		// Of this subset of objects, Notify them that this ID was deleted
		go func() {
			objects.Units().Notify(UnitNotifyDelete, destroyed.object)
		}()
	}

	if wasPlayer {
		delete(m.playerlist, destroyed.object.GUID())
	}

	// go func() {
	// 	visibilityRange := m.VisibilityDistance()
	// 	destroyedObjectPosition := destroyedObject.Movement().Position.C3()

	// 	inRangeObjects := m.ObjectsNearPosition(destroyedObjectPosition, visibilityRange)
	// 	inRangeObjects.Units().
	// 		Notify(UnitNotifyDelete, destroyedObject)
	// }()
}

func (m *Map) DebugMapInfo(s *Session) {
	s.Kv("Phase", "'%s'", m.Phase.ID)
	s.Kv(" Maps in phase", "%d", len(m.Phase.maps))
	for k, v := range m.Phase.maps {
		s.Kv("  Map", "%d, %d players", k, len(v.playerlist))
	}
	s.Kv("Map", "ID %d, '%s'", m.ID, m.Definition().MapName_lang)
	param := m.terrainMapParam()
	for x := 0; x < int(param.BlockSize.X); x++ {
		for y := 0; y < int(param.BlockSize.Y); y++ {
			index1d := (x * int(param.BlockSize.X)) + y
			if index1d < len(m.blocks) {
				block := m.blocks[index1d]
				if block != nil {
					s.Kv(" Block", "%d,%d", x, y)
					s.Kv(" Content Loaded", "%t", block.contentLoaded)
					for _, v := range block.objects {
						s.Kv("  Object", "%s", v)
					}
				}
			}
		}
	}
}
