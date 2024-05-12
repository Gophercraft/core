package input

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func ReadText(message string, length int) (text string, err error) {
	input := survey.Input{
		Message: message,
	}
	if err = survey.AskOne(&input, &text, survey.WithValidator(func(ans any) error {
		str := ans.(string)
		if len(str) > length {
			return fmt.Errorf("too long (%d/%d)", len(str), length)
		}
		return nil
	})); err != nil {
		return
	}

	return
}
