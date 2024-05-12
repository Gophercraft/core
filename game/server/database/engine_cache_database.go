package database

import "github.com/Gophercraft/phylactery/database"

func (engine *Engine) Cache() *database.Container {
	return engine.cache_db_container
}
