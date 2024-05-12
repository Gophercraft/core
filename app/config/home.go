package config

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/text"
)

var (
	ErrAlreadyExists = fmt.Errorf("config: already exists")
)

type HomeServiceConfig struct {
	Service HomeServiceID
	Address string
	Build   version.Build
}

type HomeServiceEndpoint struct {
	Service HomeServiceID
	URL     string
}

// HomeFile is the .txt file containing home server settings
type HomeFile struct {
	// The list of services to by hosted by the home server
	HostServices []HomeServiceConfig
	// This is a list of service endpoints
	ServiceEndpoints map[HomeServiceID]string
	DatabaseEngine   string
	DatabasePath     string
	OpenRegistration bool
	// true = user must confirm their email to use the server
	EmailVerificationRequired bool

	// Allowed methods
	TwoFactorAuthMethods []string

	// true = user will be required to set a 2FA mechanism after first login
	// false = user is free to use the system without setting a two-factor auth mechanism
	TwoFactorAuthRequired bool
}

type Home struct {
	Directory   string
	Certificate tls.Certificate
	File        HomeFile
}

func LoadHome(config_directory string) (home_config *Home, err error) {
	home_config = new(Home)
	home_config.Directory = config_directory

	// Load config file
	var config_text []byte
	config_text, err = os.ReadFile(filepath.Join(config_directory, "home.txt"))
	if err != nil {
		return
	}
	err = text.Unmarshal(config_text, &home_config.File)
	if err != nil {
		err = fmt.Errorf("config: could not decode home.txt: %w", err)
		return
	}

	// Load ECDSA keypair
	home_config.Certificate, err = tls.LoadX509KeyPair(
		filepath.Join(home_config.Directory, "cert.pem"),
		filepath.Join(home_config.Directory, "key.pem"),
	)

	if err != nil {
		return nil, err
	}

	return
}

func GenerateKeyPair(directory string) error {
	return generate_ecdsa_keypair(
		filepath.Join(directory, "cert.pem"),
		filepath.Join(directory, "key.pem"))
}

func (h *Home) Fingerprint() string {
	f, err := protocol.GetCertFileFingerprint(filepath.Join(h.Directory, "cert.pem"))
	if err != nil {
		panic(err)
	}
	return f.String()
}

// List all home configs in the chosen Gophercraft root directory.
// On a typical deployment, it will return []string{"Home"}, nil
func ListHomeConfigs(dir string) ([]string, error) {
	list, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var names []string

	for _, entry := range list {
		if entry.IsDir() {
			_, err := os.Stat(filepath.Join(dir, entry.Name(), "Home.txt"))
			if err == nil {
				names = append(names, entry.Name())
			}
		}
	}

	return names, nil
}
