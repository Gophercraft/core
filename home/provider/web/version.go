package web

import (
	api_models "github.com/Gophercraft/core/home/service/web/models"
	"github.com/Gophercraft/core/version"
)

func (provider *service_provider) GetVersionInfo() (version_info *api_models.VersionInfo, err error) {
	version_info = &api_models.VersionInfo{
		CoreVersion: version.GophercraftVersion.String(),
		Brand:       "Gophercraft",
		ProjectURL:  "https://github.com/Gophercraft",
	}
	return
}
