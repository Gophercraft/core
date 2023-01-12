package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

const (
	splashQuit = iota
	splashSetupHomeServer
	splashSetupWorldServer
)

var splashMsg = []string{
	"Quit",
	"Create or modify a Home server",
	"Create or modify a World server",
}

func SplashScreen(w *Wizard, prev WizFunc) WizFunc {
	w.Ok("Hello, I'm the", color.FgHiBlue, "Gophercraft Wizard!", color.Reset, "I will be your guide through the magical land of", color.FgGreen, "Gophercraft.", color.Reset)

	qs := &survey.Select{
		Message: "What would you like to do",
		Options: splashMsg,
	}

	var opt int
	err := survey.AskOne(qs, &opt)
	if err != nil {
		return w.Fatal(err)
	}

	switch opt {
	case splashQuit:
		return nil
	case splashSetupHomeServer:
		return SetupHomeserver(true)
	case splashSetupWorldServer:
		return SetupWorldServer
	}

	return nil
}
