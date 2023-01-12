package wizard

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Gophercraft/core/home/config"
)

const DefaultWorldTemplate = `{
	// build ID
	Version {{.Version}}

	// the internal IP address to listen on
	Listen 0.0.0.0:8085
	
	// external address (should be accessible from the client's computer)
	PublicAddress 127.0.0.1:8085

	// Redirect 0.0.0.0:9090 // Uncomment these two lines if you are trying to run version 3368
	// PublicRedirect 0.0.0.0:9090

	// The display name of your server. You can change this as you please.
	RealmName "{{.RealmName}}"

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
	RealmID {{.RealmID}}

	// The timezone changes which tab the realm appears under.
	Timezone "Development"

	// database driver
	DBDriver "{{.DBDriver}}"

	// database URL
	DBURL "{{.DBURL}}"

	// Address of RPC server
	HomeServer {{.HomeServer}}

	// RPC server fingerprint
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
	Version   string
	RealmName string
	RealmID   string

	DBDriver string
	DBURL    string

	HomeServer            string
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

func MakeDefaultWorldConfig(conf *config.World) []byte {
	return FormatWorldConfig(WorldConfigFormatted{
		Version:               fmt.Sprintf("%d", conf.Version),
		RealmName:             conf.RealmName,
		RealmID:               fmt.Sprintf("%d", conf.RealmID),
		DBDriver:              conf.DBDriver,
		DBURL:                 conf.DBURL,
		HomeServer:            conf.HomeServer,
		HomeServerFingerprint: conf.HomeServerFingerprint,
	})
}
