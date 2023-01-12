package voip

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type FeatureStatus uint8

const (
	Disabled FeatureStatus = iota
	DontAutoIgnore
	AutoIgnore
)

// features := packet.NewWorldPacket(packet.SMSG_FEATURE_SYSTEM_STATUS)
// features.WriteByte(2) // Can complain (0 = disabled, 1 = enabled, don't auto ignore, 2 = enabled, auto ignore)
// features.WriteByte(1) // voice chat toggled

type FeatureSystemStatus struct {
	Status       FeatureStatus
	VoiceEnabled bool
}

func (f *FeatureSystemStatus) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_FEATURE_SYSTEM_STATUS
	out.WriteByte(byte(f.Status))
	out.WriteBool(f.VoiceEnabled)
	return nil
}
