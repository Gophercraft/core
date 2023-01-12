package models

import (
	"time"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/vsn"
)

var RealmOnlineDelayMeansOffline = 15 * time.Second

// Realm tracks a public realm listing
type Realm struct {
	ID              uint64 `xorm:"'id' pk" json:"id"`
	Owner           uint64
	Name            string    `xorm:"'name'" json:"name"`
	ClientVersion   vsn.Build `json:"version"`
	Locked          bool
	Type            config.RealmType `json:"type"`
	Address         string           `json:"address"`
	RedirectAddress string           `json:"redirectAddress"`
	Description     string           `json:"description"`
	ActivePlayers   uint32           `json:"activePlayers"`
	Timezone        i18n.Timezone
	LastUpdated     time.Time   `json:"lastUpdated"`
	TierNeeded      rpcnet.Tier // TODO: have special hidden servers that only show up if you are of a tier
}

func (r Realm) Offline() bool {
	return time.Since(r.LastUpdated) >= RealmOnlineDelayMeansOffline
}

// EnlistedRealm tracks a realm's registration state.
// Its only purpose is to act as a marker in the ID space of the home server, so that the real metadata can be updated at will
type EnlistedRealm struct {
	ID          uint64 `xorm:"'id' pk autoincr"`
	Note        string `xorm:"'note'"`
	Owner       uint64 `xorm:"'owner'"`
	Fingerprint string `xorm:"'fingerprint'"`
}
