package realm

import (
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/log"
)

const (
	SpellMgrPtag = "Gophercraft/core/realm/spell_manager"
)

type MoveTick struct {
}

type SpellManager struct {
	PlayerMoved    chan MoveTick
	StartCast      chan *spell.Cast
	GoCast         chan *spell.Cast
	PendingCast    *spell.Cast
	RechargeTicker *time.Ticker
}

// get rid of any spell state in this Session.
func (s *Session) cleanupSpellManager() {
	close(s.SpellManager.PlayerMoved)
	close(s.SpellManager.StartCast)
	close(s.SpellManager.GoCast)
	close(s.AuraState.AuraExpire)
	s.AuraState.AuraExpire = nil
	for _, au := range s.AuraState.Auras {
		if au.ExpiryTimer != nil {
			au.ExpiryTimer.Stop()
			au.ExpiryTimer = nil
		}
	}
	s.SpellManager = nil
}

func (s *Session) initSpellManager() {
	s.SpellManager = new(SpellManager)
	s.SpellManager.PlayerMoved = make(chan MoveTick)
	s.SpellManager.StartCast = make(chan *spell.Cast)
	s.SpellManager.GoCast = make(chan *spell.Cast)
	s.AuraState.AuraExpire = make(chan *Aura, 16)
	s.SpellManager.RechargeTicker = time.NewTicker(5 * time.Second)
	s.CreateProcess(SpellMgrPtag, (*Session).spellManagerProcess)
}

func (s *Session) spellCastDirect(cc *spell.Cast) {
	if err := s.CanCast(cc); err != nil {
		s.Send(err)
		return
	}

	castRes := &spell.CastResult{
		CastID:  byte(cc.CastID),
		SpellID: cc.Spell,
		Status:  spell.Success,
	}

	s.Send(castRes)

	castDuration, err := s.Map().StartCast(cc, s)
	if err != nil {
		log.Warn(err)
		return
	}

	if castDuration == 0 {
		if err := s.Map().GoCast(cc, s); err != nil {
			log.Warn(err)
		}
	} else {
		s.SpellManager.PendingCast = cc

		go func() {
			time.Sleep(castDuration)
			if s.SpellManager != nil {
				s.SpellManager.GoCast <- cc
			}
		}()
	}
}

func (s *Session) InCombat() bool {
	return false
}

func (s *Session) shouldWeCancelOnMove(spelldata *dbdefs.Ent_Spell) bool {
	// attr, err := s.Server.GetSpellAttributes(spelldata)
	// if err != nil {
	// 	panic(err)
	// }
	// if attr.Enabled(uint32(spell.AttrEx2_AutorepeatFlag)) {

	// }

	return false
}

func (s *Session) cancelPendingCast() {
	s.SpellManager.PendingCast = nil
}

func (s *Session) spellManagerProcess(cancel <-chan bool) {
	for {
		select {
		case <-cancel:
			s.cleanupSpellManager()
			return
		case <-s.SpellManager.PlayerMoved:
			// Player moved. Check if this is a CANCELABLE offense!
			if s.SpellManager.PendingCast != nil {
				spelldata := s.Server.GetSpellData(s.SpellManager.PendingCast.Spell)
				if s.shouldWeCancelOnMove(spelldata) {
					s.cancelPendingCast()
				}
			}
		case cast := <-s.SpellManager.StartCast:
			s.spellCastDirect(cast)
		case cast := <-s.SpellManager.GoCast:
			// This is gonna fire even if player canceled the spell cast.
			// PendingCast == cast means it wasn't canceled
			if s.SpellManager.PendingCast != cast {
				s.SpellManager.PendingCast = nil
				if err := s.Map().GoCast(cast, s); err != nil {
					log.Warn(err)
				}
			}
		case <-s.SpellManager.RechargeTicker.C:
			s.rechargeHeartbeat()
		case au := <-s.AuraState.AuraExpire:
			s.Map().Unaura(s, au.ID, -1, true)
		}
	}
}
