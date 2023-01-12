package home

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/log"
)

func (c *Server) GetAccountID(user string) (uint64, error) {
	var acc []models.Account

	err := c.DB.Where("username = ?", user).Find(&acc)
	if err != nil {
		return 0, err
	}

	if len(acc) == 0 {
		return 0, fmt.Errorf("gcore: empty set")
	}

	return acc[0].ID, nil
}

func (c *Server) AccountID(user string) uint64 {
	id, err := c.GetAccountID(user)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (c *Server) GetAccount(user string) (*models.Account, []models.GameAccount, error) {
	var acc models.Account
	found, err := c.DB.Where("username = ?", strings.ToUpper(user)).Get(&acc)
	if err != nil {
		return nil, nil, err
	}

	if !found {
		return nil, nil, fmt.Errorf("account %s not found", user)
	}

	var gameAccs []models.GameAccount
	c.DB.Where("owner = ?", acc.ID).Find(&gameAccs)

	return &acc, gameAccs, nil
}
