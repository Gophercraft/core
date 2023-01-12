package realm

import (
	"github.com/Gophercraft/core/packet/raid"
)

func (s *Session) HandleRequestRaidInfo() {
	var info raid.InstanceInfo
	s.Send(&info)
}
