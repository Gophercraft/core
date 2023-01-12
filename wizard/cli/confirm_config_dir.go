package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/core/wizard"
	"github.com/fatih/color"
)

func (w *Wizard) ConfirmConfigDir() error {
	if w.AskedGophercraftDir {
		return nil
	}

	w.AskedGophercraftDir = true

	w.Ok("First thing's first: you'll want to know where I'm keeping your", color.FgHiGreen, "configuration files!\n", color.Reset,
		"Right now, they're located at", w.GophercraftDir+".\nYou can type in a new directory or just press", color.FgCyan, "Enter", color.Reset, "to continue.",
	)

	newDir := ""

	prompt := &survey.Input{
		Message: "...",

		Default: w.GophercraftDir,
		// Suggest: func(toComplete string) []string {
		// 	files, _ := filepath.Glob(toComplete + "*")
		// 	return files
		// },
	}

	err := survey.AskOne(prompt, &newDir)
	if err != nil {
		return err
	}

	w.Ok("Awesome. Okay, from now on I'll be saving files to", newDir, "until you decide to quit the wizard.")

	w.GophercraftDir = newDir
	w.Configurator = wizard.NewConfigurator(w.GophercraftDir)

	return nil
}
