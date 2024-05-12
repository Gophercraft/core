package game_utilities

import (
	"context"
	"time"

	game_utilities_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/game_utilities/v1"
	"github.com/Gophercraft/core/bnet/pb/realmlist"
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/bnet/rpc/status"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/home/provider/bnet/rpc/authentication"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

type realm_list_query struct {
	AccountTier auth.AccountTier
	ClientBuild version.Build
	SubRegionId string
}

// Returns
func (service *Service) build_realmlist(realm_query *realm_list_query) (realm_list_text []byte, err error) {
	var realm_list []models.Realm
	if err = service.home_db.Table("Realm").Where(query.Gte("RequiredTier", realm_query.AccountTier)).Find(&realm_list); err != nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "%w", err)
		return
	}

	// var realm_list []models.Realm
	// for x := 0; x < 850; x++ {
	// 	id := x + 1
	// 	port := 8000 + x
	// 	realm_list = append(realm_list, models.Realm{
	// 		ID:          uint64(id),
	// 		Name:        fmt.Sprintf("Test Realm %05d", id),
	// 		Build:       realm_query.ClientBuild,
	// 		Type:        models.RealmTypeNormal,
	// 		Address:     fmt.Sprintf("0.0.0.0:%d", port),
	// 		LastUpdated: time.Now(),
	// 		Category:    1,
	// 	})

	// 	if x%17 == 0 {
	// 		realm_list[len(realm_list)-1].Category = 2
	// 	}
	// }

	now := time.Now()

	var realm_list_updates realmlist.RealmListUpdates
	realm_list_updates.Updates = make([]*realmlist.RealmState, len(realm_list))

	for i, realm := range realm_list {
		realm_state := new(realmlist.RealmState)
		realm_state.Update = new(realmlist.RealmEntry)

		var realm_flags uint32
		realm_flags |= models.REALM_FLAG_SPECIFYBUILD

		if now.Sub(realm.LastUpdated) >= models.RealmOnlineDelayMeansOffline {
			realm_flags |= models.REALM_FLAG_OFFLINE
		}

		set_realm_address(&realm_state.Update.WowRealmAddress, NewRealmAddress(1, 1, uint16(realm.ID)))

		util.Set(&realm_state.Update.CfgConfigsID, models.RealmTypeToCfgConfigID(realm.Type))

		util.Set(&realm_state.Update.Flags, realm_flags)
		util.Set(&realm_state.Update.Name, realm.Name)
		realm_state.Update.Version = new(realmlist.ClientVersion)
		build_info := realm.Build.BuildInfo()
		var (
			major    uint32
			minor    uint32
			revision uint32
		)
		if build_info != nil {
			major = build_info.MajorVersion
			minor = build_info.MinorVersion
			revision = build_info.BugfixVersion
		}

		util.Set(&realm_state.Update.Version.VersionBuild, uint32(realm.Build))
		util.Set(&realm_state.Update.Version.VersionMajor, major)
		util.Set(&realm_state.Update.Version.VersionMinor, minor)
		util.Set(&realm_state.Update.Version.VersionRevision, revision)

		util.Set(&realm_state.Update.CfgLanguagesID, 1)
		util.Set(&realm_state.Update.CfgTimezonesID, 1)
		util.Set(&realm_state.Update.CfgCategoriesID, 1)
		util.Set(&realm_state.Update.PopulationState, 1)

		// util.Set(&realm_state.Update.CfgRealmsID, uint32(realm.ID))
		util.Set(&realm_state.Update.CfgRealmsID, 1)

		util.Set(&realm_state.Deleting, false)

		realm_list_updates.Updates[i] = realm_state
	}

	realm_list_text, err = json_marshal(&realm_list_updates)
	if err != nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "%w", err)
		return
	}

	realm_list_text = append([]byte("JSONRealmListUpdates:"), realm_list_text...)
	realm_list_text = append(realm_list_text, 0)
	log.Println(string(realm_list_text))

	return
}

func (service *Service) build_character_count_list(game_account uint64) (character_count_list_text []byte, err error) {
	var character_counts []models.CharacterCount
	err = service.home_db.Table("CharacterCount").Where(query.Eq("GameAccountID", game_account)).Find(&character_counts)
	if err != nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "%w", err)
		return
	}
	// for x := 0; x < 150; x++ {
	// 	id := x + 1
	// 	character_counts = append(character_counts, models.CharacterCount{
	// 		RealmID: uint64(id),
	// 		Count:   100,
	// 	})
	// }

	var character_count_list realmlist.RealmCharacterCountList
	for _, count := range character_counts {
		count_entry := new(realmlist.RealmCharacterCountEntry)
		set_realm_address(&count_entry.WowRealmAddress, NewRealmAddress(1, 1, uint16(count.RealmID)))
		util.Set(&count_entry.Count, uint32(count.Count))
		character_count_list.Counts = append(character_count_list.Counts, count_entry)
	}

	character_count_list_text, err = json_marshal(&character_count_list)
	if err != nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "%w", err)
		return
	}

	character_count_list_text = append([]byte("JSONRealmCharacterCountList:"), character_count_list_text...)
	character_count_list_text = append(character_count_list_text, 0)

	log.Println(string(character_count_list_text))
	return
}

func handle_realmlist_request_v1(service *Service, ctx context.Context, parameters ClientRequestParameters, response *game_utilities_v1.ClientResponse) (err error) {
	account_details := authentication.GetDetails(ctx)

	session := GetSession(ctx)
	if session == nil || account_details == nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_INVALID_IDENTITY_ARGS, "game account was not found")
		return
	}

	var query realm_list_query
	query.ClientBuild = session.Build
	query.AccountTier = account_details.AccountTier
	sub_region := parameters.Get("Command_RealmListRequest_v1")
	if sub_region != nil {
		query.SubRegionId = sub_region.GetStringValue()
	}

	var (
		realm_list           []byte
		character_count_list []byte
	)
	realm_list, err = service.build_realmlist(&query)
	if err != nil {
		return
	}

	character_count_list, err = service.build_character_count_list(session.GameAccount)
	if err != nil {
		return
	}

	append_compressed_blob_value_to_response(&response.Attribute, "Param_RealmList", realm_list)
	append_compressed_blob_value_to_response(&response.Attribute, "Param_CharacterCountList", character_count_list)

	return
}
