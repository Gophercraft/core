package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"strings"

	"github.com/Gophercraft/core/crypto/hashutil"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authCheckSeed        = []byte{0xC5, 0xC6, 0x98, 0x95, 0x76, 0x3F, 0x1D, 0xCD, 0xB6, 0xA1, 0x37, 0x28, 0xB3, 0x12, 0xFF, 0x8A}
	sessionKeySeed       = []byte{0x58, 0xCB, 0xCF, 0x40, 0xFE, 0x2E, 0xCE, 0xA6, 0x5A, 0x90, 0xB8, 0x01, 0x68, 0x6C, 0x28, 0x0B}
	continuedSessionSeed = []byte{0x16, 0xAD, 0x0C, 0xD4, 0x46, 0xF9, 0x4F, 0xB2, 0xEF, 0x7D, 0xEA, 0x2A, 0x17, 0x66, 0x4D, 0x2F}
	encryptionKeySeed    = []byte{0xE9, 0x75, 0x3C, 0x50, 0x90, 0x93, 0x61, 0xDA, 0x3B, 0x07, 0xEE, 0xFA, 0xFF, 0x9D, 0x41, 0xB8}
)

func (provider *auth_provider) authenticate_game_connection_modern(auth_game_connection *auth.AuthenticateGameConnectionRequest) (connection_response *auth.AuthenticateGameConnectionResponse, err error) {
	var (
		login_username     string
		login_account      models.Account
		account_found      bool
		game_account_found bool
		session_key        []byte
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

	client_build := version.Build(auth_game_connection.ClientBuild)

	client_build_info := client_build.BuildInfo()
	if client_build_info == nil {
		err = status.Errorf(codes.Unauthenticated, "build info for %s not found", &client_build)
		return
	}
	if len(client_build_info.Win64AuthSeed) == 0 || len(client_build_info.Mac64AuthSeed) == 0 {
		err = status.Errorf(codes.Unauthenticated, "auth seed for %s not found", client_build)
		return
	}

	localChallenge := auth_game_connection.LocalChallenge
	serverChallenge := auth_game_connection.ServerChallenge
	digest := auth_game_connection.Digest

	sessionKeyHash := sha256.New()
	skl, _ := sessionKeyHash.Write(login_account.SessionKey)
	if skl != 64 {
		panic("invalid key length")
	}

	// log.Dump("req.Digest", req.Digest)
	log.Dump("localChallenge", localChallenge)
	log.Dump("serverChallenge", serverChallenge)
	log.Dump("buildInfo.Win64AuthSeed", client_build_info.Win64AuthSeed)

	if login_account.Platform == "Wn64" {
		sessionKeyHash.Write(client_build_info.Win64AuthSeed)
	} else if login_account.Platform == "Mc64" {
		sessionKeyHash.Write(client_build_info.Mac64AuthSeed)
	} else {
		err = status.Errorf(codes.InvalidArgument, "invalid user platform %s", login_account.Platform)
		return
	}

	digestKeyHash := sessionKeyHash.Sum(nil)

	hmc := hmac.New(sha256.New, digestKeyHash)
	hmc.Write(localChallenge)  //localChallenge
	hmc.Write(serverChallenge) //serverChallenge
	hmc.Write(authCheckSeed)
	authCheckHash := hmc.Sum(nil)

	if subtle.ConstantTimeCompare(authCheckHash[:24], digest[:24]) == 0 {
		err = status.Errorf(codes.Unauthenticated, "connection authentication failure: phony connection attempt to '%s' from address %s", login_username, auth_game_connection.GetClientIpAddress())
		return
	}

	keyDataDigest := sha256.Sum256(login_account.SessionKey)

	sessionKeyHmac := hmac.New(sha256.New, keyDataDigest[:])
	sessionKeyHmac.Write(serverChallenge)
	sessionKeyHmac.Write(localChallenge)
	sessionKeyHmac.Write(sessionKeySeed)

	session_key = make([]byte, 40)
	skg := hashutil.NewSessionKeyGenerator(sha256.New, sessionKeyHmac.Sum(nil))
	skg.Read(session_key)

	log.Dump("sessionKey", session_key)

	encryptKeyGen := hmac.New(sha256.New, session_key)
	encryptKeyGen.Write(localChallenge)
	encryptKeyGen.Write(serverChallenge)
	encryptKeyGen.Write(encryptionKeySeed)

	encryptKeyHash := encryptKeyGen.Sum(nil)

	session_key = encryptKeyHash[:16]

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
		SessionKey:    session_key,
		AccountId:     login_account.ID,
		GameAccountId: game_account.ID,
		Locale:        login_account.Locale,
	}

	return
}
