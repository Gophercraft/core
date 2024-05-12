package detection

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gophercraft/core/version"
)

var (
	valid_exe_names = []string{
		"wowclient.exe",
		"wow.exe",
		"wow.exe",
		"wow-64.exe",
	}
)

// Returns the most likely build number of a WoW client directory
func DetectGame(game_directory string) (build version.Build, err error) {
	var (
		directory_info os.FileInfo
		directory_list []fs.DirEntry
	)

	// Stat directory
	directory_info, err = os.Stat(game_directory)
	if err != nil {
		return
	}

	// fail if it's not actually a directory
	if !directory_info.IsDir() {
		err = fmt.Errorf("version: game path '%s' isn't a directory", game_directory)
		return
	}

	// list directory contents
	directory_list, err = os.ReadDir(game_directory)
	if err != nil {
		return
	}

	// look for wow.app
	for _, directory_entry := range directory_list {
		name := strings.ToLower(directory_entry.Name())
		if name == "wow.app" && directory_entry.IsDir() {
			return detect_osx_app_build(filepath.Join(game_directory, directory_entry.Name()))
		}
	}

	// look for windows pe/exe files
	found_name := ""
	for _, directory_entry := range directory_list {
		name := strings.ToLower(directory_entry.Name())
		for _, possible_name := range valid_exe_names {
			if name == possible_name {
				found_name = directory_entry.Name()
				break
			}
		}
		if found_name != "" {
			break
		}
	}

	if found_name == "" {
		err = fmt.Errorf("detection: could not detect either .app or .exe in game folder")
		return
	}
	found_exe_path := filepath.Join(game_directory, found_name)
	return detect_exe_build(found_exe_path)
}
