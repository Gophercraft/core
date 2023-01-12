package realm

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
)

type UnitNotification uint8

const (
	UnitNotifyCreate UnitNotification = iota
	UnitNotifyDelete
	UnitNotifyMovement
	UnitNotifyOutOfRange
)

// Objects that can cast and receive spells
type Unit interface {
	WorldObject

	Target(guid.GUID)
	GetTarget() guid.GUID
	GetAuras() *AuraState
	Notify(UnitNotification, WorldObject, ...any)
	// NotifyMovement(packet.WorldType, WorldObject)
}

type UnitSet []Unit

func (wos WorldObjectSet) Units() UnitSet {
	var ss UnitSet

	for _, v := range wos {
		if s, ok := v.(Unit); ok {
			ss = append(ss, s)
		}
	}

	return ss
}

// will panic if set contains non-units
func (wos WorldObjectSet) mapallunits() UnitSet {
	ss := make(UnitSet, len(wos))
	for i, v := range wos {
		ss[i] = v.(Unit)
	}
	return ss
}

func (us UnitSet) Notify(notif UnitNotification, wo WorldObject, args ...any) {
	for _, v := range us {
		v.Notify(notif, wo, args)
	}
}

func (us UnitSet) NotifyMovement(movementType packet.WorldType, wo WorldObject) {
	for _, v := range us {
		v.Notify(UnitNotifyMovement, wo, movementType)
	}
}
