package auth

import (
	"context"

	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/home/provider/core/util"
	"github.com/Gophercraft/core/version"
)

var plaintext_auth_removed_update = version.V1_12_1

func (provider *auth_provider) AuthenticateGameConnection(ctx context.Context, auth_game_connection *auth.AuthenticateGameConnectionRequest) (connection_response *auth.AuthenticateGameConnectionResponse, err error) {
	// verify peer identity
	if err = util.CheckPeerIdentity(provider.home_db, auth_game_connection.RealmId, ctx); err != nil {
		return
	}

	client_build := version.Build(auth_game_connection.ClientBuild)

	switch {
	// Alpha uses plaintext authentication
	case client_build.RemovedIn(plaintext_auth_removed_update):
		return provider.authenticate_game_connection_alpha(auth_game_connection)
	case client_build.RemovedIn(version.NewAuthSystem):
		return provider.authenticate_game_connection_legacy(auth_game_connection)
	default:
		return provider.authenticate_game_connection_modern(auth_game_connection)
	}
}
