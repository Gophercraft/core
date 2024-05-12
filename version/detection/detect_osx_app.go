package detection

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gophercraft/core/version"
	"github.com/groob/plist"
)

type osx_app_info struct {
	Version string `plist:"BlizzardFileVersion"`
}

func detect_osx_app_build(osx_app_path string) (build version.Build, err error) {
	var (
		contents_directory      os.FileInfo
		contents_directory_list []fs.DirEntry
		app_directory_list      []fs.DirEntry
	)

	app_directory_list, err = os.ReadDir(osx_app_path)
	if err != nil {
		return
	}

	contents_name := ""

	for _, directory_entry := range app_directory_list {
		name := strings.ToLower(directory_entry.Name())
		if name == "contents" {
			contents_name = directory_entry.Name()
			break
		}
	}
	if contents_name == "" {
		err = fmt.Errorf("detection: malformed app (no contents)")
		return
	}

	osx_app_contents_path := filepath.Join(osx_app_path, contents_name)
	contents_directory, err = os.Stat(osx_app_contents_path)
	if err != nil {
		return
	}
	if !contents_directory.IsDir() {
		err = fmt.Errorf("detection: malformed contents")
	}
	contents_directory_list, err = os.ReadDir(osx_app_contents_path)
	if err != nil {
		return
	}
	plist_name := ""
	for _, directory_entry := range contents_directory_list {
		name := strings.ToLower(directory_entry.Name())
		if name == "info.plist" {
			plist_name = directory_entry.Name()
			break
		}
	}
	if plist_name == "" {
		err = fmt.Errorf("detection: malformed app (no info.plist)")
		return
	}

	var plist_file []byte
	osx_app_plist_path := filepath.Join(osx_app_contents_path, plist_name)
	plist_file, err = os.ReadFile(osx_app_plist_path)
	if err != nil {
		return
	}

	var info osx_app_info
	err = plist.Unmarshal(plist_file, &info)
	if err != nil {
		return
	}

	if info.Version == "" {
		err = fmt.Errorf("detection: could not read version info from info.plist")
		return
	}

	return version.ParseDBD(info.Version)
}
