package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Race uint8

const (
	AllRaces RaceMask = 0xFFFFFFFFFFFFFFFF
)

// Load with info from game database on startup
var RaceNames = map[Race]string{}

func (r Race) String() string {
	n, ok := RaceNames[r]
	if !ok {
		return fmt.Sprintf("unknown race ID: %d", r)
	}
	return n
}

type RaceMask uint64

func (c *RaceMask) All() bool {
	return *c == AllRaces
}

func (c *RaceMask) EncodeWord() (string, error) {
	if c.All() {
		return "All", nil
	}

	var enables []string

	for race := Race(1); race < 64; race++ {
		if c.Has(race) {
			enables = append(enables, fmt.Sprintf("%d", race))
		}
	}

	return strings.Join(enables, "|"), nil
}

func (c *RaceMask) DecodeWord(data string) error {
	if data == "All" {
		*c = AllRaces
		return nil
	}

	enables := strings.Split(data, "|")

	for _, enable := range enables {
		var race Race
		u, err := strconv.ParseUint(enable, 10, 8)
		if err != nil {
			return err
		}
		race = Race(u)
		c.Set(race, true)
	}

	return nil
}

func (c RaceMask) Has(race Race) bool {
	if race == 0 {
		panic("race cannot be zero")
	}

	return c&(1<<RaceMask(race-1)) != 0
}

func (c *RaceMask) Set(race Race, t bool) {
	if race == 0 {
		panic("race cannot be zero")
	}

	var flag RaceMask = 1 << RaceMask(race-1)

	if t {
		*c |= flag
	} else {
		*c &= ^flag
	}
}
