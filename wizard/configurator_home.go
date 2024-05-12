package wizard

import (
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/phylactery/database"
)

type HomeConfigurator struct {
	Configurator *Configurator
	Config       *config.Home
}

func (co *Configurator) NewHomeConfigurator() *HomeConfigurator {
	hc := new(HomeConfigurator)
	hc.Configurator = co
	hc.Config = &config.Home{}
	return hc
}

func (hc *HomeConfigurator) CreateDB() error {
	return nil
}

func (hc *HomeConfigurator) RegisterAccount(username, password string, tier auth.AccountTier) error {
	var db_path = hc.Config.File.DatabasePath
	switch hc.Config.File.DatabaseEngine {
	case "leveldb_core":
		if !filepath.IsAbs(db_path) {
			// make path relative to config path
			db_path = filepath.Join(hc.Config.Directory, hc.Config.File.DatabasePath)
		}
	}

	db, err := database.Open(db_path, database.WithEngine(hc.Config.File.DatabaseEngine))
	if err != nil {
		return err
	}

	if err = db.Table("Account").Sync(new(models.Account)); err != nil {
		return err
	}
	if err = db.Table("GameAccount").Sync(new(models.GameAccount)); err != nil {
		return err
	}

	if err = login.RegisterAccount(db, true, auth.AccountTier_ADMIN, "", username, password, tier); err != nil {
		return err
	}

	return db.Close()
}

func (hc *HomeConfigurator) GenerateConfig() error {
	if err := os.MkdirAll(hc.Config.Directory, 0700); err != nil {
		return err
	}

	data := MakeDefaultHomeConfig(hc.Config)
	if err := os.WriteFile(filepath.Join(hc.Config.Directory, "home.txt"), data, 0700); err != nil {
		return err
	}

	if err := config.GenerateKeyPair(hc.Config.Directory); err != nil {
		return err
	}

	return nil
}
