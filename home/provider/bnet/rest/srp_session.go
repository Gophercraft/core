package rest

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/service/bnet_rest"
	"github.com/Gophercraft/phylactery/database/query"
)

const max_srp_period = 60 * time.Second

// Bnet login session
type srp_session struct {
	remote_address string
	platform_id    string
	srp_session    *srp.Session
	timestamp      time.Time
}

func (provider *Provider) fix_host(remote_address string) string {
	host, _, err := net.SplitHostPort(remote_address)
	if err != nil {
		return "unknown"
	}

	return host
}

// Find a user log in session
func (provider *Provider) checkout_srp_session(user_info *bnet_rest.UserInfo, username string) (account_id uint64, user_srp_session *srp.Session, platform_id string, err error) {
	var (
		account models.Account
		found   bool
	)

	begin := time.Now()

	account_table := provider.home_db.Table("Account")

	found, err = account_table.Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/bnet/rest: account '%s' does not exist", username)
		return
	}

	key := fmt.Sprintf("%d-%s", account.ID, provider.fix_host(user_info.Address))
	log.Println("Looking up SRP session ", key)

	provider.guard_srp_sessions.Lock()

	var checked_out_session *srp_session

	// delete old sessions
	for active_session_key, active_session := range provider.srp_sessions {
		if active_session_key == key {
			// we found our session!
			checked_out_session = active_session
		} else {
			if begin.Sub(active_session.timestamp) > max_srp_period {
				delete(provider.srp_sessions, active_session_key)
			}
		}
	}

	if checked_out_session != nil {
		// remove srp session from database
		delete(provider.srp_sessions, key)
	}

	provider.guard_srp_sessions.Unlock()

	if !found {
		err = fmt.Errorf("home/provider/bnet/rest: account does not have an active srp session associated")
		return
	}
	account_id = account.ID
	user_srp_session = checked_out_session.srp_session
	platform_id = checked_out_session.platform_id

	return
}

func (provider *Provider) store_srp_session(user_info *bnet_rest.UserInfo, username, hash_function, platform_id string, srp_params *srp.Session) (err error) {
	var (
		account models.Account
		found   bool
	)

	found, err = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = fmt.Errorf("home/provider/bnet/rest: account '%s' does not exist", username)
		return
	}

	session := new(srp_session)
	key := fmt.Sprintf("%d-%s", account.ID, provider.fix_host(user_info.Address))

	log.Println("Saving SRP session as", key)

	session.remote_address = user_info.Address
	session.platform_id = platform_id
	session.srp_session = srp_params

	provider.guard_srp_sessions.Lock()
	provider.srp_sessions[key] = session
	provider.guard_srp_sessions.Unlock()

	return
}
