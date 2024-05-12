package auth

import (
	"context"

	"github.com/Gophercraft/core/home/login"
	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
)

func (provider *auth_provider) get_account(ctx context.Context, credential string) (account *home_models.Account, status auth.WebTokenStatus, err error) {
	var address string
	address, err = protocol.GetPeerAddress(ctx)
	if err != nil {
		return
	}

	return login.CheckCredential(provider.home_db, home_models.Login_Core, credential, address, "")
}
