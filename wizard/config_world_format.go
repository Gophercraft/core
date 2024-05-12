package wizard

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Gophercraft/core/app/config"
)

var _ = config.WorldFile{}

const DefaultWorldTemplate = `{
	// build ID
	Build {{.Version}}

	// the internal IP address to listen on
	BindGameAddress 0.0.0.0:8085
	
	// external address (this is the address the game client connects to)
	PublicGameAddress 127.0.0.1:8085

	BindCoreAddress 0.0.0.0:32771

	PublicCoreAddress 127.0.0.1:32771

	// Uncomment these two lines if your version needs a redirect server
	// Redirect 0.0.0.0:9090
	// RedirectEndpoint 127.0.0.1:9090

	// The display name of your server. Try not to change this
	LongName "{{.RealmName}}"
	
	ShortName "{{.RealmName}}"

	// The type of server you want to create. A server of the type "RP" will allow for an all-GM roleplaying experience.
	// Note: changing this value will also change the default WorldVars.
	// You can always override these with custom values to fine-tune your server to the desired behavior!
	Type RP	

	// Description of your server. This will appear in the Gophercraft website.
	Description "Put the description for your server here!"

	// Editing this can lead to duplicate entries in the realm list.
	ID {{.RealmID}}

	// The timezone changes which tab the realm appears under.
	Timezone "Development"

	// database driver
	WorldDatabaseEngine "leveldb_core"
	WorldDatabasePath "world_db"

	CacheDatabaseEngine "leveldb_core"
	CacheDatabasePath "cache_db"

	// Address of Home server
	HomeServer {{.HomeServerADdress}}

	// Home server fingerprint
	HomeServerFingerprint {{.HomeServerFingerprint}}

	// Uncomment to perform CPU usage profiling.
	// CPUProfile "cpu.prof"

	WorldVars
	{
		// Here go extra options.
	}
}
`

type WorldConfigFormatted struct {
	Build     string
	RealmName string
	RealmID   string

	HomeServerAddress     string
	HomeServerFingerprint string
}

func FormatWorldConfig(wcf WorldConfigFormatted) []byte {
	tmpl, err := template.New("worldconfig").Parse(DefaultWorldTemplate)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, wcf); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func MakeDefaultWorldConfig(world_config *config.World) []byte {
	return FormatWorldConfig(WorldConfigFormatted{
		Build:                 fmt.Sprintf("%d", world_config.Build),
		RealmName:             world_config.LongName,
		RealmID:               fmt.Sprintf("%d", world_config.ID),
		HomeServerAddress:     world_config.HomeServerAddress,
		HomeServerFingerprint: world_config.HomeServerFingerprint.String(),
	})
}
