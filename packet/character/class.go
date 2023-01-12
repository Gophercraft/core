package character

import "fmt"

type Class uint8

// Load with info from game database on startup
var ClassNames = map[Class]string{}

func (r Class) String() string {
	n, ok := ClassNames[r]
	if !ok {
		return fmt.Sprintf("unknown class ID: %d", r)
	}
	return n
}
