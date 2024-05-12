package spell

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type LogEffectApplication struct {
	Target           guid.GUID
	PowerTaken       uint32
	PowerType        uint32
	Multiplier       float32
	Amount           uint32
	InterruptedSpell uint32
	Item             uint32
	Slot             uint32
}

type LoggedEffect struct {
	Kind         Effect
	Applications []LogEffectApplication
}

func (eff *LoggedEffect) Encode(build version.Build, out *message.Packet) error {
	out.WriteInt32(int32(eff.Kind))
	out.WriteUint32(uint32(len(eff.Applications)))

	for _, app := range eff.Applications {
		switch eff.Kind {
		case EffectPowerDrain, EffectPowerBurn:
			app.Target.EncodePacked(build, out)
			out.WriteUint32(app.PowerTaken)
			out.WriteUint32(app.PowerType)
			out.WriteFloat32(app.Multiplier)
		case EffectAddExtraAttacks:
			app.Target.EncodePacked(build, out)
			out.WriteUint32(app.Amount)
		case EffectInterruptCast:
			app.Target.EncodePacked(build, out)
			out.WriteUint32(app.InterruptedSpell)
		case EffectDurabilityDamage:
			app.Target.EncodePacked(build, out)
			out.WriteUint32(app.Item)
			out.WriteUint32(app.Slot)
		case EffectOpenLock:
			app.Target.EncodePacked(build, out)
		case EffectCreateItem, EffectCreateRandomItem, EffectCreateItem2:
			out.WriteUint32(app.Item)
		default:
			return fmt.Errorf("spell: cannot encode unknown effect kind %d", eff.Kind)
		}
	}

	return nil
}

type LogExecute struct {
	Caster  guid.GUID
	Spell   uint32
	Effects []LoggedEffect
}

func (exe *LogExecute) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_SPELLLOGEXECUTE
	exe.Caster.EncodePacked(build, out)

	out.WriteUint32(uint32(exe.Spell))

	out.WriteUint32(uint32(len(exe.Effects)))

	for _, effect := range exe.Effects {
		if err := effect.Encode(build, out); err != nil {
			return err
		}
	}

	return nil
}
