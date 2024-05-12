package home

import (
	"fmt"
	"os"

	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var create_cmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Gophercraft Home server configuration in the current directory",
	Long:  `The Gophercraft Home server creates a configuration the current working directory`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

func create() {
	var (
		err error
	)
	configurator := wizard.New()
	home_configurator := configurator.NewHomeConfigurator()
	home_configurator.Config.File.DatabaseEngine = "leveldb_core"
	home_configurator.Config.File.DatabasePath = "home_db"
	home_configurator.Config.Directory, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	if err := home_configurator.GenerateConfig(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New Home config generated at", home_configurator.Config.Directory)

}
