package rest

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"
	"time"

	"github.com/Gophercraft/core/home/login"
)

func generate_random_ticket() string {
	var random_bytes [20]byte
	if _, err := io.ReadFull(rand.Reader, random_bytes[:]); err != nil {
		panic(err)
	}

	return "TC-" + strings.ToUpper(hex.EncodeToString(random_bytes[:]))
}

func (provider *Provider) create_ticket(account_id uint64) (ticket string, err error) {
	// var (
	// 	found   bool
	// 	account models.Account
	// )
	// account_table := provider.home_db.Table("Account")

	// found, err = account_table.Where(query.Eq("ID", account_id)).Get(&account)
	// if err != nil {
	// 	return
	// }

	// if !found {
	// 	err = fmt.Errorf("home/provider/bnet/rest: account '%d' does not exist", account_id)
	// 	return
	// }

	// new_ticket := generate_random_ticket()
	// var ticket_record models.LoginTicket
	// ticket_record.Ticket = new_ticket
	// ticket_record.Account = account.ID
	// ticket_record.Expiry = time.Now().Add(12 * time.Hour)

	// if err = provider.home_db.Table("LoginTicket").Insert(&ticket_record); err != nil {
	// 	return
	// }

	// ticket = new_ticket
	return login.CreateTicket(provider.home_db, account_id, time.Now().Add(12*time.Hour))
}
