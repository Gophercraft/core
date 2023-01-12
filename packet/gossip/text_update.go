package gossip

import (
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type TextUpdate struct {
	Entry  uint32
	Locale i18n.Locale
	Text   *models.NPCText
}

func (tu *TextUpdate) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_NPC_TEXT_UPDATE
	out.WriteUint32(tu.Entry)

	if tu.Text == nil {
		for x := 0; x < 8; x++ {
			out.WriteFloat32(0)
			out.WriteCString("Hail, $r.")
			out.WriteCString("Hail, $r.")
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
		}
	} else {
		x := 0

		for ; x < len(tu.Text.Opts); x++ {
			opt := tu.Text.Opts[x]
			out.WriteFloat32(opt.Prob)
			text := opt.Text.GetLocalized(tu.Locale)
			out.WriteCString(text)
			out.WriteCString(text)
			out.WriteUint32(opt.Lang)

			em := 0

			for ; em < len(opt.Emote); em += 2 {
				e := opt.Emote[em]
				out.WriteUint32(e.Delay)
				out.WriteUint32(e.ID)
			}

			for ; em < 6; em += 2 {
				out.WriteUint32(0)
				out.WriteUint32(0)
			}

		}

		for ; x < 8; x++ {
			out.WriteFloat32(0)
			out.WriteCString("")
			out.WriteCString("")
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
		}
	}

	return nil
}
