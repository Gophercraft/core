// Code generated by "stringer -type=MessageType"; DO NOT EDIT.

package grunt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LogonChallenge-0]
	_ = x[LogonProof-1]
	_ = x[ReconnectChallenge-2]
	_ = x[ReconnectProof-3]
	_ = x[RealmList-16]
	_ = x[XferInitiate-48]
	_ = x[XferData-49]
	_ = x[XferAccept-50]
	_ = x[XferResume-51]
	_ = x[XferCancel-52]
}

const (
	_MessageType_name_0 = "LogonChallengeLogonProofReconnectChallengeReconnectProof"
	_MessageType_name_1 = "RealmList"
	_MessageType_name_2 = "XferInitiateXferDataXferAcceptXferResumeXferCancel"
)

var (
	_MessageType_index_0 = [...]uint8{0, 14, 24, 42, 56}
	_MessageType_index_2 = [...]uint8{0, 12, 20, 30, 40, 50}
)

func (i MessageType) String() string {
	switch {
	case i <= 3:
		return _MessageType_name_0[_MessageType_index_0[i]:_MessageType_index_0[i+1]]
	case i == 16:
		return _MessageType_name_1
	case 48 <= i && i <= 52:
		i -= 48
		return _MessageType_name_2[_MessageType_index_2[i]:_MessageType_index_2[i+1]]
	default:
		return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}