package input

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func ConfirmImportant(question string) (confirmation bool, err error) {
	var yes_no string
	if err = survey.AskOne(&survey.Input{
		Message: question,
		Default: "[yes/no]",
	}, &yes_no); err != nil {
		return
	}

	if yes_no != "yes" && yes_no != "no" {
		err = fmt.Errorf("must either be 'yes' or 'no'")
		return
	}

	confirmation = yes_no == "yes"
	return
}
