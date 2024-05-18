package wizard

import (
	"fmt"
	"os"

	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var home_create_cmd = &cobra.Command{
	Use:   "create",
	Short: "create a Gophercraft Home server configuration in the current directory",
	Run:   home_create_func,
}

func home_create_func(cmd *cobra.Command, args []string) {
	if err := home_create(cmd, args); err != nil {
		printwarn(err)
	}
}

func home_create(cmd *cobra.Command, args []string) (err error) {
	configurator := wizard.New()
	home_configurator := configurator.NewHomeConfigurator()
	home_configurator.Config.File.DatabaseEngine = "leveldb_core"
	home_configurator.Config.File.DatabasePath = "home_db"
	home_configurator.Config.Directory, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	if err = home_configurator.GenerateConfig(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New Home config generated at", home_configurator.Config.Directory)
	return
}
