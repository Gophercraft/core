package wizard

import (
	"context"
	"fmt"

	"github.com/Gophercraft/core/home/protocol/pb/account"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
)

func (c *Configurator) CreateAccount(tier auth.AccountTier, email string, username, password string) (id uint64, err error) {
	account_client := account.NewAccountServiceClient(c.home_connection)

	if !c.LoggedIn() {
		err = fmt.Errorf("wizard: cannot create account without being logged in")
		return
	}

	var create_account_request account.AccountCreateRequest
	create_account_request.Credential = c.login_info.WebToken
	create_account_request.Email = email
	create_account_request.Name = username
	create_account_request.Password = password
	create_account_request.Tier = tier

	var create_account_response *account.AccountCreateResponse
	create_account_response, err = account_client.CreateAccount(context.Background(), &create_account_request)
	if err != nil {
		return
	}
	id = create_account_response.Id
	return
}
