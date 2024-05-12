package content

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gophercraft/core/version"
)

type plain_directory struct {
	build     version.Build
	directory string
}

func (plain_directory *plain_directory) Build() version.Build {
	return plain_directory.build
}

func (plain_directory *plain_directory) Type() Type {
	return Plain
}

func (plain_directory *plain_directory) Close() error {
	return nil
}

func (plain_directory *plain_directory) Open(path string) (file io.ReadCloser, err error) {
	converted_path := strings.ReplaceAll(path, "\\", string(os.PathSeparator))
	realpath := filepath.Join(plain_directory.directory, converted_path)
	return os.Open(realpath)
}

func open_plain(build version.Build, path string) (directory *plain_directory, err error) {
	directory = new(plain_directory)
	directory.build = build
	directory.directory = path
	return
}
