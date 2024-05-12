package datapack

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/Gophercraft/text"
)

type text_loader_path struct {
	pack *Pack
	path string
}

// TextDatabaseLoader loads all the contents of a text database
// from multiple packs, obeying the rules of each pack
// according to load order
type TextDatabaseLoader struct {
	loader     *Loader
	file       io.Closer
	decoder    *text.Decoder
	index      int
	text_paths []text_loader_path
}

func (text_loader *TextDatabaseLoader) push_text_path(pack *Pack, path string) {
	text_loader.text_paths = append(text_loader.text_paths, text_loader_path{
		pack: pack,
		path: path,
	})
}

func NewTextDatabaseLoader(loader *Loader, table_name string) (text_loader *TextDatabaseLoader, err error) {
	text_loader = new(TextDatabaseLoader)
	text_loader.loader = loader

	// Exclude packs that come before
	ignore_packs := make(map[int]bool)

	for pack_index_in_loader, pack := range loader.packs {
		for _, override := range pack.info.OverrideTables {
			if override == table_name {
				for x := pack_index_in_loader - 1; x >= 0; x-- {
					ignore_packs[x] = true
				}
				break
			}
		}
	}

	// Treat only these packs as a valid source
	// for populating this database table
	var include_packs []*Pack

	for pack_index_in_loader, pack := range loader.packs {
		if !ignore_packs[pack_index_in_loader] {
			include_packs = append(include_packs, pack)
		}
	}

	prefix := "DB/" + table_name
	regular_text := prefix + ".txt"
	directory_prefix := prefix + "/"

	for _, pack := range include_packs {
		list := pack.List()
		for _, list_entry := range list {
			if list_entry == regular_text {
				text_loader.push_text_path(pack, list_entry)
			} else if strings.HasPrefix(list_entry, directory_prefix) {
				text_loader.push_text_path(pack, list_entry)
			}
		}
	}

	return
}

func (text_loader *TextDatabaseLoader) Load(record any) (err error) {
	if text_loader.index == len(text_loader.text_paths) {
		err = io.EOF
		return
	}

	current := &text_loader.text_paths[text_loader.index]

	if text_loader.decoder == nil {
		var file io.ReadCloser
		file, err = current.pack.Open(current.path)
		if err != nil {
			return
		}

		text_loader.file = file
		text_loader.decoder = text.NewDecoder(file)
	}

	// Attempt to decode record
	if err = text_loader.decoder.Decode(record); err == nil {
		return
	}

	if errors.Is(err, io.EOF) {
		text_loader.file.Close()
		text_loader.decoder = nil

		text_loader.index++

		if text_loader.index == len(text_loader.text_paths) {
			err = io.EOF
		} else {
			err = fmt.Errorf("datapack: unexpected EOF")
		}
		return
	}

	return
}

// // NewTextDatabaseLoader returns a TextDatabaseLoader of all text files in all datapacks matching the supplied parameters
// func (ld *Loader) NewTextDatabaseLoader(prefix, name string) (*TextDatabaseLoader, error) {
// 	tl := &TextDatabaseLoader{}
// 	tl.Loader = ld

// 	// Whether a pack table is overridden depends on load order.
// 	packDisabled := map[int]bool{}

// 	for i, pack := range ld.Volumes {
// 		for _, ot := range pack.OverrideTables {
// 			if ot == name {
// 				for x := i - 1; x >= 0; x-- {
// 					packDisabled[x] = true
// 				}
// 			}
// 		}
// 	}

// 	for i, pack := range ld.Volumes {
// 		folderPath := prefix + "/" + name + "/"
// 		filePath := prefix + "/" + name + ".txt"

// 		if packDisabled[i] {
// 			continue
// 		}

// 		// DB/Table/*.txt
// 		if pack.FolderExists(folderPath) {
// 			list := pack.FolderList(folderPath)
// 			for _, file := range list {
// 				tl.Paths = append(tl.Paths, packText{
// 					PackIndex: i,
// 					Path:      file,
// 				})
// 			}
// 		}

// 		// DB/Table.txt
// 		if pack.Exists(filePath) {
// 			tl.Paths = append(tl.Paths, packText{
// 				PackIndex: i,
// 				Path:      filePath,
// 			})
// 		}

// 	}

// 	return tl, nil
// }

// func (tl *TextDatabaseLoader) Scan(out interface{}) error {
// 	for {
// 		if tl.Decoder == nil {
// 			if len(tl.Paths) == 0 {
// 				return io.EOF
// 			}

// 			// Next
// 			nextFile := tl.Paths[0]
// 			tl.Paths = tl.Paths[1:]
// 			file, err := tl.Loader.Volumes[nextFile.PackIndex].ReadFile(nextFile.Path)
// 			if err != nil {
// 				return err
// 			}

// 			tl.Decoder = text.NewDecoder(file)
// 		}

// 		err := tl.Decoder.Decode(out)
// 		if err != nil {
// 			if err == io.EOF {
// 				tl.Decoder = nil
// 				continue
// 			} else {
// 				return err
// 			}
// 		}
// 		return nil
// 	}
// }
