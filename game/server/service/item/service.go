package item

import (
	"github.com/Gophercraft/core/game/network"
	"github.com/Gophercraft/core/game/server/database"
)

type Service struct {
	// Item service requires a database engine to store its data
	engine *database.Engine
	// Actions with item service have effects elsewhere.
	// reactor
}

func NewService(engine *database.Engine) (service *Service) {
	service.engine = engine
}

func (service *Service) Init(server *network.Server) {

}
