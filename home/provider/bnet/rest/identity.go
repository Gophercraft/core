package rest

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/Gophercraft/core/bnet/pb/login"
)

const max_identity_age = 10 * time.Minute

type identity_session struct {
	guard                sync.Mutex
	timestamp            time.Time
	account_ID           uint64
	identified           bool
	authenticated        bool
	pending_login_result *login.LoginResult
}

func generate_random_identity_session_id() string {
	var random_bytes [16]byte
	if _, err := io.ReadFull(rand.Reader, random_bytes[:]); err != nil {
		panic(err)
	}
	return hex.EncodeToString(random_bytes[:])
}

func (provider *Provider) web_domain() string {
	host, _, err := net.SplitHostPort(provider.config.Endpoint)
	if err != nil {
		panic(err)
	}
	return host
}

func (provider *Provider) make_new_identity() (identity_cookie *http.Cookie, err error) {
	session := new(identity_session)
	session.timestamp = time.Now()

	id := generate_random_identity_session_id()

	provider.guard_identity_sessions.Lock()

	provider.identity_sessions[id] = session

	provider.guard_identity_sessions.Unlock()
	identity_cookie = new(http.Cookie)

	identity_cookie.Path = "/bnet"
	identity_cookie.Name = "JSESSIONID"
	identity_cookie.Value = id
	identity_cookie.Domain = provider.web_domain()
	identity_cookie.Secure = true
	identity_cookie.HttpOnly = true
	identity_cookie.SameSite = http.SameSiteNoneMode
	// identity_cookie.MaxAge = -1

	return
}

func (provider *Provider) checkout_identity(id string) (checked_out_session *identity_session) {
	begin := time.Now()

	provider.guard_identity_sessions.Lock()

	// delete old sessions
	for active_session_key, active_session := range provider.identity_sessions {
		if active_session_key == id {
			// we found our session!
			checked_out_session = active_session
		} else {
			if begin.Sub(active_session.timestamp) > max_identity_age {
				delete(provider.identity_sessions, active_session_key)
			}
		}
	}
	provider.guard_identity_sessions.Unlock()

	if checked_out_session == nil {
		return
	}

	if begin.Sub(checked_out_session.timestamp) > max_identity_age {
		checked_out_session = nil
		delete(provider.identity_sessions, id)
	}

	return
}

func (provider *Provider) NewSessionCookie() (identity_cookie *http.Cookie, err error) {
	identity_cookie, err = provider.make_new_identity()
	return
}

func (provider *Provider) IsValidSessionID(id string) bool {
	return provider.checkout_identity(id) != nil
}
