package account

import (
	"context"
	"fmt"
	"strings"

	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/account"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
	"github.com/davecgh/go-spew/spew"
)

func (provider *account_provider) CreateAccount(ctx context.Context, account_create_request *account.AccountCreateRequest) (account_create_response *account.AccountCreateResponse, err error) {
	var (
		caller_account *models.Account
		status         auth.WebTokenStatus
	)

	log.Println("Attempting login", spew.Sdump(account_create_request))

	caller_account, status, err = provider.get_account(ctx, account_create_request.Credential)
	if err != nil {
		return
	}
	if status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/core/account: not authenticated (%s)", status)
		return
	}

	err = login.RegisterAccount(provider.home_db, true, caller_account.Tier, account_create_request.Email, account_create_request.Name, account_create_request.Password, account_create_request.Tier)

	account_create_response = new(account.AccountCreateResponse)
	if err == nil {
		var (
			found       bool
			new_account models.Account
		)
		found, _ = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(account_create_request.Name))).Get(&new_account)
		if found {
			account_create_response.Id = new_account.ID
			account_create_response.Name = new_account.Username
		}
	}

	return
}
