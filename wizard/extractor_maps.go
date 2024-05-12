package wizard

import (
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

	pack, err := ex.AuthorPack(tempPackDir, datapack.PackConfig{
		Name:           "Gophercraft Base Maps",
		Author:         exAuthor,
		Description:    "The base map data required by Gophercraft Core." + ex.generationNotice(),
		PackVersion:    "1.0",
		MinCoreVersion: version.GophercraftVersion.String(),
		Depends:        ex.dependencies(),
	})

	if err != nil {
		return err
	}

	if err := pack.ZipToFile(filepath.Join(ex.Dir, MapPackName)); err != nil {
		ex.removePack(tempPackDir)
		return err
	}

	return ex.removePack(tempPackDir)
}
