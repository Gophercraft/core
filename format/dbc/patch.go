package dbc

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"

	"github.com/Gophercraft/core/format/dbc/dbd"
	"github.com/Gophercraft/core/i18n"
)

func (t *Table) makeStringRef(str string) int {
	ch := 0
	var strCursor string
	var next int
	var err error
	// First, search in StringBlock. The string may already exist, so in the name of efficiency we avoid the same string appearing more than once
	for ch < len(t.StringBlock) {
		strCursor, next, err = t.stringRef(ch)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if strCursor == str {
			return ch
		}

		ch = next + 1
	}

	// String is not present in block, add it now
	strBytes := append([]byte(str), 0x00)
	t.StringBlock = append(t.StringBlock, strBytes...)
	return ch
}

// Resizes the table buffer to hold i rows
func (t *Table) setNumRows(i int) {
	numRowsCurrent := int(t.Header.RecordCount)
	switch {
	case i == numRowsCurrent:
		return
	case i < numRowsCurrent:
		t.Records = t.Records[:int(t.Header.RecordSize)*i]
	case i > numRowsCurrent:
		t.Records = append(t.Records, make([]byte, i-numRowsCurrent)...)
	}

	t.Header.RecordCount = uint32(i)
}

func (t *Table) Append(recordSlice interface{}) error {
	slice := reflect.ValueOf(recordSlice)
	if slice.Kind() == reflect.Ptr {
		slice = slice.Elem()
	}
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf("dbc: can't append non-slice")
	}
	return t.appendRecords(slice)
}

func (t *Table) appendRecords(records reflect.Value) error {
	start := int(t.Header.RecordCount)
	t.setNumRows(int(t.Header.RecordCount) + records.Len())
	for i := 0; i < records.Len(); i++ {
		err := t.setIndex(start+i, records.Index(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) setArrayField(writeRecord []byte, valid bool, field reflect.Value, gCol *dbd.ColumnDefinition, lCol *dbd.LayoutColumn) ([]byte, error) {
	var err error

	for i := 0; i < lCol.ArraySize; i++ {
		idx := field
		if i < field.Len() {
			idx = field.Index(i)
			valid = idx.IsValid()
		} else {
			valid = false
		}
		writeRecord, err = t.setField(writeRecord, valid, idx, gCol, lCol)
		if err != nil {
			return nil, err
		}
	}

	return writeRecord, nil
}

func (t *Table) setField(writeRecord []byte, valid bool, field reflect.Value, gCol *dbd.ColumnDefinition, lCol *dbd.LayoutColumn) ([]byte, error) {
	if !valid {
		// The entity struct lacks this field, so make a default one so its empty value can be encoded.
		field = reflect.New(reflectTypeField(gCol)).Elem()
	}

	switch gCol.Type {
	case dbd.Bool:
		if field.Bool() {
			writeRecord[0] = 1
		} else {
			writeRecord[0] = 0
		}
		writeRecord = writeRecord[1:]
	case dbd.Uint, dbd.Int, dbd.Float:
		bits := lCol.Bits

		if bits == 0 {
			bits = 32
		}

		switch bits {
		case 8:
			switch gCol.Type {
			case dbd.Uint:
				writeRecord[0] = uint8(field.Uint())
				writeRecord = writeRecord[1:]
			case dbd.Int:
				writeRecord[0] = uint8(int8(field.Int()))
				writeRecord = writeRecord[1:]
			default:
				return nil, fmt.Errorf("dbc: can't encode gCol.Type %v in 8 bits", gCol.Type)
			}
		case 16:
			u16 := uint16(0)
			switch gCol.Type {
			case dbd.Uint:
				u16 = uint16(field.Uint())
			case dbd.Int:
				u16 = uint16(int16(field.Int()))
			default:
				return nil, fmt.Errorf("dbc: can't encode gCol.Type %v in 16 bits", gCol.Type)
			}
			binary.LittleEndian.PutUint16(writeRecord[0:2], u16)
			writeRecord = writeRecord[2:]
		case 32:
			u32 := uint32(0)
			switch gCol.Type {
			case dbd.Uint:
				u32 = uint32(field.Uint())
			case dbd.Int:
				u32 = uint32(int32(field.Int()))
			case dbd.Float:
				u32 = math.Float32bits(float32(field.Float()))
			default:
				return nil, fmt.Errorf("dbc: can't encode gCol.Type %v in 32 bits", gCol.Type)
			}
			binary.LittleEndian.PutUint32(writeRecord[0:4], u32)
			writeRecord = writeRecord[4:]
		case 64:
			u64 := uint64(0)
			switch gCol.Type {
			case dbd.Uint:
				u64 = field.Uint()
			case dbd.Int:
				u64 = uint64(field.Int())
			case dbd.Float:
				u64 = math.Float64bits(field.Float())
			default:
				return nil, fmt.Errorf("dbc: can't encode gCol.Type %v in 64 bits", gCol.Type)
			}
			binary.LittleEndian.PutUint64(writeRecord[0:8], u64)
			writeRecord = writeRecord[8:]
		default:
			panic(gCol.Type)
		}
	case dbd.String:
		strRef := t.makeStringRef(field.String())
		binary.LittleEndian.PutUint32(writeRecord[0:4], uint32(strRef))
		writeRecord = writeRecord[4:]
	case dbd.LocString:
		txt := field.Interface().(i18n.Text)
		locSz, err := LocStringSize(t.DB.Build)
		if err != nil {
			return nil, err
		}
		bitmaskSize := 0
		if locSz > 1 {
			bitmaskSize = 1
			locSz--
		}
		locArray := make([]uint32, locSz)
		for i := i18n.English; int(i) < len(locArray); i++ {
			if localizedString, ok := txt[i]; ok {
				locArray[int(i)] = uint32(t.makeStringRef(localizedString))
			} else {
				locArray[int(i)] = uint32(t.makeStringRef(""))
			}
		}
		if bitmaskSize != 0 {
			// TODO: figure out what flags do!
			locArray = append(locArray, 0xFFFFFFFF)
		}
		for _, i := range locArray {
			binary.LittleEndian.PutUint32(writeRecord[0:4], i)
			writeRecord = writeRecord[4:]
		}
	default:
		return nil, fmt.Errorf("dbc: unknown gCol.Type %d", gCol.Type)
	}

	return writeRecord, nil
}

func (t *Table) setIndex(i int, value reflect.Value) error {
	if i >= int(t.Header.RecordCount) {
		return fmt.Errorf("dbc: attempt to set index outside of bounds of file Table(size:%d).setIndex(%d, ...)", t.Header.RecordCount, i)
	}

	recordData := t.Records[i*int(t.Header.RecordSize) : (i+1)*int(t.Header.RecordSize)]
	writeRecord := recordData

	for i := range t.Layout.Columns {
		lCol := &t.Layout.Columns[i]
		var err error
		gCol := t.Definition.Column(lCol.Name)
		if gCol == nil {
			return fmt.Errorf("dbc: no column definition found for layout column %s", lCol.Name)
		}

		curField := value.FieldByName(gCol.Name)
		valid := curField.IsValid()

		if lCol.ArraySize > -1 {
			writeRecord, err = t.setArrayField(writeRecord, valid, curField, gCol, lCol)
			if err != nil {
				return err
			}
		} else {
			writeRecord, err = t.setField(writeRecord, valid, curField, gCol, lCol)
			if err != nil {
				return err
			}
		}
	}

	if len(writeRecord) > 0 {
		return fmt.Errorf("dbc: %d bytes remain unset in record data", len(writeRecord))
	}

	return nil
}
