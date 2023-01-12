package models

import "github.com/Gophercraft/core/i18n"

type Broadcast struct {
	Message i18n.Text
	Weight  int
}
