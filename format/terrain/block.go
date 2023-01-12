package terrain

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Gophercraft/core/format/chunked"
	"github.com/Gophercraft/core/tempest"
)

type MapBlock struct {
	// MVER
	Version uint32

	// MHDR
	Flags uint32

	// MCIN
	ChunkInfo [16 * 16]ChunkInfo

	// MTEX
	Textures []string

	// MMDX
	MapModels []byte // m2 files

	// MMID
	MapModelFilenameOffsets []uint32

	// MWMO
	MapObjects []byte // wmo files

	// MWID
	MapObjectFilenameOffsets []uint32

	// MDDF
	DoodadDefs []DoodadDef

	// MODF
	MapObjectDefs []WMODef

	// MCNK * 16 * 16
	ChunkData [16][16]*ChunkData
}

// ChunkInfo Pointers to MCNK chunks and their sizes.
type ChunkInfo struct {
	Offset uint32
	Size   uint32
	Flags  uint32
	Pad    uint32
}

// DoodadDef Placement information for doodads (M2 models). Additional to this, the models to render are referenced in each MCRF chunk.
type DoodadDef struct {
	// NameID references an entry in the MMID chunk, specifying the model to use.
	NameID   uint32 // 0x00
	UniqueID uint32 // 0x04

	Position tempest.C3Vector // 0x08
	Rotation tempest.C3Vector // 0x14

	Scale uint16 // 0x20
	Flags uint16 // 0x22
}

type WMODef struct {
	NameID   uint32
	UniqueID uint32

	Position tempest.C3Vector
	Rotation tempest.C3Vector

	Extent    tempest.CAaBox
	Flags     uint16
	DoodadSet uint16
	NameSet   uint16
	Scale     uint16
}

func getStringList(chunk []byte) []string {
	out := []string{}
	strs := bytes.Split(chunk, []byte{0})
	for _, str := range strs {
		if len(str) > 0 {
			out = append(out, string(str))
		}
	}
	return out
}

