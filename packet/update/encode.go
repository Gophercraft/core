package update

import (
	"fmt"
	"io"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

//go:generate gcraft_stringer -type=BlockType
type BlockType int

const (
	Values BlockType = iota
	Movement
	CreateObject
	SpawnObject
	DeleteFarObjects
	DeleteNearObjects
)

type BlockData interface {
	Type() BlockType
	WriteData(*Encoder, VisibilityFlags, bool) error
}

type Encoder struct {
	io.Writer
	Build      vsn.Build
	Descriptor *Descriptor
	blockTypes BlockTypeDescriptor
}

func NewEncoder(version vsn.Build, writer io.Writer, numBlocks int) (*Encoder, error) {
	e := &Encoder{
		Writer: writer,
		Build:  version,
	}
	err := vsn.QueryDescriptors(version, Descriptors, &e.Descriptor)
	if err != nil {
		return nil, fmt.Errorf("update: problem looking for descriptor, %s: %s", version, err)
	}
	if err := vsn.QueryDescriptors(version, BlockTypeDescriptors, &e.blockTypes); err != nil {
		return nil, fmt.Errorf("update: problem looking for block types, %s: %s", version, err)
	}

	writeUint32(writer, uint32(numBlocks))

	if e.Descriptor.DescriptorOptions&DescriptorOptionHasHasTransport != 0 {
		writeBool(writer, false)
	}

	return e, nil
}

func (enc *Encoder) EncodeGUID(id guid.GUID) error {
	if enc.Descriptor.DescriptorOptions&DescriptorOptionAlpha != 0 {
		return id.EncodeUnpacked(enc.Build, enc)
	}
	id.EncodePacked(enc.Build, enc)
	return nil
}

func (enc *Encoder) EncodeBlockType(bt BlockType) error {
	value, ok := enc.blockTypes[bt]
	if !ok {
		panic("no type found for " + bt.String())
	}

	// if value > 3 {
	// 	panic(value)
	// }

	log.Warn("encoding", bt, "as ", value)

	return writeUint8(enc, value)
}

func (enc *Encoder) AddBlock(id guid.GUID, data BlockData, viewMask VisibilityFlags) error {
	blockType := data.Type()

	if err := enc.EncodeBlockType(blockType); err != nil {
		return err
	}

	if id != guid.Nil {
		if err := enc.EncodeGUID(id); err != nil {
			return err
		}
	}

	return data.WriteData(enc, viewMask, false)
}

func (enc *Encoder) alpha() bool {
	return vsn.Range(0, vsn.Alpha).Contains(enc.Build)
}
