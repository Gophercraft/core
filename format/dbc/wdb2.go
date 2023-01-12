package dbc

import (
	"encoding/binary"
	"io"
)

type WDB2Header struct {
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

func (t *Table) readWDB2Header(rd io.Reader) error {
	var wdbcHeader WDB2Header
	err := binary.Read(rd, binary.LittleEndian, &wdbcHeader)
	if err != nil {
		return err
	}
	t.Header.RecordCount = wdbcHeader.RecordCount
	t.Header.FieldCount = wdbcHeader.FieldCount
	t.Header.RecordSize = wdbcHeader.RecordSize
	t.Header.StringBlockSize = wdbcHeader.StringBlockSize
	t.Header.TableHash = wdbcHeader.TableHash
	t.Header.Build = wdbcHeader.Build
	t.Header.TimestampLastWritten = wdbcHeader.TimestampLastWritten
	t.Header.MinID = wdbcHeader.MinID
	t.Header.MaxID = wdbcHeader.MaxID
	t.Header.CopyTableSize = wdbcHeader.CopyTableSize
	return nil
}
