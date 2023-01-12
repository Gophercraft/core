package wizard

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/dbsupport"
	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
	"xorm.io/xorm"
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
	return dbsupport.Create(hc.Config.DBDriver, hc.Config.DBURL)
}

func (hc *HomeConfigurator) RegisterAccount(username, password string, tier rpcnet.Tier) error {
	engi, err := xorm.NewEngine(hc.Config.DBDriver, hc.Config.DBURL)
	if err != nil {
		return err
	}

	if err = engi.Sync(
		new(models.Account),
		new(models.GameAccount),
	); err != nil {
		return err
	}

	if err = login.RegisterAccount(engi, username, password, tier); err != nil {
		return err
	}

	return engi.Close()
}

func (hc *HomeConfigurator) SetDirName(dirName string) {
	hc.Config.Dir = filepath.Join(hc.Configurator.Dir, dirName)
}

func (hc *HomeConfigurator) GenerateConfig() error {
	if err := os.MkdirAll(hc.Config.Dir, 0700); err != nil {
		return err
	}

	data := MakeDefaultHomeConfig(hc.Config)
	if err := ioutil.WriteFile(filepath.Join(hc.Config.Dir, "Home.txt"), data, 0700); err != nil {
		return err
	}

	if err := hc.Config.GenerateKeyPair(); err != nil {
		return err
	}

	return nil
}
