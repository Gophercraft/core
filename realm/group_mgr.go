package realm

import (
	"sync"

	"github.com/Gophercraft/core/packet/chat"
	"github.com/Gophercraft/core/packet/party"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
)

func (s *Session) SendPartyResult(operation party.Operation, memberName string, result party.Result) {
	s.Send(&party.CommandResult{
		Operation: operation,
		Member:    memberName,
		Result:    result,
	})
}

type Group struct {
	sync.Mutex
	GroupType     party.GroupType
	Server        *Server
	Leader        guid.GUID
	Members       []guid.GUID
	LootMethod    party.LootMethod
	LootThreshold models.ItemQuality
}

func (g *Group) Disband() {
	// Delete party membership from database
	g.SetLeader(guid.Nil)

	for _, member := range g.Members {
		sess, err := g.Server.GetSessionByGUID(member)
		if err != nil {
			continue
		}

		sess.SendGroupDestroyed()
		sess.Group = nil
		sess.GroupInvite = guid.Nil
	}
}

func (g *Group) Empty() bool {
	g.Lock()
	ln := len(g.Members)
	g.Unlock()

	return ln == 0
}

func (g *Group) Add(session *Session) {
	g.Members = append(g.Members, session.GUID())

	for _, member := range g.Members {
		g.Server.DB.Where("id = ?", member.Counter()).Cols("leader").Update(&models.Character{
			Leader: g.Leader.Counter(),
		})
	}

	g.UpdateList()
}

func (g *Group) RemoveMember(id guid.GUID) {
	for i, gid := range g.Members {
		if gid == id {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
			g.UpdateList()
			return
		}
	}
}

func (s *Session) HasYouIgnored(g guid.GUID) bool {
	if s.IsAdmin() {
		return false
	}

	return false
}

func (s *Session) SendGroupInvite(from string) {
	// p := packet.NewWorldPacket(packet.SMSG_GROUP_INVITE)
	// p.WriteCString(from)
	// s.SendPacket(p)
	s.Send(&party.Invitation{
		From: from,
	})
}

func (s *Session) HandleGroupInvite(ir *party.InviteRequest) {
	playerName := ir.To

	receiver, err := s.Server.GetSessionByPlayerName(playerName)
	if err != nil {
		s.SendPartyResult(party.Invite, playerName, party.BadPlayerName)
		return
	}

	if receiver.GUID() == s.GUID() {
		s.SendPartyResult(party.Invite, playerName, party.BadPlayerName)
		return
	}

	if s.HasYouIgnored(receiver.GUID()) {
		s.SendPartyResult(party.Invite, playerName, party.IgnoringYou)
		return
	}

	// Disallow cross-factional parties unless you are a GM OR the server has explicitly allowed them.
	if !s.IsGM() && !s.Server.BoolVar("PVP.CrossFactionGroups") && s.Map().Hostile(s, receiver) {
		s.SendPartyResult(party.Invite, playerName, party.WrongFaction)
		return
	}

	if s.Group != nil {
		s.Group.Lock()
		if s.GUID() != s.Group.Leader {
			s.SendPartyResult(party.Invite, playerName, party.NotLeader)
			return
		}
		s.Group.Unlock()
	} else {
		s.GroupInvite = guid.Nil
		if s.Group != nil {
			s.Group.RemoveMember(s.GUID())
		}

		s.Group = new(Group)
		s.Group.Server = s.Server
		s.Group.Leader = s.GUID()
		s.Group.Members = append(s.Group.Members, s.GUID())
	}

	s.SendPartyResult(party.Invite, playerName, party.OK)

	receiver.GroupInvite = s.GUID()
	receiver.SendGroupInvite(s.PlayerName())
}

