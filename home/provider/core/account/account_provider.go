package account

import (
	"github.com/Gophercraft/core/home/protocol/pb/account"
	"github.com/Gophercraft/phylactery/database"
)

type account_provider struct {
	account.UnimplementedAccountServiceServer
	home_db *database.Container
}

func New(db *database.Container) (provider *account_provider) {
	provider = new(account_provider)
	provider.home_db = db
	return
}
