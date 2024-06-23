package admin

import (
	"context"
	"fmt"
	"time"

	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/admin"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (provider *admin_provider) find_administered_game_account(admin_account *home_models.Account, id uint64) (administered_game_account *home_models.GameAccount, err error) {
	administered_game_account = new(home_models.GameAccount)
	var (
		found                bool
		administered_account *home_models.Account
	)

	found, err = provider.home_db.Table("GameAccount").Where(query.Eq("ID", id)).Get(administered_game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to connect to ")
		return
	}
	if !found {
		err = fmt.Errorf("home/provider/core/admin: game account %d not found", id)
		return
	}

	administered_account, err = provider.find_administered_account(admin_account, administered_account.ID)
	if err != nil {
		return
	}

	return
}

func make_game_account_ban_status(account *home_models.GameAccount) (ban_status *admin.BanStatus) {
	ban_status = new(admin.BanStatus)
	ban_status.Id = account.ID
	ban_status.Banned = account.Banned
	ban_status.Suspended = account.Suspended
	if account.Suspended {
		ban_status.UnsuspendTime = timestamppb.New(account.UnsuspendAt)
	}
	return
}

func (provider *admin_provider) BanGameAccount(ctx context.Context, ban_request *admin.AccountRequest) (ban_status *admin.BanStatus, err error) {
	var (
		admin_account       *home_models.Account
		banned_game_account *home_models.GameAccount
		updated             uint64
	)
	admin_account, err = provider.validate_moderator_access(ctx, ban_request.Credential)
	if err != nil {
		return
	}

	banned_game_account, err = provider.find_administered_game_account(admin_account, ban_request.Id)
	if err != nil {
		return
	}

	banned_game_account.Banned = true

	updated, err = provider.home_db.Table("GameAccount").Where(query.Eq("ID", banned_game_account.ID)).Columns("Banned").Update(banned_game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	ban_status = make_game_account_ban_status(banned_game_account)
	return
}

func (provider *admin_provider) UnbanGameAccount(ctx context.Context, unban_request *admin.AccountRequest) (ban_status *admin.BanStatus, err error) {
	var (
		admin_account         *home_models.Account
		unbanned_game_account *home_models.GameAccount
		updated               uint64
	)
	admin_account, err = provider.validate_moderator_access(ctx, unban_request.Credential)
	if err != nil {
		return
	}

	unbanned_game_account, err = provider.find_administered_game_account(admin_account, unban_request.Id)
	if err != nil {
		return
	}

	unbanned_game_account.Banned = false

	updated, err = provider.home_db.Table("GameAccount").Where(query.Eq("ID", unbanned_game_account.ID)).Columns("Banned").Update(unbanned_game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	ban_status = make_game_account_ban_status(unbanned_game_account)
	return
}

func (provider *admin_provider) SuspendGameAccount(ctx context.Context, suspend_request *admin.SuspendAccountRequest) (ban_status *admin.BanStatus, err error) {
	var (
		admin_account          *home_models.Account
		suspended_game_account *home_models.GameAccount
		updated                uint64
	)
	admin_account, err = provider.validate_moderator_access(ctx, suspend_request.Credential)
	if err != nil {
		return
	}

	suspended_game_account, err = provider.find_administered_game_account(admin_account, suspend_request.Id)
	if err != nil {
		return
	}

	suspension_duration := time.Duration(suspend_request.SuspensionDuration)
	if suspension_duration < 0 {
		suspension_duration = 292 * 365 * 24 * time.Hour
	}

	suspended_game_account.Suspended = true
	suspended_game_account.UnsuspendAt = time.Now().Add(suspension_duration)

	updated, err = provider.home_db.Table("GameAccount").Where(query.Eq("ID", suspended_game_account.ID)).Columns("Banned").Update(suspended_game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	ban_status = make_game_account_ban_status(suspended_game_account)
	return
}

func (provider *admin_provider) UnsuspendGameAccount(ctx context.Context, unsuspend_request *admin.AccountRequest) (ban_status *admin.BanStatus, err error) {
	var (
		admin_account            *home_models.Account
		unsuspended_game_account *home_models.GameAccount
		updated                  uint64
	)
	admin_account, err = provider.validate_moderator_access(ctx, unsuspend_request.Credential)
	if err != nil {
		return
	}

	unsuspended_game_account, err = provider.find_administered_game_account(admin_account, unsuspend_request.Id)
	if err != nil {
		return
	}

	unsuspended_game_account.Suspended = false
	unsuspended_game_account.UnsuspendAt = time.Time{}

	updated, err = provider.home_db.Table("GameAccount").Where(query.Eq("ID", unsuspended_game_account.ID)).Columns("Suspended", "UnsuspendAt").Update(unsuspended_game_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	ban_status = make_game_account_ban_status(unsuspended_game_account)
	return
}
