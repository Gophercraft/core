package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func DefaultLocation() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "Gophercraft")
	default:
		return filepath.Join(os.Getenv("HOME"), ".local", "share", "Gophercraft")
	}
}
