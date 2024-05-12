package web

import (
	"fmt"
	"time"

	home_models "github.com/Gophercraft/core/home/models"
	api_models "github.com/Gophercraft/core/home/service/web/models"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *service_provider) clean_credential() {
	if time.Since(provider.last_credential_sweep) > 10*time.Minute {
		now := time.Now()

		provider.home_db.Table("WebToken").Where(query.Lt("Expiry", now)).Delete()
		provider.last_credential_sweep = now
	}
}

func (provider *service_provider) token_is_expired(web_token *home_models.WebToken) bool {
	now := time.Now()
	return web_token.Expiry.Before(now)
}

func (provider *service_provider) expire_token(web_token *home_models.WebToken) {
	provider.home_db.Table("WebToken").Where(query.Eq("Token", web_token.Token)).Delete()
}

func (provider *service_provider) GetCredentialStatus(user_info *api_models.UserInfo) (credential_status *api_models.CredentialStatus, err error) {
	provider.clean_credential()

	var (
		web_token home_models.WebToken
		found     bool
		account   home_models.Account
	)
	found, err = provider.home_db.Table("WebToken").Where(query.Eq("Token", user_info.Token)).Get(&web_token)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error getting token from database")
		return
	}
	if found {
		// if token was found, also try to find the underlying account
		found, err = provider.home_db.Table("Account").Where(query.Eq("ID", web_token.Account)).Get(&account)
		if err != nil {
			log.Warn(err)
			err = fmt.Errorf("home/provider/web: error getting token account from database")
			return
		}
		if found {
			if account.Banned {
				// The account must have been banned while we were logged in
				// we are not giving service to any banned accounts, so delete the token
				provider.home_db.Table("WebToken").Where(query.Eq("Token", user_info.Token)).Delete()

				credential_status = &api_models.CredentialStatus{
					CredentialIsValid: found,
				}
				return
			}
		}
	}

	credential_status = &api_models.CredentialStatus{
		CredentialIsValid: found,
		WebTokenStatus:    web_token.Status.String(),
	}

	if found {
		if provider.token_is_expired(&web_token) {
			credential_status.CredentialIsValid = false
		}
	}

	return
}
