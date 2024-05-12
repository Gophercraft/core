package home

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/log"
)

var err_no_config = fmt.Errorf("app/home: no Homeserver configuration found in the current directory. One can be created by running 'gophercraft create home'")

func get_config() (home_config *config.Home, err error) {
	var path string

	if _, err = os.Stat("home.txt"); err == nil {
		path, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		err = err_no_config
		return
	}

	home_config, err = config.LoadHome(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = err_no_config
		}
		return
	}

	var fingerprint protocol.Fingerprint
	fingerprint, err = protocol.GetCertFileFingerprint(filepath.Join(home_config.Directory, "cert.pem"))
	if err != nil {
		return
	}

	log.Println("This home server's fingerprint is", fingerprint)

	return
}
