package home

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/vsn"
	"github.com/dchest/captcha"
	"github.com/gorilla/mux"
)

func (r *Request) ScanJSON(v interface{}) error {
	b, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func (c *Server) GetAuthStatus(r *Request) {
	getAuthStatusRequest := struct {
		Token string `json:"token"`
	}{}

	type getAuthStatusResponse struct {
		Valid   bool        `json:"valid"`
		Account string      `json:"account"`
		Tier    rpcnet.Tier `json:"tier"`
	}

	err := r.ScanJSON(&getAuthStatusRequest)
	if err != nil {
		r.Respond(http.StatusBadRequest, "could not read json")
		return
	}

	var wt models.WebToken
	found, _ := c.DB.Where("token = ?", getAuthStatusRequest.Token).Get(&wt)
	if !found {
		r.Encode(http.StatusOK, getAuthStatusResponse{
			Valid: false,
		})
		return
	}

	if time.Since(wt.Expiry) > 0 {
		c.DB.Delete(&wt)
		r.Encode(http.StatusOK, getAuthStatusResponse{
			Valid: false,
		})
		return
	}

	var acc models.Account
	found, err = c.DB.Where("id = ?", wt.Account).Get(&acc)
	if !found {
		panic(err)
	}

	resp := getAuthStatusResponse{
		Valid:   true,
		Account: strings.ToLower(acc.Username),
		Tier:    acc.Tier,
	}

	r.Encode(http.StatusOK, resp)
}

func (c *Server) ResetAccount(user, pass string, tier rpcnet.Tier) error {
	return login.RegisterAccount(c.DB, user, pass, tier)
}

func (c *Server) NewCaptcha(r *Request) {
	cp := captcha.New()
	r.Encode(http.StatusOK, CaptchaResponse{
		Status:    http.StatusOK,
		CaptchaID: cp,
	})
}

func (c *Server) UserExists(r *Request) {
	t := strings.ToUpper(r.Vars["username"])
	var acc []models.Account
	err := c.DB.Where("username = ?", t).Find(&acc)
	if err != nil {
		r.Respond(http.StatusInternalServerError, "Internal server error")
		return
	}

	r.Encode(http.StatusOK, UserExistsResponse{
		Status:     http.StatusOK,
		UserExists: len(acc) == 1,
	})
}

func (c *Server) Register(r *Request) {
	var rr models.WebRegisterRequest

	if err := r.ScanJSON(&rr); err != nil {
		r.Respond(http.StatusBadRequest, "malformed json")
		return
	}

	if !captcha.VerifyString(rr.CaptchaID, rr.CaptchaSolution) {
		r.Encode(http.StatusOK, models.WebRegisterResponse{
			Error:        "Captcha failed.",
			ResetCaptcha: true,
		})
		return
	}

	if err := login.RegisterAccount(c.DB, rr.Username, rr.Password, rpcnet.Tier_NormalPlayer); err != nil {
		r.Encode(http.StatusOK, models.WebRegisterResponse{
			Error: err.Error(),
		})
		return
	}

	_, token, err := c.HandleWebLogin(rr.Username, rr.Password)
	if err != nil {
		r.Encode(http.StatusOK, models.WebRegisterResponse{
			Error: err.Error(),
		})
		return
	}

	r.Encode(http.StatusOK, models.WebRegisterResponse{
		WebToken: token,
	})
}

func (c *Server) PublishRealmInfo(r models.Realm) uint64 {
	r.LastUpdated = time.Now()
	var rinf []models.Realm
	err := c.DB.Where("name = ?", r.Name).Find(&rinf)
	if err != nil {
		panic(err)
	}
	if len(rinf) == 0 {
		if _, err := c.DB.Insert(&r); err != nil {
			panic(err)
		}
	} else {
		if _, err := c.DB.AllCols().Update(&r); err != nil {
			panic(err)
		}
	}

	return r.ID
}

func (c *Server) RealmState() []models.Realm {
	var r []models.Realm
	c.DB.Find(&r)
	return r
}

func (c *Server) RealmList(r *Request) {
	r.Encode(http.StatusOK, map[string]interface{}{
		"status":  200,
		"listing": c.RealmState(),
	})
}

func (c *Server) Intercept(required int, fn RequestHandler) *Interceptor {
	return &Interceptor{required, c, fn}
}

type Interceptor struct {
	requiredLevel int
	core          *Server
	fn            RequestHandler
}

func (s *Interceptor) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	s.fn(&Request{
		Writer:  rw,
		Request: req,
		Vars:    mux.Vars(req),
	})
}

func (c *Server) InfoHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`<p><a href="https://github.com/Gophercraft/core">Gophercraft ` + vsn.GophercraftVersion.String() + `<a/></p>`))
	})
	return r
}
