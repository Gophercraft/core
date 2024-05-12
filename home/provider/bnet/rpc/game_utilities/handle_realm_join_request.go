package game_utilities

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	game_utilities_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/game_utilities/v1"
	"github.com/Gophercraft/core/bnet/pb/realmlist"
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/bnet/rpc/status"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/provider/bnet/rpc/authentication"
	"github.com/Gophercraft/phylactery/database/query"
)

func (service *Service) build_realm_addresses(requested_realm *models.Realm) (realm_address_text []byte, err error) {
	var (
		host        string
		port_string string
		port        uint64
	)

	host, port_string, err = net.SplitHostPort(requested_realm.Address)
	if err != nil {
		return
	}
	port, err = strconv.ParseUint(port_string, 10, 32)
	if err != nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "%w", err)
		return
	}

	var addresses realmlist.RealmListServerIPAddresses

	address_family := new(realmlist.RealmIPAddressFamily)
	util.Set(&address_family.Family, 1)
	addresses.Families = append(addresses.Families, address_family)

	address := new(realmlist.IPAddress)
	util.Set(&address.Ip, host)
	util.Set(&address.Port, uint32(port))

	address_family.Addresses = append(address_family.Addresses, address)

	realm_address_text, err = json_marshal(&addresses)
	if err != nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "%w", err)
		return
	}

	realm_address_text = append([]byte("JSONRealmListServerIPAddresses:"), realm_address_text...)
	return
}

func handle_realm_join_request_v1(service *Service, ctx context.Context, parameters ClientRequestParameters, response *game_utilities_v1.ClientResponse) (err error) {
	auth_details := authentication.GetDetails(ctx)
	if auth_details == nil {
		return status.Errorf(codes.ERROR_USER_SERVER_NOT_PERMITTED_ON_REALM, "no authentication details")
	}
	session := GetSession(ctx)
	if session == nil {
		return status.Errorf(codes.ERROR_USER_SERVER_NOT_PERMITTED_ON_REALM, "no game utilities session")
	}

	join_realm_ID := parameters["Param_RealmAddress"].GetUintValue()
	log.Println("Join request for Realm", join_realm_ID)

	var realm_list []models.Realm
	if err = service.home_db.Table("Realm").Where(query.Gte("RequiredTier", auth_details.AccountTier)).Find(&realm_list); err != nil {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "%w", err)
		return
	}

	var requested_realm *models.Realm

	for i := range realm_list {
		realm := &realm_list[i]
		if realm.ID == join_realm_ID {
			requested_realm = realm
			break
		}
	}

	if requested_realm == nil {
		err = status.Errorf(codes.ERROR_USER_SERVER_NOT_PERMITTED_ON_REALM, "requested realm not found")
		return
	}

	if requested_realm.Build != auth_details.Build {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "client build mismatch (%d) with server build (%d)", auth_details.Build, requested_realm.Build)
		return
	}
	var (
		json_address_text []byte
	)

	json_address_text, err = service.build_realm_addresses(requested_realm)
	if err != nil {
		return
	}

	var server_secret [32]byte
	var secret_key [64]byte
	if _, err = io.ReadFull(rand.Reader, server_secret[:]); err != nil {
		panic(err)
	}
	if len(session.ClientSecret) != 32 {
		err = status.Errorf(codes.ERROR_UTIL_SERVER_FAILED_TO_SERIALIZE_RESPONSE, "client secret (len %d) wasn't the correct length", len(session.ClientSecret))
		return
	}

	copy(secret_key[:32], session.ClientSecret)
	copy(secret_key[32:], server_secret[:])

	// TODO save os and locale
	var account models.Account
	var account_found bool
	if account_found, err = service.home_db.Table("Account").Where(query.Eq("ID", auth_details.AccountID)).Get(&account); err != nil {
		err = status.Errorf(codes.ERROR_INTERNAL, "%w", err)
		return
	}
	if !account_found {
		err = status.Errorf(codes.ERROR_INTERNAL, "account not found")
		return
	}

	account.SessionKey = secret_key[:]

	if _, err = service.home_db.Table("Account").Where(query.Eq("ID", auth_details.AccountID)).Update(&account); err != nil {
		err = status.Errorf(codes.ERROR_INTERNAL, "%w", err)
		return
	}

	realm_join_ticket := fmt.Sprintf("%s:%d", strings.ToUpper(account.Username), session.GameAccount)

	append_blob_value_to_response(&response.Attribute, "Param_RealmJoinTicket", []byte(realm_join_ticket))
	append_compressed_blob_value_to_response(&response.Attribute, "Param_ServerAddresses", []byte(json_address_text))
	append_blob_value_to_response(&response.Attribute, "Param_JoinSecret", server_secret[:])

	return
}
