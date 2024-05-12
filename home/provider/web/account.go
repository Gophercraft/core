package web

import (
	"fmt"
	"strconv"

	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	api_models "github.com/Gophercraft/core/home/service/web/models"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

// get the account associated with this token
func (provider *service_provider) get_account(user_info *api_models.UserInfo) (account *home_models.Account, status auth.WebTokenStatus, err error) {
	return login.CheckCredential(provider.home_db, home_models.Login_Web, user_info.Token, user_info.Address, user_info.UserAgent)
}

func (provider *service_provider) GetAccountStatus(user_info *api_models.UserInfo) (account_status *api_models.AccountStatus, err error) {
	var account *home_models.Account
	account, _, err = provider.get_account(user_info)
	if err != nil {
		return
	}

	var (
		game_accounts []home_models.GameAccount
	)

	err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account.ID)).Find(&game_accounts)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error getting game accounts from database")
		return
	}

	account_status = new(api_models.AccountStatus)
	account_status.Username = account.Username
	account_status.ID = fmt.Sprintf("%d", account.ID)
	account_status.Email = account.Email
	account_status.Tier = account.Tier.String()
	account_status.CreationDate = strconv.FormatInt(account.CreatedAt.Unix(), 10)
	account_status.Locked = account.Locked
	account_status.Suspended = account.Suspended
	if account_status.Suspended {
		account_status.SuspensionLiftDate = strconv.FormatInt(account.UnsuspendAt.Unix(), 10)
	}

	account_status.Authenticator = account.Authenticator

	account_status.GameAccounts = make([]api_models.GameAccountStatus, len(game_accounts))

	for i := range game_accounts {
		game_account := &game_accounts[i]
		status := &account_status.GameAccounts[i]
		status.ID = fmt.Sprintf("%d", game_account.ID)
		status.Name = game_account.Name
		status.Active = game_account.Active
		status.Banned = game_account.Banned
		status.Suspended = game_account.Suspended
		if status.Suspended {
			status.SuspensionLiftDate = strconv.FormatInt(game_account.UnsuspendAt.Unix(), 10)
		}

		var character_counts []home_models.CharacterCount
		err = provider.home_db.Table("CharacterCount").Where(query.Eq("GameAccountID", account.ID)).Find(&character_counts)
		if err != nil {
			panic(err)
		}

		for _, count := range character_counts {
			status.Characters += count.Count
		}

	}

	return
}

func (provider *service_provider) NewGameAccount(user_info *api_models.UserInfo, new_game_account_request *api_models.NewGameAccountRequest) (response *api_models.NewGameAccountResponse, err error) {
	var (
		account          *home_models.Account
		game_account     home_models.GameAccount
		web_token_status auth.WebTokenStatus
		count            uint64
	)
	account, web_token_status, err = provider.get_account(user_info)
	if err != nil {
		return
	}
	if web_token_status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/web: not authenticated")
		return
	}

	if len(new_game_account_request.Name) > int(provider.config.MaxGameAccountNameLength) {
		err = fmt.Errorf("home/provider/web: game account name is too long")
		return
	}

	// count existing game accounts
	count, err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account.ID)).Count()
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error counting number of game accounts")
		return
	}

	if count > uint64(provider.config.MaxGameAccounts) {
		err = fmt.Errorf("home/provider/web: error counting number of game accounts")
		return
	}

	game_account.Name = new_game_account_request.Name
	game_account.Owner = account.ID
	err = provider.home_db.Table("GameAccount").Insert(&game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error inserting game account into database")
		return
	}
	response = new(api_models.NewGameAccountResponse)
	response.ID = fmt.Sprintf("%d", game_account.ID)

	return
}

func (provider *service_provider) ActivateGameAccount(user_info *api_models.UserInfo, entity string) (err error) {
	var (
		account          *home_models.Account
		game_account_id  uint64
		found            bool
		web_token_status auth.WebTokenStatus
		game_account     home_models.GameAccount
	)
	account, web_token_status, err = provider.get_account(user_info)
	if err != nil {
		return
	}
	account_id := account.ID
	if web_token_status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/web: not authenticated")
		return
	}

	game_account_id, err = strconv.ParseUint(entity, 10, 64)
	if err != nil {
		return
	}

	// make sure game account exists
	found, err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account_id), query.Eq("ID", game_account_id)).Get(&game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error getting game account from database")
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/web: could not find game account owned by account %d with the id %d", account_id, game_account_id)
		return
	}

	// de-activate all game accounts that are not this one
	game_account.Active = false
	if _, err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account_id), query.Not(query.Eq("ID", game_account_id))).Columns("Active").Update(&game_account); err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error deactivating other game_accounts")
		return
	}

	// Activate game accounts that are this one
	game_account.Active = true
	if _, err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account_id), query.Eq("ID", game_account_id)).Columns("Active").Update(&game_account); err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error deactivating other game_accounts")
		return
	}

	return
}

func (provider *service_provider) DeleteGameAccount(user_info *api_models.UserInfo, entity string) (err error) {
	var (
		account          *home_models.Account
		game_account_id  uint64
		deleted          uint64
		web_token_status auth.WebTokenStatus
	)
	account, web_token_status, err = provider.get_account(user_info)
	if err != nil {
		return
	}
	account_id := account.ID
	if web_token_status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/web: not authenticated")
		return
	}

	game_account_id, err = strconv.ParseUint(entity, 10, 64)
	if err != nil {
		return
	}

	deleted, err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account_id), query.Eq("ID", game_account_id)).Delete()
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error deleting game account")
		return
	}

	if deleted == 0 {
		err = fmt.Errorf("home/provider/web: could not find game account %d to be", game_account_id)
		return
	}

	return
}

func (provider *service_provider) RenameGameAccount(user_info *api_models.UserInfo, entity string, rename_request *api_models.RenameGameAccountRequest) (err error) {
	var (
		game_account_id  uint64
		updated          uint64
		account          *home_models.Account
		web_token_status auth.WebTokenStatus
	)
	account, web_token_status, err = provider.get_account(user_info)
	if err != nil {
		return
	}
	account_id := account.ID
	if web_token_status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/web: not authenticated")
		return
	}

	game_account_id, err = strconv.ParseUint(entity, 10, 64)
	if err != nil {
		return
	}

	if len(rename_request.Name) > int(provider.config.MaxGameAccountNameLength) {
		err = fmt.Errorf("home/provider/web: game account name is too long")
		return
	}

	var game_account models.GameAccount
	game_account.Name = rename_request.Name

	updated, err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account_id), query.Eq("ID", game_account_id)).Columns("Name").Update(&game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: error updating game account name")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/web: could not find game account %d to update", game_account_id)
		return
	}

	return
}
