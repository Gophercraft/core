package voip

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type FeatureStatus uint8

const (
	Disabled FeatureStatus = iota
	DontAutoIgnore
	AutoIgnore
)

// features := packet.NewWorldPacket(packet.SMSG_FEATURE_SYSTEM_STATUS)
// features.WriteUint8(2) // Can complain (0 = disabled, 1 = enabled, don't auto ignore, 2 = enabled, auto ignore)
// features.WriteUint8(1) // voice chat toggled

type FeatureSystemStatus struct {
	Status       FeatureStatus
	VoiceEnabled bool
}

func (f *FeatureSystemStatus) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_FEATURE_SYSTEM_STATUS
	out.WriteUint8(byte(f.Status))
	out.WriteBool(f.VoiceEnabled)
	return nil
}
