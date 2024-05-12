package web

import "io/fs"

type ServiceConfig struct {
	Address string
	WebApp  fs.FS
}