func (mr *MapReader) ReadBlock(x, y int) (*MapBlock, error) {
	block := new(MapBlock)
	blockFile, err := mr.ReadFile(fmt.Sprintf("World/Maps/%s/%s_%d_%d.adt", mr.Name, mr.Name, x, y))
	if err != nil {
		return nil, err
	}

	chunkIndex := 0

	defer blockFile.Close()

	cr := &chunked.Reader{blockFile}

	for {
		id, chunk, err := cr.ReadChunk()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		switch id {
		case MapVersion:
			block.Version = binary.LittleEndian.Uint32(chunk[:4])
		case BlockHeader:
			block.Flags = binary.LittleEndian.Uint32(chunk[:4])
		case MapChunkInfo:
			for x := 0; x < len(block.ChunkInfo); x++ {
				infoBase := x * 16
				mcin := ChunkInfo{
					Offset: binary.LittleEndian.Uint32(chunk[:infoBase+4]),
					Size:   binary.LittleEndian.Uint32(chunk[infoBase+4 : infoBase+8]),
					Flags:  binary.LittleEndian.Uint32(chunk[infoBase+8 : infoBase+12]),
					Pad:    binary.LittleEndian.Uint32(chunk[infoBase+12 : infoBase+16]),
				}
				block.ChunkInfo[x] = mcin
			}
		case MapTextures:
			block.Textures = getStringList(chunk)
		case MapModels:
			block.MapModels = chunk
			// ?
		case MapModelFilenamesOffsets:
			block.MapModelFilenameOffsets = make([]uint32, len(chunk)/4)
			for x := 0; x < len(block.MapModelFilenameOffsets); x++ {
				block.MapModelFilenameOffsets[x] = binary.LittleEndian.Uint32(chunk[4*x : (4*x)+4])
			}
		case MapObject:
			block.MapObjects = chunk
			// ?
		case MapObjectFilenameOffsets:
			block.MapObjectFilenameOffsets = make([]uint32, len(chunk)/4)
			for x := 0; x < len(block.MapObjectFilenameOffsets); x++ {
				block.MapObjectFilenameOffsets[x] = binary.LittleEndian.Uint32(chunk[4*x : (4*x)+4])
			}
		case DoodadDefs:
			chunkReader := bytes.NewReader(chunk)

			sizeof := 36

			numDefs := len(chunk) / sizeof

			block.DoodadDefs = make([]DoodadDef, numDefs)

			for x := 0; x < numDefs; x++ {
				err := binary.Read(chunkReader, binary.LittleEndian, &block.DoodadDefs[x])
				if err != nil {
					return nil, err
				}
			}
		case MapObjectDefs:
			chunkReader := bytes.NewReader(chunk)

			sizeof := 64

			numDefs := len(chunk) / sizeof

			block.MapObjectDefs = make([]WMODef, numDefs)

			for x := 0; x < numDefs; x++ {
				err := binary.Read(chunkReader, binary.LittleEndian, &block.MapObjectDefs[x])
				if err != nil {
					return nil, err
				}
			}
		case Chunk:
			cReader := bytes.NewReader(chunk)

			// Read header
			var cd ChunkData
			err := binary.Read(cReader, binary.LittleEndian, &cd.ChunkHeader)
			if err != nil {
				return nil, err
			}

			// Chunk (MCNK) contains several chunked sub-messages
			// subChunkReader := &chunked.Reader{cReader}

			var subchunkID chunked.Tag
			var subchunk []byte
			for {
				err = binary.Read(cReader, binary.LittleEndian, &subchunkID)
				if err != nil {
					return nil, err
				}

				if subchunkID == chunked.Nil {
					break
				}

				var subchunksize uint32

				// Strange chunk format, size is not included
				err = binary.Read(cReader, binary.LittleEndian, &subchunksize)
				if err != nil {
					return nil, fmt.Errorf("terrain: while reading subchunk size for %s: %s", subchunkID, err)
				}

				if subchunkID == Liquids {
					subchunksize = cd.SizeLiquid - 8
				} else if subchunkID == Alpha {
					subchunksize = cd.SizeAlpha - 8
				}

				if subchunksize == 0 {
					continue
				}

				// if subchunkID == Alpha && cd.SizeAlpha != 8 {
				// 	continue
				// }

				subchunk = make([]byte, subchunksize)

				subchunkBytes, err := cReader.Read(subchunk[:])
				if err != nil {
					return nil, fmt.Errorf("terrain: while reading %d bytes of subchunk data for %s: %s", len(subchunk), subchunkID, err)
				}

				if subchunkBytes != len(subchunk) {
					return nil, fmt.Errorf("terrain: MCNK subchunk '%s' only had %d/%d bytes", subchunkID, subchunkBytes, len(subchunk))
				}

				switch subchunkID {
				case Normals:
					scr := bytes.NewReader(subchunk)
					err = binary.Read(scr, binary.LittleEndian, &cd.Normals)
					if err != nil {
						return nil, err
					}

					// For some unknown reason, there is 13 bytes of extraneous data outside of the chunk stream following MCNR
					cReader.Read(cd.NormalsPad[:])
				case Heights:
					if err = binary.Read(bytes.NewReader(subchunk), binary.LittleEndian, &cd.Heights); err != nil {
						return nil, err
					}
				case Layer:
					cd.Layer = make([]ChunkLayer, len(subchunk)/16)
					if err = binary.Read(bytes.NewReader(subchunk), binary.LittleEndian, &cd.Layer); err != nil {
						return nil, err
					}
				case CollisionObjects:
					scr := bytes.NewReader(subchunk)
					cd.CollisionDoodads = make([]uint32, cd.NumDoodadRefs)
					if err = binary.Read(scr, binary.LittleEndian, &cd.CollisionDoodads); err != nil {
						return nil, err
					}
					cd.CollisionWMOs = make([]uint32, cd.NumMapObjRefs)
					if err = binary.Read(scr, binary.LittleEndian, &cd.CollisionWMOs); err != nil {
						return nil, err
					}
				case ShadowMap:
					if err = binary.Read(bytes.NewReader(subchunk), binary.LittleEndian, &cd.ShadowMap); err != nil {
						return nil, err
					}
				case Alpha:
					numAlphaLayers := len(cd.Layer) - 1

					if numAlphaLayers <= 0 {
						break
						// return nil, fmt.Errorf("terrain: alpha subchunk has no layers? %d", numAlphaLayers)
					}

					cd.AlphaMaps = make([]ChunkAlphaMap, numAlphaLayers)

					for layer := 0; layer < numAlphaLayers; layer++ {
						if mr.Index.HeaderFlags&ADTHasBigAlpha != 0 || cd.Layer[layer+1].Flags&0x200 != 0 {
							copy(cd.AlphaMaps[layer][:], subchunk[:4096])
						} else {
							for x := 0; x < 2048; x++ {
								byte := subchunk[x]
								cd.AlphaMaps[layer][x*2] = byte & 0xF
								cd.AlphaMaps[layer][(x*2)+1] = byte * 0x10
							}
						}
					}
				case Liquids:
					if len(subchunk) == 0 {
						continue
					}
					if err = binary.Read(bytes.NewReader(subchunk), binary.LittleEndian, &cd.Liquids); err != nil {
						return nil, fmt.Errorf("terrain: while reading liquids %d: %s", len(subchunk), err)
					}
				case SoundEmitters:
					cd.OldSoundEmitters = make([]ChunkOldSoundEmitter, len(subchunk)/52)
					// TODO: handle new format
					if err = binary.Read(bytes.NewReader(subchunk), binary.LittleEndian, &cd.OldSoundEmitters); err != nil {
						return nil, err
					}
				default:
					return nil, fmt.Errorf("terrain: unhandled subchunk in MCNK: %s", subchunkID)
				}
			}

			xindex := chunkIndex / 16
			yindex := chunkIndex % 16

			block.ChunkData[xindex][yindex] = &cd
			chunkIndex++
		default:
			return nil, fmt.Errorf("terrain: unhandled chunk in ADT file: %s", id)
		}
	}

	return block, nil
}
