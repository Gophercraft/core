package models

import "fmt"

type SpawnKind uint8

const (
	SpawnKindCreature SpawnKind = iota
	SpawnKindGameObject
)

func (ok *SpawnKind) EncodeWord() (string, error) {
	switch *ok {
	case SpawnKindCreature:
		return "Creature", nil
	case SpawnKindGameObject:
		return "GameObject", nil
	default:
		return "", fmt.Errorf("%d", *ok)
	}
}

func (ok *SpawnKind) DecodeWord(str string) error {
	switch str {
	case "Creature":
		*ok = SpawnKindCreature
	case "GameObject":
		*ok = SpawnKindGameObject
	default:
		return fmt.Errorf("invalid Spawn kind %s", str)
	}
	return nil
}
