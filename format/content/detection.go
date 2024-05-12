package content

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Type uint8

const (
	// volume is a flat file directory
	Plain Type = iota
	// volume's files are contained in MPQ archives
	MPQ
	// volume's files are contained in a CASC / NGDP container
	NGDP
)

func detect_content_type(volume_path string) (content_type Type, content_path string, err error) {
	var (
		volume_info      os.FileInfo
		volume_data_path string
		volume_top_list  []fs.DirEntry
		volume_data_list []fs.DirEntry
	)
	if len(volume_path) == 0 {
		err = fmt.Errorf("content: invalid path")
		return
	}

	volume_info, err = os.Stat(volume_path)
	if err != nil {
		return
	}
	// the volume has to be a directory, at least
	if !volume_info.IsDir() {
		err = fmt.Errorf("content: volume '%s' is not a directory", volume_path)
		return
	}

	// Construct a path to the data folder
	volume_top_list, err = os.ReadDir(volume_path)
	if err != nil {
		return
	}

	// Look for folders with either the name "Data" or "data"
	for _, dirent := range volume_top_list {
		if strings.ToLower(dirent.Name()) == "data" {
			volume_data_path = filepath.Join(volume_path, dirent.Name())
			break
		}
	}

	if volume_data_path == "" {
		// the "data" directory was not found
		// it may be that the directory is INVALID,
		// but this is also the condition of a fully extracted MPQ set.
		content_type = Plain
		content_path = volume_path
		return
	}

	// Begin with  the assumption that .../data/ is a CAS/NGDP container.
	content_type = NGDP
	content_path = volume_data_path

	volume_data_list, err = os.ReadDir(volume_data_path)
	if err != nil {
		return
	}

	for _, dirent := range volume_data_list {
		name := strings.ToLower(dirent.Name())
		// reject the NGDP assumption when we encounter anything with a ".MPQ" filetype.
		if strings.HasSuffix(name, ".mpq") {
			content_type = MPQ
			content_path = volume_data_path
			return
		}
	}

	return
}
