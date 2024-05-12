package datapack

import (
	"archive/zip"
	"fmt"
	"io"
	"strings"
)

var (
	ErrFileNotFound = fmt.Errorf("datapack: file not found")
)

// archive implements Driver

// .zip
type archive struct {
	// the path to the archive
	path string
	// the path of the directory where the pack actually resides.
	// This is found by scanning where "Pack.txt" is.
	// When you download a ZIP from GitHub or other sources your pack is probably put into its own little container directory
	// so prefix = "container/"
	// if we have a zipped file called "container/Pack.txt"
	prefix     string
	zip_reader *zip.ReadCloser
}

func (a *archive) find_prefix() (err error) {
	for _, file := range a.zip_reader.File {
		if file.Name == "Pack.txt" {
			return
		}

		if file.FileInfo().IsDir() {
			// Top level directory
			if strings.HasSuffix(file.Name, "/") && strings.Count(file.Name, "/") == 1 {
				a.prefix = file.Name
				break
			}
		}
	}

	return
}

func open_archive(path string) (reader reader, err error) {
	a := new(archive)
	a.path = path
	a.zip_reader, err = zip.OpenReader(a.path)
	if err != nil {
		return
	}

	err = a.find_prefix()
	return
}

func (a *archive) Open(path string) (file io.ReadCloser, err error) {
	prefixed_path := a.prefix + path

	for _, zipped_entry := range a.zip_reader.File {
		if prefixed_path == zipped_entry.Name {
			return zipped_entry.Open()
		}
	}

	err = fmt.Errorf("datapack: (*archive).Open: cannot find file with path '%s'", path)
	return
}

func (a *archive) Close() error {
	// nothing to do here, really
	return nil
}

func (a *archive) List() (path_list []string) {
	for _, zipped_entry := range a.zip_reader.File {
		if !zipped_entry.FileInfo().IsDir() {
			// entry is not a file

			// Remove prefix directory
			valid_filepath := strings.TrimPrefix(zipped_entry.Name, a.prefix)

			path_list = append(path_list, valid_filepath)
		}
	}

	return
}
