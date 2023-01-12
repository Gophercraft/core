package wizard

import (
	"io/ioutil"
	"path/filepath"
)

func (wc *WorldConfigurator) SetRegistrar(reg Registrar) error {
	wc.Registrar = reg
	return nil
}

func (wc *WorldConfigurator) RegistrarBegin(location string) error {
	return wc.Registrar.Begin(wc, location)
}

func (w *WorldConfigurator) RegistrarAuth(username, password string) error {
	return w.Registrar.Auth(username, password)
}

func (w *WorldConfigurator) RegistrarCheck() (*ValidationCheck, error) {
	return w.Registrar.Check()
}

func (w *WorldConfigurator) RegistrarConfirm(vc *ValidationCheck) error {
	return w.Registrar.Confirm(vc)
}

func (w *WorldConfigurator) RegistrarMustAuth() bool {
	return w.Registrar.MustAuth()
}

func (w *WorldConfigurator) RegistrarEnlist() error {
	// Register fingerprint and get a realm ID
	if err := w.Registrar.Enlist(); err != nil {
		return err
	}

	// rename world folder to World_(realmID)
	worldTxt := filepath.Join(w.Config.Dir, "World.txt")
	ioutil.WriteFile(worldTxt, MakeDefaultWorldConfig(w.Config), 0700)

	return nil
}
