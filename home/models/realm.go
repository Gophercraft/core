package models

import (
	"time"

	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/version"
)

type RealmType uint8

const (
	RealmTypeNone   = 0
	RealmTypePvP    = 1
	RealmTypeNormal = 4
	RealmTypeRP     = 6
	RealmTypeRP_PvP = 8
)

const (
	REALM_FLAG_NONE             = 0x00
	REALM_FLAG_VERSION_MISMATCH = 0x01
	REALM_FLAG_OFFLINE          = 0x02
	REALM_FLAG_SPECIFYBUILD     = 0x04
	REALM_FLAG_UNK1             = 0x08
	REALM_FLAG_UNK2             = 0x10
	REALM_FLAG_RECOMMENDED      = 0x20
	REALM_FLAG_NEW              = 0x40
	REALM_FLAG_FULL             = 0x80
)

var (
	realm_types = []uint32{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	}
)

func (rt RealmType) String() string {
	switch rt {
	case RealmTypeNone:
		return "None"
	case RealmTypePvP:
		return "PvP"
	case RealmTypeNormal:
		return "Normal"
	case RealmTypeRP:
		return "RP"
	case RealmTypeRP_PvP:
		return "RP-PvP"
	default:
		return ""
	}
}

var RealmOnlineDelayMeansOffline = 15 * time.Second

// Realm tracks a public realm listing
type Realm struct {
	ID              uint64 `database:"1:index,exclusive"`
	Owner           uint64
	Name            string
	Build           version.Build
	Locked          bool
	Type            RealmType
	Address         string
	RedirectAddress string
	Description     string
	ActivePlayers   uint32
	Category        uint32
	LastUpdated     time.Time
	RequiredTier    auth.AccountTier // TODO: have special hidden servers that only show up if you are of a tier
}

// EnlistedRealm tracks a realm's registration state.
// Its only purpose is to act as a marker in the ID space of the home server, so that the real metadata can be updated at will
type EnlistedRealm struct {
	ID          uint64 `database:"1:auto_increment,index,exclusive"`
	Note        string
	Owner       uint64
	Fingerprint protocol.Fingerprint
}

func RealmTypeToCfgConfigID(realm_type RealmType) uint32 {
	if int(realm_type) >= len(realm_types) {
		return 1
	}
	return realm_types[int(realm_type)]
}
