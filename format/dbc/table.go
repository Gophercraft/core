package dbc

import (
	"fmt"
	"io"
	"reflect"
	"unicode/utf8"

	"github.com/Gophercraft/core/format/dbc/dbd"
	"github.com/Gophercraft/core/version"
	"github.com/davecgh/go-spew/spew"
)

const (
	WDBC = "WDBC"
	WDB2 = "WDB2"
)

type FileHeader struct {
	Version              string
	RecordCount          uint32
	FieldCount           uint32
	RecordSize           uint32
	StringBlockSize      uint32
	TableHash            uint32
	Build                uint32
	TimestampLastWritten uint32
	MinID                uint32
	MaxID                uint32
	CopyTableSize        uint32
}

type Table struct {
	DB          *DB
	Name        string
	Header      FileHeader
	Definition  *dbd.Definition
	Layout      *dbd.Layout
	Records     []byte
	StringBlock []byte
}

func (t *Table) Len() int {
	return int(t.Header.RecordCount)
}

func (d *DB) Open(name string, reader io.Reader) (*Table, error) {
	table := new(Table)
	table.DB = d
	table.Name = name
	if err := d.detectLayout(table); err != nil {
		return nil, err
	}

	if err := table.readHeader(reader); err != nil {
		return nil, err
	}
	recordSz, numCols, err := table.SizeCount(d.Build)
	if err != nil {
		return nil, err
	}
	if int(table.Header.RecordSize) != recordSz {
		return nil, fmt.Errorf("dbc: failed to open %s, disagreement on record size between file (%d) and definition (%d)", table.Name, table.Header.RecordSize, recordSz)
	}

	if int(table.Header.FieldCount) != numCols {
		return nil, fmt.Errorf("dbc: failed to open %s, disagreement on number of fields (array-inclusive) between file (%d) and definition (%d)", table.Name, table.Header.FieldCount, numCols)
	}

	table.Records = make([]byte, table.Header.RecordSize*table.Header.RecordCount)

	_, err = io.ReadFull(reader, table.Records)
	if err != nil {
		return nil, err
	}

	table.StringBlock = make([]byte, table.Header.StringBlockSize)
	_, err = io.ReadFull(reader, table.StringBlock)
	if err != nil {
		return nil, err
	}
	return table, nil
}

// func (t *Table) FindRowByID(id int, ent interface{}) error {

// }

func (t *Table) Range(handler interface{}) error {
	fnType := reflect.TypeOf(handler)
	fn := reflect.ValueOf(handler)
	if fnType.NumIn() != 1 {
		return fmt.Errorf("dbc: invalid Range handler, needs only 1 argument")
	}
	if fnType.NumOut() != 1 || fnType.Out(0).Kind() != reflect.Bool {
		return fmt.Errorf("dbc: invalid Range handler, needs to return a boolean that, if false, halts the range")
	}
	ent := fnType.In(0)
	if ent.Kind() != reflect.Ptr || ent.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dbc: handler argument needs to be a pointer to an entity struct")
	}

	cursor := reflect.New(ent.Elem())

	for i := 0; i < int(t.Header.RecordCount); i++ {
		err := t.indexRow(i, cursor.Elem())
		if err != nil {
			return err
		}
		out := fn.Call([]reflect.Value{cursor})
		if !out[0].Bool() {
			return nil
		}
	}
	return nil
}

func (t *Table) ID(id int, ent interface{}) error {
	if id <= 0 {
		return fmt.Errorf("dbc: IDs start at 1")
	}
	col := t.Layout.IDColumn()
	if col == nil {
		return fmt.Errorf("dbc: table has no ID")
	}
	// index based IDs
	if col.HasOption("noninline") {
		return t.Index(id-1, ent)
	}
	entity := reflect.ValueOf(ent)
	if entity.Kind() == reflect.Ptr {
		entity = entity.Elem()
	}
	for i := 0; i < int(t.Header.RecordCount); i++ {
		err := t.indexRow(i, entity)
		if err != nil {
			return err
		}
		cursorID := int(entity.FieldByName(col.Name).Int())
		if cursorID == id {
			return nil
		}
	}
	return fmt.Errorf("dbc: couldn't find record matching ID %d", id)
}

func (t *Table) Index(i int, ent interface{}) error {
	entity := reflect.ValueOf(ent)
	if entity.Kind() == reflect.Ptr {
		entity = entity.Elem()
	}
	return t.indexRow(i, entity)
}

