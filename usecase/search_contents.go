package usecase

import (
	"aozorarodoku-service/database"
	"context"

	"github.com/jmoiron/sqlx"
)

func SearchContents(ctx context.Context, db *sqlx.DB, searchText string) ([]database.Content, error) {
	var contents []database.Content
	err := database.DoTransaction(ctx, db, func(tx *sqlx.Tx) error {
		var err error
		contents, err = database.FindContents(ctx, db, searchText)
		return err
	})
	if err != nil {
		return nil, err
	}
	return contents, nil
}
