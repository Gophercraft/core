package wizard

import (
	"fmt"

	"github.com/Gophercraft/core/app/input"
	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var connect_cmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a Gophercraft Home server",
	Long: `Connect to remote server, .
This tool 
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: connect_func,
}

func display_connection_challenge(challenge *wizard.ConnectionChallenge) {
	printok("Gophercraft Home fingerprint is", challenge.Fingerprint.String())
	if challenge.Trusted {
		fmt.Println("(trusted)")
	}
}

func connect_func(cmd *cobra.Command, args []string) (err error) {
	wiz := wizard.New()

	if err = wiz.Init(); err != nil {
		printwarn(err)
		return nil
	}

	var challenge *wizard.ConnectionChallenge
	if len(args) == 0 {
		challenge, err = wiz.ConnectDefault()
	} else {
		challenge, err = wiz.Connect(args[0])
	}
	if err != nil {
		printwarn(err)
		return nil
	}

	//
	if !challenge.Trusted {
		display_connection_challenge(challenge)

		var confirmed bool
		confirmed, err = input.ConfirmImportant(
			fmt.Sprintf(
				"The authenticity of this server '%s (%s)' cannot be established.\nAre you sure you want to continue connecting?",
				challenge.Host,
				challenge.Address))
		if err != nil || !confirmed {
			printwarn("host key verification failed")
			return nil
		}

		if err = wiz.ConfirmConnection(challenge); err != nil {
			return
		}
		printfok("Added host '%s' to list of known hosts\n", challenge.Host)
		return
	} else {
		printfok("Connected to trusted host '%s' (%s)", challenge.Host, challenge.Address)
		return wiz.ConfirmConnection(challenge)
	}
}
