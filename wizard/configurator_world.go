package wizard

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/format/content"
	"github.com/Gophercraft/core/version/detection"
)

type WorldOptions struct {
	RealmName string
}

type WorldConfigurator struct {
	Configurator *Configurator
	Config       *config.World
	GamePath     string
	ConfigName   string
}

func (co *Configurator) NewWorldConfigurator() *WorldConfigurator {
	wc := new(WorldConfigurator)
	wc.Configurator = co
	wc.Config = &config.World{}
	return wc
}

func (w *WorldConfigurator) ConfigExists() bool {
	_, err := os.Stat(w.Config.Directory)
	return err == nil
}

func (w *WorldConfigurator) RemoveConfigDir() error {
	if w.Config.Directory == "" {
		return fmt.Errorf("wizard: can't remove empty directory")
	}

	return os.RemoveAll(w.Config.Directory)
}

func (w *WorldConfigurator) GenerateConfigDir() error {
	if err := os.MkdirAll(w.Config.Directory, 0700); err != nil {
		return err
	}

	if err := os.MkdirAll(
		filepath.Join(w.Config.Directory, "Datapacks"),
		0700); err != nil {
		return err
	}

	if err := config.GenerateKeyPair(w.Config.Directory); err != nil {
		return err
	}

	return nil
}

func (wc *WorldConfigurator) SetVolume(path string) error {
	var err error
	wc.GamePath = path
	wc.Config.Build, err = detection.DetectGame(wc.GamePath)
	return err
}

func (wc *WorldConfigurator) SetName(name string) {
	wc.Config.LongName = name
}

func (wc *WorldConfigurator) SetRealmType(t config.RealmType) error {
	wc.Config.Type = t
	return nil
}

// func (wc *WorldConfigurator) BeginEnlistLocal() error {
// 	wc.Registrar = &LocalHomeserver{
// 		wc.Configurator,
// 	}
// 	return nil
// }

func (wc *WorldConfigurator) CreateDB() error {
	return nil
	// return dbsupport.Create(wc.Config.DBDriver, &wc.Config.DBURL)
}

func (w *WorldConfigurator) NewExtractor() (*Extractor, error) {
	ex := new(Extractor)
	ex.Dir = filepath.Join(w.Config.Directory, "Datapacks")
	var err error
	ex.Source, err = content.Open(w.GamePath)
	if err != nil {
		return nil, err
	}
	return ex, nil
}
