package models

import (
	"fmt"
	"strconv"
	"strings"
)

type BodyType uint8

const (
	BodyType1 BodyType = iota
	BodyType2
)

// Load with info from game database on startup
var BodyTypeNames = map[BodyType]string{}

func (r BodyType) String() string {
	n, ok := BodyTypeNames[r]
	if !ok {
		return fmt.Sprintf("unknown BodyType ID: %d", r)
	}
	return n
}

type BodyTypeMask uint8

const (
	AllBodyTypes BodyTypeMask = 0xFF
)

func (c *BodyTypeMask) All() bool {
	return *c == AllBodyTypes
}

func (c *BodyTypeMask) EncodeWord() (string, error) {
	if c.All() {
		return "All", nil
	}

	var enables []string

	for BodyType := BodyType(1); BodyType < 9; BodyType++ {
		if c.Has(BodyType) {
			enables = append(enables, fmt.Sprintf("%d", BodyType))
		}
	}

	return strings.Join(enables, "|"), nil
}

func (c *BodyTypeMask) DecodeWord(data string) error {
	if data == "All" {
		*c = AllBodyTypes
		return nil
	}

	enables := strings.Split(data, "|")

	for _, enable := range enables {
		var bodyType BodyType
		u, err := strconv.ParseUint(enable, 10, 8)
		if err != nil {
			return err
		}
		bodyType = BodyType(u)
		c.Set(bodyType, true)
	}

	return nil
}

func (c BodyTypeMask) Has(bodyType BodyType) bool {
	return c&(1<<BodyTypeMask(bodyType)) != 0
}

func (c *BodyTypeMask) Set(bodyType BodyType, t bool) {
	var flag BodyTypeMask = 1 << BodyTypeMask(bodyType)

	if t {
		*c |= flag
	} else {
		*c &= ^flag
	}
}
