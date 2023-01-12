package realm

import (
	"context"
	"time"

	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"

	"github.com/golang/protobuf/ptypes/empty"
)

// repeatedly posts AnnounceRealmMsg into the mothership
func (ws *Server) phoneHome() {
	gc, err := rpcnet.DialConn(
		ws.Config.HomeServer,
		ws.Config.HomeServerFingerprint,
		&ws.Config.Certificate,
	)

	if err != nil {
		panic(err)
	}

	cl := rpcnet.NewHomeServiceClient(gc)

	ws.HomeServiceClient = cl

	vi, err := cl.GetVersionData(context.Background(), &empty.Empty{})
	if err != nil {
		log.Warn(err)
	} else {
		if vi.CoreVersion != vsn.GophercraftVersion.String() {
			log.Warn("Your authentication server is using Gophercraft", vi.CoreVersion, ", and this server is using", vsn.GophercraftVersion, ". Try updating if you experience problems.")
		}
	}

	for {
		st, err := cl.AnnounceRealm(context.Background(), &rpcnet.AnnounceRealmMsg{
			RealmID:          ws.Config.RealmID,
			Type:             uint32(ws.Config.RealmType),
			RealmName:        ws.Config.RealmName,
			RealmDescription: ws.Config.RealmDescription,
			Build:            uint32(ws.Config.Version),
			Address:          ws.Config.PublicAddress,
			ActivePlayers:    uint32(len(ws.PlayerList)),
			RedirectAddress:  ws.Config.PublicRedirect,
		})
		if err != nil {
			log.Warn(err)
		}
		if st != nil && st.Status != rpcnet.Status_OK {
			log.Warn("Recieved non-ok status", st.Status)
		}
		time.Sleep(8 * time.Second)
	}
}
