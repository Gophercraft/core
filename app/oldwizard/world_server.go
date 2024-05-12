package wizard

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/text"
)

var (
	goMainMenu                = "Return to main menu"
	createNewWorldserver      = "Create a new World server"
	modifyExistingWorldserver = "Modify an existing World server"

	enterAnotherContentVolume = "Try to enter another content volume"
	quitToWorldServerSetup    = "Quit to World server setup menu"

	enterAnotherRegistrar = "Try to enter another registrar location"

	registrarLocalConfig  = "Use local Home config to enlist this World"
	registrarRemoteServer = "Use a remote Home server to enlist this World"
)

func (w *Wizard) ValidateContentVolumePath(str any) error {
	return nil
}

func GetContentVolumeWorldServer(w *Wizard) WizFunc {
	for {
		var cvPath string

		cvPathPrompt := &survey.Input{
			Message: "Enter the path to your Content Volume.",
		}

		if err := survey.AskOne(cvPathPrompt, &cvPath, survey.WithValidator(w.ValidateContentVolumePath)); err != nil {
			return w.Fatal(err)
		}

		var err error

		err = w.WorldConfigurator.SetVolume(cvPath)
		if err != nil {
			w.Warn("Oh no! There was a problem detecting that volume: ", err)

			msg := []string{
				enterAnotherContentVolume,
				quitToWorldServerSetup,
			}

			whatToDo := &survey.Select{
				Message: "What now?",
				Options: msg,
			}

			var doWhat string

			if err = survey.AskOne(whatToDo, &doWhat); err != nil {
				return w.Fatal(err)
			}

			switch doWhat {
			case enterAnotherContentVolume:
				continue
			case quitToWorldServerSetup:
				return SetupWorldServer
			}
		}

		return nil
	}
}

func RunExtractorWorldServer(w *Wizard) WizFunc {
	ex, err := w.WorldConfigurator.NewExtractor()
	if err != nil {
		return w.Fatal(err)
	}

	data, err := text.Marshal(w.WorldConfigurator.GamePath)
	if err != nil {
		return w.Fatal(err)
	}

	ioutil.WriteFile(filepath.Join(w.WorldConfigurator.Config.Directory, "Content.txt"), data, 0700)

	w.Ok("Do you want to extract client DB datapacks? If you don't know, just press Enter.")

	var exClientDBs, exMaps bool

	exClientDBsPrompt := &survey.Confirm{
		Message: "Extract client DB?",
		Default: true,
	}

	if err := survey.AskOne(exClientDBsPrompt, &exClientDBs); err != nil {
		return w.Fatal(err)
	}

	if exClientDBs {
		if err := ex.ExtractDatabases(); err != nil {
			return w.Fatal(err)
		}
	}

	w.Ok("Do you want to extract map datapacks?", "If you don't know, just press Enter.")

	exMapsPrompt := &survey.Confirm{
		Message: "Extract map data files?",
		Default: true,
	}

	if err := survey.AskOne(exMapsPrompt, &exMaps); err != nil {
		return w.Fatal(err)
	}

	if exMaps {
		if err := ex.ExtractMaps(); err != nil {
			return w.Fatal(err)
		}
	}

	return nil
}

var testConfigName = regexp.MustCompile("^([a-zA-Z0-9_]{1,32})$")

func (w *Wizard) ValidateConfigName(configNameA any) error {
	configName := configNameA.(string)

	ok := testConfigName.MatchString(configName)
	if !ok {
		return fmt.Errorf("invalid config name %s", configName)
	}

	return nil
}
