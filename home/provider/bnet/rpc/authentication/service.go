package authentication

import (
	"context"
	"crypto/rand"
	"io"

	"github.com/Gophercraft/core/bnet"
	"github.com/Gophercraft/core/bnet/pb/bgs/protocol"
	authentication_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/authentication/v1"
	challenge_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/challenge/v1"
	"github.com/Gophercraft/core/bnet/rpc/codes"
	"github.com/Gophercraft/core/bnet/rpc/status"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"

	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
)

type context_key uint8

const (
	details_context_key context_key = iota
)

type Details struct {
	AccountID   uint64
	AccountTier auth.AccountTier
	Build       version.Build
}

func GetDetails(ctx context.Context) (details *Details) {
	value := ctx.Value(details_context_key)
	if value == nil {
		return
	}
	details = value.(*Details)
	return
}

type Service struct {
	config  *ServiceConfig
	home_db *database.Container

	authentication_v1.UnimplementedAuthenticationServiceServer
}

type ServiceConfig struct {
	// host:port form address of the public URL of the RESTful bnet login server
	RESTLoginEndpoint string
}

func New(service_config *ServiceConfig, home_db *database.Container) (service *Service) {
	service = new(Service)
	service.config = service_config
	service.home_db = home_db
	return
}

func (service *Service) lookup_web_credentials(web_credentials []byte) (found bool, account_id uint64, err error) {
	web_credential_string := string(web_credentials)

	var login_ticket models.LoginTicket
	found, err = service.home_db.Table("LoginTicket").Where(query.Eq("Ticket", web_credential_string)).Get(&login_ticket)
	if err != nil {
		return
	}

	account_id = login_ticket.Account
	return
}

func (service *Service) lookup_account(id uint64) (found bool, account *models.Account, err error) {
	account = new(models.Account)
	found, err = service.home_db.Table("Account").Where(query.Eq("ID", id)).Get(account)
	return
}

func (service *Service) lookup_game_accounts(id uint64) (game_accounts []models.GameAccount, err error) {
	err = service.home_db.Table("GameAccount").Where(query.Eq("Owner", id)).Find(&game_accounts)
	return
}

func (service *Service) verify_web_credentials(ctx context.Context, web_credentials []byte) (*protocol.NoData, error) {
	found, account_id, err := login.VerifyTicket(service.home_db, string(web_credentials))
	if err != nil {
		return nil, err
	}

	var logon_result authentication_v1.LogonResult
	if !found {

		// Request OnLogonComplete from client
		util.Set(&logon_result.ErrorCode, uint32(codes.ERROR_DENIED))

		// invalid := uint32(ERROR_DENIED)
		// conn.AuthenticationListener_Request_OnLogonComplete(&v1.LogonResult{
		// 	ErrorCode: &invalid,
		// })

		// conn.SendResponseCode(token, ERROR_OK)
		// conn.Close() // We want to show the invalid code, however NOT closing the connection causes the Login button to be greyed out.
		// "you have been disconnected"
	} else {
		var session_key [64]byte
		if _, err = io.ReadFull(rand.Reader, session_key[:]); err != nil {
			panic(err)
		}

		found, account, err := service.lookup_account(account_id)
		if err != nil {
			return nil, err
		}

		if !found {
			util.Set(&logon_result.ErrorCode, uint32(codes.ERROR_DENIED))
		} else {
			// Associate authentication details with this connection if we haven't done so already
			if GetDetails(ctx) == nil {
				peer_connection := bnet.GetConnection(ctx)
				if peer_connection != nil {
					peer_connection.Set(details_context_key, &Details{
						AccountID:   account_id,
						AccountTier: account.Tier,
					})
				}
			}

			util.Set(&logon_result.ErrorCode, uint32(codes.ERROR_OK))

			// set account entity ID
			var account_entity_id protocol.EntityId
			util.Set(&account_entity_id.High, uint64(0x100000000000000))
			util.Set(&account_entity_id.Low, uint64(account_id))
			logon_result.SessionKey = session_key[:]
			logon_result.AccountId = &account_entity_id

			game_accounts, err := service.lookup_game_accounts(account_id)
			if err != nil {
				return nil, err
			}

			// set game account IDs
			logon_result.GameAccountId = make([]*protocol.EntityId, len(game_accounts))
			for i := range game_accounts {
				game_account_entity_id := new(protocol.EntityId)
				util.Set(&game_account_entity_id.High, uint64(0x200000200576F57))
				util.Set(&game_account_entity_id.Low, uint64(game_accounts[i].ID))
				logon_result.GameAccountId[i] = game_account_entity_id
			}
		}
	}

	go func() {
		// log.Println("OnLogonCOmplete calling")
		// Setup response to client
		peer_connection := bnet.GetConnection(ctx)
		if peer_connection != nil {
			authentication_listener_client := authentication_v1.NewAuthenticationListenerClient(peer_connection)

			// Request OnLogonComplete from client
			if _, err := authentication_listener_client.OnLogonComplete(context.TODO(), &logon_result); err != nil {
				log.Warn(err)
			}
		}
	}()
	return nil, nil
}

func (service *Service) Logon(ctx context.Context, logon_request *authentication_v1.LogonRequest) (*protocol.NoData, error) {
	// verify program
	if logon_request.GetProgram() != "WoW" {
		return nil, status.Errorf(codes.ERROR_BAD_PROGRAM, "bad program")
	}

	if logon_request.CachedWebCredentials != nil {
		_, err := service.verify_web_credentials(ctx, logon_request.GetCachedWebCredentials())
		if err != nil {
			return nil, err
		}

		details := GetDetails(ctx)
		if details != nil {
			details.Build = version.Build(logon_request.GetApplicationVersion())
		}
		return nil, nil
	}

	// conn.locale = args.GetLocale()
	// conn.platform = args.GetPlatform()

	go func() {
		log.Println("Trying to send request")
		connection := bnet.GetConnection(ctx)
		challenge_listener_client := challenge_v1.NewChallengeListenerClient(connection)
		log.Println("created client")

		var request challenge_v1.ChallengeExternalRequest
		util.Set(&request.PayloadType, "web_auth_url")
		request.Payload = []byte("https://" + service.config.RESTLoginEndpoint + "/bnet/login")
		log.Println("created payload")

		// log.Println("Sending request")
		// log.Dump("request", request)
		log.Println("calling external challenge")
		_, err := challenge_listener_client.OnExternalChallenge(context.TODO(), &request)
		log.Println("called external challenge")
		if err != nil {
			log.Warn(err)
		} else {
			log.Println("Sent request")
		}
	}()

	return nil, nil
}

func (service *Service) VerifyWebCredentials(ctx context.Context, verify_web_credentials_request *authentication_v1.VerifyWebCredentialsRequest) (*protocol.NoData, error) {
	return service.verify_web_credentials(ctx, verify_web_credentials_request.GetWebCredentials())
}
