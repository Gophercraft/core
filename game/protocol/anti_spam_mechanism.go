package protocol

type TraceKind uint8

const (
	NewConnection = iota
)

type TraceEvent struct {
	Kind  TraceKind
	Param any
}

type CounterAction uint8

const (
	Proceed CounterAction = iota
	Terminate
	Suspend
)

type AntiSpamMechanism interface {
	Trace(connection *Connection, te *TraceEvent) CounterAction
}
