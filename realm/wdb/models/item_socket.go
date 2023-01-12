package models

import (
	"fmt"
	"strings"
)

type SocketColor uint8

const (
	SocketMeta SocketColor = 1 << iota
	SocketRed
	SocketYellow
	SocketBlue
)

func (sc *SocketColor) EncodeWord() (string, error) {
	var enables []string
	c := *sc

	if c&SocketMeta != 0 {
		enables = append(enables, "Meta")
	}

	if c&SocketRed != 0 {
		enables = append(enables, "Red")
	}

	if c&SocketYellow != 0 {
		enables = append(enables, "Yellow")
	}

	if c&SocketBlue != 0 {
		enables = append(enables, "Blue")
	}

	return strings.Join(enables, "|"), nil
}

func (sc *SocketColor) DecodeWord(data string) error {
	enables := strings.Split(data, "|")

	for _, enable := range enables {
		switch enable {
		case "Meta":
			*sc |= SocketMeta
		case "Red":
			*sc |= SocketRed
		case "Yellow":
			*sc |= SocketYellow
		case "Blue":
			*sc |= SocketBlue
		default:
			return fmt.Errorf("models: unknown socket color %s", enable)
		}
	}

	return nil
}

type ItemSocket struct {
	Color   SocketColor
	Content int32
}
