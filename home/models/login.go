package models

import (
	"time"

	"github.com/Gophercraft/core/home/protocol/pb/auth"
)

type LoginMethod uint8

const (
	Login_Web LoginMethod = iota
	Login_Core
)

// A login ticket proves that a user is who they say they are - this ticket can be sent to realm servers safely without risking the integrity of the account
type LoginTicket struct {
	Ticket  string `database:"1:index,exclusive"`
	Account uint64
	Expiry  time.Time
}

// Website and Core GRPC protocol
type WebToken struct {
	Token     string `database:"1:index,exclusive"`
	Status    auth.WebTokenStatus
	Account   uint64
	Method    LoginMethod
	Address   string
	UserAgent string
	Expiry    time.Time
}
