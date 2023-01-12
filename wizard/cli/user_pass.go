package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type Credentials struct {
	Username string
	Password string
}

func (wiz *Wizard) ReadLogin(kind string, creds *Credentials) error {
	var err error
	var username, password string

	usernamePrompt := &survey.Input{
		Message: fmt.Sprintf("%s username", kind),
	}

	if err = survey.AskOne(usernamePrompt, &username); err != nil {
		return err
	}

	passwordPrompt := &survey.Password{
		Message: fmt.Sprintf("%s password", kind),
	}

	if err = survey.AskOne(passwordPrompt, &password); err != nil {
		return err
	}

	creds.Username = username
	creds.Password = password

	return nil
}
