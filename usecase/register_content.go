package usecase

import (
	"aozorarodoku-service/aozorarodokuweb"
	"aozorarodoku-service/database"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func csvContentToDbContent(c aozorarodokuweb.Content) (database.Content, error) {
	return database.Content{Id: uuid.NewString(),
		TitleRuby: c.TitleRuby, Title: c.Title,
		AuthorRuby: c.AuthorRuby, Author: c.Author,
		SpeakerRuby: c.SpeakerRuby, Speaker: c.Speaker,
		FileName: c.FileName, NewArrivalDate: c.NewArrivalDate,
		Time: c.Time,
	}, nil

}

func csvContentsToDbContents(csvContents []aozorarodokuweb.Content) ([]database.Content, error) {
	contents := make([]database.Content, len(csvContents))
	for i, v := range csvContents {
		if c, err := csvContentToDbContent(v); err != nil {
			return nil, err
		} else {
			contents[i] = c
		}
	}
	return contents, nil
}

func RegisterContentsToDB(ctx context.Context, db *sqlx.DB) error {
	csvContents, err := aozorarodokuweb.FetchContents()
	if err != nil {
		return err
	}
	contents, err := csvContentsToDbContents(csvContents)
	if err != nil {
		return err
	}

	return database.DoTransaction(ctx, db, func(tx *sqlx.Tx) error {
		database.TruncateContents(ctx, tx)
		return database.InsertContent(ctx, tx, contents...)
	})
}
