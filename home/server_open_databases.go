package home

import (
	"reflect"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/phylactery/database"
)

func (server *Server) open_databases() (err error) {
	// Open phylactery database
	server.home_database, err = database.Open(server.home_config.File.DatabasePath, database.WithEngine(server.home_config.File.DatabaseEngine))
	if err != nil {
		return
	}

	// Sync table structures
	schemas := []struct {
		table  string
		schema reflect.Type
	}{
		{"Account", reflect.TypeFor[models.Account]()},
		{"GameAccount", reflect.TypeFor[models.GameAccount]()},
		{"Realm", reflect.TypeFor[models.Realm]()},
		{"EnlistedRealm", reflect.TypeFor[models.EnlistedRealm]()},
		{"LoginTicket", reflect.TypeFor[models.LoginTicket]()},
		{"WebToken", reflect.TypeFor[models.WebToken]()},
		{"CharacterCount", reflect.TypeFor[models.CharacterCount]()},
		{"LastCharacterLoggedIn", reflect.TypeFor[models.LastCharacterLoggedIn]()},
	}

	for _, schema := range schemas {
		if err = server.home_database.Table(schema.table).SyncType(schema.schema); err != nil {
			return
		}
	}

	return
}
