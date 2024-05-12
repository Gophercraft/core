package dbd

import (
	"bufio"
	"io"
	"strings"

	"github.com/Gophercraft/core/version"
)

type ColumnType uint8

const (
	Uint ColumnType = iota
	Int
	Float
	Bool
	String
	LocString
)

type ColumnDefinition struct {
	Name          string
	Type          ColumnType
	ForeignRecord string
	ForeignKey    string
	Verified      bool
	HintArray     bool // Set to true if any layout treats the field as an array.
	HintBits      int
}

type Definition struct {
	Name    string
	Columns []ColumnDefinition
	Layouts []Layout
}

func (d *Definition) Column(name string) *ColumnDefinition {
	for i := range d.Columns {
		v := &d.Columns[i]
		if v.Name == name {
			return v
		}
	}
	return nil
}

func DecodeDefinition(name string, rd io.Reader) (*Definition, error) {
	decoder := &decoder{Reader: bufio.NewReader(rd)}
	return decoder.Decode(name)
}

type Layout struct {
	Hashes         []string
	VerifiedBuilds []version.Build
	BuildRanges    []version.BuildRange
	Columns        []LayoutColumn
}

func (l *Layout) Column(name string) *LayoutColumn {
	for i := 0; i < len(l.Columns); i++ {
		v := &l.Columns[i]
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (l *Layout) IDColumn() *LayoutColumn {
	for i := 0; i < len(l.Columns); i++ {
		v := &l.Columns[i]
		if v.HasOption("id") {
			return v
		}
	}
	return nil
}

type LayoutColumn struct {
	Options   []string
	Name      string
	Bits      int
	Signed    bool
	ArraySize int
}

func (lc LayoutColumn) HasOption(name string) bool {
	for _, opt := range lc.Options {
		if opt == name {
			return true
		}
	}
	return false
}

func ParseBuildRange(str string) (br version.BuildRange, err error) {
	var els []string
	els = strings.SplitN(str, "-", 2)
	var v version.Build
	v, err = version.ParseDBD(els[0])
	if err != nil {
		return
	}
	br[0] = v
	v, err = version.ParseDBD(els[1])
	if err != nil {
		return
	}
	br[1] = v
	return
}
