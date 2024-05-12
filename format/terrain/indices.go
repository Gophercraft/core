package terrain

import (
	"fmt"

	"github.com/Gophercraft/core/tempest"
	"github.com/arl/math32"
)

// Since documentation is often lacking or contradictory, a definition of terms:
// Map: a continuum of world geometry data. A 2 dimensional model storing data. Typically consisting of 64x64 blocks. (in the future, with modded clients, we hope to be able to push this to an even greater size)
// Block: 16x16 list within a Map. Confusingly, some (Noggit and others) would also refer to this as a "tile".
// Chunks: The little bits of map geometry that comprise a Block

const (
	BlockSize = 533.33333
	ChunkSize = BlockSize / 16
)

type BlockIndex tempest.C2iVector
type ChunkIndex tempest.C2iVector

type MapParam struct {
	BlockSize tempest.C2iVector
}

var (
	DefaultMap = MapParam{
		BlockSize: tempest.C2iVector{
			X: 64,
			Y: 64,
		},
	}
)

// Using map parameters, calculate what the block index of a position is. Will error if pos is out of bounds.
//
//		index, _ := terrain.CalcBlockIndex(&terrain.DefaultMap, tempest.C2Vector{
//				X:
//	})
func CalcBlockIndex(m *MapParam, pos tempest.C2Vector) (bi BlockIndex, err error) {
	var (
		// Suppose BlockSizes are 64 (standard)
		// mapWidth = 64 * 533.33333 = 34133.33312
		// mapHeight = 64 * 533.33333 = 34133.33312
		mapWidth  = float32(m.BlockSize.X) * BlockSize
		mapLength = float32(m.BlockSize.Y) * BlockSize

		max = tempest.C2Vector{
			// X = 34133.33312 / 2 = 17066.66656
			X: mapWidth / 2,
			// Y = 34133.33312 / 2 = 17066.66656
			Y: mapLength / 2,
		}

		min = tempest.C2Vector{
			X: 0 - max.X,
			Y: 0 - max.Y,
		}
	)

	if pos.X >= max.X || pos.Y >= max.Y {
		err = fmt.Errorf("terrain: position is out of bounds pos = %s, max = %s", pos, max)
		return
	}

	if pos.Y <= min.X || pos.Y <= min.Y {
		err = fmt.Errorf("terrain: position is out of bounds pos = %s, min = %s", pos, min)
		return
	}

	blockZeroPoint := tempest.C2Vector{
		// 64 / 2 = 32
		X: float32(m.BlockSize.X) / 2,
		Y: float32(m.BlockSize.Y) / 2,
	}

	xBlock := math32.Floor(blockZeroPoint.X - (pos.X / BlockSize))
	yBlock := math32.Floor(blockZeroPoint.Y - (pos.Y / BlockSize))

	bi.X = int32(xBlock)
	bi.Y = int32(yBlock)

	// // Client coordinats
	// // now, rangePos.X is now between 0 and mapWidth
	// clientCoordinates := tempest.C2Vector{
	// 	X: pos.X + max.X,
	// 	Y: pos.Y + max.Y,
	// }

	// widthFraction := rangePos.X / mapWidth
	// lengthFraction := rangePos.Y / mapLength

	// bi.X = int32(widthFraction * float32(m.BlockSize.X))
	// bi.Y = int32(lengthFraction * float32(m.BlockSize.Y))

	return
}

// TODO: invalid fix
// Using a BlockIndex, return the in-world XY position of the corner of this.
// Note that this does not return the CENTER of the block, only the lowest corner. For instance, 32,32 -> 0,0
func CalcBlockCornerPos(m *MapParam, bi BlockIndex) (pos tempest.C2Vector, err error) {
	if bi.X < 0 || bi.X >= m.BlockSize.X {
		err = fmt.Errorf("terrain: CalcBlockCornerPos: X value %d out of bounds 0-%d", bi.X, m.BlockSize.X)
		return
	}

	if bi.Y < 0 || bi.Y >= m.BlockSize.Y {
		err = fmt.Errorf("terrain: CalcBlockCornerPos: Y value %d out of bounds 0-%d", bi.Y, m.BlockSize.Y)
		return
	}

	var (
		// 64 * 533.33333 = 34133.33312
		mapWidth  = float32(m.BlockSize.X) * BlockSize
		mapLength = float32(m.BlockSize.Y) * BlockSize
	)

	relativeOffset := tempest.C2Vector{
		// 34133.33312 / 2 = 17066.66656
		X: mapWidth / 2,
		Y: mapLength / 2,
	}

	pos = tempest.C2Vector{
		X: relativeOffset.X - (float32(bi.X) * BlockSize),
		Y: relativeOffset.Y - (float32(bi.Y) * BlockSize),
	}

	return
}

// Using 2d position vector, determine what chunk slot this position falls into.
// Pos must be a valid position before calling. Call CalcBlockIndex to be sure.
func CalcChunkIndex(m *MapParam, pos tempest.C2Vector) (ci ChunkIndex, err error) {
	var (
		mapWidth  = float32(m.BlockSize.X) * BlockSize
		mapLength = float32(m.BlockSize.Y) * BlockSize
	)

	clientCoordinates := tempest.C2Vector{
		X: pos.X + (mapWidth / 2),
		Y: pos.Y + (mapLength / 2),
	}

	blockRelPos := tempest.C2Vector{
		X: math32.Mod(clientCoordinates.X, BlockSize),
		Y: math32.Mod(clientCoordinates.Y, BlockSize),
	}

	widthFraction := blockRelPos.X / BlockSize
	lengthFraction := blockRelPos.Y / BlockSize

	ci.X = int32(widthFraction * 16)
	ci.Y = int32(lengthFraction * 16)

	return
}
