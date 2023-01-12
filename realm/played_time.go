package realm

import (
	"time"

	"github.com/Gophercraft/core/packet/account"
)

func (s *Session) HandlePlayedTimeRequest() {
	var pt account.PlayedTime

	nao := time.Now()

	pt.TotalTime = s.Char.TotalPlayedTime + nao.Sub(s.JoinedWorldAt)
	pt.LevelTime = s.Char.LevelPlayedTime + nao.Sub(s.LevelJoinTime)

	s.Send(&pt)
}

func (s *Session) savePlayedTime() {
	if s.HasState(InWorld) {
		nao := time.Now()
		timePlayed := nao.Sub(s.JoinedWorldAt)
		s.Char.TotalPlayedTime += timePlayed
		timePlayedAtLevel := nao.Sub(s.LevelJoinTime)
		s.Char.LevelPlayedTime += timePlayedAtLevel
		s.DB().Where("id = ?", s.Char.ID).Cols("playtime_total", "playtime_level").Update(s.Char)
	}
}
