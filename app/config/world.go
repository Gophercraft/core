package config

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/text"
)

type WorldVars map[string]string

var (
	Presets map[RealmType]WorldVars
)

func init() {
	Presets = map[RealmType]WorldVars{
		RealmTypeRP: {
			"XP.Rate": "1.0",

			"Weather.On": "true",

			"Sync.VisibilityRange": "250.0",

			"PvP.Deathmatch": "false",
			"PvP.AtWar":      "false",

			"Char.StartLevel": "255",
			"Char.Cinematic":  "true",

			"Chat.HostAnnounceChannel": "true",
		},
	}
}

type RealmType uint8

const (
	RealmTypeNone   = 0
	RealmTypePvP    = 1
	RealmTypeNormal = 4
	RealmTypeRP     = 6
	RealmTypeRP_PvP = 8
)

func (rt RealmType) EncodeWord() (string, error) {
	switch rt {
	case RealmTypeNone:
		return "None", nil
	case RealmTypePvP:
		return "PvP", nil
	case RealmTypeNormal:
		return "Normal", nil
	case RealmTypeRP:
		return "RP", nil
	case RealmTypeRP_PvP:
		return "RP-PvP", nil
	default:
		return "", fmt.Errorf("gcore: unknown RealmType %d", rt)
	}
}

func (rt *RealmType) DecodeWord(wdata string) error {
	data := strings.ToLower(wdata)

	switch data {
	case "none", "":
		*rt = RealmTypeNone
	case "pvp":
		*rt = RealmTypePvP
	case "normal":
		*rt = RealmTypeNormal
	case "rp":
		*rt = RealmTypeRP
	case "rp-pvp":
		*rt = RealmTypeRP_PvP
	default:
		return fmt.Errorf("gcore: unrecognized realm type %s", wdata)
	}

	return nil
}

type WorldFile struct {
	// The enlisted ID of the world server
	ID uint64
	// The address of the Home server
	HomeServerAddress string
	// The trusted fingerprint of the Home server
	HomeServerFingerprint protocol.Fingerprint
	// The supported client build
	Build version.Build
	// The TCP address to listen the actual world server to
	BindGameAddress string
	// The world server's game address as it appears in the realm list
	PublicGameAddress string
	// the TCP address to listen the redirect server on
	BindRedirectAddress string
	// The world server's redirect address as it appears in the realm list
	PublicRedirect string
	// The TCP address to listen to core protocol server
	BindCoreAddress string
	// The address of the core server (visible from the home server's network)
	PublicCoreAddress string
	// The Timezone (category) of the server as it appears in the realm list
	Timezone i18n.Timezone
	// The type enum of the server
	Type RealmType
	// Long-form name e.g. "The Beluga Boys Fanclub Ltd."
	// How it appears in the realmlist (should be relatively unchanging to avoid connection issues)
	LongName string
	// Short-form no-spaced name e.g. "BelugaBoys"
	ShortName string
	// Multiline description
	Description string
	// Engine used for the cache database
	CacheDatabaseEngine string
	// path to cache database
	CacheDatabasePath string
	// Engine used for the world database
	WorldDatabaseEngine string
	// path to world database
	WorldDatabasePath string
	// Message of the day (Sent in SMSG_MOTD packet)
	SystemMOTD string
	// Message of the day sent in chat
	ChatMOTD string
	// Enable Warden loading activity
	WardenEnabled bool
	// World configuration variables that differ from the Type-configured presets
	WorldVars WorldVars
}

type World struct {
	Directory string
	WorldFile
	Certificate tls.Certificate
}

func WorldExists(at string) bool {
	fi, err := os.Stat(at)
	if err != nil {
		return false
	}

	if !fi.IsDir() {
		return false
	}

	txtName := filepath.Join(at, "World.txt")
	if _, err = os.Stat(txtName); err != nil {
		return false
	}

	return true
}

func LoadWorld(at string) (*World, error) {
	wc := new(World)

	wc.Directory = at

	c, err := os.ReadFile(filepath.Join(wc.Directory, "world.txt"))
	if err != nil {
		return nil, err
	}

	err = text.Unmarshal(c, &wc.WorldFile)
	if err != nil {
		return nil, fmt.Errorf("%s: unmarshaling text in world config failed", err)
	}

	wc.Certificate, err = tls.LoadX509KeyPair(
		filepath.Join(wc.Directory, "cert.pem"),
		filepath.Join(wc.Directory, "key.pem"),
	)

	if err != nil {
		return nil, err
	}

	return wc, nil
}

func (c *World) Fingerprint() string {
	f, err := protocol.GetCertFileFingerprint(filepath.Join(c.Directory, "cert.pem"))
	if err != nil {
		panic(err)
	}
	return f.String()
}

// List all World configs in the chosen Gophercraft root directory.
func ListWorldConfigs(dir string) ([]string, error) {
	list, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var names []string

	for _, entry := range list {
		if entry.IsDir() {
			_, err := os.Stat(filepath.Join(dir, entry.Name(), "World.txt"))
			if err == nil {
				names = append(names, entry.Name())
			}
		}
	}

	return names, nil
}