func (s *Session) HandleGroupDecline() {
	if s.GroupInvite == guid.Nil {
		return
	}

	inviter, err := s.Server.GetSessionByGUID(s.GroupInvite)
	if err != nil {
		return
	}

	if inviter.Group == nil {
		return
	}

	if inviter.Group.Empty() {
		inviter.Group = nil
	}

	inviter.Send(&party.Declination{
		From: s.PlayerName(),
	})
}

func (s *Session) HandleGroupAccept() {
	if s.GroupInvite == guid.Nil {
		return
	}

	player, err := s.Server.GetSessionByGUID(s.GroupInvite)
	if err != nil {
		return
	}

	group := player.Group
	if group == nil {
		return
	}

	s.Group = group

	group.Add(s)
}

func (s *Session) SendSetLeader(leaderName string) {
	s.Send(&party.SetLeader{
		leaderName,
	})
}

func (s *Session) SendGroupList() {
	if s.Group == nil {
		return
	}

	s.Group.Lock()
	defer s.Group.Unlock()

	if s.GUID() == s.Group.Leader {
		s.SendSetLeader(s.PlayerName())
	}

	list := &party.GroupList{
		GroupType: s.Group.GroupType,
	}

	for _, member := range s.Group.Members {
		if member == s.GUID() {
			continue
		}

		var lmem party.GroupMember
		lmem.GUID = member

		str, err := s.Server.GetPlayerNameByGUID(member)
		if err != nil {
			log.Warn(err)
			str = "???"
			lmem.Status |= party.MemberOffline
		} else {
			lmem.Status |= party.MemberOnline
		}

		lmem.Name = str

		if member == s.Group.Leader {
			s.SendSetLeader(str)
		}

		list.Members = append(list.Members, lmem)
	}

	list.LootThreshold = s.Group.LootThreshold
	list.LootMethod = s.Group.LootMethod

	s.Send(list)
}

func (g *Group) SetLeader(id guid.GUID) {
	char := models.Character{
		Leader: id.Counter(),
	}
	if _, err := g.Server.DB.Where("leader = ?", g.Leader.Counter()).Cols("leader").Update(&char); err != nil {
		panic(err)
	}
}

func (g *Group) UpdateList() {
	for _, member := range g.Members {
		sess, err := g.Server.GetSessionByGUID(member)
		if err == nil {
			sess.SendGroupList()
		} else {
			log.Warn(err)
		}
	}
}

func (s *Session) SendGroupDestroyed() {
	p := packet.NewWorldPacket(packet.SMSG_GROUP_DESTROYED)
	s.SendPacket(p)
}

func (s *Session) LeaveGroup() {
	char := models.Character{
		Leader: 0,
	}
	s.DB().Where("id = ?", s.GUID().Counter()).Cols("leader").Update(&char)
	if s.Group != nil {
		// if len(s.Group.Members) > 2 {
		var newLeaderGUID guid.GUID
		// Set next player to leader
		for _, member := range s.Group.Members {
			if member != s.GUID() {
				newLeaderGUID = member
				break
			}
		}

		if newLeaderGUID != guid.Nil {
			s.Group.SetLeader(newLeaderGUID)
		}

		s.Group.RemoveMember(s.GUID())
		// } else {
		// 	s.Group.Disband()
		// }

		s.Group.UpdateList()
		s.Group = nil
	}

	s.SendGroupDestroyed()
}

func (s *Session) InitGroup() {
	if s.Char.Leader != 0 {
		var members []models.Character
		s.DB().Where("leader = ?", s.Char.Leader).Find(&members)

		// See if group already exists.
		for _, member := range members {
			sess, err := s.Server.GetSessionByGUID(guid.RealmSpecific(guid.Player, s.Server.RealmID(), member.ID))
			if err == nil {
				if sess.Group != nil {
					s.Group = sess.Group
					break
				}
			}
		}

		if s.Group == nil {
			// No one appears to be online, instantiate a new Group object.
			s.Group = &Group{
				Server: s.Server,
				// TODO: what should happen if the leader character was deleted?
				Leader: guid.RealmSpecific(guid.Player, s.Server.RealmID(), s.Char.Leader),
			}

			for _, memb := range members {
				s.Group.Members = append(s.Group.Members, guid.RealmSpecific(guid.Player, s.Server.RealmID(), memb.ID))
			}
		}

		s.Group.UpdateList()
	}
}

