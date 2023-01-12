package main

import (
	"fmt"
	"reflect"

	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/text"
)

type NPCText struct {
	ID      uint32  `xorm:"'id'"`
	Text0_0 string  `xorm:"'text0_0'"`
	Text0_1 string  `xorm:"'text0_1'"`
	Lang0   uint32  `xorm:"'lang0'"`
	Prob0   float32 `xorm:"'prob0'"`
	Em0_0   uint32  `xorm:"'em0_0'"`
	Em0_1   uint32  `xorm:"'em0_1'"`
	Em0_2   uint32  `xorm:"'em0_2'"`
	Em0_3   uint32  `xorm:"'em0_3'"`
	Em0_4   uint32  `xorm:"'em0_4'"`
	Em0_5   uint32  `xorm:"'em0_5'"`
	Text1_0 string  `xorm:"'text1_0'"`
	Text1_1 string  `xorm:"'text1_1'"`
	Lang1   uint32  `xorm:"'lang1'"`
	Prob1   float32 `xorm:"'prob1'"`
	Em1_0   uint32  `xorm:"'em1_0'"`
	Em1_1   uint32  `xorm:"'em1_1'"`
	Em1_2   uint32  `xorm:"'em1_2'"`
	Em1_3   uint32  `xorm:"'em1_3'"`
	Em1_4   uint32  `xorm:"'em1_4'"`
	Em1_5   uint32  `xorm:"'em1_5'"`
	Text2_0 string  `xorm:"'text2_0'"`
	Text2_1 string  `xorm:"'text2_1'"`
	Lang2   uint32  `xorm:"'lang2'"`
	Prob2   float32 `xorm:"'prob2'"`
	Em2_0   uint32  `xorm:"'em2_0'"`
	Em2_1   uint32  `xorm:"'em2_1'"`
	Em2_2   uint32  `xorm:"'em2_2'"`
	Em2_3   uint32  `xorm:"'em2_3'"`
	Em2_4   uint32  `xorm:"'em2_4'"`
	Em2_5   uint32  `xorm:"'em2_5'"`
	Text3_0 string  `xorm:"'text3_0'"`
	Text3_1 string  `xorm:"'text3_1'"`
	Lang3   uint32  `xorm:"'lang3'"`
	Prob3   float32 `xorm:"'prob3'"`
	Em3_0   uint32  `xorm:"'em3_0'"`
	Em3_1   uint32  `xorm:"'em3_1'"`
	Em3_2   uint32  `xorm:"'em3_2'"`
	Em3_3   uint32  `xorm:"'em3_3'"`
	Em3_4   uint32  `xorm:"'em3_4'"`
	Em3_5   uint32  `xorm:"'em3_5'"`
	Text4_0 string  `xorm:"'text4_0'"`
	Text4_1 string  `xorm:"'text4_1'"`
	Lang4   uint32  `xorm:"'lang4'"`
	Prob4   float32 `xorm:"'prob4'"`
	Em4_0   uint32  `xorm:"'em4_0'"`
	Em4_1   uint32  `xorm:"'em4_1'"`
	Em4_2   uint32  `xorm:"'em4_2'"`
	Em4_3   uint32  `xorm:"'em4_3'"`
	Em4_4   uint32  `xorm:"'em4_4'"`
	Em4_5   uint32  `xorm:"'em4_5'"`
	Text5_0 string  `xorm:"'text5_0'"`
	Text5_1 string  `xorm:"'text5_1'"`
	Lang5   uint32  `xorm:"'lang5'"`
	Prob5   float32 `xorm:"'prob5'"`
	Em5_0   uint32  `xorm:"'em5_0'"`
	Em5_1   uint32  `xorm:"'em5_1'"`
	Em5_2   uint32  `xorm:"'em5_2'"`
	Em5_3   uint32  `xorm:"'em5_3'"`
	Em5_4   uint32  `xorm:"'em5_4'"`
	Em5_5   uint32  `xorm:"'em5_5'"`
	Text6_0 string  `xorm:"'text6_0'"`
	Text6_1 string  `xorm:"'text6_1'"`
	Lang6   uint32  `xorm:"'lang6'"`
	Prob6   float32 `xorm:"'prob6'"`
	Em6_0   uint32  `xorm:"'em6_0'"`
	Em6_1   uint32  `xorm:"'em6_1'"`
	Em6_2   uint32  `xorm:"'em6_2'"`
	Em6_3   uint32  `xorm:"'em6_3'"`
	Em6_4   uint32  `xorm:"'em6_4'"`
	Em6_5   uint32  `xorm:"'em6_5'"`
	Text7_0 string  `xorm:"'text7_0'"`
	Text7_1 string  `xorm:"'text7_1'"`
	Lang7   uint32  `xorm:"'lang7'"`
	Prob7   float32 `xorm:"'prob7'"`
	Em7_0   uint32  `xorm:"'em7_0'"`
	Em7_1   uint32  `xorm:"'em7_1'"`
	Em7_2   uint32  `xorm:"'em7_2'"`
	Em7_3   uint32  `xorm:"'em7_3'"`
	Em7_4   uint32  `xorm:"'em7_4'"`
	Em7_5   uint32  `xorm:"'em7_5'"`
}

func (nt NPCText) TableName() string {
	return "npc_text"
}

func (nt NPCText) GetOption(idx int) models.NPCTextOption {
	var to models.NPCTextOption
	val := reflect.ValueOf(nt)
	if val.IsZero() {
		return to
	}

	to.Text = i18n.GetEnglish(val.FieldByName(fmt.Sprintf("Text%d_0", idx)).String())
	if to.Text == nil {
		to.Text = i18n.GetEnglish(val.FieldByName(fmt.Sprintf("Text%d_1", idx)).String())
	}

	to.Lang = uint32(val.FieldByName(fmt.Sprintf("Lang%d", idx)).Uint())
	to.Prob = float32(val.FieldByName(fmt.Sprintf("Prob%d", idx)).Float())

	for x := 0; x < 6; x += 2 {
		emoteDelay := uint32(val.FieldByName(fmt.Sprintf("Em%d_%d", idx, x)).Uint())

		strName := fmt.Sprintf("Em%d_%d", idx, x+1)
		emote := uint32(val.FieldByName(strName).Uint())

		if emoteDelay != 0 || emote != 0 {
			to.Emote = append(to.Emote, models.NPCTextOptionEmote{
				Delay: emoteDelay,
				ID:    emote,
			})
		}
	}

	return to
}

func extractNPCText() {
	fl := openFile("DB/NPCText.txt")
	wr := text.NewEncoder(fl)

	var npcText []NPCText

	err := DB.Find(&npcText)
	if err != nil {
		panic(err)
	}

	for _, text := range npcText {
		var nt models.NPCText

		nt.ID = fmt.Sprintf("nt:%d", text.ID)
		for x := 0; x < 8; x++ {
			opt := text.GetOption(x)

			if reflect.ValueOf(opt).IsZero() == false {
				nt.Opts = append(nt.Opts, opt)
			}
		}

		if err := wr.Encode(nt); err != nil {
			panic(err)
		}
	}

	fl.Close()
}
