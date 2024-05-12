package wizard

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/text"
)

type ConnectionChallenge struct {
	Trusted     bool
	Host        string
	Address     string
	Fingerprint protocol.Fingerprint
}

func (c *Configurator) ConnectDefault() (challenge *ConnectionChallenge, err error) {
	return c.Connect("localhost:32777")
}

func (c *Configurator) Connect(hostport string) (challenge *ConnectionChallenge, err error) {
	challenge = new(ConnectionChallenge)

	var (
		address *net.TCPAddr
	)

	address, err = net.ResolveTCPAddr("tcp", hostport)
	if err != nil {
		return
	}

	challenge.Host = hostport
	challenge.Address = address.String()

	known_host_fingerprint, known := c.known_hosts[hostport]

	challenge.Fingerprint, err = protocol.FingerprintServer(hostport)
	if err != nil {
		return
	}

	if known {
		if !protocol.FingerprintsEqual(challenge.Fingerprint, known_host_fingerprint) {
			fmt.Println("Server fingerprint:", challenge.Fingerprint)
			fmt.Println("Trusted fingerprint:", known_host_fingerprint)
			err = fmt.Errorf("!!! WARNING !!! Home server fingerprint does not match the one on record for '%s'.\nSomeone may be trying to do something nasty.\nIf you know *for a certainty* this is not the case, run: gophercraft remove-known-host %s", hostport, hostport)
			return
		}
		challenge.Trusted = true
	} else {
		challenge.Trusted = false
	}

	return
}

func (c *Configurator) ConfirmConnection(challenge *ConnectionChallenge) (err error) {
	c.connection_info.Connected = true
	c.connection_info.Host = challenge.Host
	c.connection_info.Address = challenge.Address
	c.known_hosts[challenge.Host] = challenge.Fingerprint
	err = c.write_known_hosts()
	if err != nil {
		return
	}
	err = c.write_connection_info()
	if err != nil {
		return
	}
	return
}

func (c *Configurator) Close() (err error) {
	if c.home_connection != nil {
		err = c.home_connection.Close()
		c.home_connection = nil
	}
	return
}

func (c *Configurator) Disconnect() (err error) {
	c.Close()
	c.connection_info.Connected = false
	c.connection_info.Address = ""
	c.connection_info.Host = ""
	c.login_info = login_info{}

	if err = c.write_connection_info(); err != nil {
		return
	}

	if err = c.write_login_info(); err != nil {
		return
	}
	return
}

func (c *Configurator) RemoteHost() string {
	return c.connection_info.Host
}

func (c *Configurator) RemoteAddress() string {
	return c.connection_info.Address
}

type connection_info struct {
	Connected bool
	Host      string
	Address   string
}

func (c *Configurator) read_connection_info() (err error) {
	var connection_file []byte
	connection_txt := filepath.Join(c.directory, "connection_info.txt")
	connection_file, err = os.ReadFile(connection_txt)
	if err == nil {
		err = text.Unmarshal(connection_file, &c.connection_info)
		if err != nil {
			return
		}
	}
	err = nil
	return
}

func (c *Configurator) write_connection_info() (err error) {
	var connection_file []byte
	connection_file, err = text.Marshal(&c.connection_info)
	if err != nil {
		return
	}
	connection_txt := filepath.Join(c.directory, "connection_info.txt")
	return os.WriteFile(connection_txt, connection_file, 0700)
}

func (c *Configurator) Forget(host string) (err error) {
	_, known := c.known_hosts[host]
	if !known {
		err = fmt.Errorf("wizard: host '%s' is unknown", host)
		return
	}

	delete(c.known_hosts, host)
	err = c.write_known_hosts()
	return
}
