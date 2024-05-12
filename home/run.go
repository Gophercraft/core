package home

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/log"
)

func Run(home_config *config.Home) (err error) {
	server := new(Server)
	server.home_config = home_config

	// cd to home config directory
	if err = os.Chdir(server.home_config.Directory); err != nil {
		return
	}

	// validate configuration
	if err = server.validate_config(); err != nil {
		return
	}

	// open database container
	if err = server.open_databases(); err != nil {
		return
	}

	// launch network services listed in config
	if err = server.launch_services(); err != nil {
		return
	}

	// Wait for Ctrl-C
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.Warn("Shutting down...")

	// Stop services
	if err = server.stop_services(); err != nil {
		return
	}

	// Close database
	if err = server.close_databases(); err != nil {
		return
	}

	return
}
