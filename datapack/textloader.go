package datapack

import (
	"io"

	"github.com/Gophercraft/text"
)

type packText struct {
	PackIndex int
	Path      string
}

type TextLoader struct {
	Loader  *Loader
	Decoder *text.Decoder
	Paths   []packText
}

// NewTextLoader returns a TextLoader of all text files in all datapacks matching the supplied parameters
func (ld *Loader) NewTextLoader(prefix, name string) (*TextLoader, error) {
	tl := &TextLoader{}
	tl.Loader = ld

	// Whether a pack table is overridden depends on load order.
	packDisabled := map[int]bool{}

	for i, pack := range ld.Volumes {
		for _, ot := range pack.OverrideTables {
			if ot == name {
				for x := i - 1; x >= 0; x-- {
					packDisabled[x] = true
				}
			}
		}
	}

	for i, pack := range ld.Volumes {
		folderPath := prefix + "/" + name + "/"
		filePath := prefix + "/" + name + ".txt"

		if packDisabled[i] {
			continue
		}

		// DB/Table/*.txt
		if pack.FolderExists(folderPath) {
			list := pack.FolderList(folderPath)
			for _, file := range list {
				tl.Paths = append(tl.Paths, packText{
					PackIndex: i,
					Path:      file,
				})
			}
		}

		// DB/Table.txt
		if pack.Exists(filePath) {
			tl.Paths = append(tl.Paths, packText{
				PackIndex: i,
				Path:      filePath,
			})
		}

	}

	return tl, nil
}

func (tl *TextLoader) Scan(out interface{}) error {
	if tl.Decoder == nil {
		if len(tl.Paths) == 0 {
			return io.EOF
		}

		// Next
		nextFile := tl.Paths[0]
		tl.Paths = tl.Paths[1:]
		file, err := tl.Loader.Volumes[nextFile.PackIndex].ReadFile(nextFile.Path)
		if err != nil {
			return err
		}

		tl.Decoder = text.NewDecoder(file)
	}

	err := tl.Decoder.Decode(out)
	if err != nil {
		if err == io.EOF {
			tl.Decoder = nil
			// Next call may return EOF
			return nil
		} else {
			return err
		}
	}
	return nil
}
