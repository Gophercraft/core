package wizard

import (
	"errors"
	"fmt"

	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var logout_cmd = &cobra.Command{
	Use:   "logout",
	Short: "log out from a Gophercraft Home server",
	Run:   logout_func,
}

func logout_func(cmd *cobra.Command, args []string) {
	if err := logout(cmd, args); err != nil {
		printwarn(err)
	}
	return
}

func logout(cmd *cobra.Command, args []string) (err error) {
	wiz := wizard.New()

	err = wiz.Init()

	if errors.Is(err, wizard.ErrDisconnected) {
		fmt.Println(err)
		err = nil
	} else if err != nil {
		return
	}

	if !wiz.LoggedIn() {
		return fmt.Errorf("you are not logged in")
	}

	printfok("logged out")

	err = wiz.Logout()

	wiz.Close()

	return
}
