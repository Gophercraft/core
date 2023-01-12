package cli

import (
	"io/ioutil"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/text"
)

var (
	modifyWorldDatapack = "Regenerate datapacks?"
)

func RegenerateDatapackInternal(w *Wizard, exDB, exMaps bool) error {
	//  cv, err := content.Open(gamePath)
	//  if err != nil {
	// 	return err
	// }

	ex, err := w.WorldConfigurator.NewExtractor()
	if err != nil {
		return err
	}

	if exDB {
		if err := ex.ExtractDatabases(); err != nil {
			return err
		}
	}

	if exMaps {
		if err := ex.ExtractMaps(); err != nil {
			return err
		}
	}

	return nil
}

func RegenerateDatapack(w *Wizard, prev WizFunc) WizFunc {
	cvPath := filepath.Join(w.WorldConfigurator.Config.Dir, "Content.txt")

	var cv string

	cvData, err := ioutil.ReadFile(cvPath)
	if err == nil {
		if err := text.Unmarshal(cvData, &cv); err != nil {
			return w.Fatal(err)
		}
	}

	getGamePath := &survey.Input{
		Message: "Enter a new Content Volume or press Enter if this is correct.",
		Default: cv,
	}

	if err := survey.AskOne(getGamePath, &cv); err != nil {
		return w.Fatal(err)
	}

	w.WorldConfigurator.GamePath = cv

	w.Ok("Do you want to extract client DB datapacks? If you don't know, just press Enter.")

	var exClientDBs, exMaps bool

	exClientDBsPrompt := &survey.Confirm{
		Message: "Extract client DB?",
		Default: true,
	}

	if err := survey.AskOne(exClientDBsPrompt, &exClientDBs); err != nil {
		return w.Fatal(err)
	}

	w.Ok("Do you want to extract map datapacks?", "If you don't know, just press Enter.")

	exMapsPrompt := &survey.Confirm{
		Message: "Extract map data files?",
		Default: true,
	}

	if err := survey.AskOne(exMapsPrompt, &exMaps); err != nil {
		return w.Fatal(err)
	}

	if err := RegenerateDatapackInternal(w, exClientDBs, exMaps); err != nil {
		return w.Fatal(err)
	}

	return nil
}

func ModifyWorldserver(w *Wizard, prev WizFunc) WizFunc {
	if err := w.ConfirmConfigDir(); err != nil {
		return w.Fatal(err)
	}

	worlds, err := w.Configurator.ListWorldConfigs()
	if err != nil {
		return w.Fatal(err)
	}

	if len(worlds) == 0 {
		return SetupWorldServer
	}

	quit := len(worlds)

	worlds = append(worlds, quitToWorldServerSetup)

	selectWorldName := &survey.Select{
		Message: "What world do you want to modify?",
		Options: worlds,
	}

	var windex int

	if err := survey.AskOne(selectWorldName, &windex); err != nil {
		return w.Fatal(err)
	}

	if windex == quit {
		return SetupWorldServer
	}

	world := worlds[windex]

	w.WorldConfigurator = w.Configurator.NewWorldConfigurator()
	if err := w.WorldConfigurator.LoadConfig(world); err != nil {
		return w.Fatal(err)
	}

	doWhat := &survey.Select{
		Message: "What now?",
		Options: []string{
			modifyWorldDatapack,
			quitToWorldServerSetup,
		},
	}

	var what string

	if err := survey.AskOne(doWhat, &what); err != nil {
		return w.Fatal(err)
	}
	switch what {
	case modifyWorldDatapack:
		if err := RegenerateDatapack(w, SetupWorldServer); err != nil {
			return err

		}
	case quitToWorldServerSetup:
		return SetupWorldServer
	}

	return SetupWorldServer
}
