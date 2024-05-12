package wizard

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/text"
	"google.golang.org/grpc"
)

type Configurator struct {
	directory       string
	known_hosts     map[string]protocol.Fingerprint
	connection_info connection_info
	login_info      login_info
	home_connection *grpc.ClientConn
}

func New() (c *Configurator) {
	c = new(Configurator)
	c.directory = Directory()
	return
}

var ErrDisconnected = fmt.Errorf("wizard: disconnected from home server")

func (c *Configurator) Init() (err error) {
	err = os.MkdirAll(c.directory, 0700)
	if err != nil {
		return
	}

	err = c.read_known_hosts()
	if err != nil {
		return
	}

	err = c.read_connection_info()
	if err != nil {
		return
	}

	err = c.read_login_info()
	if err != nil {
		return
	}

	// If the wizard is connected, see if we can create a GRPC connection
	if c.connection_info.Connected {
		trusted_fingerprint, found := c.known_hosts[c.connection_info.Host]

		disconnect := false

		if found {
			c.home_connection, err = protocol.DialConn(c.connection_info.Host, trusted_fingerprint, nil)
			if err != nil {
				disconnect = true
			}
		} else {
			disconnect = true
		}

		if disconnect {
			c.connection_info.Connected = false
			c.connection_info.Host = ""
			c.login_info = login_info{}

			if err = c.write_connection_info(); err != nil {
				return
			}

			if err = c.write_login_info(); err != nil {
				return
			}

			err = fmt.Errorf("%w: %s", ErrDisconnected, c.connection_info.Host)
			return
		}
	}

	if c.connection_info.Connected {
		if c.login_info.LoggedIn {
			credential_status, err := c.GetCredentialStatus()
			login_expired := false
			if err != nil {
				log.Warn("failed to get credential status:", err)
				login_expired = true
			} else {
				if credential_status.WebTokenStatus != auth.WebTokenStatus_AUTHENTICATED {
					login_expired = false
				}
			}
			if login_expired {
				c.login_info = login_info{}
				c.write_login_info()
			}
		}
	}

	return
}

func (c *Configurator) Connected() bool {
	return c.connection_info.Connected
}

func (c *Configurator) read_known_hosts() (err error) {
	var known_hosts_file []byte
	known_hosts_txt := filepath.Join(c.directory, "known_hosts.txt")
	known_hosts_file, err = os.ReadFile(known_hosts_txt)
	if err == nil {
		err = text.Unmarshal(known_hosts_file, &c.known_hosts)
		if err != nil {
			return
		}
	} else {
		err = nil
		c.known_hosts = make(map[string]protocol.Fingerprint)
	}

	return
}

func (c *Configurator) write_known_hosts() (err error) {
	var known_hosts_file []byte
	known_hosts_file, err = text.Marshal(&c.known_hosts)
	if err != nil {
		return
	}
	known_hosts_txt := filepath.Join(c.directory, "known_hosts.txt")
	return os.WriteFile(known_hosts_txt, known_hosts_file, 0700)
}

type login_info struct {
	LoggedIn bool
	WebToken string
}

func (c *Configurator) read_login_info() (err error) {
	var login_file []byte
	login_txt := filepath.Join(c.directory, "login_info.txt")
	login_file, err = os.ReadFile(login_txt)
	if err == nil {
		err = text.Unmarshal(login_file, &c.login_info)
		if err != nil {
			return
		}
	}
	err = nil
	return
}

func (c *Configurator) write_login_info() (err error) {
	var login_file []byte
	login_file, err = text.Marshal(&c.login_info)
	if err != nil {
		return
	}
	login_txt := filepath.Join(c.directory, "login_info.txt")
	return os.WriteFile(login_txt, login_file, 0700)
}
