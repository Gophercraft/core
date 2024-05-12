package character

import "fmt"

type Race uint8

// Load with info from game database on startup
var RaceNames = map[Race]string{}

func (r Race) String() string {
	n, ok := RaceNames[r]
	if !ok {
		return fmt.Sprintf("unknown race ID: %d", r)
	}
	return n
}
