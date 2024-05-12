package rest

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *Provider) update_account_data(username, platform_id string, session_key []byte) (err error) {
	var (
		found   bool
		account models.Account
	)
	account_table := provider.home_db.Table("Account")

	found, err = account_table.Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/bnet/rest: account '%s' does not exist", username)
		return
	}

	account.Platform = platform_id
	account.SessionKey = session_key

	_, err = account_table.Where(query.Eq("ID", account.ID)).
		Columns(
			"Platform",
			"SessionKey",
		).Update(&account)

	return
}
