package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/core/home/dbsupport"
	"github.com/fatih/color"
)

type DB struct {
	Driver string
	URL    string
}

func (w *Wizard) ValidateDBURL(value any) error {
	// str := value.(string)
	return nil
}

func (wiz *Wizard) AskDB(d *DB) error {
	wiz.Ok("Please select a database backend")

	promptBackend := &survey.Select{
		Message: "choose a database backend",
		Options: dbsupport.Supported,
	}

	err := survey.AskOne(promptBackend, &d.Driver)
	if err != nil {
		return err
	}

	wiz.Ok(
		"You picked",
		color.FgCyan, d.Driver, color.Reset,
		"as your database driver. Nice!\n",
		"Next, you must enter a database URL. Database urls for",
		color.FgCyan, d.Driver, color.Reset,
		"Look like this: ",
		color.FgGreen, dbsupport.PathFormat[d.Driver], color.Reset,
	)

	promptDburl := &survey.Input{
		Message: "Enter a database URL.",
	}

	err = survey.AskOne(
		promptDburl,
		&d.URL,
		survey.WithValidator(wiz.ValidateDBURL))

	if err != nil {
		return err
	}

	return err
}
