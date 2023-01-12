package alphalist

import (
	"errors"
	"fmt"
	"net"

	"github.com/Gophercraft/log"
	"github.com/superp00t/etc"
)

const (
	onlineColor  = "FF00FF00"
	offlineColor = "FFEE2200"
)

var ErrTooManyRealms = errors.New("alphalist: too many realms")

type Session struct {
	Server *Server
	c      net.Conn
}

func (s *Session) handle() {
	conn := s.c.(*net.TCPConn)

	homeList := s.Server.Backend.ListRealms()

	list := new(List)
	list.Realms = make([]Realm, 0, len(homeList))

	for _, listing := range homeList {
		if listing.RedirectAddress != "" {
			var aRealm Realm
			// aRealm.ID = uint32(listing.ID)
			switch {
			case listing.Offline():
				aRealm.Name = fmt.Sprintf("|c%s%s (Offline)|r", offlineColor, listing.Name)
			default:
				aRealm.Name = fmt.Sprintf("|c%s%s|r", onlineColor, listing.Name)
			}
			aRealm.RedirectAddress = listing.RedirectAddress
			aRealm.Players = listing.ActivePlayers
			// aRealm.Players = 1
			list.Realms = append(list.Realms, aRealm)
		}
	}

	msg := etc.NewBuffer()
	if err := list.Encode(msg); err != nil {
		panic(err)
	}

	if _, err := conn.Write(msg.Bytes()); err != nil {
		log.Warn(err)
	}

	s.c.(*net.TCPConn).Close()
}
