package home

import (
	"fmt"
	"os"
	"strings"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/app/input"
	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var register_cmd = &cobra.Command{
	Use:   "register",
	Short: "Register an account with the database in the current home config directory. The server must not be running for this to work properly.",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var (
			email, username, password string
			tier                      auth.AccountTier
		)
		tier, err = input.ReadAccountTier()
		if err != nil {
			return
		}
		email, username, password, err = input.ReadRegistrationLogin()
		if err != nil {
			return
		}

		return register(tier, email, username, password)
	},
}

func register(tier auth.AccountTier, email, username, password string) (err error) {
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
	if err = db.Table("Account").Sync(new(models.Account)); err != nil {
		return
	}

	if err = db.Table("GameAccount").Sync(new(models.GameAccount)); err != nil {
		return
	}

	if err = login.RegisterAccount(db, true, auth.AccountTier_ADMIN, email, username, password, tier); err != nil {
		return
	}

	var account models.Account
	db.Table("Account").Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
	fmt.Println("Registered!")
	fmt.Println(spew.Sdump(account))

	return
}
