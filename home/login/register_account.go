package login

import (
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
	"xorm.io/xorm"
)

func RegisterAccount(db *xorm.Engine, user, pass string, tier rpcnet.Tier) error {
	if user == "" {
		return ErrUsernameCannotBeEmpty
	}

	var acc models.Account
	found, err := db.Where("username = ?", user).Get(&acc)
	if err != nil {
		return err
	}

	if err := SetAccount(&acc, user, pass, tier); err != nil {
		return err
	}

	if found {
		if _, err := db.Where("id = ?", acc.ID).AllCols().Update(&acc); err != nil {
			return err
		}
		return nil
	} else {
		if _, err := db.Insert(&acc); err != nil {
			return err
		}
		_, err = db.Insert(&models.GameAccount{
			Name:   "Main",
			Active: true,
			Owner:  acc.ID,
		})
		return err
	}
}
