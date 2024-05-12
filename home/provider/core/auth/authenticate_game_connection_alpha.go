package auth

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func sanitize_line(line string) string {
	return strings.TrimSpace(line)
}

func parse_alpha_session_file(file string) (ticket string, err error) {
	ticket = sanitize_line(file)
	return
}

func (provider *auth_provider) authenticate_game_connection_alpha(auth_game_connection *auth.AuthenticateGameConnectionRequest) (connection_response *auth.AuthenticateGameConnectionResponse, err error) {
	ses_file := string(auth_game_connection.LocalChallenge)

	var (
		login_ticket       string
		login_account      models.Account
		account_id         uint64
		account_found      bool
		game_account_found bool
		valid              bool
	)

	login_ticket, err = parse_alpha_session_file(ses_file)
	if err != nil {
		return nil, err
	}

	valid, account_id, err = login.VerifyTicket(provider.home_db, login_ticket)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, fmt.Errorf("home/provider/core/auth: user logon ticket is not valid")
	}

	// credentials := login.Credentials(login_username, login_password)

	account_found, err = provider.home_db.Table("Account").Where(query.Eq("ID", account_id)).Get(&login_account)
	if err != nil {
		log.Warn(err)
		return nil, status.Errorf(codes.Internal, "database error")
	}

	if !account_found {
		return nil, status.Errorf(codes.Unauthenticated, "account not found")
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
		AccountId:     login_account.ID,
		GameAccountId: game_account.ID,
		Locale:        login_account.Locale,
	}

	return
}
