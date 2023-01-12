package realm

import (
	"fmt"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"
)

func (s *Session) KnowsAbility(spellID uint32) bool {
	if !s.HasState(InWorld) {
		return false
	}
	n, err := s.DB().Where("player = ?", s.PlayerID()).Where("spell = ?", spellID).Count(new(models.LearnedAbility))
	if err != nil {
		panic(err)
	}
	return n >= 1
}

func (s *Session) LearnAbility(spellID uint32) error {
	if s.KnowsAbility(spellID) {
		return fmt.Errorf("realm: ability %d already known!", spellID)
	}

	var spelldata *dbdefs.Ent_Spell

	s.DB().Lookup(wdb.BucketKeyUint32ID, spellID, &spelldata)

	if spelldata == nil {
		return fmt.Errorf("realm: ability %d does not exist", spellID)
	}

	s.DB().Insert(models.LearnedAbility{
		Player: s.PlayerID(),
		Spell:  spellID,
		Active: true,
	})

	s.Send(&spell.Learned{
		Page: spell.Page{
			Spell: spellID,
		},
	})

	return nil
}

func (s *Session) SendSpellList() {
	// s.shex(packet.SMSG_INITIAL_SPELLS, "0060002a8500006f8200006c820000d77d000074760000bb620000b2620000b06200009a620000946200008f620000896200007c62000073620000705d00000b5600009454000093540000087400004650000045500000064f0000863100002b8500004850000047500000412d0000a52300009c230000752300009469000076230000510000009006000078620000370c0000cb00000089040000ee020000d501000047000000c6000000e63c0000c5000000c40000007c860000c80000006b000000c7000000670300000a0100009e020000621c00009c0400001a5900000a020000ca0000009d0200007e140000c2200000530d0000e30000009a090000a4020000630100003a2d0000cb1900009313000000080000ca0b00004e0900009909000008010000ea0b00007f0a0000050a0000cc0a0000070a0000cc000000af090000cb0c0000250d0000b7060000b514000059180000a20200006618000067180000957600004d1900004e190000212200009a190000631c000043480000bb1c00000000")

	var spells []models.LearnedAbility
	s.DB().Where("player = ?", s.PlayerID()).Find(&spells)

	// to talk

	spells = append(spells, models.LearnedAbility{
		Player: s.PlayerID(),
		Spell:  668,
	}, models.LearnedAbility{
		Player: s.PlayerID(),
		Spell:  669,
	})

	book := &spell.Book{}

	book.Spells = make([]spell.Page, len(spells))

	for i := range spells {
		book.Spells[i].Spell = spells[i].Spell
		book.Spells[i].Slot = uint16(spells[i].Slot)
	}

	s.Send(book)
}

func (s *Session) HandleNewSpellSlot(sl *spell.NewSlot) {
	log.Dump("new slot", sl)
	var sp models.LearnedAbility
	sp.Player = s.PlayerID()
	sp.Slot = int(sl.Index)
	sp.Spell = sl.Spell

	s.DB().Where("player = ?", s.PlayerID()).Where("spell = ?", sl.Spell).Cols("slot").Update(&sp)
}
