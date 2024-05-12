package wizard

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Gophercraft/core/app/config"
)

const DefaultHomeTemplate = `{
	// This block configures the services hosted
	// by Gophercraft Home, and their addresses
	HostServices
	{
		{
			// The HTTP API for your Home server.
			Service web
			Address 0.0.0.0:32776 
		}

		{
			// The GRPC server used to communicate
			// between services in a Gophercraft network
			Service core
			Address 0.0.0.0:32777
		}

		{
			// The Grunt login/realmlist service
			Service grunt
			Address 0.0.0.0:3724
		}

		{
			// The ancient realm list server.
			// This should be disabled if you aren't hosting a very old game
			Service old_realmlist
			Address 0.0.0.0:9100
			// Only display realms with this build number
			Build 3368
		}

		{
			// The Protobuf/TLS-based RPC protocol used by newer clients to request realmlists and find the REST API
			Service bnet_rpc
			Address 0.0.0.0:1119
		}

		{
			// The RESTful login API used by BNet to authenticate newer clients
			Service bnet_rest
			Address 0.0.0.0:1120
		}
	}

	// This block stores the public addresses of
	// services 
	ServiceEndpoints
	{
		web http://127.0.0.1:32776/

		grunt 127.0.0.1:3724
		
		bnet_rpc 127.0.0.1:1119

		bnet_rest 127.0.0.1:1120

		old_realmlist 127.0.0.1:9100
	}

	// the Phylactery storage engine to use
	DatabaseEngine {{.DatabaseEngine}}

	// the path of the Phylactery database
	DatabasePath {{.DatabasePath}}

	// true = anyone can register using the web UI
	// false = only administrators can register accounts using the 'gophercraft' command
	OpenRegistration false

	// true = user must confirm their email to use the server
	EmailVerificationRequired false

	// Methods 
	TwoFactorAuthMethods
	{
		TOTP
	}

	// true = user will be required to set a 2FA mechanism after first login
	// false = user is free to 
	TwoFactorAuthRequired false
}`

type home_config_template struct {
	DatabaseEngine   string
	DatabasePath     string
	OpenRegistration string
}

func format_home_config(config_template home_config_template) []byte {
	tmpl, err := template.New("homeconfig").Parse(DefaultHomeTemplate)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, config_template); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func MakeDefaultHomeConfig(conf *config.Home) []byte {
	return format_home_config(home_config_template{
		DatabaseEngine:   conf.File.DatabaseEngine,
		DatabasePath:     conf.File.DatabasePath,
		OpenRegistration: fmt.Sprintf("%t", conf.File.OpenRegistration),
	})
}
