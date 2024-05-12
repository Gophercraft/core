package character

import (
	"github.com/Gophercraft/core/game/protocol/message/packet/character"
	"github.com/Gophercraft/core/guid"
)

type ServiceProvider interface {
	// Return a slice containing all Characters belonging to the game account
	GetCharacterList(game_account uint64) ([]character.Character, error)
	// Create a character
	CreateCharacter(game_account uint64, character *character.Character) error
	//
	DeleteCharacter(id guid.GUID) error
}
