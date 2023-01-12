package wizard

import (
	"context"
	"fmt"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/rpcnet"
)

// talks to Home server over GRPC to enlist a World
type RemoteHomeserverRegistrar struct {
	WorldConfigurator *WorldConfigurator
	Address           string
	ValidFingerprint  string

	Token  string
	Client rpcnet.HomeServiceClient
}

func (rhr *RemoteHomeserverRegistrar) Begin(wc *WorldConfigurator, address string) error {
	rhr.Address = address
	rhr.WorldConfigurator = wc
	return nil
}

func (rhr *RemoteHomeserverRegistrar) MustAuth() bool {
	return true
}

func (rhr *RemoteHomeserverRegistrar) Auth(username, password string) error {
	cc, err := rpcnet.DialConn(rhr.Address, rhr.ValidFingerprint, nil)
	if err != nil {
		return err
	}

	rhr.Client = rpcnet.NewHomeServiceClient(cc)

	sInfo, err := rhr.Client.SignIn(context.Background(), &rpcnet.Credentials{
		Account:  username,
		Password: password,
	})

	if err != nil {
		return err
	}

	if sInfo.Status != rpcnet.Status_OK {
		return fmt.Errorf("wizard: sign in failed: %s", sInfo.Status)
	}

	rhr.Token = sInfo.WebToken

	return nil
}

func (rhr *RemoteHomeserverRegistrar) Check() (*ValidationCheck, error) {
	finger, err := rpcnet.FingerprintServer(rhr.Address)
	if err != nil {
		return nil, err
	}

	return &ValidationCheck{
		Question: fmt.Sprintf("Is this fingerprint received from %s correct?", rhr.Address),
		Object:   finger,
	}, nil
}

func (rhr *RemoteHomeserverRegistrar) Confirm(check *ValidationCheck) error {
	rhr.ValidFingerprint = check.Object
	return nil
}

func (rhr *RemoteHomeserverRegistrar) Enlist() error {
	var world *config.World = rhr.WorldConfigurator.Config

	world.HomeServer = rhr.Address
	world.HomeServerFingerprint = rhr.ValidFingerprint

	sInfo, err := rhr.Client.EnlistRealm(context.Background(), &rpcnet.EnlistRealmRequest{
		WebToken:         rhr.Token,
		RealmName:        world.RealmName,
		RealmFingerprint: world.Fingerprint(),
	})

	if err != nil {
		return err
	}

	if sInfo.Status != rpcnet.Status_OK {
		return fmt.Errorf("wizard: enlist realm failed: %s", sInfo.Status)
	}

	world.RealmID = sInfo.RealmID
	world.HomeServerFingerprint = rhr.ValidFingerprint

	return nil
}
