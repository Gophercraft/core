package realm

import (
	"sync"

	"github.com/Gophercraft/core/bnet"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/packet/auth"
)

type WaitQueue struct {
	sync.Mutex
	Q []*Session
}

func (wq *WaitQueue) EnqueueSession(s *Session) {
	wq.Lock()
	index := len(wq.Q)
	wq.Q = append(wq.Q, s)
	s.SendInitWaitQueue(index)
	wq.Unlock()
}

func (wq *WaitQueue) Dequeue() {
	wq.Lock()
	defer wq.Unlock()

	if len(wq.Q) == 0 {
		return
	}

	dq := wq.Q[0]
	dq.CompleteLogin()
	wq.removeSessionPtr(dq)
}

func (wq *WaitQueue) DeleteSession(s *Session) {
	wq.Lock()
	defer wq.Unlock()
	wq.removeSessionPtr(s)
}

func (wq *WaitQueue) removeSessionPtr(s *Session) {
	index := -1
	for i, qs := range wq.Q {
		if qs == s {
			index = i
			break
		}
	}

	if index < 0 {
		return
	}

	wq.Q = append(wq.Q[:index], wq.Q[index+1:]...)
}

func (s *Session) EnterWaitQueue() {
	s.SetState(Waiting)

	if s.Tier >= rpcnet.Tier_Privileged {
		s.CompleteLogin()
		return
	}

	if !s.Server.AtCapacity() {
		s.CompleteLogin()
		return
	}

	s.Server.WaitQueue.EnqueueSession(s)
}

func (s *Session) RemoveFromWaitQueue() {
	s.Server.WaitQueue.DeleteSession(s)
}

func (ws *Server) PlayerCapacity() int {
	return 2 //
}

func (rs *Server) ActivePlayers() int {
	rs.GuardPlayerList.Lock()
	n := len(rs.PlayerList)
	rs.GuardPlayerList.Unlock()
	return n
}

func (ws *Server) AtCapacity() bool {
	cap := ws.PlayerCapacity()
	// No limit on players, no queue
	if cap <= 0 {
		return false
	}

	return ws.ActivePlayers() >= cap
}

func (s *Session) SendInitWaitQueue(position int) {
	s.Send(&auth.Response{
		Result:      bnet.ERROR_OK,
		SuccessInfo: &auth.SuccessInfo{},
		WaitInfo: &auth.WaitInfo{
			WaitCount: uint32(position),
		},
	})
}

func (s *Session) SendUpdateWaitQueuePosition(position uint32) {
	s.Send(&auth.Response{
		Result: bnet.ERROR_OK,
		WaitInfo: &auth.WaitInfo{
			WaitCount: uint32(position),
		},
	})
}
