package auth_home_rpcnet

import (
	"context"

	game_auth "github.com/Gophercraft/core/game/server/service/auth"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	pb_auth "github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/i18n"
)

type provider struct {
	client auth.AuthServiceClient
}

func (provider *provider) AuthenticateGameConnection(request *game_auth.AuthenticateGameConnectionRequest) (response *game_auth.AuthenticateGameConnectionResponse, err error) {
	// This connects back to the home server, checking to see if this client is actually a registered user.
	// The auth server will perform calculations, and if valid, return to us a session key.
	verify_world_query := &pb_auth.AuthenticateGameConnectionRequest{
		RealmId:         request.RealmID,
		ClientBuild:     uint32(request.Build),
		ClientAccount:   request.Account,
		ClientIpAddress: request.IPAddress,
		Digest:          request.Digest,
		ServerChallenge: request.Challenge,
		LocalChallenge:  request.LocalChallenge,
	}

	verify_world_response, err := provider.client.AuthenticateGameConnection(context.Background(), verify_world_query)
	if err != nil {
		return
	}

	response = &game_auth.AuthenticateGameConnectionResponse{
		Tier:        verify_world_response.Tier,
		SessionKey:  verify_world_response.SessionKey,
		Account:     verify_world_response.AccountId,
		GameAccount: verify_world_response.GameAccountId,
	}
	response.Locale, err = i18n.LocaleFromString(verify_world_response.Locale)
	return
}

func New(auth_service_client auth.AuthServiceClient) *provider {
	return &provider{
		client: auth_service_client,
	}
}
