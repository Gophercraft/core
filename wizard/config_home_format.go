package wizard

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Gophercraft/core/home/config"
)

const DefaultHomeTemplate = `{
	// the TCP/IP address to listen the Gophercraft website on.
	// You can reverse proxy this however you like.
	HTTPInternal 0.0.0.0:8086

	// The public hostname of your Gophercraft API server.
	// if left uncommented, it will be set to localhost
	// this is needed to tell the client where the REST logon service is located.
	// 
	// HostExternal gcraft.example.com

	// The TCP/IP addresses to listen Auth/Realmlist servers on.
	// Keep these unchanged, unless you really know what you're doing.
	AuthListen 0.0.0.0:3724 // The address which serves the legacy realmlist server, as well as the GRPC server.
	BnetListen 0.0.0.0:1119
	BnetRESTListen 0.0.0.0:1120

	// Database options
	// the go-xorm SQL driver to use.
	DBDriver {{.DBDriver}}

	// the go-xorm SQL URL to use.
	DBURL {{.DBURL}}

	OpenRegistration {{.OpenRegistration}}

	// Alpha: uncomment this to use the Alpha protocol.
	// AlphaRealmlistListen 0.0.0.0:9100
}`

type HomeConfigFormatted struct {
	DBDriver string

	DBURL string

	OpenRegistration string
}

func FormatHomeConfig(hcf HomeConfigFormatted) []byte {
	tmpl, err := template.New("homeconfig").Parse(DefaultHomeTemplate)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, hcf); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func MakeDefaultHomeConfig(conf *config.Home) []byte {
	return FormatHomeConfig(HomeConfigFormatted{
		DBDriver:         conf.DBDriver,
		DBURL:            conf.DBURL,
		OpenRegistration: fmt.Sprintf("%t", conf.OpenRegistration),
	})
}
