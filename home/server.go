// Package home implements the behavior of a Gophercraft home server
package home

import (
	"log"
	"net/http"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/i18n"
	"github.com/gorilla/mux"

	"github.com/superp00t/etc"

	"github.com/Gophercraft/core/home/config"
	"xorm.io/xorm"
)

type Server struct {
	Config       *config.Home
	WebDirectory string
	DB           *xorm.Engine
	Router       *mux.Router
}

func RunServer(cfg *config.Home) error {
	db, err := xorm.NewEngine(cfg.DBDriver, cfg.DBURL)
	if err != nil {
		return err
	}

	server := &Server{DB: db, Config: cfg}
	server.WebDirectory = etc.Import("github.com/Gophercraft/core/home/webapp").Render()

	schemas := []any{
		new(models.SessionKey),
		new(models.Account),
		new(models.GameAccount),
		new(models.Realm),
		new(models.EnlistedRealm),
		new(models.LoginTicket),
		new(models.WebToken),
		new(models.CVar),
	}

	for _, v := range schemas {
		err = server.DB.Sync2(v)
		if err != nil {
			return err
		}
	}

	_, err = server.DB.Count(new(models.Account))
	if err != nil {
		return err
	}

	for user, pass := range cfg.Admin {
		if err := server.ResetAccount(user, pass, rpcnet.Tier_Admin); err != nil {
			return err
		}
	}

	go server.cleanup()

	go server.multiprotoListen()

	go server.bnetListen()

	server.listenAlphaList()

	log.Println("Starting HTTP server at", server.Config.HTTPInternal)
	mux := server.WebServer()
	return http.ListenAndServe(server.Config.HTTPInternal, mux)
}

func (c *Server) StoreKey(user, locale, platform string, K []byte) {
	id := c.AccountID(user)

	loc, err := i18n.LocaleFromString(locale)
	if err != nil {
		// Fallback
		loc = i18n.English
	}

	c.DB.Where("id = ?", id).Cols("locale", "platform").Update(&models.Account{
		Locale:   loc,
		Platform: platform,
	})

	c.DB.Where("id = ?", id).Delete(new(models.SessionKey))

	c.DB.Insert(&models.SessionKey{
		ID: id,
		K:  K,
	})
}

// func (r Realm) Offline() bool {
// 	return (time.Now().UnixNano() - r.LastUpdated.UnixNano()) > (time.Second * 15).Nanoseconds()
// }

func (c *Server) ListRealms() []models.Realm {
	var rlmState []models.Realm
	if err := c.DB.Find(&rlmState); err != nil {
		panic(err)
	}
	return rlmState
}

func (c *Server) GetCVar(k string) string {
	var cv []models.CVar
	c.DB.Where("key = ?", k).Where("server_id = ?", 0).Find(&cv)
	if len(cv) == 0 {
		return ""
	}
	return cv[0].Value
}

func (c *Server) SetCVar(k, v string) {
	cvar := &models.CVar{0, k, v}

	c.DB.Where("key = ?", k).AllCols().Update(cvar)
}
