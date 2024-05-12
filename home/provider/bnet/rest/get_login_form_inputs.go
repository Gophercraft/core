package rest

import (
	"fmt"

	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
)

func (provider *Provider) GetLoginFormInputs() (form_inputs *login.FormInputs) {
	srp_url := fmt.Sprintf("https://%s/bnet/login/srp", provider.config.Endpoint)

	form_inputs = make_login_form_inputs()
	util.Set(&form_inputs.SrpUrl, srp_url)

	return
}
