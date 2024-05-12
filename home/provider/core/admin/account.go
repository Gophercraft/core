package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/Gophercraft/core/home/login"
	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/core/home/protocol/pb/admin"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Find the account
func (provider *admin_provider) validate_access(ctx context.Context, credential string) (account *home_models.Account, err error) {
	var (
		web_token_status auth.WebTokenStatus
		address          string
	)

	address, err = protocol.GetPeerAddress(ctx)
	if err != nil {
		return
	}

	account, web_token_status, err = login.CheckCredential(provider.home_db, home_models.Login_Core, credential, address, "")
	if err != nil {
		return
	}

	if web_token_status != auth.WebTokenStatus_AUTHENTICATED {
		err = fmt.Errorf("home/provider/core/admin: credential is expired, or has not been authenticated fully")
		return
	}

	if account.Tier < auth.AccountTier_MODERATOR {
		err = fmt.Errorf("home/provider/core/admin: you must be moderator or higher to use the admin service")
		return
	}

	return
}

// find a
func (provider *admin_provider) find_administered_account(admin_account *home_models.Account, id uint64) (administered_account *home_models.Account, err error) {
	administered_account = new(home_models.Account)
	var found bool

	found, err = provider.home_db.Table("Account").Where(query.Eq("ID", id)).Get(administered_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: cannot connect to database")
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/core/admin: cannot find account to be administered: %d", id)
		return
	}

	if administered_account.Tier >= admin_account.Tier {
		err = fmt.Errorf("home/provider/core/admin: you are not allowed to administer an account with same or higher tier clearance")
		return
	}

	return
}

func make_ban_status(account *home_models.Account) (ban_status *admin.BanStatus) {

	ban_status = new(admin.BanStatus)
	ban_status.Id = account.ID
	ban_status.Banned = account.Banned
	ban_status.Suspended = account.Suspended
	if account.Suspended {
		ban_status.UnsuspendTime = timestamppb.New(account.UnsuspendAt)
	}
	return
}

func (provider *admin_provider) BanAccount(ctx context.Context, ban_request *admin.AccountRequest) (ban_status *admin.BanStatus, err error) {
	var (
		admin_account  *home_models.Account
		banned_account *home_models.Account
		updated        uint64
	)
	admin_account, err = provider.validate_access(ctx, ban_request.Credential)
	if err != nil {
		return
	}

	banned_account, err = provider.find_administered_account(admin_account, ban_request.Id)
	if err != nil {
		return
	}

	banned_account.Banned = true

	updated, err = provider.home_db.Table("Account").Where(query.Eq("ID", banned_account.ID)).Columns("Banned").Update(banned_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	ban_status = make_ban_status(banned_account)

	return
}

func (provider *admin_provider) UnbanAccount(ctx context.Context, unban_request *admin.AccountRequest) (ban_status *admin.BanStatus, err error) {

	var (
		admin_account    *home_models.Account
		unbanned_account *home_models.Account
		updated          uint64
	)
	admin_account, err = provider.validate_access(ctx, unban_request.Credential)
	if err != nil {
		return
	}

	unbanned_account, err = provider.find_administered_account(admin_account, unban_request.Id)
	if err != nil {
		return
	}

	unbanned_account.Banned = false

	updated, err = provider.home_db.Table("Account").Where(query.Eq("ID", unbanned_account.ID)).Columns("Banned").Update(unbanned_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	ban_status = make_ban_status(unbanned_account)

	return
}

func (provider *admin_provider) LockAccount(ctx context.Context, lock_request *admin.AccountRequest) (lock_status *admin.LockStatus, err error) {
	var (
		admin_account  *home_models.Account
		locked_account *home_models.Account
		updated        uint64
	)
	admin_account, err = provider.validate_access(ctx, lock_request.Credential)
	if err != nil {
		return
	}

	locked_account, err = provider.find_administered_account(admin_account, lock_request.Id)
	if err != nil {
		return
	}

	locked_account.Locked = true

	updated, err = provider.home_db.Table("Account").Where(query.Eq("ID", locked_account.ID)).Columns("Banned").Update(locked_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	lock_status = new(admin.LockStatus)
	lock_status.Locked = locked_account.Locked
	return
}

func (provider *admin_provider) UnlockAccount(ctx context.Context, unlock_request *admin.AccountRequest) (lock_status *admin.LockStatus, err error) {
	var (
		admin_account    *home_models.Account
		unlocked_account *home_models.Account
		updated          uint64
	)
	admin_account, err = provider.validate_access(ctx, unlock_request.Credential)
	if err != nil {
		return
	}

	unlocked_account, err = provider.find_administered_account(admin_account, unlock_request.Id)
	if err != nil {
		return
	}

	unlocked_account.Locked = true

	updated, err = provider.home_db.Table("Account").Where(query.Eq("ID", unlocked_account.ID)).Columns("Banned").Update(unlocked_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	lock_status = new(admin.LockStatus)
	lock_status.Locked = unlocked_account.Locked
	return
}

func (provider *admin_provider) SuspendAccount(ctx context.Context, suspend_request *admin.SuspendAccountRequest) (ban_status *admin.BanStatus, err error) {
	var (
		admin_account     *home_models.Account
		suspended_account *home_models.Account
		updated           uint64
	)
	admin_account, err = provider.validate_access(ctx, suspend_request.Credential)
	if err != nil {
		return
	}

	suspended_account, err = provider.find_administered_account(admin_account, suspend_request.Id)
	if err != nil {
		return
	}

	suspension_duration := time.Duration(suspend_request.SuspensionDuration)
	if suspension_duration < 0 {
		suspension_duration = 292 * 365 * 24 * time.Hour
	}

	suspended_account.Suspended = true
	suspended_account.UnsuspendAt = time.Now().Add(suspension_duration)
	updated, err = provider.home_db.Table("Account").Where(query.Eq("ID", suspended_account.ID)).Columns("Suspended", "UnsuspendAt").Update(suspended_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update suspension status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update suspension status")
		return
	}

	ban_status = make_ban_status(suspended_account)
	return
}

func (provider *admin_provider) UnsuspendAccount(ctx context.Context, unsuspend_request *admin.AccountRequest) (ban_status *admin.BanStatus, err error) {

	var (
		admin_account       *home_models.Account
		unsuspended_account *home_models.Account
		updated             uint64
	)
	admin_account, err = provider.validate_access(ctx, unsuspend_request.Credential)
	if err != nil {
		return
	}

	unsuspended_account, err = provider.find_administered_account(admin_account, unsuspend_request.Id)
	if err != nil {
		return
	}

	unsuspended_account.Suspended = false
	unsuspended_account.UnsuspendAt = time.Time{}

	updated, err = provider.home_db.Table("Account").Where(query.Eq("ID", unsuspended_account.ID)).Columns("Suspended", "UnsuspendAt").Update(unsuspended_account)
	if err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	if updated == 0 {
		err = fmt.Errorf("home/provider/core/admin: failed to update ban status")
		return
	}

	ban_status = make_ban_status(unsuspended_account)
	return
}
