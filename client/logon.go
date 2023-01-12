package client

import (
	"github.com/Gophercraft/core/auth"
	"github.com/Gophercraft/core/vsn"
)

const NewBNet vsn.Build = vsn.V8_3_0

func (c *Client) LoginGrunt(username,password string) error {
	grunt, err := auth.Login(c.Config.Build, c.Config.Realmlist,c.Config.Username, c.Config.Password)
	if err != nil {
		return err
	}

	rls, err := grunt.GetRealmlist()
}

func (c *Client) Login() error {
	switch {
	case c.Config.Build < NewBattleNet:
		return c.LoginGrunt()
	case c.Config.Build >= NewBattleNet:
		return c.LoginBattleNet()
	}
}
