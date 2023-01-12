package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/core/wizard"
	"github.com/fatih/color"
)

const (
	goBackWorldServer       = "Return to World server setup"
	overwriteWorldServer    = "Overwrite existing World config (Cannot be undone!)"
	chooseAnotherConfigName = "Choose another config name"
)

func CreateWorldServer(showHelp bool) WizFunc {
	return func(w *Wizard, prev WizFunc) WizFunc {
		if err := w.ConfirmConfigDir(); err != nil {
			return w.Fatal(err)
		}

		w.Ok("So you want to create a", futureItemColor, "World server,", color.Reset, "eh?")
		w.Ok("I can help you with that! Let's start by picking a", itemRequiredColor, "Config Name.")
		w.Ok("You'll use this to launch the World server later on.")

		w.WorldConfigurator = w.Configurator.NewWorldConfigurator()

		var configName string

		// User is expected to enter a config name for their new world server
		// If it exists, they can try again under a different name
		for {
			askConfigName := &survey.Input{
				Message: "Enter a name for this World config.",
			}
			if err := survey.AskOne(
				askConfigName,
				&configName,
				survey.WithValidator(w.ValidateConfigName)); err != nil {
				return w.Fatal(err)
			}

			w.WorldConfigurator.SetConfigName(configName)

			if w.WorldConfigurator.ConfigExists() {
				w.Ok("A config named", configName, "already exists. What should I do?")

				alreadyExistsChoice := &survey.Select{
					Message: "What should I do?",
					Options: []string{
						goBackWorldServer,
						chooseAnotherConfigName,
						overwriteWorldServer,
					},
				}

				var choice string

				if err := survey.AskOne(
					alreadyExistsChoice,
					&choice); err != nil {
					return w.Fatal(err)
				}

				switch choice {
				case goBackWorldServer:
					return SetupWorldServer
				case overwriteWorldServer:
					w.WorldConfigurator.RemoveConfigDir()
					break
				case chooseAnotherConfigName:
					continue
				}
			} else {
				break
			}
		}

		w.Ok("You might want to set a", itemRequiredColor, "Realm Name", color.Reset, "that is different from your", itemCreationColor, "Config Name.", color.Reset)

		realmNamePrompt := &survey.Input{
			Message: "Enter a realm name?",
			Default: configName,
		}

		if err := survey.AskOne(realmNamePrompt, &w.WorldConfigurator.Config.RealmName); err != nil {
			return w.Fatal(err)
		}

		if err := w.WorldConfigurator.GenerateConfigDir(); err != nil {
			return w.Fatal(err)
		}

		for {
			var db DB
			if err := w.AskDB(&db); err != nil {
				return w.Fatal(err)
			}

			w.WorldConfigurator.Config.DBDriver = db.Driver
			w.WorldConfigurator.Config.DBURL = db.URL

			if err := w.WorldConfigurator.CreateDB(); err != nil {
				w.Warn(err)
				continue
			}
			break
		}

		w.Ok("Do you want to use a", itemRequiredColor, "Content Volume?", color.Reset)
		w.Ok("A", itemRequiredColor, "Content Volume", color.Reset, "allows you to easily generate the", futureItemColor, "Base Datapacks", color.Reset, "that Gophercraft depends on.")
		w.Ok("However, you can answer no if you are able to supply your own", itemRequiredColor, "Base Datapacks.", color.Reset)

		useCVPrompt := &survey.Confirm{
			Message: "Use a Content Volume to configure this server? (recommended)",
			Default: true,
		}

		var useCV bool

		if err := survey.AskOne(useCVPrompt, &useCV); err != nil {
			return w.Fatal(err)
		}

		if useCV {
			if wizErr := GetContentVolumeWorldServer(w); wizErr != nil {
				return wizErr
			}
		} else {
			w.Ok("I cannot scan a volume, so tell me what client build this server is for.")

			clientBuildPrompt := &survey.Input{
				Message: "Enter a client build number",
			}

			var cliBuild uint32

			if err := survey.AskOne(clientBuildPrompt, &cliBuild); err != nil {
				return w.Fatal(err)
			}

			w.WorldConfigurator.Config.Version = vsn.Build(cliBuild)
		}

		w.Ok("Now, it's time to register this World server with the Home server.")

		selectRegistrarPrompt := &survey.Select{
			Message: "Choose how you wish to enlist this World.",
			Options: []string{
				registrarLocalConfig,
				registrarRemoteServer,
			},
			Default: registrarLocalConfig,
		}
		var registrarMode string
		if err := survey.AskOne(selectRegistrarPrompt, &registrarMode); err != nil {
			return w.Fatal(err)
		}

		switch registrarMode {
		case registrarLocalConfig:
			w.WorldConfigurator.SetRegistrar(&wizard.LocalHomeConfig{})
		case registrarRemoteServer:
			w.WorldConfigurator.SetRegistrar(&wizard.RemoteHomeserverRegistrar{})
		}

		for {
			var registrarLocation string

			// The menu is different based on what registrar mode
			switch registrarMode {
			// Local config choosers get a list of home configs to choose from
			case registrarLocalConfig:
				homes, err := w.Configurator.ListHomeConfigs()
				if err != nil {
					return w.Fatal(err)
				}

				if len(homes) == 1 {
					registrarLocation = homes[0]
				} else {
					localHomeConfigName := &survey.Select{
						Message: "Select the Home config to use",
						Options: homes,
					}

					if err := survey.AskOne(localHomeConfigName, &registrarLocation); err != nil {
						return w.Fatal(err)
					}
				}
			// Remote server choosers must enter a host:port combination to connect.
			case registrarRemoteServer:
				homeServerAddressPrompt := &survey.Input{
					Message: "Enter the address of a Home server",
					Default: "localhost:3274",
				}

				if err := survey.AskOne(homeServerAddressPrompt, &registrarLocation); err != nil {
					return w.Fatal(err)
				}
			}

			if err := w.WorldConfigurator.RegistrarBegin(registrarLocation); err != nil {
				return w.Fatal(err)
			}

			check, err := w.WorldConfigurator.RegistrarCheck()
			if err != nil {
				w.Warn("Error using \""+registrarLocation+"\"", err.Error())

				msg := []string{
					enterAnotherRegistrar,
					quitToWorldServerSetup,
				}

				whatToDo := &survey.Select{
					Message: "What now?",
					Options: msg,
					Default: enterAnotherRegistrar,
				}

				var doWhat string

				if err = survey.AskOne(whatToDo, &doWhat); err != nil {
					return w.Fatal(err)
				}

				switch doWhat {
				case enterAnotherRegistrar:
					continue
				case quitToWorldServerSetup:
					return SetupWorldServer
				}
			}

			if check == nil {
				break
			}

			validateCheckPrompt := &survey.Confirm{
				Message: fmt.Sprintf("%s: %s", check.Question, check.Object),
				Default: true,
			}

			var confirmCheck bool

			if err := survey.AskOne(validateCheckPrompt, &confirmCheck); err != nil {
				return w.Fatal(err)
			}

			if !confirmCheck {
				continue
			}

			if err := w.WorldConfigurator.RegistrarConfirm(check); err != nil {
				return w.Fatal(err)
			}
		}

		// read credentials until authenticated
		if w.WorldConfigurator.RegistrarMustAuth() {
			for {
				var creds Credentials

				if err := w.ReadLogin("RealmEnlist", &creds); err != nil {
					return w.Fatal(err)
				}

				if err := w.WorldConfigurator.RegistrarAuth(creds.Username, creds.Password); err != nil {
					w.Warn(err)
					continue
				}
			}
		}

		if err := w.WorldConfigurator.RegistrarEnlist(); err != nil {
			return w.Fatal(err)
		}

		w.Ok("Hurray! Your World has been enlisted on the home server. Your new ID is", itemCreationColor, w.WorldConfigurator.Config.RealmID, color.Reset)

		if useCV {
			if wizErr := RunExtractorWorldServer(w); wizErr != nil {
				return wizErr
			}
		}

		return SetupWorldServer
	}
}
