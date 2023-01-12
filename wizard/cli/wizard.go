package cli

import (
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/core/wizard"
)

type WizFunc func(w *Wizard, prev WizFunc) WizFunc

type Wizard struct {
	Menu                WizFunc
	AskedGophercraftDir bool
	GophercraftDir      string
	Configurator        *wizard.Configurator
	HomeConfigurator    *wizard.HomeConfigurator
	WorldConfigurator   *wizard.WorldConfigurator
	CachedCredentials   map[string]*Credentials

	logPrefix string
}

func NewWizard(loc string) *Wizard {
	w := new(Wizard)
	w.GophercraftDir = loc
	w.initIO()
	return w
}

func (w *Wizard) Run() {
	vsn.PrintBanner()

	w.Menu = SplashScreen

	var prevMenu WizFunc

	for {
		if w.Menu == nil {
			break
		}

		currentMenu := w.Menu

		w.Menu = currentMenu(w, prevMenu)

		prevMenu = currentMenu
	}
}
