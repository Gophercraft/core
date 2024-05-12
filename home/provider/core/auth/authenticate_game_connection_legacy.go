package auth

import (
	"crypto/sha1"
	"crypto/subtle"
	"strings"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (provider *auth_provider) authenticate_game_connection_legacy(auth_game_connection *auth.AuthenticateGameConnectionRequest) (connection_response *auth.AuthenticateGameConnectionResponse, err error) {
	var (
		login_username     string
		login_account      models.Account
		account_found      bool
		game_account_found bool
	)

	login_username = strings.ToUpper(auth_game_connection.ClientAccount)

	account_found, err = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(login_username))).Get(&login_account)
	if err != nil {
		log.Warn(err)
		return nil, status.Errorf(codes.Internal, "database error")
	}

	if !account_found {
		return nil, status.Errorf(codes.Unauthenticated, "account not found")
	}

	hash := sha1.New()
	hash.Write([]byte(login_username))
	hash.Write([]byte{0, 0, 0, 0})
	hash.Write(auth_game_connection.LocalChallenge)
	hash.Write(auth_game_connection.ServerChallenge)
	hash.Write(login_account.SessionKey)

	digest := hash.Sum(nil)

	if subtle.ConstantTimeCompare(digest, auth_game_connection.Digest) == 0 {
		err = status.Errorf(codes.Unauthenticated, "connection authentication failure: phony connection attempt to '%s' from address %s", login_username, auth_game_connection.GetClientIpAddress())
		return
	}

	var game_account models.GameAccount
	game_account_found, err = provider.home_db.Table("GameAccount").
		Where(
			query.Eq("Owner", login_account.ID),
			query.Eq("Active", true),
		).Get(&game_account)
	if err != nil {
		log.Warn(err)
		return nil, status.Errorf(codes.Internal, "database error")
	}
	if !game_account_found {
		return nil, status.Errorf(codes.Unauthenticated, "no active game account!")
	}

	connection_response = &auth.AuthenticateGameConnectionResponse{
		Tier:          login_account.Tier,
		SessionKey:    login_account.SessionKey,
		AccountId:     login_account.ID,
		GameAccountId: game_account.ID,
		Locale:        login_account.Locale,
	}

	return
}
