package database

import "github.com/Gophercraft/phylactery/database"

func (engine *Engine) World() *database.Container {
	return engine.world_db_container
}
