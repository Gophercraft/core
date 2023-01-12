package cli

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/fatih/color"
)

var goBack = "Main menu"
var modifyExistingHomeserver = "Modify an extant Home server"
var createNewHomeserver = "Create a new Home server."

func SetupHomeserver(showMsg bool) WizFunc {
	return func(wiz *Wizard, prev WizFunc) WizFunc {
		if showMsg {
			wiz.Ok("The", color.FgGreen, "Home server", color.Reset, "is the heart of any", color.FgGreen, "Gophercraft network.", color.Reset)
			wiz.Ok("It's where you'll register, log in, and manage your realm list.")
		}

		var msg = []string{
			goBack,
			createNewHomeserver,
			modifyExistingHomeserver,
		}

		prompt := &survey.Select{
			Message: "What do you want to do?",

			Options: msg,
		}

		var action string

		err := survey.AskOne(prompt, &action)
		if err != nil {
			return wiz.Fatal(err)
		}

		switch action {
		case goBack:
			return SplashScreen
		case createNewHomeserver:
			return CreateHomeserver
		case modifyExistingHomeserver:
			return ModifyHomeserver
		}

		return nil
	}
}

func CreateHomeserver(wiz *Wizard, prev WizFunc) WizFunc {
	if err := wiz.ConfirmConfigDir(); err != nil {
		return wiz.Fatal(err)
	}

	wiz.HomeConfigurator = wiz.Configurator.NewHomeConfigurator()

	wiz.Ok(
		"By default, the Gophercraft Home server will read from the config named \"Home\" in your config directory.\n",
		"This might be useful for developers. If this is a typical deployment, then just press", color.FgCyan, "Enter", color.Reset, "to continue.")

	promptCustomName := &survey.Input{
		Message: "Enter a Home server directory name",
		Default: "Home",
	}

	var homeName string

	if err := survey.AskOne(promptCustomName, &homeName); err != nil {
		return wiz.Fatal(err)
	}

	wiz.HomeConfigurator.SetDirName(homeName)

	if _, err := os.Stat(wiz.HomeConfigurator.Config.Dir); err == nil {
		confirmDeletePrompt := &survey.Confirm{
			Message: fmt.Sprintf("Home config \"%s\" already exists. Delete?", homeName),
			Default: true,
		}

		var confirm bool

		err = survey.AskOne(confirmDeletePrompt, &confirm)
		if err != nil {
			return wiz.Fatal(err)
		}

		if !confirm {
			return SetupHomeserver(false)
		}

		if err = os.RemoveAll(wiz.HomeConfigurator.Config.Dir); err != nil {
			return wiz.Fatal(err)
		}
	}

	var db DB
	if err := wiz.AskDB(&db); err != nil {
		return wiz.Fatal(err)
	}

	wiz.HomeConfigurator.Config.DBDriver = db.Driver
	wiz.HomeConfigurator.Config.DBURL = db.URL

	if err := wiz.HomeConfigurator.CreateDB(); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok("Okay, I made a database. Now, enter a username and a password for your admin account.\nDon't worry about losing the password - You can always reset it later!")

	var cred Credentials

	if err := wiz.ReadLogin("Admin", &cred); err != nil {
		return wiz.Fatal(err)
	}

	if err := wiz.HomeConfigurator.RegisterAccount(cred.Username, cred.Password, rpcnet.Tier_Admin); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok("Account", color.FgCyan, cred.Username, color.Reset, "was successfully registered.")

	wiz.Ok("Are you okay with Gophercraft allowing anyone to sign up to your home server and create a bottom-tier account?")

	promptOpenRegistration := &survey.Confirm{
		Message: "Allow anybody to register an account with the Home server web portal?",
		Default: false,
	}

	if err := survey.AskOne(promptOpenRegistration, &wiz.HomeConfigurator.Config.OpenRegistration); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok("Now creating your Home server config.")

	if err := wiz.HomeConfigurator.GenerateConfig(); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok(
		"Home configuration has been generated at",
		color.FgCyan, wiz.HomeConfigurator.Config.Dir, color.Reset,
		"Fingerprint:",
		wiz.HomeConfigurator.Config.Fingerprint())

	return SetupHomeserver(false)
}

func ModifyHomeserver(wiz *Wizard, prev WizFunc) WizFunc {
	return nil
}
