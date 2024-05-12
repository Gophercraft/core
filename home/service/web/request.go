package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Gophercraft/core/home/service/web/models"
)

// TODO: replace with reverse proxy header if available
func (service *Service) get_remote_address(r *http.Request) (remote_address string) {
	remote_address = r.RemoteAddr
	return
}

func (service *Service) get_user_info(r *http.Request) (user_info *models.UserInfo) {
	user_info = &models.UserInfo{
		Token:     r.Header.Get("X-GC-Credential"),
		Address:   service.get_remote_address(r),
		UserAgent: r.UserAgent(),
	}
	return
}

func read_request(r *http.Request, arguments any) (err error) {
	var json_data []byte
	json_data, err = io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(json_data, arguments)
	return
}

func respond(rw http.ResponseWriter, status int, result any) {
	var json_data []byte
	var err error
	if json_data, err = json.Marshal(result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	rw.Write(json_data)
}
