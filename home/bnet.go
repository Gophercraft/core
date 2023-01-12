package home

import (
	"time"

	"github.com/Gophercraft/core/bnet"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/log"
)

func (c *Server) StoreLoginTicket(user, ticket string, expiry time.Time) {
	c.DB.Insert(&models.LoginTicket{
		user, ticket, expiry,
	})
}

func (c *Server) GetTicket(ticket string) (string, time.Time) {
	var ticks []models.LoginTicket

	c.DB.Where("ticket = ?", ticket).Find(&ticks)
	if len(ticks) == 0 {
		return "", time.Time{}
	}

	if time.Since(ticks[0].Expiry) > 0 {
		c.DB.Where("ticket = ?", ticket).Delete(new(models.LoginTicket))
		return "", time.Time{}
	}

	return ticks[0].Account, ticks[0].Expiry
}

func (s *Server) bnetListen() {
	log.Println("Starting bnet server at", s.Config.BnetListen)
	lst, err := bnet.Listen(s.Config.BnetListen, s.Config.BnetRESTListen, s.Config.HostExternal)
	if err != nil {
		log.Fatal(err)
	}

	lst.Backend = s

	log.Fatal(lst.Serve())
}
