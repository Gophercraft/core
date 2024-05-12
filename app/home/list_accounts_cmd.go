package home

import (
	"fmt"
	"os"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/phylactery/database"
	"github.com/spf13/cobra"
)

var list_accounts_cmd = &cobra.Command{
	Use:   "accounts",
	Short: "show a list of accounts registered in the home database",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return list_accounts()
	},
}

func list_accounts() (err error) {
	var (
		wd          string
		home_config *config.Home
		db          *database.Container
	)

	wd, err = os.Getwd()
	if err != nil {
		return
	}
	home_config, err = config.LoadHome(wd)
	if err != nil {
		return
	}

	db, err = database.Open(home_config.File.DatabasePath, database.WithEngine(home_config.File.DatabaseEngine))
	if err != nil {
		return
	}

	defer db.Close()
	db.Table("Account").Iterate(func(cursor *models.Account) bool {
		fmt.Println(cursor.Username, cursor.Tier.String())
		return true
	})

	return
}
