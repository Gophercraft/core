package realm

import (
	"encoding/hex"
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/superp00t/etc"
)

func (s *Session) shex(pt packet.WorldType, data string) {
	hx, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	p := packet.NewWorldPacket(pt)
	p.Buffer = etc.FromBytes(hx)
	s.SendPacket(p)
}

func (s *Session) Cast(spellID uint32, target Object) error {
	return nil
}

func (s *Session) UpdateAuraDuration(slot int) {
	state := s.GetAuras()

	if slot >= len(state.Auras) {
		return
	}

	au := state.Auras[slot]
	timeLeft := au.ExpiryTime.Sub(au.AppliedTime)

	// s.Send(&spell.AuraDurationUpdate{
	// 	Slot:     uint8(slot),
	// 	Duration: uint32(au.Duration() / time.Millisecond),
	// })

	s.Send(&spell.AuraDurationUpdate{
		Slot:     uint8(slot),
		Duration: uint32(timeLeft / time.Millisecond),
	})
}

func (s *Session) UpdateAllAuraDurations() {
	state := s.GetAuras()

	auraSize := len(state.Auras)

	for x := 0; x < auraSize; x++ {
		s.UpdateAuraDuration(x)
	}
}

func (m *Map) LookupSpellInfo(id uint32, spelldata **dbdefs.Ent_Spell, spellvisual **dbdefs.Ent_SpellVisual, casttimes **dbdefs.Ent_SpellCastTimes) {
	m.Phase.Server.DB.Lookup(wdb.BucketKeyUint32ID, id, spelldata)

	if *spelldata != nil {
		vis := (*spelldata).SpellVisualID[0]
		if vis != 0 {
			if spellvisual != nil {
				m.Phase.Server.DB.Lookup(wdb.BucketKeyUint32ID, uint32(vis), spellvisual)
			}
		}

		timeIndex := (*spelldata).CastingTimeIndex
		if timeIndex != 0 {
			if casttimes != nil {
				m.Phase.Server.DB.Lookup(wdb.BucketKeyUint32ID, uint32(timeIndex), casttimes)
			}
		}
	}
}

func (s *Server) GetSpellData(id uint32) *dbdefs.Ent_Spell {
	var spelldata *dbdefs.Ent_Spell
	s.DB.Lookup(wdb.BucketKeyUint32ID, id, &spelldata)
	return spelldata
}

func (s *Server) GetSpellAttributes(sp *dbdefs.Ent_Spell) (*update.Bitmask, error) {
	return spell.LoadAttributes(s.Build(), sp)
}
