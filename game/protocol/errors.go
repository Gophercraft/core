package protocol

import "errors"

var (
	ErrNoServerConfig = errors.New("No server configuration provided")
	ErrNoConnection   = errors.New("No attempt to connect was made")
)
