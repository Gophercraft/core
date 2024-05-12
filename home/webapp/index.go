package webapp

import (
	"embed"
	"io/fs"
	"log"
)

// assets hold our web assets
//
//go:embed public/*
var webapp_public embed.FS

func FileServer() fs.FS {
	filesystem, err := fs.Sub(webapp_public, "public")
	if err != nil {
		log.Fatal(err)
	}
	return filesystem
}
