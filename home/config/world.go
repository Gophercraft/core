package config

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/vsn"
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
	Version               vsn.Build
	Listen                string
	Redirect              string
	PublicRedirect        string
	RealmID               uint64
	Timezone              i18n.Timezone
	RealmType             RealmType
	RealmName             string
	RealmDescription      string
	FederalSovereignty    bool
	DBDriver              string
	DBURL                 string
	PublicAddress         string
	WardenEnabled         bool
	ShowSQL               bool
	HomeServer            string
	HomeServerFingerprint string
	CPUProfile            string
	WorldVars             WorldVars
}

type World struct {
	Dir string
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

	wc.Dir = at

	c, err := ioutil.ReadFile(filepath.Join(wc.Dir, "World.txt"))
	if err != nil {
		return nil, err
	}

	err = text.Unmarshal(c, &wc.WorldFile)
	if err != nil {
		return nil, fmt.Errorf("%s: unmarshaling text in world config failed", err)
	}

	wc.Certificate, err = tls.LoadX509KeyPair(
		filepath.Join(wc.Dir, "cert.pem"),
		filepath.Join(wc.Dir, "key.pem"),
	)

	if err != nil {
		return nil, err
	}

	return wc, nil
}

const DefaultWorld = `{
	// build ID
	Version %d

	// the internal IP address to listen on
	Listen 0.0.0.0:8085
	
	// external address (should be accessible from the client's computer)
	PublicAddress 127.0.0.1:8085

	// Redirect 0.0.0.0:9090 // Uncomment these two lines if you are trying to run version 3368
	// PublicRedirect 0.0.0.0:9090

	// The display name of your server. You can change this as you please.
	RealmName "%s"

	// The type of server you want to create. A server of the type "RP" will allow for an all-GM roleplaying experience.
	// Note: changing this value will also change the default WorldVars.
	// You can always override these with custom values to fine-tune your server to the desired behavior!
	RealmType RP	

	// Description of your server. This will appear in the Gophercraft website.
	RealmDescription "Put the description for your server here!"

	// If true, auth server admins are not admins on your world.
	// It's recommended to keep this value false for most servers. 
	FederalSovereignty false

	// Editing this can lead to duplicate entries in the realm list.
	RealmID %d

	// The timezone changes which tab the realm appears under.
	Timezone "Development"

	// database driver
	DBDriver %s

	// database URL
	DBURL "%s"

	// Address of RPC server (replace 127.0.0.1 with host_external in gcraft_auth/Config.txt)
	AuthServer %s

	// RPC server fingerprint
	AuthServerFingerprint %s

	// Game files
	ContentSource "%s"

	// Uncomment to perform CPU usage profiling.
	// CPUProfile "cpu.prof"

	WorldVars
	{
		// 
	}
}
`

// func GenerateDefaultWorld(version uint32, name string, id uint64, sqlDriver, sqlDB string, at string, authServer, serverFingerprint string) error {
// 	path := etc.ParseSystemPath(at)

// 	if path.IsExtant() {
// 		return ErrAlreadyExists
// 	}

// 	if err := path.MakeDir(); err != nil {
// 		return err
// 	}

// 	wfile := fmt.Sprintf(DefaultWorld, version, name, id, sqlDriver, sqlDB, authServer, serverFingerprint)

// 	path.Concat("World.txt").WriteAll([]byte(wfile))
// 	path.Concat("Vars.txt").WriteAll([]byte(DefaultVars))

// 	path.Concat("Datapacks").MakeDir()

// 	if err := GenerateTLSKeyPair(path.Render()); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (a *Home) GenerateDefaultWorld(version uint32, name string, id uint64, sqlDriver, sqlDB string, at string) error {
// 	asf, err := sys.GetCertFileFingerprint(a.Path.Concat("cert.pem").Render())
// 	if err != nil {
// 		return err
// 	}

// 	if err = GenerateDefaultWorld(version, name, id, sqlDriver, sqlDB, at, "127.0.0.1:3724", asf); err != nil {
// 		return nil
// 	}

// 	wsf, err := sys.GetCertFileFingerprint(etc.ParseSystemPath(at).Concat("cert.pem").Render())
// 	if err != nil {
// 		return err
// 	}

// 	realms := RealmsFile{}
// 	realmsFile := a.Path.Concat("Realms.txt")

// 	if realmsFile.IsExtant() {
// 		rdata, err := realmsFile.ReadAll()
// 		if err != nil {
// 			return err
// 		}

// 		err = text.Unmarshal(rdata, &realms)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		realms.Realms = make(map[uint64]Realm)
// 	}

// 	realms.Realms[id] = Realm{
// 		FP: wsf,
// 	}

// 	rdata, err := text.Marshal(realms)
// 	if err != nil {
// 		return err
// 	}

// 	return realmsFile.WriteAll(rdata)
// }

func (c *World) String() string {
	return "[" + c.RealmName + "] @" + c.PublicAddress
}

func (c *World) GenerateKeyPair() error {
	return GenerateTLSKeyPair(c.Dir)
}

func (c *World) Fingerprint() string {
	f, err := rpcnet.GetCertFileFingerprint(filepath.Join(c.Dir, "cert.pem"))
	if err != nil {
		panic(err)
	}
	return f
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
