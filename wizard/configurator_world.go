package wizard

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/format/content"
	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/dbsupport"
	"github.com/Gophercraft/core/vsn/detection"
)

type WorldOptions struct {
	RealmName string
}

type WorldConfigurator struct {
	Configurator *Configurator
	Config       *config.World
	GamePath     string
	ConfigName   string
	Registrar    Registrar
}

func (co *Configurator) NewWorldConfigurator() *WorldConfigurator {
	wc := new(WorldConfigurator)
	wc.Configurator = co
	wc.Config = &config.World{}
	return wc
}

func (w *WorldConfigurator) LoadConfig(confName string) (err error) {
	w.Config, err = config.LoadWorld(filepath.Join(w.Configurator.Dir, confName))
	if err != nil {
		return
	}

	return
}

func (w *WorldConfigurator) SetConfigName(confName string) {
	w.Config.Dir = filepath.Join(w.Configurator.Dir, confName)
	w.ConfigName = confName
}

func (w *WorldConfigurator) ConfigExists() bool {
	_, err := os.Stat(w.Config.Dir)
	return err == nil
}

func (w *WorldConfigurator) RemoveConfigDir() error {
	if w.Config.Dir == "" {
		return fmt.Errorf("wizard: can't remove empty directory")
	}

	return os.RemoveAll(w.Config.Dir)
}

func (w *WorldConfigurator) GenerateConfigDir() error {
	if err := os.MkdirAll(w.Config.Dir, 0700); err != nil {
		return err
	}

	if err := os.MkdirAll(
		filepath.Join(w.Config.Dir, "Datapacks"),
		0700); err != nil {
		return err
	}

	if err := w.Config.GenerateKeyPair(); err != nil {
		return err
	}

	return nil
}

func (wc *WorldConfigurator) SetVolume(path string) error {
	var err error
	wc.GamePath = path
	wc.Config.Version, err = detection.DetectGame(wc.GamePath)
	return err
}

func (wc *WorldConfigurator) SetName(name string) {
	wc.Config.RealmName = name
}

func (wc *WorldConfigurator) SetRealmType(t config.RealmType) error {
	wc.Config.RealmType = t
	return nil
}

// func (wc *WorldConfigurator) BeginEnlistLocal() error {
// 	wc.Registrar = &LocalHomeserver{
// 		wc.Configurator,
// 	}
// 	return nil
// }

func (wc *WorldConfigurator) CreateDB() error {
	return dbsupport.Create(wc.Config.DBDriver, wc.Config.DBURL)
}

func (w *WorldConfigurator) NewExtractor() (*Extractor, error) {
	ex := new(Extractor)
	ex.Dir = filepath.Join(w.Config.Dir, "Datapacks")
	var err error
	ex.Source, err = content.Open(w.GamePath)
	if err != nil {
		return nil, err
	}
	return ex, nil
}
