package world

import (
	"errors"
	"fmt"
	"os"

	"github.com/Gophercraft/core/app/config"
)

var err_no_config = fmt.Errorf("app/world: no world configuration found in the current directory. Go to docs/SETUP.md to read about creating a world configuration'")

func get_config() (world_config *config.World, err error) {
	var path string

	if _, err = os.Stat("world.txt"); err == nil {
		path, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		err = err_no_config
		return
	}

	world_config, err = config.LoadWorld(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = err_no_config
		}
		return
	}
	return
}
