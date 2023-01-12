package realm

import (
	"time"

	"github.com/Gophercraft/core/packet/synctime"
)

func (s *Session) HandleQueryTime(q *synctime.QueryTime) {
	s.Send(&synctime.QueryTimeResponse{
		Unix: int32(time.Now().Unix()),
	})
}
