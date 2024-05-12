package login

import (
	"fmt"
	"strings"
	"time"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
)

// Replaces an existing account or creates a new one
func RegisterAccount(db *database.Container, can_reset bool, registrator_tier auth.AccountTier, email, name, password string, tier auth.AccountTier) error {
	if name == "" {
		return ErrUsernameCannotBeEmpty
	}

	var account models.Account
	found, err := db.Table("Account").
		Where(query.Eq("Username", strings.ToLower(name))).
		Get(&account)
	if err != nil {
		return err
	}

	if tier != auth.AccountTier_NORMAL && registrator_tier != auth.AccountTier_ADMIN {
		return fmt.Errorf("home/login: you cannot register an account with an elevated tier if you are not an admin")
	}

	if err := SetAccount(&account, email, name, password, tier); err != nil {
		return err
	}

	if found {
		if !can_reset {
			return fmt.Errorf("home/login: account '%s' already exists", name)
		}

		if registrator_tier < auth.AccountTier_MODERATOR {
			return fmt.Errorf("home/login: you must be moderator or higher to register an account")
		}

		if registrator_tier <= account.Tier && registrator_tier != auth.AccountTier_ADMIN {
			return fmt.Errorf("home/login: you cannot reset an account with a higher tier than you")
		}

		if _, err := db.Table("Account").
			Where(query.Eq("ID", account.ID)).
			// Update only the columns that got changed
			Columns(
				"Email",
				"Identity_SHA_1",
				"Identity_SHA_256",
				"Identity_SHA_512",
				"Identity_PBKDF2_SHA_512",
				"Identity_PBKDF2_Salt",
				"Identity_Bcrypt",
				"Tier",
			).
			Update(&account); err != nil {
			return err
		}
		return nil
	} else {
		account.CreatedAt = time.Now()
		if err = db.Table("Account").
			Insert(&account); err != nil {
			return err
		}
		err = db.Table("GameAccount").
			Insert(&models.GameAccount{
				Name:   "Main",
				Active: true,
				Owner:  account.ID,
			})
		return err
	}
}
