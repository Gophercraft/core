package update

import (
	"fmt"
	"io"
	"reflect"
	"sync"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

var (
	MaxBlockCount uint32 = 2048
)

// ValuesBlock acts a wrapper container for reflection-based update field storage and serialization
type ValuesBlock struct {
	sync.Mutex

	TypeMask guid.TypeMask
	// Descriptor describes the exact structure of a particular version's SMSG_UPDATE_OBJECT.
	Descriptor *Descriptor
	// Points to chunks that have been updated
	ChangeMask *Bitmask
	// StorageDescriptor is the version and type-specific structure. Descriptors can be found in packages labelled d<build number>
	// Write your own descriptors to support a lesser known version if there is a change in SMSG_UPDATE_OBJECT
	StorageDescriptor reflect.Value
}

// ValuesBlock data is copied into a new region of memory where it can be changed without ruining the original, at the cost of GC
func (v *ValuesBlock) Copy() *ValuesBlock {
	nv := &ValuesBlock{}
	nv.TypeMask = v.TypeMask
	nv.Descriptor = v.Descriptor
	// Copy changemask bits
	cm := v.ChangeMask.Copy()
	nv.ChangeMask = &cm
	nv.StorageDescriptor = reflect.New(nv.StorageDescriptor.Type().Elem())
	// Copy big update fields into a different mem region
	nv.StorageDescriptor.Elem().Set(nv.StorageDescriptor.Elem())
	return nv
}

// Decoder decodes an SMSG_UPDATE_OBJECT input stream into various sub-structures.
type Decoder struct {
	*Descriptor
	Build                  version.Build
	HasTransport           bool
	Map                    uint16
	SmoothDeleteStartIndex uint16
	BlockCount             uint32
	CurrentBlockIndex      uint32
	CurrentBlockType       BlockType
	Reader                 io.Reader
}

// Start chewing up blocks.
func NewDecoder(version version.Build, reader io.Reader) (*Decoder, error) {
	decoder := new(Decoder)
	decoder.Build = version
	decoder.Reader = reader

	err := version.QueryDescriptors(version, Descriptors, &decoder.Descriptor)
	if err != nil {
		return nil, err
	}
	decoder.BlockCount, err = readUint32(reader)
	if err != nil {
		return nil, err
	}

	if decoder.DescriptorOptions&DescriptorOptionHasHasTransport != 0 {
		decoder.HasTransport, err = readBool(reader)
		if err != nil {
			return nil, err
		}
	}

	if decoder.BlockCount > MaxBlockCount {
		return nil, fmt.Errorf("update: block count (%d) > update.MaxBlockCount (%d)", decoder.BlockCount, MaxBlockCount)
	}

	return decoder, nil
}

// DecodeBlockType decodes the BlockType. This function tries to resolve differences between protocol revisions.
func (decoder *Decoder) DecodeBlockType() (BlockType, error) {
	bt, err := readUint8(decoder.Reader)
	if err != nil {
		return 0, err
	}

	var desc BlockTypeDescriptor
	if err := version.QueryDescriptors(decoder.Build, BlockTypeDescriptors, &desc); err != nil {
		return 0, err
	}

	for typeSpec, val := range desc {
		if typeSpec == SpawnObject && val == bt {
			if desc[CreateObject] == bt {
				return CreateObject, nil
			}
		}

		if val == bt {
			return typeSpec, nil
		}
	}

	// todo: support new format

	return BlockType(bt), nil
}

// if true, more blocks lie ahead for us to eat
func (decoder *Decoder) MoreBlocks() bool {
	if decoder.CurrentBlockIndex >= decoder.BlockCount {
		return false
	}

	return true
}

// What is this next block tagged as? It tells us which data to decode after this
func (decoder *Decoder) NextBlock() (BlockType, error) {
	if !decoder.MoreBlocks() {
		return 0, io.EOF
	}

	var err error
	decoder.CurrentBlockType, err = decoder.DecodeBlockType()
	if err != nil {
		return 0, err
	}

	decoder.CurrentBlockIndex++

	return decoder.CurrentBlockType, nil
}
