package home

import (
	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/phylactery/database"
)

type Service interface {
	ID() config.HomeServiceID
	Start() error
	Stop() error
}

type Server struct {
	home_config   *config.Home
	home_database *database.Container
	services      []Service
}
