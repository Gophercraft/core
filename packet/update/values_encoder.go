package update

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/log"
)

type VisibilityFlags uint32

const (
	None        VisibilityFlags = 0
	Owner       VisibilityFlags = 0x01
	PartyMember VisibilityFlags = 0x02
	UnitAll     VisibilityFlags = 0x04
	Empath      VisibilityFlags = 0x08
)

type ValuesEncoder struct {
	Encoder        *Encoder
	ViewMask       VisibilityFlags
	Create         bool
	ValuesBlock    *ValuesBlock
	CurrentBitmask *Bitmask
	ChunkPos       uint32
	BitPos         uint32
	NextChunk      [4]byte
	// WritePos: the current chunk position being recorded to the stream.
	WritePos uint32
}

func (v *ValuesEncoder) Write(b []byte) (int, error) {
	v.WritePos += uint32(len(b))
	return v.Encoder.Write(b)
}

func (valenc *ValuesEncoder) SetCreateBits() error {
	value := valenc.ValuesBlock.StorageDescriptor.Elem()

	if err := valenc.setCreateBitsFor(value, ""); err != nil {
		return err
	}

	// panic(valenc.CurrentBitmask)

	valenc.BitPos = 0
	valenc.ChunkPos = 0
	valenc.clearNextChunk()

	return nil
}

func (valenc *ValuesEncoder) clearNextChunk() {
	valenc.NextChunk[0] = 0
	valenc.NextChunk[1] = 0
	valenc.NextChunk[2] = 0
	valenc.NextChunk[3] = 0
}

func (valenc *ValuesEncoder) includeValue(tag FieldTag) bool {
	if tag.IsPrivate() {
		if valenc.ViewMask&Owner != 0 {
			return true
		}
		return false
	}

	if tag.IsParty() {
		if valenc.ViewMask&PartyMember != 0 {
			return true
		}
		return false
	}

	return true
}

