// Package content offers a simplified API for accessing game data archives
package content

import (
	"fmt"
	"io"

	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/core/version/detection"
)

type Volume interface {
	// returns the content.Type
	Type() Type
	// returns the build associated with this Volume
	Build() (build version.Build)
	// opens a file contained within this Volume
	Open(path string) (file io.ReadCloser, err error)
	Close() (err error)
}

func Open(path string) (volume Volume, err error) {
	var (
		build        version.Build
		content_type Type
		content_path string
	)

	build, err = detection.DetectGame(path)
	if err != nil {
		return nil, err
	}

	if build == 0 {
		return nil, fmt.Errorf("cannot read from a game with version: %d", build)
	}

	content_type, content_path, err = detect_content_type(path)
	if err != nil {
		return nil, err
	}

	switch content_type {
	case NGDP:
		return nil, fmt.Errorf("NGDP not net implemented")
	case MPQ:
		return open_mpq(build, content_path)
	case Plain:
		return open_plain(build, content_path)
	default:
		return nil, fmt.Errorf("unknown folder type")
	}
}
