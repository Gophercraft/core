package auth

import (
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/phylactery/database"
)

type auth_provider struct {
	auth.UnimplementedAuthServiceServer
	home_db *database.Container
}

func New(db *database.Container) (provider *auth_provider) {
	provider = new(auth_provider)
	provider.home_db = db
	return
}
