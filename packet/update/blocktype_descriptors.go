package update

import "github.com/Gophercraft/core/vsn"

type BlockTypeDescriptor map[BlockType]uint8

var BlockTypeDescriptors = map[vsn.BuildRange]BlockTypeDescriptor{
	{0, 3368}: {
		Values:            0,
		Movement:          1,
		CreateObject:      2,
		SpawnObject:       2,
		DeleteFarObjects:  3,
		DeleteNearObjects: 4,
	},

	{5875, 6005}: {
		Values:            0,
		Movement:          1,
		CreateObject:      2,
		SpawnObject:       3,
		DeleteFarObjects:  4,
		DeleteNearObjects: 5,
	},

	{8606, vsn.Max}: {
		Values:            0,
		Movement:          1,
		CreateObject:      2,
		SpawnObject:       3,
		DeleteFarObjects:  4,
		DeleteNearObjects: 5,
	},
}
