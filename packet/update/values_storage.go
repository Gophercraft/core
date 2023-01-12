package update

import (
	"fmt"
	"reflect"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/vsn"
)

func NewValuesBlock(build vsn.Build, mask guid.TypeMask) (*ValuesBlock, error) {
	var descriptor *Descriptor
	if err := vsn.QueryDescriptors(build, Descriptors, &descriptor); err != nil {
		return nil, err
	}

	storageDescriptorType := descriptor.ObjectDescriptors[mask]
	if storageDescriptorType == nil {
		return nil, fmt.Errorf("update: cannot find storage descriptor for type %s in descriptor %s", mask, build)
	}

	newStorageDescriptor := reflect.New(storageDescriptorType)

	vBlock := &ValuesBlock{
		TypeMask:          mask,
		Descriptor:        descriptor,
		ChangeMask:        NewBitmask(),
		StorageDescriptor: newStorageDescriptor,
	}

	typeUint, err := mask.Resolve(build)
	if err != nil {
		return nil, err
	}

	vBlock.SetUint32("Type", typeUint)

	return vBlock, nil
}

func (vb *ValuesBlock) ClearChanges() {
	vb.ChangeMask.Clear()
}

func (vb *ValuesBlock) ClearChangesAndUnlock() {
	vb.ClearChanges()
	vb.Unlock()
}

type Value struct {
	Block            *ValuesBlock
	Bitmask          *Bitmask
	ChangeMaskOffset uint32
	Value            reflect.Value
}

func (vb *ValuesBlock) Value() *Value {
	val := &Value{
		Block:   vb,
		Bitmask: vb.ChangeMask,
		Value:   vb.StorageDescriptor.Elem(),
	}

	return val
}

func (vb *ValuesBlock) Get(name string) *Value {
	v := vb.Value()
	if v == nil {
		panic("no valuesblock value?")
	}
	return v.Find(name)
}

// Returns any substructure with a reflect Name == field
func (v *Value) Find(field string) *Value {
	if v.Value.Kind() != reflect.Struct {
		panic("you cannot find on a non-struct")
		return nil
	}

	// fmt.Println("lookin for", field, "in ", v.Value.Type())

	t := v.Value.Type()

	// Quick-search for named fields.
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == field {
			return v.FieldIndex(i)
		}
	}

	// Recurse through all field indexes.
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Struct {
			found := v.FieldIndex(i).Find(field)
			if found != nil {
				return found
			}
		}
	}

	return nil
}

// func (v *Value) Get() {

// }

// Returns a list of fields contained within a Struct value.
func (v *Value) FieldNames() []string {
	t := v.Value.Type()

	names := make([]string, t.NumField())
	for i := 0; i < len(names); i++ {
		names[i] = t.Field(i).Name
	}

	return names
}

// Take a value with an already computed base changemask offset, building until the index is reached
func (v *Value) FieldIndex(n int) *Value {
	nextValue := &Value{}
	nextValue.Block = v.Block
	// todo: in modern protocol, bitmasks exist at several levels throughout the packet
	nextValue.Bitmask = v.Bitmask
	// If field index in struct == 0, then nextvalue.ChangeMaskOffset == v.ChangeMaskOffset.

	var chunkOffset, bitOffset uint32

	chunkOffset = v.ChangeMaskOffset

	for i := 0; i < n; i++ {
		fieldValue := &Value{
			Block:            v.Block,
			Bitmask:          v.Bitmask,
			ChangeMaskOffset: v.ChangeMaskOffset,
			Value:            v.Value.Field(i),
		}
		fieldValue.nextField(&chunkOffset, &bitOffset)
	}

	nextValue.ChangeMaskOffset = chunkOffset
	nextValue.Value = v.Value.Field(n)

	return nextValue
}

func (v *Value) Field(name string) *Value {
	for i := 0; i < v.Value.NumField(); i++ {
		field := v.Value.Type().Field(i)
		// Check this field to find its absolute offset
		if field.Name == name {
			return v.FieldIndex(i)
		}
	}

	return nil
}

func (v *Value) Index(index int) *Value {
	if v.Value.Len() <= index {
		panic("value index out of bounds")
		return nil
	}

	nextValue := &Value{}
	nextValue.Block = v.Block
	nextValue.Bitmask = v.Bitmask
	nextValue.ChangeMaskOffset = v.ChangeMaskOffset

	var chunkOffset, bitOffset uint32

	chunkOffset = v.ChangeMaskOffset

	for i := 0; i < v.Value.Len(); i++ {
		nextValue.Value = v.Value.Index(i)
		if i == index {
			nextValue.ChangeMaskOffset = chunkOffset
			return nextValue
		}

		nextValue.nextField(&chunkOffset, &bitOffset)
	}

	return nil
}

// Ensure proper type inference by setting with these functions

func (v *Value) SetUint32(value uint32) {
	v.Block.Lock()
	v.Value.SetUint(uint64(value))
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Block.Unlock()
}

func (vb *ValuesBlock) SetUint32(glob string, nvalue uint32) {
	value := vb.Get(glob)
	if value == nil {
		panic(fmt.Errorf("Could not Get %s in %s", glob, vb.Value().Value.Type()))
	}
	value.SetUint32(nvalue)
}

func (v *Value) Len() int {
	return v.Value.Len()
}

func (v *Value) SetGUID(value guid.GUID) {
	v.Block.Lock()
	v.Value.Set(reflect.ValueOf(value))
	// Todo: in, modern protocol all guids are packed and only consume 1 bit in change mask.
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Bitmask.Set(v.ChangeMaskOffset+1, true)
	v.Block.Unlock()
}

func (vb *ValuesBlock) SetGUID(glob string, value guid.GUID) {
	vb.Get(glob).SetGUID(value)
}

func (v *Value) SetFloat32(value float32) {
	v.Block.Lock()
	v.Value.SetFloat(float64(value))
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Block.Unlock()
}

func (vb *ValuesBlock) SetFloat32(glob string, value float32) {
	vb.Get(glob).SetFloat32(value)
}

func (v *Value) SetBit(value bool) {
	v.Block.Lock()
	v.Value.SetBool(value)
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Block.Unlock()
}

func (vb *ValuesBlock) SetBit(glob string, value bool) {
	vb.Get(glob).SetBit(value)
}

func (v *Value) SetUint16(value uint16) {
	v.Block.Lock()
	v.Value.SetUint(uint64(value))
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Block.Unlock()
}

func (v *Value) SetInt16(value int16) {
	v.Block.Lock()
	v.Value.SetInt(int64(value))
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Block.Unlock()
}

func (v *Value) SetByte(value uint8) {
	v.Block.Lock()
	v.Value.SetUint(uint64(value))
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Block.Unlock()
}

func (vb *ValuesBlock) SetByte(glob string, value uint8) {
	vb.Get(glob).SetByte(value)
}

func (v *Value) SetInt32(value int32) {
	v.Block.Lock()
	v.Value.SetInt(int64(value))
	v.Bitmask.Set(v.ChangeMaskOffset, true)
	v.Block.Unlock()
}

func (vb *ValuesBlock) SetInt32(glob string, value int32) {
	vb.Get(glob).SetInt32(value)
}

func (v *Value) Uint32() uint32 {
	return uint32(v.Value.Uint())
}

func (v *Value) Int32() int32 {
	return int32(v.Value.Int())
}

func (v *Value) Byte() uint8 {
	return uint8(v.Value.Uint())
}

func (v *Value) GUID() guid.GUID {
	return v.Value.Interface().(guid.GUID)
}
