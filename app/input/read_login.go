package input

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
)

func validate_username(input any) error {
	return login.UsernameValidate(input.(string))
}

func validate_password(input any) error {
	return login.PasswordValidate(input.(string))
}

func validate_username_login(input any) error {
	if input.(string) == "" {
		return fmt.Errorf("username cannot be left empty")
	}
	return nil
}

func validate_password_login(input any) error {
	if input.(string) == "" {
		return fmt.Errorf("password cannot be left empty")
	}
	return nil
}

func ReadLogin() (username, password string, err error) {
	username_input := survey.Input{
		Message: "Enter account name",
	}
	password_input := survey.Password{
		Message: "Enter password",
	}
	if err = survey.AskOne(&username_input, &username, survey.WithValidator(validate_username_login)); err != nil {
		return

	}
	if err = survey.AskOne(&password_input, &password, survey.WithValidator(validate_password_login)); err != nil {
		return
	}

	return
}

func ReadRegistrationLogin() (email, username, password string, err error) {
	var password_confirmation string
	email_input := survey.Input{
		Message: "Enter email address (or leave blank)",
	}
	username_input := survey.Input{
		Message: "Enter account name",
	}
	password_input := survey.Password{
		Message: "Enter password",
	}
	password_input_again := survey.Password{
		Message: "Enter the same password again",
	}
	if err = survey.AskOne(&email_input, &email); err != nil {
		return

	}

	if err = survey.AskOne(&username_input, &username, survey.WithValidator(validate_username)); err != nil {
		return

	}
	if err = survey.AskOne(&password_input, &password, survey.WithValidator(validate_password)); err != nil {
		return
	}

	if err = survey.AskOne(&password_input_again, &password_confirmation); err != nil {
		return
	}

	if password != password_confirmation {
		err = fmt.Errorf("app/input: password mismatch")
		return
	}

	return
}

func ReadAccountTier() (tier auth.AccountTier, err error) {
	tiers := []string{
		"regular",
		"privileged",
		"gamemaster",
		"moderator",
		"admin",
	}

	prompt := &survey.Select{
		Message: "What is the tier of this account?",
		Help:    "The tier determines how this account is allowed to ",
		Options: tiers,
		Default: "admin",
	}
	var tier_index int
	err = survey.AskOne(prompt, &tier_index)
	if err != nil {
		return
	}

	tier = auth.AccountTier(tier_index)
	return
}
