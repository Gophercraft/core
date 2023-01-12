package spell

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

type ActionButton struct {
	Action uint32
	Type   models.ActionType
}

type ActionButtons struct {
	Buttons []ActionButton
}

type SetActionButton struct {
	Slot uint32
	ActionButton
}

func NumActionButtons(build vsn.Build) int {
	ab := 120
	if build.AddedIn(vsn.V2_0_1) {
		ab = 132
	}
	return ab
}

const (
	spellMask = 0xffffff
)

func (ab ActionButton) typeNibble() uint32 {
	switch ab.Type {
	case models.ActionSpell:
		return 0
	case models.ActionMacro:
		return 64
	case models.ActionItem:
		return 128
	}

	panic(ab.Type)
}

func (ab ActionButton) Uint32(build vsn.Build) uint32 {
	// if build == vsn.Alpha {
	// 	var sign int32
	// 	switch ab.Type {
	// 	case models.ActionItem:
	// 		sign = -1
	// 	case models.ActionSpell:
	// 		sign = 1
	// 	default:
	// 		panic(fmt.Sprintf("spell: ActionButton.Uint32: cannot encode this type %d", ab.Type))
	// 	}

	// 	data := int32(ab.Action) * sign
	// 	return uint32(data)
	// }

	if ab.Type >= 0xff {
		log.Warn("actionbutton slot too big")
		return 0
	}
	// uint24(ab.Spell) + uint8(ab.Slot)
	return (ab.Action & spellMask) | (ab.typeNibble() << 24)
}

func (abs *ActionButtons) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_ACTION_BUTTONS
	nab := NumActionButtons(build)

	if len(abs.Buttons) > nab {
		return fmt.Errorf("packet: action button overflow %d/%d", len(abs.Buttons), nab)
	}

	for ab := 0; ab < nab; ab++ {
		if len(abs.Buttons) <= ab {
			out.WriteUint32(0)
		} else {
			out.WriteUint32(abs.Buttons[ab].Uint32(build))
		}
	}

	log.Dump("buttons", out.Bytes())
	return nil
}

func convertTypeNibble(nibble uint32) (at models.ActionType, err error) {
	switch nibble {
	case 0:
		at = models.ActionSpell
	case 64:
		at = models.ActionMacro
	case 128:
		at = models.ActionItem
	default:
		err = fmt.Errorf("spell: invalid nibble value %d", nibble)
	}

	return
}

func (abs *ActionButtons) Decode(build vsn.Build, in *packet.WorldPacket) error {
	nab := in.Len() / 4
	abs.Buttons = make([]ActionButton, nab)

	// In Alpha, instead of type nibble, the sign is used to separate items
	// No macros yet, only ( < 0 == Item ) ( > 0 == Spell )
	if build == vsn.Alpha {
		for i := 0; i < nab; i++ {
			act := int32(in.ReadUint32())

			if act < 0 {
				abs.Buttons[i].Action = uint32(act * -1)
				abs.Buttons[i].Type = models.ActionItem
			} else {
				abs.Buttons[i].Action = uint32(act)
				abs.Buttons[i].Type = models.ActionSpell
			}
		}
		return nil
	}

	var err error

	for i := 0; i < nab; i++ {
		act := in.ReadUint32()
		abs.Buttons[i].Action = act & spellMask
		typeNibble := (act >> 24) & 0xff
		abs.Buttons[i].Type, err = convertTypeNibble(typeNibble)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sab *SetActionButton) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_SET_ACTION_BUTTON
	out.WriteByte(uint8(sab.Slot))
	// out.WriteUint32(sab.Action&spellMask | (sab.Type << 24))
	out.WriteUint32(sab.ActionButton.Uint32(build))
	return nil
}

func (sab *SetActionButton) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	sab.Slot = uint32(in.ReadByte())
	if sab.Slot >= uint32(NumActionButtons(build)) {
		err = fmt.Errorf("packet: SetActionButton slot %d too large for build %s", sab.Slot, build)
		return
	}

	act := in.ReadUint32()

	if build == vsn.Alpha {
		sact := int32(act)
		if sact < 0 {
			sab.Type = models.ActionItem
			sab.Action = uint32(sact * -1)
		} else {
			sab.Type = models.ActionSpell
			sab.Action = act
		}

		return
	}

	sab.Action = act & spellMask
	typeNibble := (act >> 24) & 0xff
	sab.Type, err = convertTypeNibble(typeNibble)
	if err != nil {
		return
	}

	return
}
