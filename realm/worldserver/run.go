// Package worldserver is the main ent
package worldserver

import (
	"fmt"
	"log"
	"path/filepath"

	"os"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/realm"
	"github.com/Gophercraft/core/vsn"

	_ "github.com/go-sql-driver/mysql"
	// Imports core plugins, needed to avoid an import cycle.
	_ "github.com/Gophercraft/core/realm/plugins/commands"
	_ "github.com/Gophercraft/core/realm/plugins/discord"
)

func runAt(path string) {
	cfg, err := config.LoadWorld(path)
	if err != nil {
		log.Fatal(err)
	}

	vsn.PrintBanner()

	log.Println("This server's fingerprint is", cfg.Fingerprint())

	log.Fatal(realm.Start(cfg))
}

func runMain(location string, args []string) {
	if len(args) == 0 {
		worlds, err := config.ListWorldConfigs(location)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Worlds: ")

		for _, world := range worlds {
			fmt.Println(" ", world)
		}
		return
	}
	configName := args[0]

	if len(os.Args) == 0 {
		fmt.Println(args, "<world config name>")
		os.Exit(0)
		return
	}

	// full path?
	if config.WorldExists(configName) {
		runAt(configName)
		return
	}

	path := filepath.Join(location, configName)

	if config.WorldExists(path) {
		runAt(path)
		return
	}

	worlds, err := config.ListWorldConfigs(location)
	if err != nil {
		log.Fatal(err)
	}

	for _, world := range worlds {
		log.Println("Existing worlds: ", world)
	}
}

func Run() {
	rootCmd.Execute()
}
