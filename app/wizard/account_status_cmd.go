package wizard

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Gophercraft/core/home/protocol/pb/account"
	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var account_status_cmd = &cobra.Command{
	Use:   "status [id]",
	Short: "retrieves the status of an account",
	Run:   account_status_func,
}

func account_status_func(cmd *cobra.Command, args []string) {
	if err := account_status(cmd, args); err != nil {
		printwarn(err)
	}
}

func account_status(cmd *cobra.Command, args []string) (err error) {
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

	var account_status *account.AccountStatus

	if len(args) > 0 {
		var id uint64
		id, err = strconv.ParseUint(args[0], 0, 64)
		if err != nil {
			return
		}

		account_status, err = wiz.GetAccountStatus(id)
	} else {
		account_status, err = wiz.GetLoggedInAccountStatus()
	}

	if err != nil {
		return
	}

	display_account_status(account_status)

	return
}
