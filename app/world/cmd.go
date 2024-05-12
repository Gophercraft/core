package home

import (
	"os"

	"github.com/spf13/cobra"
)

// root_cmd represents the base command when called without any subcommands
var root_cmd = &cobra.Command{
	Use:   "gophercraft_world",
	Short: "Run a Gophercraft World server",
	Long:  `The Gophercraft World server reads configuration from the working directory it is started in`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		run_main()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := root_cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.core.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
