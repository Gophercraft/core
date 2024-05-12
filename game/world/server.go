package world

// type Observer

type Server interface {
	NewObserver() Observer
}
