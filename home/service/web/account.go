package web

import (
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (service *Service) handle_get_account_status(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	registration_challenge, err := service.provider.GetAccountStatus(user_info)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, registration_challenge)
}

func (service *Service) handle_put_new_game_account(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var new_game_account_request models.NewGameAccountRequest
	if err := read_request(r, &new_game_account_request); err != nil {
		http.Error(rw, "malformed request", http.StatusBadRequest)
		return
	}

	new_game_account_response, err := service.provider.NewGameAccount(user_info, &new_game_account_request)
	if err != nil {
		respond(rw, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	respond(rw, http.StatusOK, new_game_account_response)
}

func (service *Service) handle_post_activate_game_account(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var error_message models.Error
	var status = http.StatusOK
	err := service.provider.ActivateGameAccount(user_info, r.PathValue("id"))

	if err != nil {
		error_message.Message = err.Error()
		status = http.StatusInternalServerError
	}

	respond(rw, status, &error_message)
}

func (service *Service) handle_delete_game_account(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var error_message models.Error
	var status = http.StatusOK
	err := service.provider.DeleteGameAccount(user_info, r.PathValue("id"))

	if err != nil {
		error_message.Message = err.Error()
		status = http.StatusInternalServerError
	}

	respond(rw, status, &error_message)
}

func (service *Service) handle_post_rename_game_account(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	var rename_game_account_request models.RenameGameAccountRequest
	if err := read_request(r, &rename_game_account_request); err != nil {
		http.Error(rw, "malformed request", http.StatusBadRequest)
		return
	}

	var error_message models.Error
	var status = http.StatusOK
	err := service.provider.RenameGameAccount(user_info, r.PathValue("id"), &rename_game_account_request)
	if err != nil {
		error_message.Message = err.Error()
		status = http.StatusInternalServerError
	}

	respond(rw, status, &error_message)
}
