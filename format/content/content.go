// Package content offers a simplified API for accessing game data archives
package content

import (
	"fmt"

	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/core/vsn/detection"
)

type Volume interface {
	Build() vsn.Build
	ListFiles() ([]string, error)
	ReadFile(at string) ([]byte, error)
	Close() error
}

func Open(path string) (Volume, error) {
	v, err := detection.DetectGame(path)
	if err != nil {
		return nil, err
	}

	if v == 0 {
		return nil, fmt.Errorf("cannot read from a game with version: %d", v)
	}

	vt, path2, err := detection.DetectVolumeLocation(path)
	if err != nil {
		return nil, err
	}

	switch vt {
	case detection.NGDP:
		return nil, fmt.Errorf("NGDP nyi")
	case detection.MPQ:
		return openMpq(v, path2)
	default:
		return nil, fmt.Errorf("unknown folder type")
	}
}
