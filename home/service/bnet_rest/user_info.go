package bnet_rest

import (
	"net/http"

	"github.com/Gophercraft/log"
)

type UserInfo struct {
	Address   string
	UserAgent string
	SessionID string
}

func (service *Service) get_user_info(rw http.ResponseWriter, r *http.Request) (user_info *UserInfo, err error) {
	user_info = new(UserInfo)
	user_info.Address = r.RemoteAddr
	user_info.UserAgent = r.UserAgent()
	cookies := r.Cookies()
	log.Dump("cookies", cookies)

	cookie, err := r.Cookie("JSESSIONID")
	if err == nil && cookie.Value != "" && service.provider.IsValidSessionID(cookie.Value) {
		// get cookie data
		user_info.SessionID = cookie.Value
	} else {
		//
		var new_cookie *http.Cookie
		new_cookie, err = service.provider.NewSessionCookie()
		if err != nil {
			return
		}
		log.Warn("Setting new session cookie", new_cookie)
		http.SetCookie(rw, new_cookie)
		user_info.SessionID = new_cookie.Value
	}

	return
}
