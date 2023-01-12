package rpcnet

func Code(s Status) *StatusMsg {
	return &StatusMsg{
		Status: s,
	}
}
