package datapack

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// directory implements Reader
type directory struct {
	path string
}

func open_directory(path string) (reader reader, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(path)
	if err != nil {
		return
	}
	if !fi.IsDir() {
		err = fmt.Errorf("datapack: open_directory() must be directory: %s", path)
		return
	}

	d := new(directory)
	d.path = path
	reader = d
	return
}

func list_directory(path string) (list []string, err error) {
	err = filepath.Walk(path, func(walked_path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			list_path := strings.TrimPrefix(walked_path, path)
			list_path = strings.TrimLeft(list_path, "\\/")
			if runtime.GOOS == "windows" {
				list_path = strings.Replace(list_path, "\\", "/", -1)
			}
			list = append(list, list_path)
		}
		return nil
	})

	return
}

func (d *directory) List() []string {
	files, err := list_directory(d.path)

	if err != nil {
		panic(err)
	}

	return files
}

func (d *directory) Open(path string) (file io.ReadCloser, err error) {
	realpath := filepath.Join(d.path, path)
	return os.Open(realpath)
}

func (d *directory) Close() (err error) {
	return
}
