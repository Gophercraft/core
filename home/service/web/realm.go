package web

import (
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_get_realm_status(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	realm_status_list, err := service.provider.GetRealmStatusList(user_info)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, realm_status_list)
}
