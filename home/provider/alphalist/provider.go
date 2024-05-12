package alphalist

import (
	"fmt"
	"time"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/service/alphalist"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
)

const (
	color_green   = "FF00FF00"
	color_cyan    = "FF00FFDD"
	color_orange  = "FFEB5B34"
	color_magenta = "FFFB00FF"
)

type ServiceConfig struct {
	// The build that the alphalist needs to display
	Build version.Build
}

type service_provider struct {
	config *ServiceConfig
	db     *database.Container
}

func New(config *ServiceConfig, db *database.Container) (provider *service_provider) {
	provider = new(service_provider)
	provider.config = config
	provider.db = db
	return
}

func (provider *service_provider) format_realm_name(realm *models.Realm) (name string, err error) {
	type_name := realm.Type.String()
	type_color := color_green
	suffix := ""

	switch realm.Type {
	case config.RealmTypeRP:
		type_color = color_cyan
	case config.RealmTypePvP:
		type_color = color_orange
	case config.RealmTypeNormal:
		type_color = color_green
	}

	// check if offline:
	if time.Since(realm.LastUpdated) > models.RealmOnlineDelayMeansOffline {
		suffix = " |cFFFF0000(Offline)|r"
	}

	name = fmt.Sprintf("|c%s[%s]| %s%s", type_color, type_name, realm.Name, suffix)
	return
}

func (provider *service_provider) ListRealms() (formatted_list []alphalist.Realm, err error) {
	var list []models.Realm
	// Query realms from database
	if err = provider.db.Table("Realm").
		Where(query.Eq("Build", provider.config.Build)).
		OrderBy("ID", false).
		Find(&list); err != nil {
		return
	}

	// Build database realm info into a list that gets showed in the alpha realm list screen
	formatted_list = make([]alphalist.Realm, len(list))
	for index := range list {
		realm := &list[index]
		formatted_realm := &formatted_list[index]
		formatted_realm.Name, err = provider.format_realm_name(realm)
		if err != nil {
			return
		}

		formatted_realm.Players = realm.ActivePlayers
		formatted_realm.RedirectAddress = realm.RedirectAddress
	}

	return
}
