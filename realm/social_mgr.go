package realm

import (
	"sort"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/chat"
	"github.com/Gophercraft/core/packet/social"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

func (ws *Server) GetFriendStatus(id guid.GUID) social.Status {
	_, err := ws.GetSessionByGUID(id)
	if err != nil {
		return social.FriendOffline
	}

	return social.FriendOnline
}

// ---- Begin Vanilla

func (s *Session) SendFriendList() {
	var contacts []models.Contact
	if err := s.DB().Where("player = ?", s.GUID().Counter()).Where("friended = 1").Find(&contacts); err != nil {
		panic(err)
	}
	list := &social.PlayerList{
		Type:    packet.SMSG_FRIEND_LIST,
		Players: make([]guid.GUID, len(contacts)),
	}
	for i, v := range contacts {
		list.Players[i] = guid.RealmSpecific(guid.Player, s.Server.RealmID(), v.Friend)
	}

	s.Send(list)
}

func (s *Session) SendIgnoreList() {
	var contacts []models.Contact
	if err := s.DB().Where("player = ?", s.GUID().Counter()).Where("ignored = 1").Find(&contacts); err != nil {
		panic(err)
	}
	list := &social.PlayerList{
		Type:    packet.SMSG_IGNORE_LIST,
		Players: make([]guid.GUID, len(contacts)),
	}
	for i, v := range contacts {
		list.Players[i] = guid.RealmSpecific(guid.Player, s.Server.RealmID(), v.Friend)
	}

	s.Send(list)
}

// ---- End Vanilla

// Starting in protocol 8606, the friends list and the ignore list are merged into a single packet.
func (s *Session) SendSocialList() {
	if s.Build().RemovedIn(vsn.V2_4_3) {
		s.SendFriendList()
		s.SendIgnoreList()
		return
	}

	p := &social.ContactList{
		Flags: social.DefaultContactFlags,
	}

	var contacts []models.Contact

	s.DB().Where("player = ?", s.GUID().Counter()).Find(&contacts)

	for _, mcontact := range contacts {
		var scontact social.Contact

		if mcontact.Friended {
			scontact.Flags |= social.FlagFriend
		}
		if mcontact.Muted {
			scontact.Flags |= social.FlagMuted
		}
		if mcontact.Ignored {
			scontact.Flags |= social.FlagIgnored
		}

		p.Contacts = append(p.Contacts, scontact)
	}

	s.Send(p)
}

func (s *Session) SendDanceMoves() {
	s.Send(&chat.LearnedDanceMoves{})
}

func (s *Session) HandleWho(wr *social.WhoRequest) {
	w := &social.Who{}
	var usernames []string

	s.Server.GuardPlayerList.Lock()
	for k := range s.Server.PlayerList {
		usernames = append(usernames, k)
	}

	sort.Strings(usernames)
	whoMatches := make([]social.WhoMatch, len(usernames))

	for _i, user := range usernames {
		playerPtr := s.Server.PlayerList[user]

		whoMatches[_i] = social.WhoMatch{
			PlayerName: user,
			GuildName:  "",
			Level:      uint32(playerPtr.GetLevel()),
			Class:      uint32(playerPtr.GetPlayerClass()),
			Race:       uint32(playerPtr.GetPlayerRace()),
			ZoneID:     playerPtr.ZoneID,
		}
	}

	w.WhoMatches = whoMatches
	s.Server.GuardPlayerList.Unlock()

	s.Send(w)
}

func (s *Session) SendFriendStatus(result social.FriendResult, id guid.GUID, note string, status social.Status, area, level, class uint32) {
	// data := packet.NewWorldPacket(packet.SMSG_FRIEND_STATUS)
	// data.WriteByte(result)
	// id.EncodeUnpacked(s.Build(), data)

	// switch result {
	// case social.FriendAddedOnline, social.FriendOnline:
	// 	data.WriteByte(uint8(status))
	// 	data.WriteUint32(area)
	// 	data.WriteUint32(level)
	// 	data.WriteUint32(class)
	// default:
	// }

	// s.SendPacket(data)

	var fs social.FriendStatus
	fs.Result = result
	fs.ID = id
	fs.Note = note
	fs.Status = status
	fs.AreaID = area
	fs.Level = level
	fs.Class = class

	s.Send(&fs)
}

func (s *Session) GetContact(friend guid.GUID) *models.Contact {
	var contact models.Contact
	found, err := s.DB().Where("player = ?", s.GUID().Counter()).Where("friend = ?", friend.Counter()).Get(&contact)
	if err != nil {
		panic(err)
	}
	if !found {
		contact.Player = s.GUID().Counter()
		contact.Friend = friend.Counter()
		if _, err := s.DB().Insert(&contact); err != nil {
			panic(err)
		}
		return &contact
	}
	return &contact
}

func (s *Session) HandleFriendAdd(add *social.Add) {
	name := add.Name
	id, err := s.Server.GetGUIDByPlayerName(name)
	if err != nil {
		s.SendFriendStatus(social.FriendNotFound, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	if id == s.GUID() {
		s.SendFriendStatus(social.FriendSelf, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	contact := s.GetContact(id)
	if contact.Friended {
		s.SendFriendStatus(social.FriendAlready, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	contact.Friended = true
	s.DB().Where("player = ?", s.GUID().Counter()).Where("friend = ?", id.Counter()).Cols("friended").Update(contact)

	status := s.Server.GetFriendStatus(id)

	var area, level, class uint32
	if status&social.FriendOnline != 0 {
		session, err := s.Server.GetSessionByGUID(id)
		if err != nil {
			s.SendFriendStatus(social.FriendDBError, guid.Nil, "", 0, 0, 0, 0)
			return
		}

		area = session.ZoneID
		level = uint32(session.GetLevel())
		class = uint32(session.GetPlayerClass())
		s.SendFriendStatus(social.FriendAddedOnline, id, contact.Note, status, area, level, class)
	} else {
		s.SendFriendStatus(social.FriendAddedOffline, id, contact.Note, status, area, level, class)
	}
}

func (s *Session) HandleFriendDelete(del *social.Delete) {
	id := del.ID

	if id == s.GUID() {
		s.SendFriendStatus(social.FriendSelf, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	var contact models.Contact
	found, err := s.DB().Where("player = ?", s.GUID().Counter()).Where("friend = ?", id.Counter()).Get(&contact)
	if err != nil {
		panic(err)
	}

	if !found {
		s.SendFriendStatus(social.FriendNotFound, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	if contact.Friended == false {
		s.SendFriendStatus(social.FriendAlready, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	contact.Friended = false
	s.DB().Where("player = ?", s.GUID().Counter()).Where("friend = ?", id.Counter()).Cols("friended").Update(&contact)
	s.SendFriendStatus(social.FriendRemoved, id, "", 0, 0, 0, 0)
}

func (s *Session) HandleFriendListRequest() {
	s.SendFriendList()
	s.SendIgnoreList()
}

func (s *Session) HandleSocialListRequest() {
	s.SendSocialList()
}

func (s *Session) BroadcastStatus(status social.Status) {
	player := s.GUID().Counter()

	var contacts []models.Contact
	s.DB().Where("friend = ?", player).Where("friended = 1").Find(&contacts)

	for _, contact := range contacts {
		id := guid.RealmSpecific(guid.Player, s.Server.RealmID(), contact.Player)
		sess, _ := s.Server.GetSessionByGUID(id)
		if sess != nil {
			var result social.FriendResult
			var area, level, class uint32
			if status&social.FriendOnline != 0 {
				result = social.FriendOnline
				area = s.ZoneID
				level = uint32(s.GetLevel())
				class = uint32(s.GetPlayerClass())
			} else {
				result = social.FriendOffline
			}

			sess.SendFriendStatus(result, s.GUID(), contact.Note, status, area, level, class)
		}
	}
}

func (s *Session) HandleIgnoreAdd(add *social.Add) {
	name := add.Name

	var char models.Character
	found, err := s.DB().Where("name = ?", name).Get(&char)
	if err != nil {
		panic(err)
	}

	if !found {
		s.SendFriendStatus(social.FriendIgnoreNotFound, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	id := guid.RealmSpecific(guid.Player, s.Server.RealmID(), char.ID)

	if id == s.GUID() {
		s.SendFriendStatus(social.FriendIgnoreSelf, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	contact := s.GetContact(id)
	if contact.Ignored == true {
		s.SendFriendStatus(social.FriendIgnoreAlready, id, "", 0, 0, 0, 0)
		return
	}

	contact.Ignored = true

	if _, err := s.DB().Where("player = ?", s.GUID().Counter()).Where("friend = ?", id.Counter()).Cols("ignored").Update(contact); err != nil {
		panic(err)
	}

	s.SendFriendStatus(social.FriendIgnoreAdded, id, "", 0, 0, 0, 0)
}

func (s *Session) HandleIgnoreDelete(del *social.Delete) {
	id := del.ID

	if id == s.GUID() {
		s.SendFriendStatus(social.FriendIgnoreSelf, guid.Nil, "", 0, 0, 0, 0)
		return
	}

	var contact models.Contact
	found, err := s.DB().Where("player = ?", s.GUID().Counter()).Where("friend = ?", id.Counter()).Get(&contact)
	if err != nil {
		panic(err)
	}

	if !found {
		s.SendFriendStatus(social.FriendIgnoreNotFound, id, "", 0, 0, 0, 0)
		return
	}

	if contact.Ignored == false {
		s.SendFriendStatus(social.FriendIgnoreAlready, id, "", 0, 0, 0, 0)
		return
	}

	contact.Ignored = false
	s.DB().Where("player = ?", s.GUID().Counter()).Where("friend = ?", id.Counter()).Cols("ignored").Update(&contact)
	s.SendFriendStatus(social.FriendIgnoreRemoved, id, "", 0, 0, 0, 0)
}
