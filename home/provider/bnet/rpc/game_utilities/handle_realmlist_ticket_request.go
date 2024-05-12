package game_utilities

import (
	"context"

	"github.com/Gophercraft/core/bnet"
	game_utilities_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/game_utilities/v1"
	"github.com/Gophercraft/core/bnet/pb/realmlist"
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/bnet/rpc/status"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/provider/bnet/rpc/authentication"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

func handle_realmlist_ticket_request_v1(service *Service, ctx context.Context, parameters ClientRequestParameters, response *game_utilities_v1.ClientResponse) (err error) {
	auth_details := authentication.GetDetails(ctx)
	if auth_details == nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_ACCOUNT_DENIED, "could not find authentication details associated with connection")
		return
	}

	var (
		game_account       models.GameAccount
		game_account_found bool
	)
	var json_text []byte

	identity := parameters.Get("Param_Identity")
	if identity != nil {
		log.Println("parsing", string(identity.BlobValue))

		_, json_text, err = split_blob_parameter(identity.BlobValue)
		if err != nil {
			return
		}
		json_text = remove_null_byte(json_text)

		log.Println("parsing json", string(json_text))

		var data realmlist.RealmListTicketIdentity
		err = json_unmarshal(json_text, &data)
		if err != nil {
			err = status.Errorf(codes.ERROR_UTIL_SERVER_INVALID_IDENTITY_ARGS, "could not unmarshal: %w", err)
			return
		}

		game_account_found, err = service.home_db.
			Table("GameAccount").
			Where(
				query.Eq("Owner", uint64(auth_details.AccountID)),
				query.Eq("ID", uint64(data.GetGameAccountID())),
			).
			Get(&game_account)
		if err != nil {
			return
		}
	}

	if !game_account_found {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_INVALID_IDENTITY_ARGS, "game account was not found")
		return
	}

	if game_account.Banned {
		err = status.Errorf(codes.ERROR_GAME_ACCOUNT_BANNED, "game account is permabanned")
		return
	}

	if game_account.Suspended {
		err = status.Errorf(codes.ERROR_GAME_ACCOUNT_SUSPENDED, "game account is suspended")
		return
	}

	var session *Session
	if session = GetSession(ctx); session == nil {
		session = new(Session)
		session.GameAccount = game_account.ID
		var connection *bnet.Connection
		connection = bnet.GetConnection(ctx)
		if connection == nil {
			err = status.Errorf(codes.ERROR_INTERNAL, "no connection associated with this RPC")
			return
		}
		connection.Set(session_context_key, session)
	}

	client_info := parameters.Get("Param_ClientInfo")
	if client_info != nil {
		var client_ticket_info realmlist.RealmListTicketClientInformation
		log.Println(string(client_info.BlobValue))
		_, json_text, err = split_blob_parameter(client_info.BlobValue)
		if err != nil {
			return
		}
		json_text = remove_null_byte(json_text)
		err = json_unmarshal(json_text, &client_ticket_info)
		if err != nil {
			return
		}
		session.Build = version.Build(client_ticket_info.GetInfo().GetVersion().GetVersionBuild())
		session.ClientSecret = client_ticket_info.GetInfo().Secret
	}

	// Append dummy ticket value
	append_blob_value_to_response(&response.Attribute, "Param_RealmListTicket", []byte("AuthRealmListTicket"))

	return
}
