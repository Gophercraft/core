package autocode

import (
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/vsn"
)

type Generator struct {
	Build vsn.Build
	Dir   string

	layouts []*layoutTarget
}

func NewGenerator(build vsn.Build, path string) *Generator {
	g := new(Generator)
	g.Build = build
	g.Dir = path
	return g
}

func (g *Generator) path(path string) string {
	return filepath.Join(g.Dir, path)
}

func (g *Generator) ensurePath(relpath string) error {
	path := g.path(relpath)

	dirContainingPath := filepath.Dir(path)

	if _, err := os.Stat(dirContainingPath); err != nil {
		return os.MkdirAll(dirContainingPath, 0700)
	}

	return nil
}

func (g *Generator) Generate() error {
	if err := g.generateLayouts(); err != nil {
		return err
	}

	if err := g.generateLayoutReaders(); err != nil {
		return err
	}

	if err := g.generateStaticLoader(); err != nil {
		return err
	}

	return nil
}
