package web

import (
	"fmt"

	"github.com/Gophercraft/core/home/service/web/models"
)

func (provider *service_provider) GetCaptchaImage(user_info *models.UserInfo, id string) (content_type string, data []byte, err error) {
	if !provider.config.UseCaptchas {
		err = fmt.Errorf("captchas not in use")
		return
	}
	content_type = "image/png"
	data, err = provider.captcha_manager.get_image(id)
	return
}
