package cli

import "github.com/AlecAivazis/survey/v2"

func SetupWorldServer(w *Wizard, prev WizFunc) WizFunc {
	var msg = []string{
		goMainMenu,
		createNewWorldserver,
		modifyExistingWorldserver,
	}

	prompt := &survey.Select{
		Message: "What do you want to do?",
		Options: msg,
	}

	var opt string

	if err := survey.AskOne(prompt, &opt); err != nil {
		return w.Fatal(err)
	}

	switch opt {
	case goMainMenu:
		return SplashScreen
	case createNewWorldserver:
		return CreateWorldServer(true)
	case modifyExistingWorldserver:
		return ModifyWorldserver
	}

	return nil
}
