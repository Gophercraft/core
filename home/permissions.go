package home

import (
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
)

func (s *Server) CanEnlistRealm(acc *models.Account) bool {
	return acc.Tier >= rpcnet.Tier_Admin
}
