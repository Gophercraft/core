package login

import (
	"errors"
	"regexp"
)

var alnum = regexp.MustCompile("[[:alpha:]]")

var (
	ErrUsernameCannotBeEmpty = errors.New("login: username cannot be empty")
	ErrUsernameTooLong       = errors.New("login: username too long")
	ErrUsernameTooShort      = errors.New("login: username too short")
	ErrUsernameInvalidChars  = errors.New("login: invalid characters")

	ErrPasswordCannotBeEmpty = errors.New("login: password cannot be empty")
	ErrPasswordTooLong       = errors.New("login: password too long")
	ErrPasswordTooShort      = errors.New("login: password too short")
)

func UsernameValidate(input string) error {
	if input == "" {
		return ErrUsernameCannotBeEmpty
	}

	if len(input) > 16 {
		return ErrUsernameTooLong
	}

	if len(input) < 3 {
		return ErrUsernameTooShort
	}

	if !alnum.MatchString(input) {
		return ErrUsernameInvalidChars
	}

	return nil
}

func PasswordValidate(in string) error {
	if in == "" {
		return ErrPasswordCannotBeEmpty
	}

	input := []rune(in)

	if len(input) > 128 {
		return ErrPasswordTooLong
	}

	if len(input) < 6 {
		return ErrPasswordTooShort
	}

	return nil
}
