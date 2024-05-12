package auth

import (
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/version"
)

type AuthenticateGameConnectionRequest struct {
	RealmID        uint64
	Build          version.Build
	Account        string
	GameAccount    string
	IPAddress      string
	Digest         []byte
	Challenge      []byte
	LocalChallenge []byte
}

type AuthenticateGameConnectionResponse struct {
	Tier        auth.AccountTier
	SessionKey  []byte
	Account     uint64
	GameAccount uint64
	Locale      i18n.Locale
}

type ServiceProvider interface {
	AuthenticateGameConnection(request *AuthenticateGameConnectionRequest) (*AuthenticateGameConnectionResponse, error)
}
