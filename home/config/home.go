package config

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/text"
)

var (
	ErrAlreadyExists = fmt.Errorf("config: already exists")
)

type HomeFile struct {
	HTTPInternal         string
	HostExternal         string
	AuthListen           string
	BnetListen           string
	BnetRESTListen       string
	AlphaRealmlistListen string
	DBDriver             string
	DBURL                string
	OpenRegistration     bool
	Admin                map[string]string
}

type RealmsFile struct {
	Realms map[uint64]Realm
}

type Realm struct {
	FP     string
	Armory string
}

type Home struct {
	Dir         string
	Certificate tls.Certificate
	HomeFile
}

func LoadHome(at string) (*Home, error) {
	ac := new(Home)
	ac.Dir = at

	c, err := ioutil.ReadFile(filepath.Join(at, "Home.txt"))
	if err != nil {
		return nil, err
	}

	err = text.Unmarshal(c, &ac.HomeFile)
	if err != nil {
		return nil, err
	}

	ac.Certificate, err = tls.LoadX509KeyPair(
		filepath.Join(ac.Dir, "cert.pem"),
		filepath.Join(ac.Dir, "key.pem"),
	)

	if ac.HostExternal == "" {
		ac.HostExternal = "localhost"
	}

	if err != nil {
		return nil, err
	}

	return ac, nil
}

func GenerateTLSKeyPair(at string) error {
	return genPair(
		filepath.Join(at, "cert.pem"),
		filepath.Join(at, "key.pem"))
}

// func GenerateDefaultHome(dbDriver, dbURL, user, pass, at string) error {
// 	path := etc.ParseSystemPath(at)

// 	if path.IsExtant() {
// 		return ErrAlreadyExists
// 	}

// 	if err := path.MakeDir(); err != nil {
// 		return err
// 	}

// 	path.Concat("Home.txt").WriteAll([]byte(fmt.Sprintf(DefaultHome, dbDriver, dbURL, user, pass)))

// 	return GenerateTLSKeyPair(path.Render())
// }

func (h *Home) GenerateKeyPair() error {
	return GenerateTLSKeyPair(h.Dir)
}

func (h *Home) Fingerprint() string {
	f, err := rpcnet.GetCertFileFingerprint(filepath.Join(h.Dir, "cert.pem"))
	if err != nil {
		panic(err)
	}
	return f
}

// func (h *Home) RealmsFile() (*RealmsFile, error) {
// 	b, err := ioutil.ReadFile(filepath.Join(h.Dir, "Realms.txt"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	rf := new(RealmsFile)
// 	err = text.Unmarshal(b, rf)
// 	return rf, err
// }

// func (h *Home) WriteRealmsFile(rf *RealmsFile) error {
// 	data, err := text.Marshal(rf)
// 	if err != nil {
// 		return err
// 	}

// 	return ioutil.WriteFile(filepath.Join(h.Dir, "Realms.txt"), data, 0700)
// }

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
