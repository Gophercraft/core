package dbc

import (
	"encoding/binary"
	"io"
)

type WDBCHeader struct {
	RecordCount     uint32
	FieldCount      uint32
	RecordSize      uint32
	StringBlockSize uint32
}

func (t *Table) readWDBCHeader(rd io.Reader) error {
	var wdbcHeader WDBCHeader
	err := binary.Read(rd, binary.LittleEndian, &wdbcHeader)
	if err != nil {
		return err
	}
	t.Header.RecordCount = wdbcHeader.RecordCount
	t.Header.FieldCount = wdbcHeader.FieldCount
	t.Header.RecordSize = wdbcHeader.RecordSize
	t.Header.StringBlockSize = wdbcHeader.StringBlockSize
	return nil
}
