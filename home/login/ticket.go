package login

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
)

func ticket_is_expired(login_ticket *models.LoginTicket) bool {
	now := time.Now()
	return login_ticket.Expiry.Before(now)
}

func expire_ticket(home_db *database.Container, login_ticket *models.LoginTicket) {
	home_db.Table("LoginTicket").Where(query.Eq("Token", login_ticket.Ticket)).Delete()
}

func VerifyTicket(home_db *database.Container, ticket string) (valid bool, account_id uint64, err error) {
	var (
		login_ticket models.LoginTicket
		account      models.Account
	)
	valid, err = home_db.Table("LoginTicket").Where(query.Eq("Ticket", ticket)).Get(&login_ticket)
	if err != nil {
		return
	}

	if !valid {
		return
	}

	if ticket_is_expired(&login_ticket) {
		expire_ticket(home_db, &login_ticket)
		return
	}

	account_id = login_ticket.Account

	valid, err = home_db.Table("Account").Where(query.Eq("ID", login_ticket.Account)).Get(&account)
	if err != nil {
		return
	}

	if !valid {
		return
	}

	// Banned accounts are not allowed to use the account service
	if account.Banned {
		expire_ticket(home_db, &login_ticket)
		err = fmt.Errorf("home/provider/web: ticket is expired or invalid")
		return
	}

	// Suspended accounts may continue to use the web API
	if account.Suspended {
		// Check if suspension has expired and update account
		now := time.Now()
		suspension_expired := account.UnsuspendAt.Compare(now) == -1

		if suspension_expired {
			account.Suspended = false
			account.UnsuspendAt = time.Time{}
			if _, err = home_db.Table("Account").Where(query.Eq("ID", account.ID)).Columns("Suspended", "UnsuspendAt").Update(&account); err != nil {
				panic(err)
			}
		}
	}

	return
}

func generate_random_ticket() string {
	var random_bytes [20]byte
	if _, err := io.ReadFull(rand.Reader, random_bytes[:]); err != nil {
		panic(err)
	}

	return "TC-" + strings.ToUpper(hex.EncodeToString(random_bytes[:]))
}

func CreateTicket(home_db *database.Container, account_id uint64, expiry time.Time) (ticket string, err error) {
	var (
		found   bool
		account models.Account
	)
	account_table := home_db.Table("Account")

	found, err = account_table.Where(query.Eq("ID", account_id)).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/bnet/rest: account '%d' does not exist", account_id)
		return
	}

	new_ticket := generate_random_ticket()
	var ticket_record models.LoginTicket
	ticket_record.Ticket = new_ticket
	ticket_record.Account = account.ID
	ticket_record.Expiry = time.Now().Add(12 * time.Hour)

	if err = home_db.Table("LoginTicket").Insert(&ticket_record); err != nil {
		return
	}

	ticket = new_ticket
	return
}
