package realm

import (
	"github.com/Gophercraft/core/format/terrain"
	"github.com/Gophercraft/core/tempest"
)

func blockIndex1d(blockIndex terrain.BlockIndex, param *terrain.MapParam) int {
	return int((blockIndex.X * param.BlockSize.X) + blockIndex.Y)
}

type MapBlock struct {
	objects WorldObjectSet
	// todo: map block will have maps, vmaps associated
	contentLoaded bool
}

// important: avoid calling emplaceobject more than once.
func (m *MapBlock) emplaceobject(wo WorldObject) {
	// id  := wo.GUID()

	for _, o := range m.objects {
		if o == wo {
			panic("duplicate emplace")
		}
	}

	m.objects = append(m.objects, wo)
}

func (m *MapBlock) removeobject(wo WorldObject) bool {
	index := -1

	for i, object := range m.objects {
		if object == wo {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	m.objects = append(m.objects[:index], m.objects[index+1:]...)
	return true
}

func (m *Map) initMapBlock(blockIndex terrain.BlockIndex) *MapBlock {
	param := m.terrainMapParam()
	mb := new(MapBlock)
	m.blocks[blockIndex1d(blockIndex, param)] = mb
	return mb
}

func (m *Map) getBlockByPosition(pos tempest.C2Vector) *MapBlock {
	param := m.terrainMapParam()
	blockIndex, err := terrain.CalcBlockIndex(param, pos)
	// panic because the position value should already have been validated
	if err != nil {
		panic(err)
	}

	index1d := blockIndex1d(blockIndex, param)

	if len(m.blocks) > int(index1d) {
		return m.blocks[index1d]
	}

	return nil
}
