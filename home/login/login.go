package login

import (
	"strings"

	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
	"golang.org/x/crypto/bcrypt"
)

func Credentials(username, password string) []byte {
	I := strings.ToUpper(username)
	P := strings.ToUpper(password)
	return []byte(I + ":" + P)
}

// Returns true if matches hash
func Verify(username, password string, againstHash []byte) bool {
	webLogin := Credentials(username, password)
	ok := bcrypt.CompareHashAndPassword(againstHash, webLogin) == nil
	return ok
}

func SetAccount(acc *models.Account, username, password string, tier rpcnet.Tier) error {
	if err := PasswordValidate(password); err != nil {
		return err
	}

	if err := UsernameValidate(username); err != nil {
		return err
	}

	if password == "" {
		return ErrPasswordCannotBeEmpty
	}

	acc.Username = username
	acc.Tier = tier
	acc.IdentityHash = srp.HashCredentials(
		username, password,
	)
	var err error
	acc.WebLoginHash, err = bcrypt.GenerateFromPassword(
		Credentials(username, password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	return nil
}
