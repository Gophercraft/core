package wizard

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
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
	wiz.Ok("Is this where you want to install your Home server config?")
	wiz.HomeConfigurator = wiz.Configurator.NewHomeConfigurator()
	promptCustomName := &survey.Input{
		Message: "Enter a Home server directory",
		Default: wiz.HomeConfigurator.Config.Directory,
	}
	var home_directory string
	if err := survey.AskOne(promptCustomName, &home_directory); err != nil {
		return wiz.Fatal(err)
	}
	wiz.HomeConfigurator.Config.Directory = home_directory

	if _, err := os.Stat(wiz.HomeConfigurator.Config.Directory); err == nil {
		confirmDeletePrompt := &survey.Confirm{
			Message: fmt.Sprintf("Home config \"%s\" already exists. Delete?", home_directory),
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

		wiz.Ok("removing directory...", wiz.HomeConfigurator.Config.Directory)

		if err = os.RemoveAll(wiz.HomeConfigurator.Config.Directory); err != nil {
			return wiz.Fatal(err)
		}
	}

	wiz.HomeConfigurator.Config.File.DatabaseEngine = "leveldb_core"
	wiz.HomeConfigurator.Config.File.DatabasePath = "home_db"

	if err := wiz.HomeConfigurator.CreateDB(); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok("Okay, I made a database. Now, enter a username and a password for your admin account.\nDon't worry about losing the password - You can always reset it later!")

	var cred Credentials

	if err := wiz.ReadLogin("Admin", &cred); err != nil {
		return wiz.Fatal(err)
	}

	if err := wiz.HomeConfigurator.RegisterAccount(cred.Username, cred.Password, auth.AccountTier_Admin); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok("Account", color.FgCyan, cred.Username, color.Reset, "was successfully registered.")

	wiz.Ok("Are you okay with Gophercraft allowing anyone to sign up to your home server and create a bottom-tier account?")

	promptOpenRegistration := &survey.Confirm{
		Message: "Allow anybody to register an account with the Home server web portal?",
		Default: false,
	}

	if err := survey.AskOne(promptOpenRegistration, &wiz.HomeConfigurator.Config.File.OpenRegistration); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok("Now creating your Home server config.")

	if err := wiz.HomeConfigurator.GenerateConfig(); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok(
		"Home configuration has been generated at",
		color.FgCyan, wiz.HomeConfigurator.Config.Directory, color.Reset,
		"Fingerprint:",
		wiz.HomeConfigurator.Config.Fingerprint())

	return SetupHomeserver(false)
}

const (
	registerAccount = "Register an account"
)

func ModifyHomeserver(wiz *Wizard, prev WizFunc) WizFunc {
	wiz.Ok("Is this the config you want to modify?")
	wiz.HomeConfigurator = wiz.Configurator.NewHomeConfigurator()
	promptCustomName := &survey.Input{
		Message: "Enter a Home server directory",
		Default: wiz.HomeConfigurator.Config.Directory,
	}
	var home_directory string
	if err := survey.AskOne(promptCustomName, &home_directory); err != nil {
		return wiz.Fatal(err)
	}

	var err error
	if wiz.HomeConfigurator.Config, err = config.LoadHome(home_directory); err != nil {
		return wiz.Fatal(err)
	}

	wiz.HomeConfigurator.Config.Directory = home_directory
	msg := []string{
		registerAccount,
		goMainMenu,
	}

	prompt := &survey.Select{
		Message: "How do you want to modify the Home server? ",
		Options: msg,
	}

	var action string
	err = survey.AskOne(prompt, &action)
	if err != nil {
		return wiz.Fatal(err)
	}

	switch action {
	case registerAccount:
		return RegisterAccount
	case goMainMenu:
		return nil
	}

	return nil
}

func RegisterAccount(wiz *Wizard, prev WizFunc) WizFunc {
	wiz.Ok("Okay, enter credentials for this new account!")
	var credentials Credentials
	if err := wiz.ReadLogin("", &credentials); err != nil {
		return wiz.Fatal(err)
	}

	tiers := make([]string, len(auth.AccountTier_value))
	for k, v := range auth.AccountTier_name {
		tiers[int(k)] = v
	}

	prompt := &survey.Select{
		Message: "What tier should this account occupy?",
		Options: tiers,
		Default: "Admin",
	}
	var tier_index int
	err := survey.AskOne(prompt, &tier_index)
	if err != nil {
		return wiz.Fatal(err)
	}

	if err := wiz.HomeConfigurator.RegisterAccount(credentials.Username, credentials.Password, auth.AccountTier(tier_index)); err != nil {
		return wiz.Fatal(err)
	}

	wiz.Ok("account", credentials.Username, "registered!")

	return nil
}