func (s *Session) HandleGroupDisband() {
	// will disband if only two players are currently there.
	s.LeaveGroup()
}

const statsMaskAll = party.GroupUpdateNone |
	party.GroupUpdateStatus |
	party.GroupUpdateCurrentHealth |
	party.GroupUpdateMaxHealth |
	party.GroupUpdatePowerType |
	party.GroupUpdateCurrentPower |
	party.GroupUpdateMaxPower |
	party.GroupUpdateLevel |
	party.GroupUpdateZone |
	party.GroupUpdatePosition | party.GroupUpdateAuras

func (s *Session) HandleRequestPartyMemberStats(rq *party.RequestMemberStats) {
	id := rq.Member

	if s.Group == nil {
		return
	}

	s.Group.Lock()

	found := false

	for _, member := range s.Group.Members {
		if member == id {
			found = true
			break
		}
	}

	if !found {
		return
	}

	s.Group.Unlock()

	player, err := s.Server.GetSessionByGUID(id)
	if err != nil {
		s.Send(&party.MemberStats{
			ChangeMask: party.GroupUpdateStatus,
			Status:     party.MemberOffline,
		})
		return
	}

	var mask uint32
	mask = statsMaskAll

	if player.Pet() != nil {
		mask |= party.GroupUpdatePetGUID |
			party.GroupUpdatePetName |
			party.GroupUpdatePetModelID |
			party.GroupUpdatePetCurrentHP |
			party.GroupUpdatePetMaxHP |
			party.GroupUpdatePetPowerType |
			party.GroupUpdatePetCurrentPower |
			party.GroupUpdatePetMaxPower |
			party.GroupUpdatePetAuras |
			party.GroupUpdatePet
	}

	s.SendPartyMemberStats(mask, player)
}

func (s *Session) PartyMemberAuraState() []party.MemberAuraState {
	return nil
}

func (s *Session) SendPartyMemberStats(mask uint32, player *Session) {
	if s == nil || s.State() == Ended {
		s.Send(&party.MemberStats{
			ChangeMask: mask,
			Member:     s.GUID(),
			Status:     party.MemberOffline,
		})
		return
	}

	s.Send(&party.MemberStats{
		ChangeMask: mask,
		Member:     s.GUID(),
		Status:     party.MemberOnline,
		PowerType:  s.GetPowerType(),
		Power:      uint16(s.Power()),
		Level:      uint16(s.Char.Level),
		ZoneID:     uint16(s.ZoneID),
		PositionX:  uint16(s.Position().X),
		PositionY:  uint16(s.Position().Y),
		Auras:      s.PartyMemberAuraState(),
	})
}

func (s *Session) HandlePartyMessage(party *chat.ClientMessage) {
	if s.Group == nil {
		return
	}

	serverMessage := &chat.ServerMessage{
		Type:       chat.MsgParty,
		SenderGUID: s.GUID(),
		Body:       party.Body,
	}

	s.Group.Lock()

	for _, v := range s.Group.Members {
		partyMember, err := s.Server.GetSessionByGUID(v)
		if err == nil {
			partyMember.Send(serverMessage)
		}
	}

	s.Group.Unlock()
}

func (s *Session) NotifyGroup() {
	if s.Group != nil {
		s.Group.Lock()
		for _, v := range s.Group.Members {
			if v != s.GUID() {
				if partyMember, err := s.Server.GetSessionByGUID(v); err == nil {
					partyMember.SendPartyMemberStats(statsMaskAll, s)
				}
			}
		}
		s.Group.Unlock()
	}
}
