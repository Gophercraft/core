package home

import (
	"html/template"
	"net/http"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/webapp"
)

type navElement struct {
	Active bool
	Link   string
	Name   string
}

type pageBase struct {
	Title       string
	Brand       string
	NavElements []*navElement
	PageBody    template.HTML
}

func (s *Server) newPageBase(activeName string) *pageBase {
	navBar := []*navElement{
		{false, "/", "Home"},
		{false, "/armory", "Armory"},
		{false, "/account", "Account"},
	}

	for _, v := range navBar {
		if v.Name == activeName {
			v.Active = true
		}
	}

	brandName := "Gophercraft"

	return &pageBase{
		Title:       brandName + " " + activeName,
		Brand:       brandName,
		NavElements: navBar,
	}
}

func (s *Server) getTemplate(named string) *template.Template {
	data, err := webapp.ReadEmbedded(webapp.Templates, named)
	if err != nil {
		panic(err)
	}

	tpl, err := template.New(named).Parse(string(data))
	if err != nil {
		panic(err)
	}

	return tpl
}

func (s *Server) loadHtml(name string) template.HTML {
	data, err := webapp.ReadEmbedded(webapp.Templates, name)
	if err != nil {
		panic(err)
	}

	return template.HTML(data)
}

func (s *Server) Home(rw http.ResponseWriter, r *http.Request) {
	pg := s.newPageBase("Home")

	base := s.getTemplate("base.html")

	pg.PageBody = s.loadHtml("home.html")

	if err := base.Execute(rw, pg); err != nil {
		panic(err)
	}
}

func (s *Server) SignUp(rw http.ResponseWriter, r *http.Request) {
	pg := s.newPageBase("Sign Up")

	base := s.getTemplate("base.html")

	pg.PageBody = s.loadHtml("signUp.html")

	if err := base.Execute(rw, pg); err != nil {
		panic(err)
	}
}

func (s *Server) HandleSignIn(r *Request) {
	var signInRequest models.WebSignInRequest
	var signInResponse models.WebSignInResponse

	if err := r.ScanJSON(&signInRequest); err != nil {
		r.Respond(http.StatusBadRequest, "malformed json")
		return
	}

	_, token, err := s.HandleWebLogin(signInRequest.Username, signInRequest.Password)
	if err != nil {
		r.Respond(http.StatusBadRequest, err.Error())
		return
	}

	signInResponse.WebToken = token
	r.Encode(http.StatusOK, &signInResponse)
}
