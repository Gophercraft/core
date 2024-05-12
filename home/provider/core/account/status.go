package account

import (
	"context"
	"fmt"

	"github.com/Gophercraft/core/home/login"
	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/home/protocol/pb/account"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (provider *account_provider) get_account(ctx context.Context, credential string) (account *home_models.Account, status auth.WebTokenStatus, err error) {
	var address string
	address, err = protocol.GetPeerAddress(ctx)
	if err != nil {
		return
	}

	return login.CheckCredential(provider.home_db, home_models.Login_Core, credential, address, "")
}

var ErrNoViewStatusPermission = fmt.Errorf("home/provider/core/account: you do not have permission to view the status of other accounts")

func (provider *account_provider) GetAccountStatus(ctx context.Context, account_status_request *account.AccountStatusRequest) (account_status *account.AccountStatus, err error) {
	// First, find the account of the caller
	var (
		game_accounts    []home_models.GameAccount
		caller_account   *home_models.Account
		viewed_account   *home_models.Account
		web_token_status auth.WebTokenStatus
		can_view_account = false
		found            bool
	)
	caller_account, web_token_status, err = provider.get_account(ctx, account_status_request.Credential)
	if err != nil {
		return
	}

	if web_token_status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/core/account: cannot view account without being authenticated")
		return
	}

	// Now, determine if the caller has permission to view the status of this account
	if caller_account.ID == account_status_request.Id {
		// Caller is always allowed to view their own account status
		can_view_account = true
	} else {
		// TODO: replace with permission table check
		if caller_account.Tier >= auth.AccountTier_MODERATOR {
			can_view_account = true
		}
	}

	if !can_view_account {
		err = ErrNoViewStatusPermission
		return
	}

	if caller_account.ID == account_status_request.Id {
		viewed_account = caller_account
	} else {
		viewed_account = new(home_models.Account)
		found, err = provider.home_db.Table("Account").Where(query.Eq("ID", account_status_request.Id)).Get(viewed_account)
		if err != nil {
			log.Warn(err)
			err = fmt.Errorf("home/provider/core/account: error getting account from database")
			return
		}

		if !found {
			err = fmt.Errorf("home/provider/core/account: account %d not found", account_status_request.Id)
			return
		}
	}

	err = provider.home_db.Table("GameAccount").Where(query.Eq("Owner", account_status_request.Id)).Find(&game_accounts)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/account: error getting game accounts from database")
		return
	}

	account_status = new(account.AccountStatus)
	account_status.Name = viewed_account.Username
	account_status.Id = viewed_account.ID
	account_status.Email = viewed_account.Email
	account_status.Tier = viewed_account.Tier
	account_status.CreationTime = timestamppb.New(viewed_account.CreatedAt)
	account_status.Locked = viewed_account.Locked
	account_status.Suspended = viewed_account.Suspended
	if account_status.Suspended {
		account_status.SuspensionLiftTime = timestamppb.New(viewed_account.UnsuspendAt)
	}

	account_status.GameAccounts = make([]*account.GameAccountStatus, len(game_accounts))

	for i := range game_accounts {
		game_account := &game_accounts[i]
		status := new(account.GameAccountStatus)
		account_status.GameAccounts[i] = status
		status.Id = game_account.ID
		status.Name = game_account.Name
		status.Active = game_account.Active
		status.Banned = game_account.Banned
		status.Suspended = game_account.Suspended
		if status.Suspended {
			status.SuspensionLiftTime = timestamppb.New(game_account.UnsuspendAt)
		}

		var character_counts []home_models.CharacterCount
		err = provider.home_db.Table("CharacterCount").Where(query.Eq("GameAccountID", game_account.ID)).Find(&character_counts)
		if err != nil {
			panic(err)
		}

		for _, count := range character_counts {
			status.Characters += count.Count
		}

	}

	return
}
