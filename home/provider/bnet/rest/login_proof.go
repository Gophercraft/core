package rest

import (
	"fmt"
	"strings"

	"github.com/Gophercraft/core/bnet/pb/login"
	pb_login "github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/crypto/srp"
	home_login "github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/service/bnet_rest"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *Provider) check_user_info(user_info *bnet_rest.UserInfo) (err error) {
	return nil
}

func account_2fa(account *models.Account) bool {
	if account.Authenticator {
		return true
	}

	// TODO: email 2fa

	return false
}

func (provider *Provider) handle_srp_login_proof(user_info *bnet_rest.UserInfo, login_form *pb_login.LoginForm) (login_result *pb_login.LoginResult, db_account *models.Account, err error) {
	// Read login proof fields
	account_name := get_login_form_input(login_form, "account_name")
	public_A := srp.NewIntFromHexString(get_login_form_input(login_form, "public_A"))
	client_evidence_M1 := srp.NewIntFromHexString(get_login_form_input(login_form, "client_evidence_M1"))

	var (
		account_id  uint64
		srp_session *srp.Session
		platform_id string
	)
	var server_M2 *srp.Int
	account_id, srp_session, platform_id, err = provider.checkout_srp_session(user_info, account_name)
	if err != nil {
		return
	}

	session_key, valid := srp.VerifyClientEvidence(
		srp_session.Hash,
		public_A,
		srp_session.PrivateB,
		srp_session.PublicB,
		srp_session.LargeSafePrime,
		srp_session.Verifier,
		client_evidence_M1,
	)

	if !valid {
		err = fmt.Errorf("invalid client evidence")
		return
	}

	var (
		account models.Account
		found   bool
	)
	found, err = provider.home_db.Table("Account").Where(query.Eq("ID", account_id)).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = fmt.Errorf("account not found")
		return
	}
	db_account = &account

	server_M2 = srp.CalculateServerEvidence(
		srp_session.Hash,
		public_A,
		client_evidence_M1,
		session_key,
	)

	if err = provider.update_account_data(account_name, platform_id, session_key.Bytes()); err != nil {
		return
	}

	login_result = new(login.LoginResult)
	util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
	util.Set(&login_result.ServerEvidence_M2, server_M2.HexString())
	return
}

func (provider *Provider) handle_login_plaintext(user_info *bnet_rest.UserInfo, login_form *pb_login.LoginForm) (login_result *pb_login.LoginResult, db_account *models.Account, err error) {
	// Proceed with plaintext credentials
	account_name := get_login_form_input(login_form, "account_name")
	password := get_login_form_input(login_form, "password")

	var (
		account    models.Account
		found      bool
		valid      bool
		new_ticket string
	)

	found, err = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(account_name))).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = fmt.Errorf("account not found")
		return
	}

	valid = home_login.Verify(account_name, password, account.Identity_Bcrypt)
	if !valid {
		err = fmt.Errorf("incorrect username/password combination")
		return
	}

	if account_2fa(&account) {
		err = fmt.Errorf("2fa not yet implemented")
		return
	}

	new_ticket, err = provider.create_ticket(account.ID)
	if err != nil {
		return
	}

	util.Set(&login_result.LoginTicket, new_ticket)
	db_account = &account
	return
}

func (provider *Provider) Login(user_info *bnet_rest.UserInfo, login_form *pb_login.LoginForm) (login_result *pb_login.LoginResult, err error) {
	// Block unwanted users
	if err = provider.check_user_info(user_info); err != nil {
		return
	}

	// Detect if this is an SRP proof login
	use_srp := get_login_form_input(login_form, "use_srp")

	var (
		account    *models.Account
		new_ticket string
	)

	if use_srp == "true" {
		login_result, account, err = provider.handle_srp_login_proof(user_info, login_form)
	} else {
		login_result, account, err = provider.handle_login_plaintext(user_info, login_form)
	}

	if err != nil {
		return
	}

	var session *identity_session
	session = provider.checkout_identity(user_info.SessionID)
	if session == nil {
		err = fmt.Errorf("no session")
		return
	}

	session.guard.Lock()
	defer session.guard.Unlock()
	session.account_ID = account.ID
	session.identified = true

	// a second verification step is needed.
	if account_2fa(account) {
		session.pending_login_result = login_result
		login_result = new(pb_login.LoginResult)

		session.authenticated = false
		util.Set(&login_result.AuthenticationState, login.AuthenticationState_AUTHENTICATOR)
		authenticator_url := fmt.Sprintf("https://%s/bnet/login/authenticator", provider.config.Endpoint)
		util.Set(&login_result.NextUrl, authenticator_url)
	} else {
		new_ticket, err = provider.create_ticket(account.ID)
		if err != nil {
			return
		}
		session.authenticated = true

		util.Set(&login_result.AuthenticationState, login.AuthenticationState_DONE)
		util.Set(&login_result.LoginTicket, new_ticket)
	}

	return
}
