package datapack

import (
	"fmt"
	"io"
	"strings"
)

type Loader struct {
	Volumes []*Pack
}

func (ld *Loader) Exists(path string) bool {
	colon := strings.Index(path, ":")
	if colon == -1 {
		for _, pack := range ld.Volumes {
			if pack.Exists(path) {
				fmt.Println(path, "does exist in", pack.Name)
				return true
			}
			fmt.Println(path, "doesn't exist in", pack.Name)
		}
	} else {
		pack := path[:colon]
		filePath := path[colon:]

		for _, dpack := range ld.Volumes {
			if dpack.Name == pack {
				if dpack.Exists(filePath) {
					return true
				}
			}
		}
	}

	return false
}

// func (ld *Loader) Open(path string) (io.ReadCloser, error) {
// 	seg := strings.SplitN(path, ":", 2)
// 	for _, pack := range ld.Volumes {
// 		if pack.Name == seg[0] {
// 			if pack.Exists(seg[1]) {
// 				return pack.ReadFile(seg[1])
// 			}
// 		}
// 	}

// 	return nil, fmt.Errorf("file %s not found", path)
// }

func (ld *Loader) ReadFile(path string) (io.ReadCloser, error) {
	colon := strings.IndexByte(path, ':')
	if colon == -1 {
		for _, pack := range ld.Volumes {
			if pack.Exists(path) {
				return pack.ReadFile(path)
			}
		}
	} else {
		pack := path[:colon]
		filePath := path[colon:]

		for _, dpack := range ld.Volumes {
			if dpack.Name == pack {
				if dpack.Exists(filePath) {
					return dpack.ReadFile(filePath)
				}
			}
		}
	}

	return nil, fmt.Errorf("file %s not found", path)
}

func (ld *Loader) Close() {
	for _, v := range ld.Volumes {
		if err := v.Close(); err != nil {
			panic(err)
		}
	}

	ld.Volumes = nil
}

func (ld *Loader) List() []string {
	var names []string

	for _, volume := range ld.Volumes {
		for _, volumeListing := range volume.List() {
			found := false
			for _, name := range names {
				if name == volumeListing {
					found = true
				}
			}
			if !found {
				names = append(names, volumeListing)
			}
		}
	}

	return names
}
