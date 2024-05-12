package home

import (
	"github.com/Gophercraft/core/home"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
)

func run_main() {
	version.PrintBanner()

	home_config, err := get_config()
	if err != nil {
		log.Fatal(err)
	}

	if err = home.Run(home_config); err != nil {
		log.Fatal(err)
	}
}
