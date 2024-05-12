package wizard

import (
	"errors"
	"fmt"

	"github.com/Gophercraft/core/app/input"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var account_create_cmd = &cobra.Command{
	Use:   "create",
	Short: "register a new account",
	Long:  "You must be logged in and have sufficient permissions to register an account",
	Run:   account_create_func,
}

func account_create_func(cmd *cobra.Command, args []string) {
	if err := account_create(cmd, args); err != nil {
		printwarn(err)
	}
}

func account_create(cmd *cobra.Command, args []string) (err error) {
	wiz := wizard.New()

	err = wiz.Init()
	defer wiz.Close()

	if errors.Is(err, wizard.ErrDisconnected) {
		fmt.Println(err)
		err = nil
	} else if err != nil {
		return
	}

	if !(wiz.Connected() && wiz.LoggedIn()) {
		return fmt.Errorf("You must be logged in and connected to use this command.")
	}

	printfok("Select the tier of the account you want to create")

	var (
		email, username, password string
		tier                      auth.AccountTier
		id                        uint64
	)
	tier, err = input.ReadAccountTier()
	if err != nil {
		return
	}
	email, username, password, err = input.ReadRegistrationLogin()
	if err != nil {
		return
	}

	if id, err = wiz.CreateAccount(tier, email, username, password); err != nil {
		return
	}

	printfok("Created account %s#%d", username, id)

	return
}
