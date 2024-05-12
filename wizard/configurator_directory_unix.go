//go:build !windows

package wizard

import (
	"os"
	"path/filepath"
)

func get_directory() string {
	return filepath.Join(os.Getenv("HOME"), ".local", "share", "gophercraft")
}
