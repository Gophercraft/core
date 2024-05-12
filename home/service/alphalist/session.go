package alphalist

import (
	"net"

	"github.com/Gophercraft/log"
)

type Session struct {
	service    *Service
	connection net.Conn
}

func (session *Session) handle() {
	realm_list, err := session.service.list_provider.ListRealms()
	if err != nil {
		log.Warn(err)
		return
	}

	realm_list_bytes, err := EncodeList(realm_list)
	if err != nil {
		log.Warn(err)
		return
	}

	if _, err = session.connection.Write(realm_list_bytes); err != nil {
		log.Warn(err)
		return
	}
}
