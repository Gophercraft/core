package grunt

import (
	"fmt"

	"github.com/Gophercraft/log"
)

func (session *Session) handle_realm_list(request *RealmList_Client) (err error) {
	if session.state != state_authorized {
		err = fmt.Errorf("grunt: cannot get realm list without authorization")
		return
	}
	var realm_list RealmList_Server
	if realm_list.Realms, err = session.server.service_provider.GetRealmList(
		session.logon_info.AccountName,
		session.logon_info.Build,
	); err != nil {
		log.Warn("failed to get realmlist", err)
		return
	}

	log.Println("realm list requested: ", request.Request)

	if err = WriteMessageType(session, RealmList); err != nil {
		return
	}
	if err = WriteRealmList_Server(session, session.logon_info.Build, &realm_list); err != nil {
		return
	}
	err = session.Send()

	return
}
