package realm

type Notification uint8

const (
	NotifyCreate Notification = iota
	NotifyDelete
	NotifyMovement
	NotifyOutOfRange
)

type Subscriber interface {
	Notify(notif Notification, object WorldObject, args ...any)
}
