package rest

import (
	"sync"

	"github.com/Gophercraft/phylactery/database"
)

type Provider struct {
	config                  *ProviderConfig
	home_db                 *database.Container
	guard_srp_sessions      sync.Mutex
	guard_identity_sessions sync.Mutex
	srp_sessions            map[string]*srp_session
	identity_sessions       map[string]*identity_session
}

type ProviderConfig struct {
	Endpoint string // host:port
}

func New(config *ProviderConfig, home_db *database.Container) (provider *Provider) {
	provider = new(Provider)
	provider.config = config
	provider.home_db = home_db
	provider.srp_sessions = make(map[string]*srp_session)
	provider.identity_sessions = make(map[string]*identity_session)

	return
}

// func (provider *Provider) Login(remote_address, username, password string) (ticket string, valid bool, err error) {
// 	var (
// 		account models.Account
// 		found   bool
// 	)

// 	found, err = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
// 	if err != nil {
// 		return
// 	}

// 	if !found {
// 		valid = false
// 		return
// 	}

// 	valid = login.Verify(username, password, account.Identity_Bcrypt)
// 	if !valid {
// 		return
// 	}

// 	new_ticket := generate_random_ticket()
// 	var ticket_record models.LoginTicket
// 	ticket_record.Account = account.ID
// 	ticket_record.Ticket = new_ticket
// 	ticket_record.Expiry = time.Now().Add(12 * time.Hour)

// 	if err = provider.home_db.Table("LoginTicket").Insert(&ticket_record); err != nil {
// 		valid = false
// 		return
// 	}

// 	ticket = new_ticket
// 	return
// }

// func (provider *Provider) LookupAccountIdentity(remote_address, username string) (identity_hash, salt []byte, err error) {
// 	var (
// 		account models.Account
// 		found   bool
// 	)
// 	found, err = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
// 	if err != nil {
// 		return
// 	}
// 	if !found {
// 		err = fmt.Errorf("home/provider/bnet/rest: account '%s' does not exist", username)
// 		return
// 	}

// 	return account.Identity_PBKDF2_SHA_512, account.Identity_PBKDF2_Salt, nil
// }

// func (provider *Provider) fix_host(remote_address string) string {
// 	host, _, err := net.SplitHostPort(remote_address)
// 	if err != nil {
// 		return "unknown"
// 	}

// 	return host
// }

// func (provider *Provider) StoreSRPSession(remote_address, username, hash_function, platform_id string, srp_session *srp.Session) (err error) {
// 	var (
// 		account models.Account
// 		found   bool
// 	)

// 	found, err = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
// 	if err != nil {
// 		return
// 	}

// 	if !found {
// 		err = fmt.Errorf("home/provider/bnet/rest: account '%s' does not exist", username)
// 		return
// 	}

// 	session := new(login_session)
// 	key := fmt.Sprintf("%d-%s", account.ID, provider.fix_host(remote_address))

// 	log.Println("Saving SRP session as", key)

// 	session.remote_address = remote_address
// 	session.platform_id = platform_id
// 	session.srp_session = srp_session

// 	provider.guard_sessions.Lock()
// 	provider.active_sessions[key] = session
// 	provider.guard_sessions.Unlock()

// 	return
// }

// func (provider *Provider) CheckoutSRPSession(remote_address, username string) (srp_session *srp.Session, platform_id string, err error) {
// 	var (
// 		account models.Account
// 		found   bool
// 	)

// 	account_table := provider.home_db.Table("Account")

// 	found, err = account_table.Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
// 	if err != nil {
// 		return
// 	}

// 	if !found {
// 		err = fmt.Errorf("home/provider/bnet/rest: account '%s' does not exist", username)
// 		return
// 	}

// 	key := fmt.Sprintf("%d-%s", account.ID, provider.fix_host(remote_address))
// 	log.Println("Looking up SRP session ", key)

// 	provider.guard_sessions.Lock()

// 	session, found := provider.active_sessions[key]
// 	if found {
// 		// remove srp session from database
// 		delete(provider.active_sessions, key)
// 	}

// 	provider.guard_sessions.Unlock()

// 	if !found {
// 		err = fmt.Errorf("home/provider/bnet/rest: account does not have an active srp session associated")
// 		return
// 	}

// 	srp_session = session.srp_session
// 	platform_id = session.platform_id

// 	return
// }

// func (provider *Provider) UpdateAccountData(username, platform_id string, session_key []byte) (err error) {
// 	var (
// 		found   bool
// 		account models.Account
// 	)
// 	account_table := provider.home_db.Table("Account")

// 	found, err = account_table.Where(query.Eq("Username", strings.ToLower(username))).Get(&account)
// 	if err != nil {
// 		return
// 	}

// 	if !found {
// 		err = fmt.Errorf("home/provider/bnet/rest: account '%s' does not exist", username)
// 		return
// 	}

// 	account.Platform = platform_id
// 	account.SessionKey = session_key

// 	_, err = account_table.Where(query.Eq("ID", account.ID)).
// 		Columns(
// 			"Platform",
// 			"SessionKey",
// 		).Update(&account)

// 	return
// }
