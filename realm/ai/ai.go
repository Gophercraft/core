package ai

import "github.com/Gophercraft/core/realm/sync"

type Action interface {
}

type Observer interface {
	Notify(notif Notification, object sync.WorldObject, args ...any)
}

// Describes how a Unit can be controlled by the computer
type UnitAI interface {
	HandleNotify()
	NextAction() Action
}

type Notification uint8

const (
	NotifyCreate Notification = iota
	NotifyDelete
	NotifyMovement
	NotifyOutOfRange
)

type Subscriber interface {
}
