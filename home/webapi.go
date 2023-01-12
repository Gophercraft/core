package home

import (
	"net/http"

	"github.com/Gophercraft/core/home/webapp"
	"github.com/dchest/captcha"
	"github.com/gorilla/mux"
)

type AccessLevel int

type WebAPI struct {
	Router *mux.Router
}

type APIFunc func(*Request)

func (s *Server) API(prefix string) *WebAPI {
	return &WebAPI{
		Router: s.Router.PathPrefix(prefix).Subrouter(),
	}
}

func (wap *WebAPI) On(route string, handler APIFunc) {
	wap.Router.HandleFunc(route, func(rw http.ResponseWriter, r *http.Request) {
		req := &Request{}
		req.Writer = rw
		req.Request = r
		req.Vars = mux.Vars(r)
		handler(req)
	})
}

func (s *Server) WebServer() http.Handler {
	wap := &WebAPI{}
	wap.Router = mux.NewRouter()
	wap.Router.HandleFunc("/", s.Home)
	wap.Router.HandleFunc("/signUp", s.SignUp)
	// wap.Router.PathPrefix("/a").

	r := mux.NewRouter()
	r.HandleFunc("/", s.Home)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.FS(webapp.Assets))))

	v2 := r.PathPrefix("/v2/").Subrouter()
	v2.Handle("/signIn", s.Intercept(0, s.HandleSignIn))
	v2.Handle("/realmList", s.Intercept(0, s.RealmList))
	v2.Handle("/getAuthStatus", s.Intercept(0, s.GetAuthStatus))
	v2.Handle("/newCaptcha", s.Intercept(0, s.NewCaptcha))
	v2.Handle("/userExists/{username}", s.Intercept(0, s.UserExists))
	v2.Handle("/register", s.Intercept(0, s.Register))
	// v2.Handle("/enlist_server_key")
	v2.PathPrefix("/captcha/").Handler(captcha.Server(captcha.StdWidth, captcha.StdHeight))

	// admin/realm RPC functions

	// r.PathPrefix("/").Handler(http.FileServer(http.Dir(os.Getenv("GOPATH") + "src/github.com/Gophercraft/core/home/webapp/public/")))
	return r
}
