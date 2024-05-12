package content

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/mpq"
)

type mpq_set struct {
	directory string
	build     version.Build
	set       *mpq.Set
}

func (set *mpq_set) Build() version.Build {
	return set.build
}

func (set *mpq_set) Type() Type {
	return MPQ
}

func (set *mpq_set) Close() error {
	return set.set.Close()
}

func (set *mpq_set) Open(filepath string) (file io.ReadCloser, err error) {
	// TODO: check for the existence of listfile-less archives
	// ie $(set.directory) + filepath + ".MPQ"
	// present in the Alpha client

	file, err = set.set.Open(filepath)
	return
}

func open_mpq(build version.Build, path string) (set *mpq_set, err error) {
	set = new(mpq_set)
	set.directory = path
	set.build = build

	var glob_chain []string

	switch {
	case build <= 3368:
		glob_chain = []string{
			"model.MPQ",
			"texture.MPQ",
			"sound.MPQ",
			"misc.MPQ",
			"interface.MPQ",
			"fonts.MPQ",
			"speech.MPQ",
			"dbc.MPQ",
		}
	case build <= 5875:
		glob_chain = []string{
			"base.MPQ",
			"model.MPQ",
			"texture.MPQ",
			"terrain.MPQ",
			"wmo.MPQ",
			"sound.MPQ",
			"misc.MPQ",
			"interface.MPQ",
			"fonts.MPQ",
			"dbc.MPQ",
			"speech.MPQ",
			"speech2.MPQ",
			"patch.MPQ",
			"patch-2.MPQ",
		}
	case build <= 12340:
		glob_chain = []string{
			"common.MPQ",
			"common-2.MPQ",
			"expansion.MPQ",
			"lichking.MPQ",
			"*/locale-*.MPQ",
			"*/speech-*.MPQ",
			"*/expansion-locale-*.MPQ",
			"*/lichking-locale-*.MPQ",
			"*/expansion-speech-*.MPQ",
			"*/lichking-speech-*.MPQ",
			"*/patch-????.MPQ",
			"*/patch-*.MPQ",
			"patch.MPQ",
			"patch-*.MPQ",
		}
	default:
		err = fmt.Errorf("content: cannot yet open build %s", build)
	}

	if err != nil {
		return
	}

	full_glob_chain := make([]string, len(glob_chain))
	for i := range glob_chain {
		full_glob_chain[i] = filepath.Join(set.directory, glob_chain[i])
	}

	set.set, err = mpq.GlobSet(full_glob_chain...)
	if err != nil {
		return
	}

	return
}