// this function is purely for setting the bitmask for created objects in the legacy protocol.
func (valenc *ValuesEncoder) setCreateBitsFor(value reflect.Value, tag FieldTag) error {
	switch value.Type() {
	case bitPadType:
		nxtBit(&valenc.ChunkPos, &valenc.BitPos)
		return nil
	case bytePadType:
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
		return nil
	case chunkPadType:
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
		return nil
	case alignPadType:
		valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
		return nil
	case guidType:
		if valenc.includeValue(tag) {
			id := value.Interface().(guid.GUID)
			var bytes [8]byte
			binary.LittleEndian.PutUint64(bytes[:], id.Classic())
			if u32(bytes[:4]) != 0 {
				valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
			}

			if u32(bytes[4:]) != 0 {
				valenc.CurrentBitmask.Set(valenc.ChunkPos+1, true)
			}
		}

		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)

		return nil
	}

	switch value.Kind() {
	case reflect.Struct:
		for x := 0; x < value.NumField(); x++ {
			if err := valenc.setCreateBitsFor(value.Field(x), FieldTag(value.Type().Field(x).Tag.Get("update"))); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		for x := 0; x < value.Len(); x++ {
			if err := valenc.setCreateBitsFor(value.Index(x), tag); err != nil {
				return err
			}
		}
	case reflect.Uint64:
		if valenc.includeValue(tag) {
			var bytes [8]byte
			binary.LittleEndian.PutUint64(bytes[:], value.Uint())
			if u32(bytes[:4]) != 0 {
				valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
			}

			if u32(bytes[4:]) != 0 {
				valenc.CurrentBitmask.Set(valenc.ChunkPos+1, true)
			}
		}
		// valenc.ChunkPos += 2
		// valenc.BitPos = 0
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Uint32:
		if value.Uint() != 0 && valenc.includeValue(tag) {
			valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Int32:
		if value.Int() != 0 && valenc.includeValue(tag) {
			valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Float32:
		if value.Float() != 0 && valenc.includeValue(tag) {
			valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Bool:
		if valenc.includeValue(tag) {
			if value.Bool() {
				valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
			}
		}
		nxtBit(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Uint8:
		if valenc.includeValue(tag) {
			if value.Uint() != 0 {
				valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
			}
		}

		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Uint16:
		if valenc.includeValue(tag) {
			if value.Uint() != 0 {
				valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
			}
		}
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Int16:
		if valenc.includeValue(tag) {
			if value.Int() != 0 {
				valenc.CurrentBitmask.Set(valenc.ChunkPos, true)
			}
		}
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
	default:
		return fmt.Errorf("update: unhandled type detected while trying to write creation bitmask: %s", value.Type())
	}

	return nil
}

func (valenc *ValuesEncoder) EncodeValue(value reflect.Value, name string, tag FieldTag) error {
	quit := true

	beginChunk := valenc.ChunkPos
	beginWritePos := valenc.WritePos
	chunkAdvanced := false

	switch value.Type() {
	case guidType:
		var bytes [8]byte
		binary.LittleEndian.PutUint64(bytes[:], value.Interface().(guid.GUID).Classic())

		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			if _, err := valenc.Write(bytes[:4]); err != nil {
				return err
			}
		}

		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos + 1) {
			if _, err := valenc.Write(bytes[4:]); err != nil {
				return err
			}
		}

		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case chunkPadType:
		terminatedMultiFieldChunk := false

		// Chunk pads can be used to terminate a multi-field chunk
		if valenc.BitPos != 0 {
			if valenc.CurrentBitmask.Enabled(beginChunk) {
				valenc.Write(valenc.NextChunk[:])
			}
			valenc.clearNextChunk()

			nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
			terminatedMultiFieldChunk = true
		}

		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			return fmt.Errorf("update: chunk %s, a padding chunk, has had its change bit set to true (chunk %d, bit %d)", name, valenc.ChunkPos, valenc.BitPos)
		}

		if !terminatedMultiFieldChunk {
			nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
		}
	case alignPadType:
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			if err := writeUint32(valenc, 0x00000000); err != nil {
				return err
			}
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case bitPadType:
		nxtBit(&valenc.ChunkPos, &valenc.BitPos)
	case bytePadType:
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
	default:
		quit = false
	}

	if quit {
		return nil
	}

	switch value.Kind() {
	case reflect.Uint16:
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			binary.LittleEndian.PutUint16(valenc.NextChunk[valenc.BitPos/8:], uint16(value.Uint()))
		}
		// valenc.BitPos += 16
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Int16:
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			binary.LittleEndian.PutUint16(valenc.NextChunk[valenc.BitPos/8:], uint16(value.Int()))
		}
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Bool:
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			// fmt.Printf("%s 0x%08X\n", name, (1 << valenc.BitPos))
			bytePos := valenc.BitPos / 8
			bitFlag := uint8(1 << (valenc.BitPos % 8))
			if value.Bool() {
				valenc.NextChunk[bytePos] |= bitFlag
			} else {
				valenc.NextChunk[bytePos] &= ^(bitFlag)
			}
		}
		// valenc.BitPos++
		nxtBit(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Uint8:
		if valenc.BitPos%8 != 0 {
			valenc.BitPos += (8 - valenc.BitPos%8)
		}
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			valenc.NextChunk[valenc.BitPos/8] = uint8(value.Uint())
		}
		// valenc.BitPos += 8
		nxtByte(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Array:
		for x := 0; x < value.Len(); x++ {
			name := fmt.Sprintf("%s[%d]", name, x)
			if err := valenc.EncodeValue(value.Index(x), name, tag); err != nil {
				return err
			}
		}
	case reflect.Struct:
		for x := 0; x < value.NumField(); x++ {
			offsetTag := value.Type().Field(x).Tag.Get("offset")

			nm := name + "." + value.Type().Field(x).Name

			if offsetTag != "" {
				numbersS := strings.SplitN(offsetTag, ",", 2)

				var ofs [2]uint32

				for i := 0; i < 2; i++ {
					ui, err := strconv.ParseUint(numbersS[i], 0, 32)
					if err != nil {
						return err
					}

					ofs[i] = uint32(ui)
				}

				if valenc.ChunkPos != ofs[0] {
					return fmt.Errorf("update: EncodeValue failed to calculate the correct chunk offset. %s should be %d, is %d", nm, ofs[0], valenc.ChunkPos)
				}

				if valenc.BitPos != ofs[1] {
					return fmt.Errorf("update: EncodeValue failed to calculate the correct bit offset. %s should be %d, is %d", nm, ofs[1], valenc.BitPos)
				}
			}

			if err := valenc.EncodeValue(value.Field(x), nm, FieldTag(value.Type().Field(x).Tag.Get("update"))); err != nil {
				return err
			}
		}
	case reflect.Uint64:
		var bytes [8]byte
		binary.LittleEndian.PutUint64(bytes[:], value.Uint())

		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			if _, err := valenc.Write(bytes[0:4]); err != nil {
				return err
			}
		}

		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos + 1) {
			if _, err := valenc.Write(bytes[4:8]); err != nil {
				return err
			}
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Uint32:
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			if err := writeUint32(valenc, uint32(value.Uint())); err != nil {
				return err
			}
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Int32:
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			if err := writeInt32(valenc, int32(value.Int())); err != nil {
				return err
			}
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	case reflect.Float32:
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			if err := writeFloat32(valenc, float32(value.Float())); err != nil {
				return err
			}
		}
		nxtChunk(&valenc.ChunkPos, &valenc.BitPos)
	default:
		return fmt.Errorf("update: unhandled attempt to encode %s kind %s", value.Type(), value.Kind())
	}

	chunkAdvanced = beginChunk < valenc.ChunkPos

	if chunkAdvanced {
		if isSubChunkType(value) {
			if valenc.CurrentBitmask.Enabled(beginChunk) {
				valenc.Write(valenc.NextChunk[:])
			}
			valenc.clearNextChunk()
		} else {
			// Check that the data has already been written.
			if valenc.CurrentBitmask.Enabled(beginChunk) {
				if valenc.WritePos == beginWritePos {
					return fmt.Errorf("update: bitmask enabled but no corresponding chunk data was emitted %s", name)
				}
			}
		}
	}

	return nil
}

