package world

import (

	"github.com/Gophercraft/log"
	"github.com/Gophercraft/core/home"
	game_server "github.com/Gophercraft/core/game/server"
)

func run_main() {
	version.PrintBanner()

	world_config, err := get_config()
	if err != nil {
		log.Fatal(err)
	}

	

	if err = home.Run(home_config); err != nil {
		log.Fatal(err)
	}
}
