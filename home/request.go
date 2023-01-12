package home

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	Vars      map[string]string
	AuthLevel int
}

type RequestHandler func(*Request)

func (r *Request) Encode(status int, v interface{}) {
	enc := json.NewEncoder(r.Writer)

	if r.Request.URL.Query().Get("fmt") != "" {
		enc.SetIndent("", " ")
	}
	r.Writer.Header().Set("Content-Type", "application/json; charset: utf-8")
	r.Writer.WriteHeader(status)
	enc.Encode(v)
}

type CaptchaResponse struct {
	Status    int    `json:"status"`
	CaptchaID string `json:"captchaID"`
}

type GenericResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type UserExistsResponse struct {
	Status     int  `json:"status"`
	UserExists bool `json:"exists"`
}

func (r *Request) Respond(status int, err string) {
	r.Encode(status, &GenericResponse{
		Status: status,
		Error:  err,
	})
}