func (ve *ValuesEncoder) ChunkCount(mask *Bitmask) int {
	var (
		chunkCount int
		slotCount  int = mask.Len() * 32
	)

	for s := 0; s < slotCount; s++ {
		if mask.Enabled(uint32(s)) {
			chunkCount++
		}
	}

	return chunkCount
}

func (valuesBlock *ValuesBlock) WriteData(e *Encoder, viewMask VisibilityFlags, create bool) error {
	valenc := &ValuesEncoder{
		Encoder:     e,
		Create:      create,
		ValuesBlock: valuesBlock,
		ViewMask:    viewMask,
	}

	if !valuesBlock.StorageDescriptor.IsValid() {
		return fmt.Errorf("update: cannot encode empty ValuesBlock.StorageDescriptor")
	}

	if valenc.Create {
		// TODO: in the future, all fields are included in the create block (without a bitmask)

		// All non-zero and non-private chunks will be included.
		valenc.CurrentBitmask = NewBitmask()
		if err := valenc.SetCreateBits(); err != nil {
			return err
		}
	} else {
		// All fields enabled in the change mask will be included, even zero and private ones.
		valenc.CurrentBitmask = valuesBlock.ChangeMask
	}

	// TODO: in the future, bitmasks are stored within the descriptor's structs.
	// valenc.CurrentBitmask will be an alias for these.

	// Save how many chunks are supposed to be emitted after this mask
	chunkCount := valenc.ChunkCount(valenc.CurrentBitmask)

	// Write uint8 len + uint32[len]
	if err := WriteBitmask(valenc.CurrentBitmask, e.Descriptor, e); err != nil {
		return err
	}

	// Write uint32 chunks, which follow every true bit in the change bitmask
	if err := valenc.EncodeValue(valuesBlock.StorageDescriptor.Elem(), valuesBlock.StorageDescriptor.Type().String(), ""); err != nil {
		return err
	}

	// Check if anything is left over.
	// This could happen if a subchunk is left unflushed at the very end of a descriptor
	if valenc.BitPos > 0 {
		if valenc.CurrentBitmask.Enabled(valenc.ChunkPos) {
			valenc.Write(valenc.NextChunk[:])
		}
		valenc.clearNextChunk()
		valenc.BitPos = 0
		valenc.ChunkPos++
	}

	// Detect stream misalignment
	if valenc.WritePos != (uint32(chunkCount) * 4) {
		// Something is fucked up in EncodeValue, not writing the correct amount of chunks
		log.Dump("Bitmask", valenc.CurrentBitmask)

		return fmt.Errorf("update: ValuesBlock.WriteData: mismatch between amount of chunks written to buffer and amount of chunks present in bitmask (Chunks written: %d (%d bytes), chunks masked: %d (%d bytes)", valenc.WritePos/4, valenc.WritePos, chunkCount, chunkCount*4)
	}

	return nil
}
