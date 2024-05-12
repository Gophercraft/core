package web

import (
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_get_version_info(rw http.ResponseWriter, r *http.Request) {
	version_info, err := service.provider.GetVersionInfo()
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, version_info)
}
