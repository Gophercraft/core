// package chat contains packets relating to the in-game chatbox feature.
package chat

import (
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc"
)

func DecodeChatString(build vsn.Build, in *etc.Buffer) string {
	if build.RemovedIn(vsn.V1_12_1) {
		return in.ReadCString()
	}
	snd := in.ReadUint32()
	return in.ReadFixedString(int(snd))
}

func EncodeChatString(build vsn.Build, out *etc.Buffer, str string) {
	if build.RemovedIn(vsn.V1_12_1) {
		out.WriteCString(str)
		return
	}
	out.WriteUint32(uint32(len(str) + 1))
	out.Write([]byte(str))
	out.WriteByte(0)
}
