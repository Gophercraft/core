package datapack

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/text"
)

// Creator is a directory
type Creator struct {
	directory string
}

func NewCreator(path string) (creator *Creator, err error) {
	_, err = os.Stat(path)
	if err == nil {
		err = fmt.Errorf("datapack: NewCreator: path already exists: %s", path)
		return
	}

	err = os.Mkdir(path, 0700)
	if err != nil {
		return
	}

	creator = new(Creator)
	creator.directory = path
	return
}

func (creator *Creator) Open(path string) (file *os.File, err error) {
	realpath := filepath.Join(creator.directory, path)
	return os.Open(realpath)
}

func (creator *Creator) WritePackInfo(pack_info *PackInfo) (err error) {
	var file *os.File
	file, err = creator.Open("Pack.txt")
	if err != nil {
		return
	}

	text_encoder := text.NewEncoder(file)
	err = text_encoder.Encode(pack_info)
	if err != nil {
		return
	}

	err = file.Close()
	return
}

func (creator *Creator) Create(archive_path string) (err error) {
	if _, err = os.Stat(archive_path); err == nil {
		if err = os.Remove(archive_path); err != nil {
			return
		}
	}

	var archive_file *os.File
	var list []string

	list, err = list_directory(creator.directory)
	if err != nil {
		return
	}

	archive_file, err = os.Create(archive_path)
	if err != nil {
		return err
	}
	defer archive_file.Close()

	zip_writer := zip.NewWriter(archive_file)
	defer zip_writer.Close()

	// Add files to zip
	for _, filepath := range list {
		var file *os.File
		file, err = creator.Open(filepath)
		if err != nil {
			return
		}

		if err = add_file_to_zip(zip_writer, filepath, file); err != nil {
			return err
		}

		file.Close()
	}
	return
}
