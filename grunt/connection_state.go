package grunt

import "github.com/Gophercraft/log"

const (
	state_unauthorized = iota
	state_logon_challenging
	state_reconnect_challenging
	state_authorized
)

func (session *Session) set_state(new_state int) {
	old_state := session.state
	session.state = new_state

	if old_state != new_state {
		log.Println("Session state changed from", old_state, "to", new_state)
	}
}
