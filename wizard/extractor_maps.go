package wizard

import (
	"fmt"
	"path/filepath"

	"github.com/Gophercraft/core/datapack"
	"github.com/Gophercraft/core/version"
)

const MapPackName = "!maps.zip"

func (ex *Extractor) ExtractMaps() error {
	const tempPackDir = "x-maps"

	if ex.packExists(MapPackName) {
		ex.removePack(MapPackName)
	}

	pack, err := ex.AuthorPack(tempPackDir, &datapack.PackInfo{
		Name: "Gophercraft Base Map Data",
		Authors: []string{
			fmt.Sprintf("Generated by %s", get_extractor_author()),
		},
		Version:            0,
		Base:               true,
		Description:        "The base map data required by Gophercraft Core (Namigator)" + ex.generationNotice(),
		MinimumCoreVersion: version.GophercraftVersion.String(),
	})

	if err != nil {
		return err
	}

	if err := pack.Create(filepath.Join(ex.Dir, MapPackName)); err != nil {
		ex.removePack(tempPackDir)
		return err
	}

	return ex.removePack(tempPackDir)
}
