package wizard

import (
	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var disconnect_cmd = &cobra.Command{
	Use:   "disconnect",
	Short: "disconnect from a Gophercraft Home server",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: disconnect_func,
}

func disconnect_func(cmd *cobra.Command, args []string) (err error) {
	wiz := wizard.New()

	if err = wiz.Init(); err != nil {
		printwarn(err)
		return nil
	}

	host := wiz.RemoteHost()
	address := wiz.RemoteAddress()

	if wiz.Connected() == false {
		printok("You are not currently connected")
		return
	}

	wiz.Disconnect()

	printfok("Disconnected from %s (%s)", host, address)
	return
}
