// Package replay will provide utilities for recording and playing back game logs
package replay

import "github.com/Gophercraft/core/version"

type Version uint8

const (
	V1 Version = '1'
)

type Header struct {
	Version
	Build version.Build
	Map   uint32
}
