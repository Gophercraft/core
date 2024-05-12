package datapack

import "io"

// Driver describes a mechanism for loading a datapack.
type reader interface {
	Open(path string) (file io.ReadCloser, err error)
	List() []string
	Close() error
}
