package wizard

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var root_cmd = &cobra.Command{
	Use:   "",
	Short: "The Gophercraft command-line wizard",
	Long:  `This utility provides a suite of commands useful with configuring and modifying a Gophercraft network`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.core.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	root_cmd.AddCommand(connect_cmd)
	root_cmd.AddCommand(login_cmd)
	root_cmd.AddCommand(logout_cmd)
	root_cmd.AddCommand(forget_cmd)
	root_cmd.AddCommand(session_cmd)
	root_cmd.AddCommand(disconnect_cmd)

	account_cmd.AddCommand(account_status_cmd)
	account_cmd.AddCommand(account_create_cmd)
	root_cmd.AddCommand(account_cmd)

	home_cmd.AddCommand(home_create_cmd)
	root_cmd.AddCommand(home_cmd)
}
