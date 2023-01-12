package player

import (
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb/models"
)

// A Character can be a flesh-and-blood human, a bot, or a Goober.
type Character struct {
	*update.ValuesBlock
	Model        models.Character
	MovementInfo *update.MovementInfo
}
