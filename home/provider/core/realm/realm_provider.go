package realm

import (
	"context"
	"time"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/home/protocol/pb/realm"
	"github.com/Gophercraft/core/home/provider/core/util"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type realm_provider struct {
	realm.UnimplementedRealmServiceServer
	home_db *database.Container
}

func (provider *realm_provider) Enlist(ctx context.Context, enlist_request *realm.EnlistRequest) (enlist_response *realm.EnlistResponse, err error) {
	var (
		token models.WebToken
		found bool
	)

	found, err = provider.home_db.Table("WebToken").Where(query.Eq("Token", enlist_request.WebToken)).Get(&token)
	if err != nil {
		log.Warn(err)
		err = status.Errorf(codes.Internal, "webtoken database error")
		return
	}
	if !found {
		err = status.Error(codes.Unauthenticated, "webtoken not found")
		return
	}

	var enlisting_account models.Account
	found, err = provider.home_db.Table("Account").Where(query.Eq("ID", token.Account)).Get(&enlisting_account)
	if err != nil {
		log.Warn(err)
		err = status.Errorf(codes.Internal, "account database error")
		return
	}
	if !found {
		err = status.Error(codes.Unauthenticated, "account referenced by web token not found")
		return
	}

	can_enlist := enlisting_account.Tier >= auth.AccountTier_ADMIN
	if !can_enlist {
		err = status.Errorf(codes.PermissionDenied, "cannot enlist")
		return
	}

	if len(enlist_request.Fingerprint) != 64 {
		err = status.Errorf(codes.InvalidArgument, "fingerprint is invalid length (%d)", len(enlist_request.Fingerprint))
		return
	}

	var enlisted_realm models.EnlistedRealm
	copy(enlisted_realm.Fingerprint[:], enlist_request.Fingerprint)
	enlisted_realm.Owner = enlisting_account.ID

	if err = provider.home_db.Table("EnlistedRealm").Insert(&enlisted_realm); err != nil {
		err = status.Errorf(codes.Internal, "enlisted realm database error")
		return
	}

	enlist_response = &realm.EnlistResponse{
		Id: enlisted_realm.ID,
	}

	return
}

func (provider *realm_provider) Announce(ctx context.Context, announce_request *realm.AnnounceRequest) (announce_response *realm.AnnounceResponse, err error) {
	if err = util.CheckPeerIdentity(provider.home_db, announce_request.GetId(), ctx); err != nil {
		return
	}

	tx, err := provider.home_db.NewTransaction()
	if err != nil {
		err = status.Error(codes.Internal, "database tx error")
		return
	}

	// delete existing realm and replace
	realm_table := tx.Table("Realm")
	realm_table.Where(query.Eq("ID", announce_request.Id)).Delete()
	realm_table.Insert(&models.Realm{
		ID:              announce_request.Id,
		Name:            announce_request.LongName,
		Build:           version.Build(announce_request.Build),
		Type:            models.RealmType(announce_request.Type),
		Address:         announce_request.Address,
		RedirectAddress: announce_request.RedirectAddress,
		Description:     announce_request.Address,
		ActivePlayers:   announce_request.ActivePlayers,
		Category:        announce_request.Category,
		LastUpdated:     time.Now(),
		RequiredTier:    announce_request.RequiredTier,
	})

	err = provider.home_db.Commit(tx)
	if err != nil {
		log.Warn(err)
		err = status.Error(codes.Internal, "error committing realm announcement")
		return
	}

	return
}

func (provider *realm_provider) SaveCharacterCount(ctx context.Context, character_count_data *realm.CharacterCountData) (response *realm.SaveResponse, err error) {
	if err = util.CheckPeerIdentity(provider.home_db, character_count_data.GetRealmId(), ctx); err != nil {
		return
	}

	if _, err = provider.home_db.Table("CharacterCount").
		Where(
			query.Eq("RealmID", character_count_data.GetRealmId()),
			query.Eq("GameAccountID", character_count_data.GameAccount),
		).Delete(); err != nil {
		return
	}

	var character_count models.CharacterCount
	character_count.RealmID = character_count_data.GetRealmId()
	character_count.GameAccountID = character_count_data.GetCharacterCount()
	character_count.Count = character_count.Count
	if err = provider.home_db.Table("CharacterCount").Insert(&character_count); err != nil {
		return
	}

	return
}

func (provider *realm_provider) SaveLastCharacterLoggedIn(ctx context.Context, character_logged_in_data *realm.CharacterLoggedInData) (response *realm.SaveResponse, err error) {
	if err = util.CheckPeerIdentity(provider.home_db, character_logged_in_data.GetRealmId(), ctx); err != nil {
		return
	}

	if _, err = provider.home_db.Table("LastCharacterLoggedIn").
		Where(
			query.Eq("RealmID", character_logged_in_data.GetRealmId()),
			query.Eq("GameAccountID", character_logged_in_data.GetGameAccountId()),
		).Delete(); err != nil {
		return
	}

	var last_character_logged_in models.LastCharacterLoggedIn
	last_character_logged_in.RealmID = character_logged_in_data.GetRealmId()
	last_character_logged_in.GameAccountID = character_logged_in_data.GetGameAccountId()
	last_character_logged_in.CharacterName = character_logged_in_data.GetCharacterName()
	last_character_logged_in.CharacterID = character_logged_in_data.GetCharacterId()
	last_character_logged_in.LastPlayedTime = time.Unix(int64(character_logged_in_data.LastPlayedTime), 0)

	if err = provider.home_db.Table("LastCharacterLoggedIn").Insert(&last_character_logged_in); err != nil {
		return
	}

	return
}

func New(db *database.Container) (provider *realm_provider) {
	provider = new(realm_provider)
	provider.home_db = db
	return
}