func (t *Table) indexRow(i int, entity reflect.Value) error {
	if i >= int(t.Header.RecordCount) {
		return fmt.Errorf("dbc: IndexRow supplied index [%d] out of bounds [%d]", i, t.Header.RecordCount)
	}

	// the size of 1 record (or row)
	recSz := int(t.Header.RecordSize)

	// the data of 1 record (or row)
	data := t.Records[i*recSz : (i+1)*recSz]

	var value reflect.Value
	var err error

	for _, col := range t.Layout.Columns {
		gCol := t.Definition.Column(col.Name)

		if gCol == nil {
			return fmt.Errorf("dbc: bad definition: could not find global column named %s", col.Name)
		}

		typ, ok := entity.Type().FieldByName(gCol.Name)
		// if !ok && gCol.Verified {
		// 	return fmt.Errorf("dbc: bad entity format for %s, lacks field %s", t.Name, gCol.Name)
		// }

		setEntityField := ok && gCol.Verified

		if col.HasOption("noninline") {
			if col.HasOption("id") {
				if gCol.Type != dbd.Int {
					return fmt.Errorf("dbc: non-inline ID type for %s is not Int", t.Name)
				}

				field := entity.FieldByName(gCol.Name)
				field.SetInt(int64(i + 1))
			}
			continue
		}

		value, data, err = t.parseField(gCol, &col, data)
		if err != nil {
			return err
		}

		if setEntityField {
			if typ.Type != value.Type() {
				fmt.Println(spew.Sdump(gCol))
				return fmt.Errorf("dbc: mismatch between generated definition type %s and entity type %s (%s)", value.Type(), typ.Type, gCol.Name)
			}

			field := entity.FieldByName(gCol.Name)
			field.Set(value)
		}
	}

	return nil
}

// Returns the size of this record layout in bytes
// This is because the size of fields has been known to change from version to version, for instance when localized strings add more locales.
func (t *Table) SizeCount(v version.Build) (int, int, error) {
	sz := int(0)
	cols := 0

	for _, col := range t.Layout.Columns {
		gCol := t.Definition.Column(col.Name)
		if gCol == nil {
			panic("no definition corresponding to " + col.Name)
		}

		if col.HasOption("noninline") {
			continue
		}

		// The type.
		fieldType := gCol.Type
		fieldSz := int(0)
		colSz := int(0)
		switch fieldType {
		// All ints must come with a bit size
		case dbd.Uint, dbd.Int:
			fieldSz = col.Bits / 8
			colSz = 1
		case dbd.Float:
			fieldSz = col.Bits / 8
			colSz = 1
		case dbd.Bool:
			fieldSz = 1
			colSz = 1
		case dbd.String:
			fieldSz = 4
			colSz = 1
		case dbd.LocString:
			lss, err := LocStringSize(v)
			if err != nil {
				return 0, 0, err
			}
			colSz = lss
			fieldSz = 4 * lss
		default:
			panic(fieldType)
		}
		if col.ArraySize >= 0 {
			fieldSz *= col.ArraySize
			colSz *= col.ArraySize
		}
		sz += int(fieldSz)
		cols += colSz
	}
	return sz, cols, nil
}

func (t *Table) stringRef(i int) (strSect string, high int, err error) {
	if i >= len(t.StringBlock) {
		err = io.EOF
		return
	}
	low := i
	high = i
	for c := i; c < len(t.StringBlock); c++ {
		if t.StringBlock[c] == 0x00 {
			high = c
			break
		}
	}
	strSect = string(t.StringBlock[low:high])
	if !utf8.ValidString(strSect) {
		err = fmt.Errorf("dbc: invalid characters in string field")
	}
	return
}

func (t *Table) StringRef(i int) (string, error) {
	if i >= len(t.StringBlock) {
		return "", fmt.Errorf("dbc: out of bounds stringref: StringBlock<%d>[%d] in table %s", len(t.StringBlock), i, t.Name)
	}

	str, _, err := t.stringRef(i)
	return str, err
}

func (t *Table) readHeader(reader io.Reader) error {
	var magic [4]byte

	if _, err := io.ReadFull(reader, magic[:]); err != nil {
		return err
	}

	t.Header.Version = string(magic[:])

	switch t.Header.Version {
	case WDBC:
		return t.readWDBCHeader(reader)
	case WDB2:
		return t.readWDB2Header(reader)
	default:
		return fmt.Errorf("dbc: unsupported dbc type: %s", t.Header.Version)
	}
}
