package wizard

import (
	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var forget_cmd = &cobra.Command{
	Use:   "forget host",
	Short: "remove a host from the trusted list of fingerprints",
	RunE:  forget_func,
}

func forget_func(cmd *cobra.Command, args []string) (err error) {
	wiz := wizard.New()

	host := args[0]

	if err = wiz.Init(); err != nil {
		printwarn(err)
		return nil
	}

	if wiz.Connected() {
		if wiz.RemoteHost() == host {
			printwarn("disconnect from the server before forgetting it")
			return nil
		}
	}

	if err = wiz.Forget(host); err != nil {
		printwarn(err)
		return nil
	}

	return
}
