package aozorarodokuweb

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

type Content struct {
	TitleRuby      string `csv:"作品ルビ"`
	Title          string `csv:"作品"`
	AuthorRuby     string `csv:"作家ルビ"`
	Author         string `csv:"作家"`
	SpeakerRuby    string `csv:"読み手ルビ"`
	Speaker        string `csv:"読み手"`
	FileName       string `csv:"mp3"`
	NewArrivalDate string `csv:"新着"`
	Time           string `csv:"時間"`
}

func ParseContents(data []byte) ([]Content, error) {
	dec, err := csvutil.NewDecoder(csv.NewReader(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	var contents []Content
	for {
		content := Content{}
		if err := dec.Decode(&content); err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}
