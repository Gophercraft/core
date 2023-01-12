// Package replay will provide utilities for recording and playing back game logs
package replay

import "github.com/Gophercraft/core/vsn"

type Version uint8

const (
	V1 Version = '1'
)

type Header struct {
	Version
	Build vsn.Build
	Map   uint32
}
