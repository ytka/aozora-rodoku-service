package aozorarodokuweb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ContentsFixture = `作品ルビ,作品,作家ルビ,作家,読み手ルビ,読み手,mp3,新着,時間
ああき,ア、秋,だざいおさむ,太宰 治,なかむらあきよ,中村 昭代,rd260.mp3,,6分52秒
ああしんど,ああしんど,いけだしょうえん,池田 焦園,いけどみか,池戸 美香,rd158.mp3,,2分40秒
ああとうきょうはくいだおれ,ああ東京は食い倒れ,ふるかわろっぱ,古川 緑波,のむらようじ,野村 洋二,rd164.mp3,2020-11-01,9分25秒
`

func Test_ParseContents(t *testing.T) {
	contents, err := ParseContents([]byte(ContentsFixture))
	assert.NoError(t, err)
	assert.Equal(t, []Content{
		{TitleRuby: "ああき", Title: "ア、秋", AuthorRuby: "だざいおさむ", Author: "太宰 治", SpeakerRuby: "なかむらあきよ", Speaker: "中村 昭代", FileName: "rd260.mp3", NewArrivalDate: "", Time: "6分52秒"},
		{TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
		{TitleRuby: "ああとうきょうはくいだおれ", Title: "ああ東京は食い倒れ", AuthorRuby: "ふるかわろっぱ", Author: "古川 緑波", SpeakerRuby: "のむらようじ", Speaker: "野村 洋二", FileName: "rd164.mp3", NewArrivalDate: "2020-11-01", Time: "9分25秒"},
	}, contents)
}
