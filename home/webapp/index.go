package webapp

import (
	"embed"
	"io"
	"io/fs"
)

// assets hold our web assets
//
//go:embed assets/*
var assets embed.FS
var Assets fs.FS

// templates holds our templates
//
//go:embed template/*
var templates embed.FS
var Templates fs.FS

func init() {
	var err error
	Templates, err = fs.Sub(templates, "template")
	if err != nil {
		panic(err)
	}

	Assets, err = fs.Sub(assets, "assets")
	if err != nil {
		panic(err)
	}
}

func ReadEmbedded(filesys fs.FS, path string) ([]byte, error) {
	f, err := filesys.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(f)
	f.Close()
	return data, err
}
