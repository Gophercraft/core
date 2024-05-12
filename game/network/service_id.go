package network

import "sync"

type ServiceID uint64

var (
	registry           = make(map[ServiceID]string)
	counter  ServiceID = 0
	guard    sync.Mutex
)

func RegisterServiceID(name string) ServiceID {
	guard.Lock()
	counter++
	id := counter
	registry[id] = name
	guard.Unlock()
	return id
}

func (id ServiceID) String() string {
	return registry[id]
}
