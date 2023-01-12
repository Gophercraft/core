package wizard

import (
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/models"
	"xorm.io/xorm"
)

// Uses locally-stored Home config to register a World
// Has the advantage that the home server doesn't have to be online to work
type LocalHomeConfig struct {
	*WorldConfigurator

	HomeConfigName string
}

func (l *LocalHomeConfig) Begin(wc *WorldConfigurator, name string) error {
	l.WorldConfigurator = wc
	l.HomeConfigName = name
	return nil
}

func (l *LocalHomeConfig) MustAuth() bool {
	return false
}

func (l *LocalHomeConfig) path() string {
	return filepath.Join(l.WorldConfigurator.Configurator.Dir, l.HomeConfigName)
}

func (l *LocalHomeConfig) Check() (*ValidationCheck, error) {
	if _, err := os.Stat(l.path()); err != nil {
		return nil, err
	}

	return nil, nil
}

func (l *LocalHomeConfig) Confirm(check *ValidationCheck) error {
	return nil
}

func (l *LocalHomeConfig) Auth(username, password string) error {
	return nil
}

func (l *LocalHomeConfig) Enlist() error {
	home, err := config.LoadHome(l.path())
	if err != nil {
		return err
	}

	l.WorldConfigurator.Config.HomeServer = home.AuthListen
	l.WorldConfigurator.Config.HomeServerFingerprint = home.Fingerprint()

	engi, err := xorm.NewEngine(home.DBDriver, home.DBURL)
	if err != nil {
		return err
	}
	defer engi.Close()

	engi.Sync(
		new(models.EnlistedRealm),
	)

	var er models.EnlistedRealm
	er.Fingerprint = l.WorldConfigurator.Config.Fingerprint()
	er.Note = l.WorldConfigurator.ConfigName
	er.Owner = 0

	if _, err := engi.Insert(&er); err != nil {
		return err
	}

	l.WorldConfigurator.Config.RealmID = er.ID
	return nil
}
