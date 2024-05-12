package web

import (
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_get_credential_status(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	credential_status, err := service.provider.GetCredentialStatus(user_info)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, credential_status)
}
