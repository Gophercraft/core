package dbc

import (
	"encoding/binary"
	"math"
	"reflect"

	"github.com/Gophercraft/core/format/dbc/dbd"
	"github.com/Gophercraft/core/i18n"
)

var (
	bitTypes = map[dbd.ColumnType]map[int]reflect.Type{
		dbd.Int: {
			8:  reflect.TypeOf(int8(0)),
			16: reflect.TypeOf(int16(0)),
			32: reflect.TypeOf(int32(0)),
			64: reflect.TypeOf(int64(0)),
		},

		dbd.Uint: {
			8:  reflect.TypeOf(uint8(0)),
			16: reflect.TypeOf(uint16(0)),
			32: reflect.TypeOf(uint32(0)),
			64: reflect.TypeOf(uint64(0)),
		},

		dbd.Float: {
			32: reflect.TypeOf(float32(0.0)),
			64: reflect.TypeOf(float64(0.0)),
		},
	}

	locStrType = reflect.TypeOf(i18n.Text{})
)

func getNumBits(gCol *dbd.ColumnDefinition) int {
	if gCol.HintBits > 0 {
		return gCol.HintBits
	}

	return 64
}

// func colTypeToKind(colType dbd.ColumnType) (kind reflect.Kind, ok bool) {
// 	switch colType {
// 	case dbd.Uint:
// 		return reflect.Uint, true
// 	case dbd.Int:
// 		return reflect.Int, true
// 	case dbd.Float:
// 		return reflect.Float32, true
// 	default:
// 		return reflect.Invalid, false
// 	}
// }

func reflectTypeField(gCol *dbd.ColumnDefinition) reflect.Type {
	bitsMap, ok := bitTypes[gCol.Type]
	if ok {
		var bits int = getNumBits(gCol)
		fieldType, ok := bitsMap[bits]
		if ok {
			return fieldType
		}
	}

	var typ reflect.Type
	switch gCol.Type {
	case dbd.Uint:
		typ = reflect.TypeOf(uint64(0))
	case dbd.Int:
		typ = reflect.TypeOf(int64(0))
	case dbd.Float:
		typ = reflect.TypeOf(float32(0))
	case dbd.Bool:
		typ = reflect.TypeOf(bool(true))
	case dbd.String:
		typ = reflect.TypeOf(string(""))
	case dbd.LocString:
		typ = reflect.TypeOf(i18n.Text{})
	default:
		panic(gCol.Type)
	}
	return typ
}

// parseSingleField: parses data according to a dbd.LayoutColum
// And converts it meet a ColumnDefinition (i.e. the definition that exists independent of specific versions)
func (t *Table) parseSingleField(gCol *dbd.ColumnDefinition, lCol *dbd.LayoutColumn, data []byte) (reflect.Value, []byte, error) {
	v := reflect.New(reflectTypeField(gCol)).Elem()

	// Note that we are parsing these values according to how they appear in the Layout
	// but we are storing the parsed values into the *Hinted* size by the generic ColumnDefiniton
	switch gCol.Type {
	case dbd.Uint:
		var u uint64
		switch lCol.Bits {
		case 8:
			u = uint64(data[0])
			data = data[1:]
		case 16:
			u = uint64(binary.LittleEndian.Uint16(data[0:2]))
			data = data[2:]
		case 32:
			u = uint64(binary.LittleEndian.Uint32(data[0:4]))
			data = data[4:]
		case 64:
			u = uint64(binary.LittleEndian.Uint64(data[0:8]))
			data = data[8:]
		default:
			panic(lCol.Bits)
		}
		v.SetUint(u)
	case dbd.Int:
		// Sometimes columns are defined as int b
		var i int64
		switch lCol.Bits {
		case 8:
			if lCol.Signed {
				i = int64(int8(data[0]))
			} else {
				i = int64(uint8(data[0]))
			}
			data = data[1:]
		case 16:
			if lCol.Signed {
				i = int64(int16(binary.LittleEndian.Uint16(data[0:2])))
			} else {
				i = int64(uint16(binary.LittleEndian.Uint16(data[0:2])))
			}
			data = data[2:]
		case 32:
			if lCol.Signed {
				i = int64(int32(binary.LittleEndian.Uint32(data[0:4])))
			} else {
				i = int64(uint32(binary.LittleEndian.Uint32(data[0:4])))
			}
			data = data[4:]
		case 64:
			i = int64(binary.LittleEndian.Uint64(data[0:8]))
			data = data[8:]
		default:
			panic(lCol.Bits)
		}
		v.SetInt(i)
	case dbd.Float:
		var f float64
		switch lCol.Bits {
		case 32:
			f = float64(math.Float32frombits(binary.LittleEndian.Uint32(data[0:4])))
			data = data[4:]
		case 64:
			f = float64(math.Float64frombits(binary.LittleEndian.Uint64(data[0:8])))
			data = data[8:]
		default:
			panic(lCol.Bits)
		}
		v.SetFloat(f)
	case dbd.Bool:
		u8 := data[0]
		data = data[1:]
		bl := u8 == 1
		v.SetBool(bl)
	case dbd.String:
		var i int
		i = int(binary.LittleEndian.Uint32(data[0:4]))
		data = data[4:]
		str, err := t.StringRef(i)
		if err != nil {
			return reflect.Value{}, nil, err
		}
		v.SetString(str)
	case dbd.LocString:
		text := make(i18n.Text)
		v.Set(reflect.ValueOf(text))

		sz, err := LocStringSize(t.DB.Build)
		if err != nil {
			return reflect.Value{}, nil, err
		}

		if sz == 1 {
			var i int
			i = int(binary.LittleEndian.Uint32(data[0:4]))
			data = data[4:]
			str, err := t.StringRef(i)
			if err != nil {
				return v, nil, err
			}

			text[i18n.English] = str

			return v, data, err
		}

		ints := make([]uint32, sz)
		for i := 0; i < sz; i++ {
			ints[i] = binary.LittleEndian.Uint32(data[i*4 : (i+1)*4])
		}

		data = data[sz*4:]

		// bitmask := ints[len(ints)-1]
		numLang := sz - 1

		for i := 0; i < numLang; i++ {
			langstring, err := t.StringRef(int(ints[i]))
			if err != nil {
				return reflect.Value{}, nil, err
			}
			if langstring != "" {
				text[i18n.Locale(i)] = langstring
			}
		}
	default:
		panic(gCol.Type)
	}

	return v, data, nil
}

func (t *Table) parseField(gCol *dbd.ColumnDefinition, col *dbd.LayoutColumn, data []byte) (reflect.Value, []byte, error) {
	if gCol.HintArray {
		// Asize is the number of array fields expected in the generic ColumnDefinition
		// i.e. the number of array elements to be expected usually by the api
		genericSize := col.ArraySize
		if genericSize == -1 {
			genericSize = 1
		}

		typ := reflect.SliceOf(reflectTypeField(gCol))

		array := reflect.MakeSlice(typ, genericSize, genericSize)

		var sField reflect.Value
		var err error

		numFieldsInRecord := col.ArraySize
		// The column's layout might not specify the the length.
		// For something that is hinted to be an array in other layouts
		// this can only mean that it is ONE FIELD.
		if col.ArraySize == -1 {
			numFieldsInRecord = 1
		}

		for i := 0; i < numFieldsInRecord; i++ {
			sField, data, err = t.parseSingleField(gCol, col, data)
			if err != nil {
				return sField, nil, err
			}
			array.Index(i).Set(sField)
		}

		return array, data, nil
	}

	return t.parseSingleField(gCol, col, data)
}
