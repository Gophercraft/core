package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Class uint8

const (
	AllClasses ClassMask = 0xFFFFFFFFFFFFFFFF
)

// Load with info from game database on startup
var ClassNames = map[Class]string{}

func (r Class) String() string {
	n, ok := ClassNames[r]
	if !ok {
		return fmt.Sprintf("unknown class ID: %d", r)
	}
	return n
}

type ClassMask uint64

func (c *ClassMask) All() bool {
	return *c == AllClasses
}

func (c *ClassMask) EncodeWord() (string, error) {
	if c.All() {
		return "All", nil
	}

	var enables []string

	for class := Class(1); class < 64; class++ {
		if c.Has(class) {
			enables = append(enables, fmt.Sprintf("%d", class))
		}
	}

	return strings.Join(enables, "|"), nil
}

func (c *ClassMask) DecodeWord(data string) error {
	if data == "All" {
		*c = AllClasses
		return nil
	}

	enables := strings.Split(data, "|")

	for _, enable := range enables {
		var class Class
		u, err := strconv.ParseUint(enable, 10, 8)
		if err != nil {
			return err
		}
		class = Class(u)
		c.Set(class, true)
	}

	return nil
}

func (c ClassMask) Has(class Class) bool {
	if class == 0 {
		panic("class cannot be zero")
	}
	return c&(1<<ClassMask(class-1)) != 0
}

func (c *ClassMask) Set(class Class, t bool) {
	if class == 0 {
		panic("class cannot be zero")
	}
	var flag ClassMask = 1 << ClassMask(class-1)

	if t {
		*c |= flag
	} else {
		*c &= ^flag
	}
}
