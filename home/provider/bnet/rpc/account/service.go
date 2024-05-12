package account

import (
	"context"

	account_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/account/v1"
	v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/account/v1"
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/bnet/rpc/status"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/provider/bnet/rpc/authentication"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
)

type Service struct {
	account_v1.UnimplementedAccountServiceServer
	home_db *database.Container
}

func New(home_db *database.Container) (service *Service) {
	service = new(Service)
	service.home_db = home_db
	return service
}

func (service *Service) GetAccountState(ctx context.Context, request *account_v1.GetAccountStateRequest) (response *account_v1.GetAccountStateResponse, err error) {
	response = new(account_v1.GetAccountStateResponse)
	response.State = new(account_v1.AccountState)

	response.Tags = new(account_v1.AccountFieldTags)

	if request.GetOptions().GetFieldPrivacyInfo() {
		response.State.PrivacyInfo = new(account_v1.PrivacyInfo)
		util.Set(&response.State.PrivacyInfo.IsUsingRid, false)
		util.Set(&response.State.PrivacyInfo.IsVisibleForViewFriends, false)
		util.Set(&response.State.PrivacyInfo.IsHiddenFromFriendFinder, true)

		util.Set(&response.Tags.PrivacyInfoTag, 0xD7CA834D)
	}

	return
}

func (service *Service) GetGameAccountState(ctx context.Context, request *account_v1.GetGameAccountStateRequest) (response *v1.GetGameAccountStateResponse, err error) {
	// Get the account ID that this connection is authenticated with
	auth_details := authentication.GetDetails(ctx)
	if auth_details == nil {
		err = status.Errorf(codes.ERROR_SERVICE_FAILURE_AUTH, "no authentication details in this connection")
		return
	}

	// find requested game account
	var (
		game_account models.GameAccount
		found        bool
	)
	found, err = service.home_db.Table("GameAccount").Where(query.Eq("Owner", auth_details.AccountID), query.Eq("ID", request.GameAccountId.GetLow())).Get(&game_account)
	if err != nil {
		err = status.Errorf(codes.ERROR_INTERNAL, "%w", err)
		return
	}

	if !found {
		err = status.Errorf(codes.ERROR_NOT_EXISTS, "%w", err)
		return
	}

	response = new(account_v1.GetGameAccountStateResponse)
	response.State = new(account_v1.GameAccountState)
	response.Tags = new(account_v1.GameAccountFieldTags)

	if request.GetOptions().GetFieldGameLevelInfo() {
		// get game level info
		response.State.GameLevelInfo = new(account_v1.GameLevelInfo)
		util.Set(&response.State.GameLevelInfo.Name, game_account.Name)
		util.Set(&response.State.GameLevelInfo.Program, 5730135)

		util.Set(&response.Tags.GameLevelInfoTag, 0x5C46D483)
	}

	if request.GetOptions().GetFieldGameStatus() {
		response.State.GameStatus = new(account_v1.GameStatus)
		util.Set(&response.State.GameStatus.IsSuspended, game_account.Suspended)
		util.Set(&response.State.GameStatus.IsBanned, game_account.Banned)
		var expiry uint64
		if !game_account.UnsuspendAt.IsZero() {
			expiry = uint64(game_account.UnsuspendAt.Unix())
		}
		util.Set(&response.State.GameStatus.SuspensionExpires, expiry)

		util.Set(&response.State.GameStatus.Program, 5730135)
		util.Set(&response.Tags.GameStatusTag, 0x98B75F99)
	}

	return
}
