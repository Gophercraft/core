package web

import (
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_get_service_addresses(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	addresses, err := service.provider.GetServiceAddresses(user_info)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, addresses)
}
