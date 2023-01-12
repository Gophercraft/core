package realm

import (
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/i18n"
)

type Relation uint8

const (
	NoRelation Relation = 0
	RelationPassive
	RelationFriendly
	RelationDisliked
	RelationHated
)

type Reason uint8

const (
	NoReason Reason = iota
	ReasonHelpedUs
	ReasonHurtUs
	ReasonAttackedUs
)

type StaticTeamRelationship struct {
	TeamID   string
	Relation Relation
	Reason   Reason
}

type TeamRelationships struct {
	Statics []StaticTeamRelationship
}

type FactionTeam struct {
	Template *dbdefs.Ent_FactionTemplate
	Faction  *dbdefs.Ent_Faction
}

func (f *FactionTeam) ID() string {
	return f.Faction.Name_lang.String()
}

type Team interface {
	ID() string
	Name() i18n.Text
	Color() string
	Relationships() *TeamRelationships
}

type TeamPlayer interface {
	GUID() guid.GUID
	Teams() []Team
}

func (ps *PlayerSession) Teams() []Team {
	return nil
}

func (m *Map) Hostile(u1, u2 TeamPlayer) bool {
	// team1 := u1.Teams()
	// team2 := u2.Teams()
	// return team1 != team2
	return true
}

func (m *Map) CanFight(u1, u2 TeamPlayer) bool {
	// if u1.

	// if m.Phase.Server
	return true
}

// type facTemplate uint32

// type TeamColors uint8

// type FactionTeam struct {
// 	Colors TeamColors
// }

// type Raid struct {
// 	Group
// }

// func (r *Raid) Init() {
// 	r.Group.GroupType = party.GroupRaid
// }

// type Team interface {
// 	Relation(s *Server, t TeamMember) Relation
// }

// func (ft facTemplate) Relation(s *Server, t TeamMember) Relation {
// 	var fac dbdefs.Ent_FactionTemplate
// 	s.DB.Lookup(wdb.BucketKeyUint32ID, uint32(ft), &fac)
// }

// func (s *Session) Team() Team {
// 	return st
// }
