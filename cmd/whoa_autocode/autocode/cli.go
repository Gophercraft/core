package autocode

import (
	"log"

	"github.com/Gophercraft/core/vsn"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "core",
	Short: "Launch Gophercraft World",
	Long:  `Launches the Gophercraft World server.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		build, err := cmd.Flags().GetUint32("build")
		if err != nil {
			log.Fatal(err)
			return
		}
		loc, err := cmd.Flags().GetString("location")
		if err != nil {
			log.Fatal(err)
			return
		}
		runMain(
			vsn.Build(build),
			loc,
		)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.core.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("location", "l", "", "The Whoa code directory")
	rootCmd.Flags().Uint32P("build", "b", 12340, "The dbc format build")
}

func Run() {
	rootCmd.Execute()
}
