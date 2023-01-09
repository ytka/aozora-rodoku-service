package database

import (
	"aozorarodoku-service/base/uuid"
	"testing"
	"time"
)

var c = Content{Id: uuid.NewMockUUID(1),
	TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"}

func TestEqualEntity(t *testing.T) {
	type args struct {
		x Content
		y Content
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same",
			args{
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
			},
			true,
		},
		{"diff",
			args{
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "xああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
			},
			false,
		},
		{"same objects",
			args{c, c},
			true,
		},
		{"same ignore CreatedAt and UpdatedAt",
			args{
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒",
					CreatedAt: time.Now(), UpdatedAt: time.Now()},
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualEntity(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("EqualEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffEntity(t *testing.T) {
	type args struct {
		x Content
		y Content
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"same",
			args{
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
			},
			"",
		},
		{"same objects",
			args{c, c},
			"",
		},
		{"same ignore CreatedAt and UpdatedAt",
			args{
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒",
					CreatedAt: time.Now(), UpdatedAt: time.Now()},
				Content{Id: uuid.NewMockUUID(1),
					TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffEntity(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("DiffEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
