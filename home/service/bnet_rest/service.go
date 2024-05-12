package bnet_rest

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/log"
)

// type ServiceProvider interface {
// 	// Associate a SRP login session with the username + remote address
// 	StoreSRPSession(remote_address, username, hash_function, platform_id string, srp_session *srp.Session) (err error)
// 	// Look for a SRP login session associated with this remote address + username, and make it inaccessible
// 	CheckoutSRPSession(remote_address, username string) (srp_session *srp.Session, platform_id string, err error)
// 	// Create a valid authentication ticket associated with the username.
// 	CreateTicket(username string) (ticket string, err error)
// 	// Attempt to look up an identity hash for this particular hash function.
// 	// return an error if the account wasn't found
// 	LookupAccountIdentity(remote_address, username string) (identity, salt []byte, err error)
// 	// Less secure login. Password is sent in cleartext over JSON
// 	Login(remote_address, username, password string) (ticket string, valid bool, err error)

// 	UpdateAccountData(username, platform_id string, session_key []byte) (err error)
// }

type ServiceProvider interface {
	IsValidSessionID(id string) bool
	NewSessionCookie() (session_cookie *http.Cookie, err error)
	GetLoginFormInputs() (login_form *login.FormInputs)
	Login(user_info *UserInfo, login_form *login.LoginForm) (result *login.LoginResult, err error)
	LoginAuthenticator(user_info *UserInfo, login_form *login.LoginForm) (result *login.LoginResult, err error)
	LoginSRP(user_info *UserInfo, login_form *login.LoginForm) (srp_challenge *login.SrpLoginChallenge, err error)
}

type Service struct {
	provider  ServiceProvider
	config    *ServiceConfig
	serve_mux *http.ServeMux
	listener  net.Listener
}

type ServiceConfig struct {
	Address string
	TLS     *tls.Config
}

func New(service_config *ServiceConfig, service_provider ServiceProvider) (service *Service) {
	if service_config.TLS == nil {
		service_config.TLS = util.GetTrinityTLSConfig()
	}
	service_config.TLS.NextProtos = []string{"http/1.1", "h2"}
	service = new(Service)
	service.provider = service_provider
	service.config = service_config
	service.serve_mux = http.NewServeMux()
	service.serve_mux.HandleFunc("GET /bnet/login", service.handle_get_login)
	service.serve_mux.HandleFunc("POST /bnet/login", service.handle_post_login)
	service.serve_mux.HandleFunc("POST /bnet/login/srp", service.handle_post_login_srp)
	service.serve_mux.HandleFunc("POST /bnet/login/authenticator", service.handle_post_login_authenticator)
	return
}

func (service *Service) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("[bnet_rest]", r.UserAgent(), r.RemoteAddr, r.Proto, r.Method, r.URL.String())
	service.serve_mux.ServeHTTP(rw, r)
}

func (service *Service) ID() config.HomeServiceID {
	return config.BNetRESTService
}

func (service *Service) Start() (err error) {
	service.listener, err = listen(service.config.Address, service.config.TLS)
	if err != nil {
		return
	}

	log.Println("bound bnet rest listener", service.listener.Addr())

	go func() {
		if err := http.Serve(service.listener, service); err != nil {
			log.Warn("home/service/bnet_rest: http.Serve()", err)
		}
	}()

	return
}

func (service *Service) Stop() (err error) {
	err = service.listener.Close()
	return
}
