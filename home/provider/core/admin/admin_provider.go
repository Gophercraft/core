package admin

import (
	"github.com/Gophercraft/core/home/protocol/pb/admin"
	"github.com/Gophercraft/phylactery/database"
)

type admin_provider struct {
	admin.UnimplementedAdminServiceServer
	home_db *database.Container
}

func New(db *database.Container) (provider *admin_provider) {
	provider = new(admin_provider)
	provider.home_db = db
	return
}
