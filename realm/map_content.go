package realm

import (
	"github.com/Gophercraft/core/format/terrain"
	"github.com/Gophercraft/core/realm/randomness"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
)

type SpawnGroupInstance struct {
	Position tempest.C4Vector
	Group    models.SpawnGroup
	Objects  WorldObjectSet
}

func (m *Map) spawnObject(sp *SpawnGroupInstance, so *models.SpawnObject) {
	if !randomness.RollPercent(so.Chance) {
		return
	}

	realPosition := sp.Position
	realPosition.X += so.RelPosition.X
	realPosition.Y += so.RelPosition.Y
	realPosition.Z += so.RelPosition.Z
	realPosition.W = so.RelPosition.W

	var object WorldObject

	switch so.Kind {
	case models.SpawnKindCreature:
		var creatureInfo *models.CreatureTemplate
		m.Phase.Server.DB.Lookup(wdb.BucketKeyStringID, so.ID, &creatureInfo)
		creature := m.Phase.Server.NewCreature(creatureInfo, realPosition)
		creature.Group = sp
		creature.Map = m
		object = creature
	case models.SpawnKindGameObject:
		var gameObjectInfo *models.GameObjectTemplate
		m.Phase.Server.DB.Lookup(wdb.BucketKeyStringID, so.ID, &gameObjectInfo)
		gameObject := m.Phase.Server.NewGameObject(gameObjectInfo, realPosition)
		gameObject.Group = sp
		gameObject.Map = m
		object = gameObject
	default:
		return
	}

	m.addObject(object)
}

func (m *Map) spawnGroup(sp *SpawnGroupInstance) {
	for _, object := range sp.Group.Objects {
		m.spawnObject(sp, &object)
	}
}

// Loads a region of the map, with all creatures and gameobjects inside.
// This gets triggered when a player first enters a zone.
func (m *Map) loadContentEnvirons(center terrain.BlockIndex) {
	ourPhase := m.Phase.ID
	ourMap := m.ID

	param := m.terrainMapParam()

	radius := m.VisibilityDistance()

	blockExtent := int32(radius / terrain.BlockSize)

	if blockExtent < 1 {
		blockExtent = 1
	}

	minBlock, maxBlock, err := m.calcBlockRange(center, blockExtent)
	if err != nil {
		panic(err)
	}

	type tmpMapBlock struct {
		block  *MapBlock
		groups []*SpawnGroupInstance
	}

	var blocksNeedingContentLoad = map[terrain.BlockIndex]*tmpMapBlock{}

	var blockCursor terrain.BlockIndex
	for blockCursor.X = minBlock.X; blockCursor.X <= maxBlock.X; blockCursor.X++ {
		for blockCursor.Y = minBlock.Y; blockCursor.Y <= maxBlock.Y; blockCursor.Y++ {
			index1d := blockIndex1d(blockCursor, param)
			block := m.blocks[index1d]
			if block == nil {
				block = m.initMapBlock(blockCursor)
			}
			if !block.contentLoaded {
				blocksNeedingContentLoad[blockCursor] = &tmpMapBlock{
					block: block,
				}
			}
		}
	}

	// Skip this expensive querying business
	if len(blocksNeedingContentLoad) == 0 {
		return
	}

	// All objects to be spawned within this square of selected mapblocks
	// We are collecting a veritable heap of objects to use.
	m.Phase.Server.DB.Range(func(sp *models.SpawnGroup) bool {
		inOurPhase := false

		inOurMap := sp.Map == ourMap

		if !inOurMap {
			return true
		}

		for _, phase := range sp.Phase {
			if phase == ourPhase {
				inOurPhase = true
				break
			}
		}

		if !inOurPhase {
			return true
		}

		weights := make([]int, len(sp.Positions))
		for k, v := range sp.Positions {
			weights[k] = v.Weight
		}

		selection := randomness.WeightedSelect(weights)

		pos := sp.Positions[selection]

		// See if we even care about this position.
		blockIndex, err := terrain.CalcBlockIndex(param, pos.Position.C2())
		if err != nil {
			panic(err)
		}

		block, doWeCare := blocksNeedingContentLoad[blockIndex]
		if !doWeCare {
			// This is in a completely different location than what we care about.
			return true
		}

		// if pos.Position.X >= minPos.X && pos.Position.Y >= minPos.X && pos.Position.X < maxPos.X

		spawnThisGroup := randomness.RollPercent(sp.Chance)
		if !spawnThisGroup {
			return true
		}

		spawnInstance := new(SpawnGroupInstance)
		spawnInstance.Group = *sp
		spawnInstance.Position = sp.Positions[selection].Position

		block.groups = append(block.groups, spawnInstance)
		return true
	})

	// Now, the trick here is to add the objects without blocking the Map routine and causing problems.

	for _, tmpBlock := range blocksNeedingContentLoad {
		// Mark content as loaded, otherwise this will become recursive
		// as spawnGroup can call addObject->updatePosition->loadContentEnvirons
		tmpBlock.block.contentLoaded = true

		for _, group := range tmpBlock.groups {
			m.spawnGroup(group)
		}
	}
}
