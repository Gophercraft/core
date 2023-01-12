package cli

import (
	"github.com/Gophercraft/core/home/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "core",
	Short: "Launch into the normal command-line wizard.",
	Long:  `Launches the Gophercraft Wizard. From there you will answer a series of questions to assist you in creating Gophercraft configs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		loc, err := cmd.Flags().GetString("location")
		if err != nil {
			panic(err)
		}
		NewWizard(loc).Run()
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.core.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("location", "l", config.DefaultLocation(), "The default Gophercraft directory that the Wizard will touch")
}
