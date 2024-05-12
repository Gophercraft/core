package web

import "net/http"

func (service *Service) handle_get_captcha(rw http.ResponseWriter, r *http.Request) {
	user_info := service.get_user_info(r)

	content_type, image, err := service.provider.GetCaptchaImage(
		user_info,
		r.PathValue("captcha_id"),
	)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", content_type)
	rw.Write(image)
}
