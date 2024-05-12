package grunt

import (
	"net"

	"github.com/Gophercraft/core/app/config"
)

// type ServiceProvider interface {
// 	GetAccount(account_name string) (*models.Account, []models.GameAccount, error)
// 	ListRealms() []models.Realm
// 	StoreKey(user, locale, platform string, K []byte)
// }

type Server struct {
	config           *ServerConfig
	service_provider ServiceProvider
	listener         net.Listener
}

func (s *Server) ID() config.HomeServiceID {
	return config.GruntService
}
