package models

import "time"

// Bnet server
type LoginTicket struct {
	Account string
	Ticket  string
	Expiry  time.Time
}

// Website
type WebToken struct {
	Token   string `xorm:"'token' pk"`
	Account uint64
	Expiry  time.Time
}

type SessionKey struct {
	ID uint64 `xorm:"'id' pk"`
	K  []byte
}
